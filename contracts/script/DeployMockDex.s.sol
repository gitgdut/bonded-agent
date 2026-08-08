// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {Script, console} from "forge-std/Script.sol";
import {MockDex} from "../src/MockDex.sol";

contract DeployMockDex is Script {
    function run() external {
        uint256 deployerKey = vm.envUint("PRIVATE_KEY");
        // Hardcode the values to avoid env var issues
        address tUSDC = 0xD5B1b42929188280631ef2502c78AA61e1A56e0a;
        uint256 rate = 100e18;

        vm.startBroadcast(deployerKey);
        MockDex dex = new MockDex(tUSDC, rate);
        console.log("MockDex:", address(dex));
        vm.stopBroadcast();
    }
}
