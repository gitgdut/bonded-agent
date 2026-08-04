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

        plan.executed = true;

        // ── Snapshot tUSDC balance of this executor before swap ─
        uint256 balanceBefore = tUSDC.balanceOf(address(this));

        // ── Execute the swap ───────────────────────────────────
        (bool success,) = plan.target.call{value: plan.inputAmount}(calldata_);

        if (success) {
            // ── Swap succeeded: measure output received by executor ─
            uint256 balanceAfter = tUSDC.balanceOf(address(this));
            uint256 actualOutput = balanceAfter - balanceBefore;

            // Forward received tUSDC to user
            if (actualOutput > 0) {
                require(tUSDC.transfer(plan.user, actualOutput), "output transfer failed");
            }

            if (actualOutput >= plan.guaranteedOutput) {
                // ── Normal: bond fully released ────────────────
                _releaseBond(planId, plan);
                emit PlanExecuted(planId, actualOutput, 0);
            } else {
                // ── Shortfall: compensate the difference ───────
                uint256 shortfall = plan.guaranteedOutput - actualOutput;
                uint256 compensation = shortfall > plan.maxCompensation
                    ? plan.maxCompensation
                    : shortfall;

                lockedBond[plan.operator] -= compensation;

                // Pay user
                require(tUSDC.transfer(plan.user, compensation), "compensation transfer failed");

                // Release remaining bond (if any) to operator
                uint256 remaining = plan.maxCompensation - compensation;
                if (remaining > 0) {
                    lockedBond[plan.operator] -= remaining;
                    require(tUSDC.transfer(plan.operator, remaining), "bond release failed");
                }

                emit ShortfallPaid(planId, plan.guaranteedOutput, actualOutput, compensation);
                emit PlanExecuted(planId, actualOutput, compensation);
            }
        } else {
            // ── Swap call reverted: refund user + pay failure compensation ──
            // Refund MON to user
            (bool refundOk,) = plan.user.call{value: plan.inputAmount}("");
            if (!refundOk) revert SwapFailed(); // should not happen for EOA

            // Pay failure compensation from bond
            uint256 comp = plan.failureCompensation;
            if (comp > 0) {
                lockedBond[plan.operator] -= comp;
                require(tUSDC.transfer(plan.user, comp), "failure compensation failed");
            }

            // Release remaining bond to operator
            uint256 remainingBond = plan.maxCompensation - comp;
            if (remainingBond > 0) {
                lockedBond[plan.operator] -= remainingBond;
                require(tUSDC.transfer(plan.operator, remainingBond), "bond release failed");
            }

            emit PlanFailed(planId, plan.inputAmount, comp);
        }
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

    // ── Internal ──────────────────────────────────────────────

    function _releaseBond(bytes32 planId, Plan storage plan) internal {
        uint256 bond = plan.maxCompensation;
        lockedBond[plan.operator] -= bond;
        require(tUSDC.transfer(plan.operator, bond), "bond release failed");
        emit BondReleased(planId, plan.operator, bond);
    }

    error Unauthorized();

    // ── Fallback ──────────────────────────────────────────────
    receive() external payable {}
}
