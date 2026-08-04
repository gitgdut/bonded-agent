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
        // Deploy
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

    // ── Helpers ──────────────────────────────────────────────

    function _createPlan(
        uint256 guaranteedOutput,
        uint256 maxCompensation,
        uint256 failureCompensation,
        uint256 deadline,
        uint256 nonce
    ) internal view returns (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) {
        calldata_ = abi.encodeWithSelector(MockDex.swap.selector, uint256(0)); // minOutput=0

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

        planId = keccak256(abi.encode(user, operator, nonce));
    }

    function _openPlan(bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) internal {
        vm.prank(operator);
        executor.openPlan(planId, plan, calldata_);
    }

    // ── Test 1: Normal — output ≥ guaranteed, bond released ──

    function test_NormalFulfillment() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 1);

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
        // Plan guarantees 95, but rate changed to 80 (shortfall = 15)
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(95e18, 20e18, 5e18, block.timestamp + 1 days, 2);

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

    // ── Test 3: Shortfall equals max compensation ────────────

    function test_ShortfallMaxCompensation() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(95e18, 10e18, 5e18, block.timestamp + 1 days, 3);

        _openPlan(planId, plan, calldata_);

        // Rate drops to 50 → shortfall = 45, but max comp = 10
        vm.prank(operator);
        dex.setRate(50e18);

        uint256 userBalanceBefore = tUSDC.balanceOf(user);

        vm.prank(user);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // User got 50 from swap + 10 from compensation (capped) = 60
        assertEq(tUSDC.balanceOf(user), userBalanceBefore + 50e18 + 10e18);
    }

    // ── Test 4: Swap fails → refund + failure compensation ──

    function test_SwapFailure_CallRevert() public {
        // Deploy a target contract that always reverts on call
        RevertingTarget revertingTarget = new RevertingTarget();

        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 4);

        // Point to the reverting target
        plan.target = address(revertingTarget);
        // Calldata hash no longer matters since the call will revert before calldata matters
        // but we need calldataHash to match — use raw bytes
        calldata_ = hex"deadbeef";
        plan.calldataHash = keccak256(calldata_);

        _openPlan(planId, plan, calldata_);

        uint256 userBalanceBefore = tUSDC.balanceOf(user);
        uint256 userMONBefore = user.balance;
        uint256 operatorBalanceBefore = tUSDC.balanceOf(operator);

        vm.prank(user);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // MON refunded
        assertEq(user.balance, userMONBefore);
        // Failure compensation paid
        assertEq(tUSDC.balanceOf(user), userBalanceBefore + 5e18);
        // Remaining bond returned to operator (20 - 5 = 15)
        assertEq(tUSDC.balanceOf(operator), operatorBalanceBefore + 15e18);
    }

    // ── Test 5: Calldata tampered → reject ───────────────────

    function test_CalldataTampered() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 5);

        _openPlan(planId, plan, calldata_);

        // Tampered calldata (different minOutput)
        bytes memory badCalldata = abi.encodeWithSelector(MockDex.swap.selector, uint256(1000e18));

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(BondedExecutor.CalldataMismatch.selector, plan.calldataHash, keccak256(badCalldata))
        );
        executor.executePlan{value: SWAP_AMOUNT}(planId, badCalldata);
    }

    // ── Test 6: Plan expired → reject ────────────────────────

    function test_PlanExpired() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 hours, 6);

        _openPlan(planId, plan, calldata_);

        // Warp past deadline
        vm.warp(block.timestamp + 2 hours);

        vm.prank(user);
        vm.expectRevert(
            abi.encodeWithSelector(BondedExecutor.PlanExpired.selector, plan.deadline, block.timestamp)
        );
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);
    }

    // ── Test 7: Replay → reject ──────────────────────────────

    function test_ReplayRejected() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 7);

        _openPlan(planId, plan, calldata_);

        // First execution succeeds
        vm.prank(user);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);

        // Second execution should fail
        vm.prank(user);
        vm.expectRevert(BondedExecutor.AlreadyExecuted.selector);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);
    }

    // ── Test 8: Nonce reused → reject ────────────────────────

    function test_NonceReuseRejected() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 8);

        _openPlan(planId, plan, calldata_);

        // Try opening another plan with same nonce
        (bytes32 planId2, BondedExecutor.Plan memory plan2, bytes memory calldata2) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 8); // same nonce

        vm.prank(operator);
        vm.expectRevert(BondedExecutor.NonceAlreadyUsed.selector);
        executor.openPlan(planId2, plan2, calldata2);
    }

    // ── Test 9: Wrong user → reject ──────────────────────────

    function test_WrongUserRejected() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 20e18, 5e18, block.timestamp + 1 days, 9);

        _openPlan(planId, plan, calldata_);

        vm.prank(stranger);
        vm.expectRevert(BondedExecutor.NotPlanUser.selector);
        executor.executePlan{value: SWAP_AMOUNT}(planId, calldata_);
    }

    // ── Test 10: Insufficient bond → reject ──────────────────

    function test_InsufficientBondRejected() public {
        (bytes32 planId, BondedExecutor.Plan memory plan, bytes memory calldata_) =
            _createPlan(90e18, 10000e18, 5e18, block.timestamp + 1 days, 10);

        // Operator only has 10000e18, needs 10000e18 bond
        // This should fail because transferFrom of 10000e18 (operator has exactly 10000e18 after approve, but it should work)

        // Let's test with a poor operator
        address poorOperator = makeAddr("poorOperator");
        vm.prank(operator);
        tUSDC.mint(poorOperator, 1e18);

        vm.prank(poorOperator);
        tUSDC.approve(address(executor), type(uint256).max);

        BondedExecutor.Plan memory planPoor = plan;
        planPoor.operator = poorOperator;
        planPoor.maxCompensation = 100e18; // needs 100e18, only has 1e18
        planPoor.nonce = 99;
        bytes32 pid = keccak256(abi.encode(planPoor.user, poorOperator, planPoor.nonce));

        vm.prank(poorOperator);
        vm.expectRevert(); // ERC20 transferFrom will revert (insufficient balance)
        executor.openPlan(pid, planPoor, calldata_);
    }
}

/// @notice Helper contract that always reverts — used to test swap failure path
contract RevertingTarget {
    fallback() external payable {
        revert("always revert");
    }
}
