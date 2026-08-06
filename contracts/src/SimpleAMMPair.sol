// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {MockUSDC} from "./MockUSDC.sol";

/**
 * @title SimpleAMMPair
 * @notice Constant-product AMM for MON/tUSDC on Monad Testnet.
 *         Implements the same swap(uint256 minOutput) interface as MockDex,
 *         so BondedExecutor requires zero changes — only the target address
 *         and calldata stay identical (selector 0x94b918de).
 *
 *         Formula: x * y = k, 0.3% fee.
 *         output = (amountIn * 997 * reserveOut) / (reserveIn * 1000 + amountIn * 997)
 */
contract SimpleAMMPair {
    MockUSDC public immutable token;

    uint256 public reserve0; // MON (tracked from address(this).balance)
    uint256 public reserve1; // tUSDC (tracked from token.balanceOf)

    bool public initialized;

    event Swap(address indexed user, uint256 amountIn, uint256 amountOut);
    event Sync(uint256 reserve0, uint256 reserve1);

    error NotInitialized();
    error AlreadyInitialized();
    error InsufficientOutput(uint256 actual, uint256 minRequired);
    error InsufficientLiquidity();

    constructor(address _token) payable {
        token = MockUSDC(_token);
    }

    /// @notice Record initial reserves. Must be called once after deployment
    ///         AND after MON + tUSDC have been sent to this contract.
    function initialize() external {
        if (initialized) revert AlreadyInitialized();
        uint256 r0 = address(this).balance;
        uint256 r1 = token.balanceOf(address(this));
        if (r0 == 0 || r1 == 0) revert InsufficientLiquidity();
        reserve0 = r0;
        reserve1 = r1;
        initialized = true;
        emit Sync(r0, r1);
    }

    /// @notice Swap MON → tUSDC. Same interface as MockDex.swap(uint256 minOutput).
    /// @param minOutput Minimum tUSDC the caller will accept (for slippage protection)
    /// @return output Amount of tUSDC transferred to caller
    function swap(uint256 minOutput) external payable returns (uint256 output) {
        if (!initialized) revert NotInitialized();

        uint256 amountIn = msg.value;
        if (amountIn == 0) revert InsufficientLiquidity();

        // Constant-product with 0.3% fee:
        // (reserve0 + amountIn * 997/1000) * (reserve1 - output) = reserve0 * reserve1
        uint256 amountInWithFee = amountIn * 997;
        uint256 numerator = amountInWithFee * reserve1;
        uint256 denominator = reserve0 * 1000 + amountInWithFee;
        output = numerator / denominator;

        if (output < minOutput) revert InsufficientOutput(output, minOutput);
        if (output == 0) revert InsufficientLiquidity();

        reserve0 += amountIn;
        reserve1 -= output;

        token.transfer(msg.sender, output);

        emit Swap(msg.sender, amountIn, output);
        emit Sync(reserve0, reserve1);

        return output;
    }

    /// @notice Compute expected output for a given MON input (view, no state change).
    ///         Used by the Go Agent to generate quotes.
    function getAmountOut(uint256 amountIn) external view returns (uint256) {
        if (!initialized || reserve0 == 0 || reserve1 == 0) return 0;
        uint256 amountInWithFee = amountIn * 997;
        uint256 numerator = amountInWithFee * reserve1;
        uint256 denominator = reserve0 * 1000 + amountInWithFee;
        return numerator / denominator;
    }

    /// @notice Sync reserves to match the contract's actual balances.
    ///         Call after directly sending MON or tUSDC to this contract.
    function sync() external {
        reserve0 = address(this).balance;
        reserve1 = token.balanceOf(address(this));
        emit Sync(reserve0, reserve1);
    }

    receive() external payable {}
}
