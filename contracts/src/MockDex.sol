// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {MockUSDC} from "./MockUSDC.sol";

/**
 * @title MockDex
 * @notice Simulated DEX on Monad Testnet.
 *         Swaps native MON for tUSDC at a configurable rate.
 *         Supports minOutput — reverts if output is below the floor.
 */
contract MockDex {
    MockUSDC public immutable tUSDC;

    /// @notice How many tUSDC per 1 MON (scaled by 1e18)
    ///         e.g. rate = 100e18 means 1 MON → 100 tUSDC
    uint256 public rate;

    address public immutable owner;

    event RateUpdated(uint256 oldRate, uint256 newRate);
    event Swapped(address indexed user, uint256 monIn, uint256 usdcOut);

    error Unauthorized();
    error InsufficientOutput(uint256 actual, uint256 minRequired);

    constructor(address _tUSDC, uint256 _initialRate) {
        tUSDC = MockUSDC(_tUSDC);
        rate = _initialRate;
        owner = msg.sender;
    }

    modifier onlyOwner() {
        if (msg.sender != owner) revert Unauthorized();
        _;
    }

    function setRate(uint256 _newRate) external onlyOwner {
        emit RateUpdated(rate, _newRate);
        rate = _newRate;
    }

    /// @notice Swap MON → tUSDC with minimum output protection
    /// @param minOutput Minimum tUSDC the caller will accept
    function swap(uint256 minOutput) external payable returns (uint256) {
        uint256 output = (msg.value * rate) / 1e18;

        if (output < minOutput) {
            revert InsufficientOutput(output, minOutput);
        }

        tUSDC.transfer(msg.sender, output);
        emit Swapped(msg.sender, msg.value, output);
        return output;
    }

    /// @notice Quote expected output (SimpleAMMPair-compatible interface)
    function getAmountOut(uint256 amountIn) external view returns (uint256) {
        return (amountIn * rate) / 1e18;
    }

    /// @notice Allow the contract to receive MON directly
    receive() external payable {}
}
