// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console} from "forge-std/Script.sol";
import {ERC8004IdentityRegistry} from "../src/ERC8004IdentityRegistry.sol";
import {ERC8004ReputationRegistry} from "../src/ERC8004ReputationRegistry.sol";

/// @notice Deploys ERC-8004 compatible registries to Monad Testnet.
contract DeployERC8004 is Script {
    function run() external {
        uint256 deployerKey = vm.envUint("PRIVATE_KEY");

        vm.startBroadcast(deployerKey);

        // 1. Deploy Identity Registry
        ERC8004IdentityRegistry identity = new ERC8004IdentityRegistry();
        console.log("IdentityRegistry:", address(identity));

        // 2. Deploy Reputation Registry (linked to Identity)
        ERC8004ReputationRegistry reputation = new ERC8004ReputationRegistry(address(identity));
        console.log("ReputationRegistry:", address(reputation));

        vm.stopBroadcast();
    }
}
