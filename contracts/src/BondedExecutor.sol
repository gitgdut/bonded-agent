// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {MockUSDC} from "./MockUSDC.sol";

/**
 * @title BondedExecutor
 * @notice Operator locks tUSDC bond to guarantee swap outcomes.
 *         V2 changes (Dev1 baseline):
 *         1. coverageFloor = guaranteedOutput - maxCompensation → swap minOutput
 *         2. Try/catch atomicity: inner call checks output, outer handles revert
 *         3. Bond deposit = max(maxCompensation, failureCompensation), paths mutual-exclusive
 *         4. Native MON pull-mode refunds (pendingRefunds)
 *         5. nonce per-operator; planId = keccak256(operator, nonce)
 *         6. Failure reason enum (outputBelowCoverage / SwapCallFailed)
 *         7. Malicious calldata guard: post-exec balanceOf(this) >= totalLockedBonds
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

    // ── Failure reason enum ──────────────────────────────────────
    enum FailureReason {
        None,                    // 0 = success (unused in events)
        OutputBelowCoverage,     // 1 = swap output < coverageFloor
        SwapCallFailed           // 2 = inner swap call reverted
    }

    // ── Plan structure ──────────────────────────────────────────
    struct Plan {
        address user;              // who the plan is for
        address operator;          // who created the plan & provides bond
        uint256 inputAmount;       // MON amount for the swap
        uint256 expectedOutput;    // simulated expected output (informational)
        uint256 guaranteedOutput;  // minimum tUSDC the user is guaranteed
        uint256 maxCompensation;   // max tUSDC the bond covers
        uint256 failureCompensation; // tUSDC paid to user if swap fails
        address target;            // target contract (MockDex)
        bytes32 calldataHash;      // keccak256(calldata) — prevents tampering
        uint256 deadline;          // block.timestamp after which plan expires
        uint256 nonce;             // per-operator nonce for planId uniqueness
        bool executed;             // prevents replay
    }

    // ── State ────────────────────────────────────────────────────
    /// @notice planId → Plan
    mapping(bytes32 => Plan) public plans;

    /// @notice operator → nonce → used
    mapping(address => mapping(uint256 => bool)) public usedNonces;

    /// @notice operator → tUSDC bond balance locked in active plans
    mapping(address => uint256) public lockedBond;

    /// @notice total tUSDC locked across all operators (for invariant check)
    uint256 public totalLockedBonds;

    /// @notice user → pending MON refund (pull pattern)
    mapping(address => uint256) public pendingRefunds;

    // ── Events ───────────────────────────────────────────────────
    event PlanOpened(
        bytes32 indexed planId,
        address indexed user,
        address indexed operator,
        uint256 guaranteedOutput,
        uint256 bondDeposited,
        uint256 coverageFloor,
        uint256 deadline
    );

    event PlanExecuted(
        bytes32 indexed planId,
        uint256 actualOutput,
        uint256 userReceived,
        uint256 shortfallPaid
    );

    event ShortfallPaid(
        bytes32 indexed planId,
        uint256 guaranteedOutput,
        uint256 actualUserReceived,
        uint256 compensationPaid
    );

    event PlanFailed(
        bytes32 indexed planId,
        FailureReason reason,
        uint256 refundedMON,
        uint256 compensationPaid
    );

    event PlanCancelled(
        bytes32 indexed planId,
        address indexed operator,
        uint256 bondReturned
    );

    event BondReleased(
        bytes32 indexed planId,
        address indexed operator,
        uint256 amount
    );

    event ServiceFeeCollected(
        bytes32 indexed planId,
        uint256 swapOutput,
        uint256 fee,
        uint256 userReceived
    );

    event ServiceFeeUpdated(uint256 oldFeeBps, uint256 newFeeBps);

    event MONRefundStored(
        bytes32 indexed planId,
        address indexed user,
        uint256 amount
    );

    // ── Errors ───────────────────────────────────────────────────
    error NotPlanUser();
    error AlreadyExecuted();
    error PlanExpired(uint256 deadline, uint256 now_);
    error CalldataMismatch(bytes32 expected, bytes32 actual);
    error InsufficientBond(uint256 required, uint256 available);
    error BondTransferFailed();
    error NonceAlreadyUsed();
    error Unauthorized();
    error FeeTooHigh();
    error InvalidSignature();
    error BelowCoverageFloor(uint256 actual, uint256 floor);
    error BondInvariantViolated(uint256 balance, uint256 locked);

    // ── Constructor ──────────────────────────────────────────────
    constructor(address _tUSDC) {
        tUSDC = MockUSDC(_tUSDC);
        DOMAIN_SEPARATOR = keccak256(abi.encode(
            EIP712_DOMAIN_TYPEHASH,
            keccak256("BondedExecutor"),
            keccak256("2"),
            block.chainid,
            address(this)
        ));
    }

    // ── Plan Lifecycle ───────────────────────────────────────────

    /**
     * @notice Operator opens a guaranteed plan and locks tUSDC bond.
     * @param planId keccak256(abi.encode(operator, nonce)) — per-operator unique
     * @param plan The plan parameters
     * @param calldata_ The raw calldata to send to target (verified against calldataHash)
     */
    function openPlan(
        bytes32 planId,
        Plan calldata plan,
        bytes calldata calldata_
    ) external {
        // Verify uniqueness (per-operator nonce domain)
        if (usedNonces[msg.sender][plan.nonce]) revert NonceAlreadyUsed();
        if (plans[planId].operator != address(0)) revert AlreadyExecuted();

        // Verify calldata
        if (keccak256(calldata_) != plan.calldataHash) {
            revert CalldataMismatch(plan.calldataHash, keccak256(calldata_));
        }

        // Bond deposit = max(maxCompensation, failureCompensation)
        uint256 bondDeposit = plan.maxCompensation;
        if (plan.failureCompensation > bondDeposit) {
            bondDeposit = plan.failureCompensation;
        }

        // Pull bond from operator
        bool ok = tUSDC.transferFrom(msg.sender, address(this), bondDeposit);
        if (!ok) revert BondTransferFailed();

        // Store plan
        plans[planId] = plan;
        plans[planId].executed = false; // explicit
        usedNonces[msg.sender][plan.nonce] = true;
        lockedBond[msg.sender] += bondDeposit;
        totalLockedBonds += bondDeposit;

        // Coverage floor: the minimum swap output before we even bother
        // If swap gives less than this, it's cheaper to refund + pay failureComp
        uint256 coverageFloor = plan.guaranteedOutput > plan.maxCompensation
            ? plan.guaranteedOutput - plan.maxCompensation
            : 0;

        emit PlanOpened(planId, plan.user, msg.sender, plan.guaranteedOutput, bondDeposit, coverageFloor, plan.deadline);
    }

    /**
     * @notice User executes a guaranteed plan by sending MON.
     */
    function executePlan(bytes32 planId, bytes calldata calldata_) external payable {
        Plan storage plan = plans[planId];

        if (msg.sender != plan.user) revert NotPlanUser();
        _executeCommon(planId, plan, calldata_);
    }

    /**
     * @notice Operator executes on user's behalf via EIP-712 signature.
     *         Operator sends MON, user signed off-chain.
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

        _executeCommon(planId, plan, calldata_);
    }

    /**
     * @notice Operator cancels an unexecuted plan and recovers bond.
     */
    function cancelPlan(bytes32 planId) external {
        Plan storage plan = plans[planId];
        if (msg.sender != plan.operator) revert Unauthorized();
        if (plan.executed) revert AlreadyExecuted();

        plan.executed = true;
        uint256 bondDeposit = plan.maxCompensation > plan.failureCompensation
            ? plan.maxCompensation
            : plan.failureCompensation;
        _releaseFullBond(planId, plan);
        emit PlanCancelled(planId, plan.operator, bondDeposit);
    }

    /**
     * @notice Anyone can cancel an expired plan — bond returns to operator.
     */
    function cancelExpiredPlan(bytes32 planId) external {
        Plan storage plan = plans[planId];
        if (plan.executed) revert AlreadyExecuted();
        if (block.timestamp <= plan.deadline) revert PlanExpired(plan.deadline, block.timestamp);
        // ^ reuse the same error to say "not yet expired"

        plan.executed = true;
        uint256 bondDeposit = plan.maxCompensation > plan.failureCompensation
            ? plan.maxCompensation
            : plan.failureCompensation;
        _releaseFullBond(planId, plan);
        emit PlanCancelled(planId, plan.operator, bondDeposit);
    }

    /**
     * @notice User withdraws pending MON refund from a failed plan.
     */
    function withdrawPendingRefund() external {
        uint256 amount = pendingRefunds[msg.sender];
        if (amount == 0) revert();
        pendingRefunds[msg.sender] = 0;
        (bool ok,) = msg.sender.call{value: amount}("");
        require(ok, "MON refund transfer failed");
    }

    // ── Admin ───────────────────────────────────────────────────

    function setServiceFee(uint256 _feeBps) external {
        if (_feeBps > 100) revert FeeTooHigh();
        uint256 old = serviceFeeBps;
        serviceFeeBps = _feeBps;
        emit ServiceFeeUpdated(old, _feeBps);
    }

    // ── Internal: Execution ──────────────────────────────────────

    /// @notice Shared guards for executePlan and executePlanWithSignature.
    function _executeCommon(bytes32 planId, Plan storage plan, bytes calldata calldata_) private {
        if (plan.executed) revert AlreadyExecuted();
        if (block.timestamp > plan.deadline) revert PlanExpired(plan.deadline, block.timestamp);
        if (keccak256(calldata_) != plan.calldataHash) {
            revert CalldataMismatch(plan.calldataHash, keccak256(calldata_));
        }
        if (msg.value != plan.inputAmount) revert();

        plan.executed = true;

        _executeSwapWithCoverage(planId, plan, calldata_);

        // ── Invariant: no malicious calldata can drain other plans' bonds ──
        if (tUSDC.balanceOf(address(this)) < totalLockedBonds) {
            revert BondInvariantViolated(tUSDC.balanceOf(address(this)), totalLockedBonds);
        }
    }

    /**
     * @notice Try/catch swap execution with coverage floor.
     *
     *         Outer: calls inner, catches revert → PlanFailed
     *         Inner: executes swap, checks output >= coverageFloor
     *
     *         coverageFloor = guaranteedOutput - maxCompensation
     *         If swap output < coverageFloor, the bond isn't enough to cover
     *         the shortfall — cheaper to refund + pay failureCompensation.
     */
    function _executeSwapWithCoverage(
        bytes32 planId,
        Plan storage plan,
        bytes calldata calldata_
    ) private {
        uint256 coverageFloor = plan.guaranteedOutput > plan.maxCompensation
            ? plan.guaranteedOutput - plan.maxCompensation
            : 0;

        uint256 balanceBefore = tUSDC.balanceOf(address(this));

        // ── Inner call ──────────────────────────────────────────
        (bool success,) =
            plan.target.call{value: plan.inputAmount}(calldata_);

        if (!success) {
            _settleSwapFailed(planId, plan, FailureReason.SwapCallFailed);
            return;
        }

        uint256 balanceAfter = tUSDC.balanceOf(address(this));
        uint256 actualOutput = balanceAfter - balanceBefore;

        // ── Check coverage floor ────────────────────────────────
        if (actualOutput < coverageFloor) {
            _settleSwapFailed(planId, plan, FailureReason.OutputBelowCoverage);
            return;
        }

        // ── Success path: settle with actual output ─────────────
        _settleSwapSuccess(planId, plan, actualOutput);
    }

    // ── Internal: Settlement ─────────────────────────────────────

    /**
     * @notice Swap succeeded, output >= coverageFloor.
     *         User receives: actualOutput - fee (or guaranteed amount via shortfall).
     */
    function _settleSwapSuccess(
        bytes32 planId,
        Plan storage plan,
        uint256 actualOutput
    ) private {
        uint256 fee = (actualOutput * serviceFeeBps) / 10000;
        uint256 userAmount = actualOutput - fee;

        if (userAmount > 0) {
            require(tUSDC.transfer(plan.user, userAmount), "output transfer failed");
        }
        if (fee > 0) {
            require(tUSDC.transfer(plan.operator, fee), "fee transfer failed");
        }

        if (userAmount >= plan.guaranteedOutput) {
            // ── Happy path: output meets or exceeds guarantee ──
            _releaseFullBond(planId, plan);
            emit ServiceFeeCollected(planId, actualOutput, fee, userAmount);
            emit PlanExecuted(planId, actualOutput, userAmount, 0);
        } else {
            // ── Shortfall: compensate user from bond ────────────
            uint256 shortfall = plan.guaranteedOutput - userAmount;
            uint256 compensation = shortfall > plan.maxCompensation
                ? plan.maxCompensation : shortfall;

            uint256 bondDeposit = plan.maxCompensation > plan.failureCompensation
                ? plan.maxCompensation
                : plan.failureCompensation;

            totalLockedBonds -= bondDeposit;
            lockedBond[plan.operator] -= bondDeposit;

            // Pay compensation to user
            require(tUSDC.transfer(plan.user, compensation), "compensation transfer failed");

            // Return remaining bond (if any) to operator
            uint256 remaining = bondDeposit - compensation;
            if (remaining > 0) {
                require(tUSDC.transfer(plan.operator, remaining), "bond release failed");
            }

            emit ShortfallPaid(planId, plan.guaranteedOutput, userAmount, compensation);
            emit BondReleased(planId, plan.operator, remaining);
            emit ServiceFeeCollected(planId, actualOutput, fee, userAmount);
            emit PlanExecuted(planId, actualOutput, userAmount, compensation);
        }
    }

    /**
     * @notice Swap failed (reverted or below coverage floor).
     *         Refund MON → pull pattern. Pay failureCompensation from bond.
     */
    function _settleSwapFailed(
        bytes32 planId,
        Plan storage plan,
        FailureReason reason
    ) private {
        // ── Refund MON via pull pattern (no revert if user can't receive) ──
        pendingRefunds[plan.user] += plan.inputAmount;
        emit MONRefundStored(planId, plan.user, plan.inputAmount);

        // ── Bond accounting: withdraw failureCompensation, release remainder ──
        uint256 bondDeposit = plan.maxCompensation > plan.failureCompensation
            ? plan.maxCompensation
            : plan.failureCompensation;

        uint256 comp = plan.failureCompensation;
        if (comp > 0) {
            totalLockedBonds -= comp;
            lockedBond[plan.operator] -= comp;
            require(tUSDC.transfer(plan.user, comp), "failure compensation failed");
        }

        // Release any remaining bond back to operator
        uint256 remaining = bondDeposit - comp;
        if (remaining > 0) {
            totalLockedBonds -= remaining;
            lockedBond[plan.operator] -= remaining;
            require(tUSDC.transfer(plan.operator, remaining), "bond release failed");
        }

        emit PlanFailed(planId, reason, plan.inputAmount, comp);
        emit BondReleased(planId, plan.operator, remaining);
    }

    /**
     * @notice Release the full bond deposit back to operator.
     */
    function _releaseFullBond(bytes32 planId, Plan storage plan) private {
        uint256 bondDeposit = plan.maxCompensation > plan.failureCompensation
            ? plan.maxCompensation
            : plan.failureCompensation;

        totalLockedBonds -= bondDeposit;
        lockedBond[plan.operator] -= bondDeposit;
        require(tUSDC.transfer(plan.operator, bondDeposit), "bond release failed");
        emit BondReleased(planId, plan.operator, bondDeposit);
    }

    // ── EIP-712 helpers ────────────────────────────────────────

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

    // ── Fallback ──────────────────────────────────────────────────
    receive() external payable {}
}
