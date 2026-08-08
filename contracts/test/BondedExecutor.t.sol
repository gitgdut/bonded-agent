// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Test, console2} from "forge-std/Test.sol";
import {MockUSDC} from "../src/MockUSDC.sol";
import {MockDex} from "../src/MockDex.sol";
import {BondedExecutor} from "../src/BondedExecutor.sol";

contract BondedExecutorTest is Test {
    MockUSDC public tUSDC;
    MockDex public dex;
    BondedExecutor public executor;

    address public user = makeAddr("user");
    address public operator = makeAddr("operator");
    address public stranger = makeAddr("stranger");
    uint256 constant STRANGER_FUND = 5 ether;

    uint256 constant INITIAL_RATE = 100e18; // 1 MON = 100 tUSDC
    uint256 constant SWAP_AMOUNT = 1 ether;

    // ── Setup ────────────────────────────────────────────────

    function setUp() public {
        vm.prank(operator);
        tUSDC = new MockUSDC();

        vm.prank(operator);
        dex = new MockDex(address(tUSDC), INITIAL_RATE);

        // Fund MockDex with tUSDC for swap liquidity
        vm.prank(operator);
        tUSDC.mint(address(dex), 1_000_000e18);

        vm.prank(operator);
        executor = new BondedExecutor(address(tUSDC));

        // Fund user with MON
        vm.deal(user, 10 ether);
        vm.deal(stranger, STRANGER_FUND);

        // Mint tUSDC to operator for bond
        vm.prank(operator);
        tUSDC.mint(operator, 10000e18);

        // Operator approves executor to pull tUSDC
        vm.prank(operator);
        tUSDC.approve(address(executor), type(uint256).max);
    }

    // ── Helpers (V2: per-operator nonce planId) ─────────────

    function _createPlan(
        uint256 guaranteedOutput,
        uint256 maxCompensation,
        uint256 failureCompensation,
        uint256 deadline,
        uint256 nonce,
        uint256 minOutput
    ) internal view returns (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) {
        calldata_ = abi.encodeWithSelector(MockDex.swap.selector, minOutput);

        plan = BondedExecutor.Plan({
            user: user,
            operator: operator,
            inputAmount: SWAP_AMOUNT,
            expectedOutput: 100e18,
            guaranteedOutput: guaranteedOutput,
            maxCompensation: maxCompensation,
            failureCompensation: failureCompensation,
            target: address(dex),
            calldataHash: keccak256(calldata_),
            deadline: deadline,
            nonce: nonce,
            executed: false
        });

        // V2: planId = keccak256(abi.encode(operator, nonce))
        planId = keccak256(abi.encode(operator, nonce));
    }

    function _openPlan(bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) internal {
        vm.prank(operator);
        executor.openPlan(planId, plan, calldata_);
    }

    // ── Test 1: Normal — output ≥ guaranteed, bond released ──

    function test_NormalFulfillment() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 1, 0);

        uint256 operatorBalanceBefore = tUSDC.balanceOf(operator);

        _openPlan(planId, plan, calldata_);

        // Operator bond should be locked
        assertEq(tUSDC.balanceOf(operator), operatorBalanceBefore - 20e18);

        // User executes
        vm.prank(user);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // User should have received ~100 tUSDC from swap
        assertGe(tUSDC.balanceOf(user), 90e18);

        // Bond should be fully released back to operator
        assertEq(tUSDC.balanceOf(operator), operatorBalanceBefore);
    }

    // ── Test 2: Shortfall — exact compensation ───────────────

    function test_ShortfallExactCompensation() public {
        // Plan guarantees 95, rate changed to 80 (shortfall = 15), coverageFloor=75
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(95e18, 20e18, 5e18, block.timestamp + 1 days, 2, 75e18);

        _openPlan(planId, plan, calldata_);

        // Manipulate rate to cause shortfall (1 MON → 80 tUSDC)
        vm.prank(operator);
        dex.setRate(80e18);

        uint256 userBalanceBefore = tUSDC.balanceOf(user);

        vm.prank(user);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // User got 80 from swap + 15 from compensation = 95 total
        assertEq(tUSDC.balanceOf(user), userBalanceBefore + 80e18 + 15e18);
    }

    // ── Test 3: Shortfall capped at max compensation ─────────

    function test_ShortfallMaxCompensation() public {
        // guaranteed=95, maxComp=10 → coverageFloor=85
        // Set rate to 85e18 so output = 85 exactly (at coverage floor boundary)
        // Shortfall = 95-85 = 10 = maxComp, compensation capped at 10
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(95e18, 10e18, 5e18, block.timestamp + 1 days, 3, 85e18);

        _openPlan(planId, plan, calldata_);

        // Rate = 85 → output = 85 exactly at coverage floor
        vm.prank(operator);
        dex.setRate(85e18);

        uint256 userBalanceBefore = tUSDC.balanceOf(user);

        vm.prank(user);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // User got 85 from swap + 10 (capped at maxComp) = 95
        assertEq(tUSDC.balanceOf(user), userBalanceBefore + 85e18 + 10e18);
    }

    // ── Test 4: Output below coverage floor → atomic revert ──

    function test_OutputBelowCoverage_Reverts() public {
        // Use LenientDex that ignores minOutput
        LenientDex lenient = new LenientDex(address(tUSDC));

        // guaranteed=95, maxComp=10 → coverageFloor=85
        // Rate=60 → output=60 < 85 → BelowCoverageFloor
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(95e18, 10e18, 5e18, block.timestamp + 1 days, 40, 85e18);

        // Point plan at lenient DEX (no minOutput enforcement)
        plan.target = address(lenient);
        // Rebuild calldata for LenientDex (same swap selector + minOutput arg)
        calldata_ = abi.encodeWithSelector(LenientDex.swap.selector, uint256(85e18));
        plan.calldataHash = keccak256(calldata_);

        _openPlan(planId, plan, calldata_);

        // Fund lenient DEX
        vm.prank(operator);
        tUSDC.mint(address(lenient), 1_000_000e18);

        // Set rate → output=60 < 85 coverageFloor
        vm.prank(operator);
        lenient.setRate(60e18);

        uint256 userMONBefore = user.balance;

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(BondedExecutor.BelowCoverageFloor.selector, 60e18, 85e18)
        );
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // ── Atomic revert: MON stayed with user, no state changed ──
        assertEq(user.balance, userMONBefore, "MON should not be spent");
        // Plan is NOT marked executed
        (, , , , , , , , , , , bool executed) = executor.plans(planId);
        assertFalse(executed);
    }

    // ── Test 5: Swap fails (DEX reverts) → pull refund + failComp ──

    function test_SwapFailure_CallRevert() public {
        // Deploy a target contract that always reverts on call
        RevertingTarget revertingTarget = new RevertingTarget();

        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 4, 0);

        // Point to the reverting target
        plan.target = address(revertingTarget);
        calldata_ = hex"deadbeef";
        plan.calldataHash = keccak256(calldata_);

        _openPlan(planId, plan, calldata_);

        uint256 userBalanceBefore = tUSDC.balanceOf(user);

        vm.prank(user);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // ── Pull-mode MON refund ──
        assertEq(executor.pendingRefunds(user), SWAP_AMOUNT, "MON refund pending");

        uint256 userMONBefore = user.balance;
        vm.prank(user);
        executor.withdrawPendingRefund();
        assertEq(user.balance, userMONBefore + SWAP_AMOUNT, "MON returned via pull");

        // ── Failure compensation paid ──
        assertEq(tUSDC.balanceOf(user), userBalanceBefore + 5e18, "failComp paid");

        // ── Plan marked executed ──
        (, , , , , , , , , , , bool executed) = executor.plans(planId);
        assertTrue(executed);
    }

    // ── Test 6: Calldata tampered → reject ───────────────────

    function test_CalldataTampered() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 5, 0);

        _openPlan(planId, plan, calldata_);

        // Tampered calldata (different minOutput)
        bytes memory badCalldata = abi.encodeWithSelector(MockDex.swap.selector, uint256(1000e18));

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(BondedExecutor.CalldataMismatch.selector, plan.calldataHash, keccak256(badCalldata))
        );
        executor.executePlan{value: SWAP_AMOUNT}(planId, badCalldata);
    }

    // ── Test 7: Plan expired → reject ────────────────────────

    function test_PlanExpired() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 hours, 6, 0);

        _openPlan(planId, plan, calldata_);

        // Warp past deadline
        vm.warp(block.timestamp + 2 hours);

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(BondedExecutor.PlanExpired.selector, plan.deadline, block.timestamp)
        );
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);
    }

    // ── Test 8: Replay → reject ──────────────────────────────

    function test_ReplayRejected() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 7, 0);

        _openPlan(planId, plan, calldata_);

        // First execution succeeds
        vm.prank(user);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // Second execution should fail
        vm.prank(user);
        vm.expectRevert(BondedExecutor.AlreadyExecuted.selector);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);
    }

    // ── Test 9: Nonce reused → reject ────────────────────────

    function test_NonceReuseRejected() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 8, 0);

        _openPlan(planId, plan, calldata_);

        // Try opening another plan with same nonce
        (bytes32 planId2, BondedExecutor.Plan memory plan2, bytes memory calldata2) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 8, 0); // same nonce

        vm.prank(operator);
        vm.expectRevert(BondedExecutor.NonceAlreadyUsed.selector);
        executor.openPlan(planId2, plan2, calldata2);
    }

    // ── Test 10: Wrong user → reject ─────────────────────────

    function test_WrongUserRejected() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 9, 0);

        _openPlan(planId, plan, calldata_);

        vm.prank(stranger);
        vm.expectRevert(BondedExecutor.NotPlanUser.selector);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);
    }

    // ── Test 11: insufficient bond → reject ──────────────────

    function test_InsufficientBondRejected() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 10000e18, 5e18, block.timestamp + 1 days, 10, 0);

        // Test with a poor operator
        address poorOperator = makeAddr("poorOperator");
        vm.prank(operator);
        tUSDC.mint(poorOperator, 1e18);

        vm.prank(poorOperator);
        tUSDC.approve(address(executor), type(uint256).max);

        BondedExecutor.Plan memory planPoor = plan;
        planPoor.operator = poorOperator;
        planPoor.maxCompensation = 100e18; // needs 100e18, only has 1e18
        planPoor.nonce = 99;
        bytes32 pid = keccak256(abi.encode(poorOperator, planPoor.nonce));

        vm.prank(poorOperator);
        vm.expectRevert(); // ERC20 transferFrom will revert (insufficient balance)
        executor.openPlan(pid, planPoor, calldata_);
    }

    // ── Test 12: cancelExpiredPlan only after expiry ─────────

    function test_CancelExpiredPlan() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 hours, 11, 0);

        _openPlan(planId, plan, calldata_);

        // Before expiry — must fail
        vm.prank(stranger);
        vm.expectRevert();
        executor.cancelExpiredPlan(planId);

        // After expiry — succeed
        vm.warp(block.timestamp + 2 hours);
        executor.cancelExpiredPlan(planId);

        (, , , , , , , , , , , bool executed) = executor.plans(planId);
        assertTrue(executed);
    }
}

/// @notice DEX that ignores minOutput — used to test OutputBelowCoverage path.
contract LenientDex {
    MockUSDC public tUSDC;
    uint256 public rate;

    event Swapped(address indexed user, uint256 monIn, uint256 usdcOut);

    constructor(address _tUSDC) {
        tUSDC = MockUSDC(_tUSDC);
        rate = 100e18;
    }

    function setRate(uint256 _rate) external { rate = _rate; }

    /// @notice Always succeeds, regardless of minOutput
    function swap(uint256 /*minOutput*/) external payable returns (uint256) {
        uint256 output = (msg.value * rate) / 1e18;
        tUSDC.transfer(msg.sender, output);
        emit Swapped(msg.sender, msg.value, output);
        return output;
    }

    receive() external payable {}
}

/// @notice Helper contract that always reverts — used to test swap failure path
contract RevertingTarget {
    fallback() external payable {
        revert("always revert");
    }
}
