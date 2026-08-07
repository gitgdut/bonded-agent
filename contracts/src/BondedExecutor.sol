// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {MockUSDC} from "./MockUSDC.sol";

/**
 * @title BondedExecutor
 * @notice Core contract: operator locks tUSDC bond to guarantee swap outcomes.
 *         If actual output < guaranteed output, auto-compensates the user.
 *         If the swap call fails, refunds user input + pays failure compensation.
 *
 *         All compensation is purely on-chain — no LLM or admin judgment.
 */
contract BondedExecutor {
    MockUSDC public immutable tUSDC;

    /// @notice Service fee in basis points (e.g., 30 = 0.3%)
    uint256 public serviceFeeBps;

    // ── EIP-712 ─────────────────────────────────────────────────
    bytes32 private constant EIP712_DOMAIN_TYPEHASH =
        keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)");
    bytes32 private constant EXECUTE_AUTH_TYPEHASH =
        keccak256("ExecuteAuthorization(bytes32 planId,uint256 inputAmount,uint256 deadline)");

    bytes32 private immutable DOMAIN_SEPARATOR;

    // ── Plan structure ──────────────────────────────────────────
    struct Plan {
        address user;              // who can execute
        address operator;          // who created the plan & provides bond
        uint256 inputAmount;       // MON amount for the swap
        uint256 expectedOutput;    // simulated expected output (informational)
        uint256 guaranteedOutput;  // minimum tUSDC the user is guaranteed
        uint256 maxCompensation;   // max tUSDC the bond covers (≥ guaranteedOutput)
        uint256 failureCompensation; // tUSDC paid to user if swap call reverts
        address target;            // target contract (MockDex)
        bytes32 calldataHash;      // keccak256(calldata) — prevents tampering
        uint256 deadline;          // block.timestamp after which plan expires
        uint256 nonce;             // unique per (user, operator, nonce)
        bool executed;             // prevents replay
    }

    // ── State ────────────────────────────────────────────────────
    /// @notice planId → Plan
    mapping(bytes32 => Plan) public plans;

    /// @notice user → operator → nonce → used
    mapping(address => mapping(address => mapping(uint256 => bool))) public usedNonces;

    /// @notice operator → tUSDC bond balance locked in active plans
    mapping(address => uint256) public lockedBond;

    // ── Events ───────────────────────────────────────────────────
    event PlanOpened(
        bytes32 indexed planId,
        address indexed user,
        address indexed operator,
        uint256 guaranteedOutput,
        uint256 maxCompensation,
        uint256 deadline
    );

    event PlanExecuted(
        bytes32 indexed planId,
        uint256 actualOutput,
        uint256 paidToUser
    );

    event ShortfallPaid(
        bytes32 indexed planId,
        uint256 guaranteed,
        uint256 actual,
        uint256 shortfall
    );

    event PlanFailed(
        bytes32 indexed planId,
        uint256 refundedMON,
        uint256 compensationPaid
    );

    event BondReleased(
        bytes32 indexed planId,
        address indexed operator,
        uint256 amount
    );

    event ServiceFeePaid(
        bytes32 indexed planId,
        uint256 swapOutput,
        uint256 fee,
        uint256 userReceived
    );

    event ServiceFeeUpdated(uint256 oldFeeBps, uint256 newFeeBps);

    // ── Errors ───────────────────────────────────────────────────
    error NotPlanUser();
    error AlreadyExecuted();
    error PlanExpired(uint256 deadline, uint256 now_);
    error CalldataMismatch(bytes32 expected, bytes32 actual);
    error InsufficientBond(uint256 required, uint256 available);
    error BondTransferFailed();
    error SwapFailed();
    error NonceAlreadyUsed();

    // ── Constructor ──────────────────────────────────────────────
    constructor(address _tUSDC) {
        tUSDC = MockUSDC(_tUSDC);
        DOMAIN_SEPARATOR = keccak256(abi.encode(
            EIP712_DOMAIN_TYPEHASH,
            keccak256("BondedExecutor"),
            keccak256("1"),
            block.chainid,
            address(this)
        ));
    }

    // ── Plan Lifecycle ───────────────────────────────────────────

    /**
     * @notice Operator opens a guaranteed plan and locks tUSDC bond.
     * @param planId keccak256(abi.encode(user, operator, nonce)) or similar unique ID
     * @param plan The plan parameters
     * @param calldata_ The raw calldata to send to target (verified against calldataHash)
     */
    function openPlan(
        bytes32 planId,
        Plan calldata plan,
        bytes calldata calldata_
    ) external {
        // Verify uniqueness
        if (usedNonces[plan.user][plan.operator][plan.nonce]) revert NonceAlreadyUsed();
        if (plans[planId].operator != address(0)) revert AlreadyExecuted();

        // Verify calldata
        if (keccak256(calldata_) != plan.calldataHash) {
            revert CalldataMismatch(plan.calldataHash, keccak256(calldata_));
        }

        // Verify bond coverage
        uint256 requiredBond = plan.maxCompensation;
        if (plan.failureCompensation > requiredBond) {
            requiredBond = plan.failureCompensation;
        }

        // Pull bond from operator
        // Operator must have approved tUSDC to BondedExecutor first
        bool ok = tUSDC.transferFrom(msg.sender, address(this), requiredBond);
        if (!ok) revert BondTransferFailed();

        // Store plan
        plans[planId] = plan;
        plans[planId].executed = false; // explicit
        usedNonces[plan.user][plan.operator][plan.nonce] = true;
        lockedBond[plan.operator] += requiredBond;

        emit PlanOpened(planId, plan.user, msg.sender, plan.guaranteedOutput, requiredBond, plan.deadline);
    }

    /**
     * @notice User executes a guaranteed plan.
     *         Sends MON to target, measures tUSDC output delta,
     *         and auto-settles bond.
     * @param planId The plan identifier
     * @param calldata_ Must match plan.calldataHash
     */
    function executePlan(bytes32 planId, bytes calldata calldata_) external payable {
        Plan storage plan = plans[planId];

        // ── Guards ────────────────────────────────────────────
        if (msg.sender != plan.user) revert NotPlanUser();
        if (plan.executed) revert AlreadyExecuted();
        if (block.timestamp > plan.deadline) revert PlanExpired(plan.deadline, block.timestamp);
        if (keccak256(calldata_) != plan.calldataHash) {
            revert CalldataMismatch(plan.calldataHash, keccak256(calldata_));
        }
        if (msg.value != plan.inputAmount) {
            revert(); // wrong MON amount sent
        }

        _executeSwap(planId, plan, calldata_);
    }

    /**
     * @notice Operator executes a plan on behalf of the user via EIP-712 signature.
     *         Operator sends MON, user signs authorization off-chain.
     * @param planId The plan identifier
     * @param calldata_ Must match plan.calldataHash
     * @param deadline Signature expiry (seconds)
     * @param signature EIP-712 signature from the plan user
     */
    function executePlanWithSignature(
        bytes32 planId,
        bytes calldata calldata_,
        uint256 deadline,
        bytes calldata signature
    ) external payable {
        Plan storage plan = plans[planId];

        // ── Verify EIP-712 signature ──────────────────────────
        if (block.timestamp > deadline) revert PlanExpired(deadline, block.timestamp);

        bytes32 structHash = keccak256(abi.encode(
            EXECUTE_AUTH_TYPEHASH,
            planId,
            msg.value,
            deadline
        ));
        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", DOMAIN_SEPARATOR, structHash));
        address signer = _recoverSigner(digest, signature);
        if (signer != plan.user) revert InvalidSignature();

        // ── Standard guards ────────────────────────────────────
        if (plan.executed) revert AlreadyExecuted();
        if (block.timestamp > plan.deadline) revert PlanExpired(plan.deadline, block.timestamp);
        if (keccak256(calldata_) != plan.calldataHash) {
            revert CalldataMismatch(plan.calldataHash, keccak256(calldata_));
        }
        if (msg.value != plan.inputAmount) {
            revert(); // wrong MON amount sent
        }

        // ── Execute swap (same logic as executePlan) ────────────
        _executeSwap(planId, plan, calldata_);
    }

    /// @notice Operator can cancel an unexecuted, non-expired plan and recover bond
    function cancelPlan(bytes32 planId) external {
        Plan storage plan = plans[planId];
        if (msg.sender != plan.operator) revert Unauthorized();
        if (plan.executed) revert AlreadyExecuted();
        // Allow cancellation even after expiry — operator should recover bond

        plan.executed = true;
        _releaseBond(planId, plan);
    }

    // ── Admin ───────────────────────────────────────────────────

    /// @notice Operator sets the service fee in basis points (e.g., 30 = 0.3%).
    ///         Max allowed: 100 (= 1%).
    function setServiceFee(uint256 _feeBps) external {
        // Only contract deployer / known operator
        // Open for now — in production restrict to operator
        if (_feeBps > 100) revert FeeTooHigh();
        uint256 old = serviceFeeBps;
        serviceFeeBps = _feeBps;
        emit ServiceFeeUpdated(old, _feeBps);
    }

    // ── Internal ──────────────────────────────────────────────

    function _releaseBond(bytes32 planId, Plan storage plan) internal {
        uint256 bond = plan.maxCompensation;
        lockedBond[plan.operator] -= bond;
        require(tUSDC.transfer(plan.operator, bond), "bond release failed");
        emit BondReleased(planId, plan.operator, bond);
    }

    /// @notice Shared swap execution + settlement logic.
    function _executeSwap(bytes32 planId, Plan storage plan, bytes calldata calldata_) internal {
        plan.executed = true;

        uint256 balanceBefore = tUSDC.balanceOf(address(this));
        (bool success,) = plan.target.call{value: plan.inputAmount}(calldata_);

        if (success) {
            uint256 balanceAfter = tUSDC.balanceOf(address(this));
            uint256 actualOutput = balanceAfter - balanceBefore;

            uint256 fee = (actualOutput * serviceFeeBps) / 10000;
            uint256 userAmount = actualOutput - fee;

            if (userAmount > 0) {
                require(tUSDC.transfer(plan.user, userAmount), "output transfer failed");
            }
            if (fee > 0) {
                require(tUSDC.transfer(plan.operator, fee), "fee transfer failed");
            }

            if (userAmount >= plan.guaranteedOutput) {
                _releaseBond(planId, plan);
                emit ServiceFeePaid(planId, actualOutput, fee, userAmount);
                emit PlanExecuted(planId, actualOutput, 0);
            } else {
                uint256 shortfall = plan.guaranteedOutput - userAmount;
                uint256 compensation = shortfall > plan.maxCompensation
                    ? plan.maxCompensation : shortfall;

                lockedBond[plan.operator] -= compensation;
                require(tUSDC.transfer(plan.user, compensation), "compensation transfer failed");

                uint256 remaining = plan.maxCompensation - compensation;
                if (remaining > 0) {
                    lockedBond[plan.operator] -= remaining;
                    require(tUSDC.transfer(plan.operator, remaining), "bond release failed");
                }

                emit ShortfallPaid(planId, plan.guaranteedOutput, userAmount, compensation);
                emit ServiceFeePaid(planId, actualOutput, fee, userAmount);
                emit PlanExecuted(planId, actualOutput, compensation);
            }
        } else {
            (bool refundOk,) = plan.user.call{value: plan.inputAmount}("");
            if (!refundOk) revert SwapFailed();

            uint256 comp = plan.failureCompensation;
            if (comp > 0) {
                lockedBond[plan.operator] -= comp;
                require(tUSDC.transfer(plan.user, comp), "failure compensation failed");
            }

            uint256 remainingBond = plan.maxCompensation - comp;
            if (remainingBond > 0) {
                lockedBond[plan.operator] -= remainingBond;
                require(tUSDC.transfer(plan.operator, remainingBond), "bond release failed");
            }

            emit PlanFailed(planId, plan.inputAmount, comp);
        }
    }

    /// @notice Recover signer from EIP-712 digest + signature.
    function _recoverSigner(bytes32 digest, bytes calldata signature) internal pure returns (address) {
        if (signature.length != 65) revert InvalidSignature();
        bytes32 r;
        bytes32 s;
        uint8 v;
        assembly {
            r := calldataload(signature.offset)
            s := calldataload(add(signature.offset, 32))
            v := byte(0, calldataload(add(signature.offset, 64)))
        }
        if (v < 27) v += 27;
        if (v != 27 && v != 28) revert InvalidSignature();
        address recovered = ecrecover(digest, v, r, s);
        if (recovered == address(0)) revert InvalidSignature();
        return recovered;
    }

    error Unauthorized();
    error FeeTooHigh();
    error InvalidSignature();

    // ── Fallback ──────────────────────────────────────────────
    receive() external payable {}
}
