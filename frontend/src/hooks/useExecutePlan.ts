"use client";

import { useState, useCallback, useEffect, useMemo } from "react";
import { useAccount, useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { decodeEventLog } from "viem";
import { planAbi, hasContract, PLAN_CONTRACT, SWAP_CALLDATA } from "@/lib/contracts";
import type { Plan } from "@/lib/types";

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

  // The reactive result: when receipt is parsed, use that; otherwise reflect executing state
  const result: ExecuteResult | null = useMemo(() => {
    if (receiptResult) return receiptResult;
    if (txHash && waitLoading) return { status: "executing", txHash };
    return null;
  }, [receiptResult, txHash, waitLoading]);

  return {
    execute,
    isConfirming: !!txHash && waitLoading,
    result,
    address,
    txHash,
  };
}
