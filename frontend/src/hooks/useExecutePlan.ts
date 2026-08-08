"use client";

import { useState, useCallback, useEffect, useMemo } from "react";
import { useAccount, useWriteContract, useWaitForTransactionReceipt, useSignTypedData } from "wagmi";
import { decodeEventLog } from "viem";
import { planAbi, hasContract, PLAN_CONTRACT, SWAP_CALLDATA } from "@/lib/contracts";
import { API_BASE_URL } from "@/lib/config";
import type { Plan } from "@/lib/types";

// ── EIP-712 domain and types (must match BondedExecutor.sol) ──

function eip712Domain(verifyingContract: string, chainId: number) {
  return {
    name: "BondedExecutor",
    version: "2",
    chainId,
    verifyingContract: verifyingContract as `0x${string}`,
  } as const;
}

const eip712Types = {
  ExecuteAuthorization: [
    { name: "planId", type: "bytes32" },
    { name: "inputAmount", type: "uint256" },
    { name: "deadline", type: "uint256" },
  ],
} as const;

export interface ExecuteResult {
  status: "executing" | "settled_ok" | "settled_shortfall" | "failed";
  actualOutput?: string;
  shortfallPaid?: string;
  compensation?: string;
  refunded?: boolean;
  txHash?: string;
}

/**
 * Parse transaction receipt logs to determine the actual swap outcome.
 * Returns null if no recognizable BondedExecutor events are found.
 */
function parseReceiptLogs(
  logs: { data: `0x${string}`; topics: `0x${string}`[] }[]
): { status: "settled_ok" | "settled_shortfall" | "failed"; actualOutput?: string; compensation?: string; shortfallPaid?: string } | null {
  let isFailed = false;
  let actualOutput: string | undefined;
  let paidToUser: string | undefined;
  let shortfallPaid: string | undefined;
  let compensationPaid: string | undefined;

  for (const log of logs) {
    try {
      const decoded = decodeEventLog({
        abi: planAbi,
        data: log.data,
        topics: log.topics as any,
      });

      if (decoded.eventName === "PlanFailed") {
        isFailed = true;
        compensationPaid = (decoded.args as any).compensationPaid?.toString();
      } else if (decoded.eventName === "PlanExecuted") {
        actualOutput = (decoded.args as any).actualOutput?.toString();
        paidToUser = (decoded.args as any).paidToUser?.toString();
      } else if (decoded.eventName === "ShortfallPaid") {
        shortfallPaid = (decoded.args as any).shortfall?.toString();
      }
    } catch {
      // Not a BondedExecutor event, skip
    }
  }

  if (isFailed) {
    return {
      status: "failed",
      compensation: compensationPaid,
    };
  }

  if (actualOutput !== undefined) {
    const hasShortfall = paidToUser && BigInt(paidToUser) > 0n;
    return {
      status: hasShortfall ? "settled_shortfall" : "settled_ok",
      actualOutput,
      compensation: hasShortfall ? paidToUser : undefined,
      shortfallPaid: hasShortfall ? (shortfallPaid ?? paidToUser) : undefined,
    };
  }

  return null;
}

/**
 * Execute a guaranteed plan (doc §12.2).
 *
 * Real mode: submits executePlan tx, then waits for the receipt and
 * parses on-chain events to determine the actual outcome.
 *
 * Demo mode: simulates chain latency and returns a mock settled_ok result.
 */
export function useExecutePlan(plan: Plan | undefined) {
  const { address } = useAccount();
  const { writeContractAsync } = useWriteContract();
  const { signTypedDataAsync } = useSignTypedData();
  const [txHash, setTxHash] = useState<`0x${string}` | undefined>();
  const [receiptResult, setReceiptResult] = useState<ExecuteResult | null>(null);

  // Wait for the transaction receipt
  const { data: receipt, isLoading: waitLoading } = useWaitForTransactionReceipt({
    hash: txHash,
    query: { enabled: !!txHash },
  });

  // When receipt arrives, parse its logs to determine the real outcome
  useEffect(() => {
    if (!receipt || !plan || receiptResult) return;

    const parsed = parseReceiptLogs(receipt.logs);

    if (parsed) {
      setReceiptResult({
        ...parsed,
        txHash: txHash,
      });
    } else if (receipt.status === "reverted") {
      // Transaction reverted without emitting our events (e.g. NotPlanUser, out of gas)
      setReceiptResult({
        status: "failed",
        txHash: txHash,
      });
    } else {
      // Receipt succeeded but no recognizable events — fallback to settled_ok
      // (should not normally happen with BondedExecutor)
      setReceiptResult({
        status: "settled_ok",
        txHash: txHash,
      });
    }
  }, [receipt, plan, receiptResult, txHash]);

  const execute = useCallback(async (): Promise<ExecuteResult> => {
    if (!plan) throw new Error("计划数据缺失");

    // ---- Real mode: contract configured ----
    if (hasContract() && planAbi.length > 0) {
      const hash = await writeContractAsync({
        address: PLAN_CONTRACT as `0x${string}`,
        abi: planAbi,
        functionName: "executePlan",
        args: [plan.planId as `0x${string}`, SWAP_CALLDATA],
        value: BigInt(plan.inputAmount),
      });
      setTxHash(hash);
      setReceiptResult(null); // reset for new attempt
      return { status: "executing", txHash: hash };
    }

    // ---- Demo mode: simulate execution & settlement ----
    await new Promise((r) => setTimeout(r, 1200));
    const fakeHash = `0x${"ab".repeat(32)}`;
    // In demo mode, randomly simulate different outcomes for testing
    const fakeResult: ExecuteResult = {
      status: "settled_ok",
      actualOutput: plan.guaranteedOutput,
      txHash: fakeHash,
    };
    setReceiptResult(fakeResult);
    return fakeResult;
  }, [plan, writeContractAsync]);

  const executeWithSignature = useCallback(async (): Promise<ExecuteResult> => {
    if (!plan || !address) throw new Error("计划数据缺失或未连接钱包");

    if (hasContract() && planAbi.length > 0) {
      // 1. Build EIP-712 typed data
      const deadline = Math.floor(Date.now() / 1000) + 3600; // 1 hour
      const domain = eip712Domain(PLAN_CONTRACT, 10143);
      const message = {
        planId: plan.planId as `0x${string}`,
        inputAmount: BigInt(plan.inputAmount),
        deadline: BigInt(deadline),
      };

      // 2. User signs off-chain (free, no gas)
      const signature = await signTypedDataAsync({
        domain,
        types: eip712Types,
        primaryType: "ExecuteAuthorization",
        message,
      });

      // 3. POST signature to operator API → operator submits tx
      const apiRes = await fetch(`${API_BASE_URL}/plans/${plan.planId}/execute-signed`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          signature,
          deadline,
        }),
      });

      if (!apiRes.ok) {
        const err = await apiRes.json().catch(() => ({}));
        throw new Error((err as any).message || "Operator 提交失败");
      }

      const { txHash: hash } = await apiRes.json();
      setTxHash(hash as `0x${string}`);
      setReceiptResult(null);
      return { status: "executing", txHash: hash as string };
    }

    // Demo mode fallback
    await new Promise((r) => setTimeout(r, 1200));
    const fakeHash = `0x${"ab".repeat(32)}`;
    const fakeResult: ExecuteResult = {
      status: "settled_ok",
      actualOutput: plan.guaranteedOutput,
      txHash: fakeHash,
    };
    setReceiptResult(fakeResult);
    return fakeResult;
  }, [plan, address, signTypedDataAsync]);

  // The reactive result: when receipt is parsed, use that; otherwise reflect executing state
  const result: ExecuteResult | null = useMemo(() => {
    if (receiptResult) return receiptResult;
    if (txHash && waitLoading) return { status: "executing", txHash };
    return null;
  }, [receiptResult, txHash, waitLoading]);

  return {
    execute,
    executeWithSignature,
    isConfirming: !!txHash && waitLoading,
    result,
    address,
    txHash,
  };
}
