// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {BondedExecutor} from "../src/BondedExecutor.sol";

/// @notice Deploys BondedExecutor V2 to Monad Testnet.
contract DeployBondedExecutor is Script {
    function run() external {
        uint256 deployerKey = vm.envUint("PRIVATE_KEY");
        address tUSDC = vm.envAddress("USDC_ADDR");

        vm.startBroadcast(deployerKey);

        BondedExecutor executor = new BondedExecutor(tUSDC);
        console.log("BondedExecutor:", address(executor));

        vm.stopBroadcast();
    }
}
