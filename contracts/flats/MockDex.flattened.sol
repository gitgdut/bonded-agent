// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

// src/MockUSDC.sol

/**
 * @title MockUSDC
 * @notice Test USDC token for Monad Testnet.
 *         Used as both the output asset and the bond collateral.
 *         Only the deployer can mint.
 */
contract MockUSDC {
    string public name = "Mock USDC";
    string public symbol = "tUSDC";
    uint8 public decimals = 18;

    uint256 public totalSupply;
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    address public immutable owner;

    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Approval(address indexed owner, address indexed spender, uint256 amount);

    error Unauthorized();

    constructor() {
        owner = msg.sender;
    }

    modifier onlyOwner() {
        if (msg.sender != owner) revert Unauthorized();
        _;
    }

    /// @notice Mint tUSDC — only deployer
    function mint(address to, uint256 amount) external onlyOwner {
        totalSupply += amount;
        balanceOf[to] += amount;
        emit Transfer(address(0), to, amount);
    }

    function transfer(address to, uint256 amount) external returns (bool) {
        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    function approve(address spender, uint256 amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(address from, address to, uint256 amount) external returns (bool) {
        allowance[from][msg.sender] -= amount;
        balanceOf[from] -= amount;
        balanceOf[to] += amount;
        emit Transfer(from, to, amount);
        return true;
    }
}

// src/MockDex.sol

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

    /// @notice Allow the contract to receive MON directly
    receive() external payable {}
}

