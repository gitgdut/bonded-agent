/**
 * 担保合约接入(文档 §7.2)
 * 合约地址和 ABI 由部署方提供
 */

import { PLAN_CONTRACT_ADDRESS } from "./config";

export function hasContract(): boolean {
  return PLAN_CONTRACT_ADDRESS.length > 0;
}

export const PLAN_CONTRACT = PLAN_CONTRACT_ADDRESS as `0x${string}` | "";

/** executePlan 的 calldata — swap(minOutput)，minOutput = coverageFloor，随每个计划参数变化。这里只是 coverageFloor=0 的特例常量，执行时需用计划对应的实际 calldata。 */
export const SWAP_CALLDATA = "0x94b918de0000000000000000000000000000000000000000000000000000000000000000";

/** BondedExecutor ABI */
export const planAbi = [
  {
    "type": "function",
    "name": "executePlan",
    "inputs": [
      { "name": "planId", "type": "bytes32", "internalType": "bytes32" },
      { "name": "calldata_", "type": "bytes", "internalType": "bytes" }
    ],
    "outputs": [],
    "stateMutability": "payable"
  },
  {
    "type": "function",
    "name": "plans",
    "inputs": [{ "name": "", "type": "bytes32", "internalType": "bytes32" }],
    "outputs": [
      { "name": "user", "type": "address" },
      { "name": "operator", "type": "address" },
      { "name": "inputAmount", "type": "uint256" },
      { "name": "expectedOutput", "type": "uint256" },
      { "name": "guaranteedOutput", "type": "uint256" },
      { "name": "maxCompensation", "type": "uint256" },
      { "name": "failureCompensation", "type": "uint256" },
      { "name": "target", "type": "address" },
      { "name": "calldataHash", "type": "bytes32" },
      { "name": "deadline", "type": "uint256" },
      { "name": "nonce", "type": "uint256" },
      { "name": "executed", "type": "bool" }
    ],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "openPlan",
    "inputs": [
      { "name": "planId", "type": "bytes32" },
      {
        "name": "plan",
        "type": "tuple",
        "components": [
          { "name": "user", "type": "address" },
          { "name": "operator", "type": "address" },
          { "name": "inputAmount", "type": "uint256" },
          { "name": "expectedOutput", "type": "uint256" },
          { "name": "guaranteedOutput", "type": "uint256" },
          { "name": "maxCompensation", "type": "uint256" },
          { "name": "failureCompensation", "type": "uint256" },
          { "name": "target", "type": "address" },
          { "name": "calldataHash", "type": "bytes32" },
          { "name": "deadline", "type": "uint256" },
          { "name": "nonce", "type": "uint256" },
          { "name": "executed", "type": "bool" }
        ]
      },
      { "name": "calldata_", "type": "bytes" }
    ],
    "outputs": [],
    "stateMutability": "nonpayable"
  },
  // ── V2 Events (BondedExecutor post-audit) ──────────────────
  {
    "type": "event",
    "name": "PlanOpened",
    "inputs": [
      { "name": "planId", "type": "bytes32", "indexed": true },
      { "name": "user", "type": "address", "indexed": true },
      { "name": "operator", "type": "address", "indexed": true },
      { "name": "guaranteedOutput", "type": "uint256", "indexed": false },
      { "name": "bondDeposited", "type": "uint256", "indexed": false },
      { "name": "coverageFloor", "type": "uint256", "indexed": false },
      { "name": "deadline", "type": "uint256", "indexed": false }
    ]
  },
  {
    "type": "event",
    "name": "PlanExecuted",
    "inputs": [
      { "name": "planId", "type": "bytes32", "indexed": true },
      { "name": "actualOutput", "type": "uint256", "indexed": false },
      { "name": "userReceived", "type": "uint256", "indexed": false },
      { "name": "shortfallPaid", "type": "uint256", "indexed": false }
    ]
  },
  {
    "type": "event",
    "name": "ShortfallPaid",
    "inputs": [
      { "name": "planId", "type": "bytes32", "indexed": true },
      { "name": "guaranteedOutput", "type": "uint256", "indexed": false },
      { "name": "actualUserReceived", "type": "uint256", "indexed": false },
      { "name": "compensationPaid", "type": "uint256", "indexed": false }
    ]
  },
  {
    "type": "event",
    "name": "PlanFailed",
    "inputs": [
      { "name": "planId", "type": "bytes32", "indexed": true },
      { "name": "reason", "type": "uint8", "indexed": false },
      { "name": "refundedMON", "type": "uint256", "indexed": false },
      { "name": "compensationPaid", "type": "uint256", "indexed": false }
    ]
  },
  {
    "type": "event",
    "name": "BondReleased",
    "inputs": [
      { "name": "planId", "type": "bytes32", "indexed": true },
      { "name": "operator", "type": "address", "indexed": true },
      { "name": "amount", "type": "uint256", "indexed": false }
    ]
  },
  {
    "type": "event",
    "name": "PlanCancelled",
    "inputs": [
      { "name": "planId", "type": "bytes32", "indexed": true },
      { "name": "operator", "type": "address", "indexed": true },
      { "name": "bondReturned", "type": "uint256", "indexed": false }
    ]
  },
  {
    "type": "error",
    "name": "AlreadyExecuted",
    "inputs": []
  },
  {
    "type": "error",
    "name": "NotPlanUser",
    "inputs": []
  },
  {
    "type": "error",
    "name": "InvalidSignature",
    "inputs": []
  }
] as const;

/** 合约未配置时的用户提示 */
export const CONTRACT_NOT_CONFIGURED = "担保合约尚未配置，请等待部署方提供地址";
