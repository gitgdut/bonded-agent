"use client";

import { useState, useCallback } from "react";
import { useAccount, useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { planAbi, hasContract, PLAN_CONTRACT, SWAP_CALLDATA } from "@/lib/contracts";
import type { Plan } from "@/lib/types";

export interface ExecuteResult {
  status: "settled_ok" | "settled_shortfall" | "failed";
  actualOutput?: string;
  shortfallPaid?: string;
  compensation?: string;
  refunded?: boolean;
  txHash?: string;
}

/**
 * 执行担保计划(文档 §12.2)
 * 真实模式:调用合约 executePlan 并等待回执
 * 演示模式:模拟链上延迟后返回结算结果(不依赖真实交易)
 */

export function useExecutePlan(plan: Plan | undefined) {
  const { address } = useAccount();
  const { writeContractAsync } = useWriteContract();
  const [txHash, setTxHash] = useState<`0x${string}` | undefined>();

  const { isLoading: waitLoading } = useWaitForTransactionReceipt({
    hash: txHash,
    query: { enabled: !!txHash },
  });

  const execute = useCallback(async (): Promise<ExecuteResult> => {
    if (!plan) throw new Error("计划数据缺失");

    // ---- 真实模式:合约已配置 ----
    if (hasContract() && planAbi.length > 0) {
      const hash = await writeContractAsync({
        address: PLAN_CONTRACT as `0x${string}`,
        abi: planAbi,
        functionName: "executePlan",
        args: [plan.planId as `0x${string}`, SWAP_CALLDATA],
        value: BigInt(plan.inputAmount),
      });
      setTxHash(hash);
      // 结算状态由事件订阅(usePlanEvents)刷新
      return { status: "settled_ok", txHash: hash };
    }

    // ---- 演示模式:模拟执行与结算 ----
    await new Promise((r) => setTimeout(r, 1200));
    const fakeHash = `0x${"ab".repeat(32)}`;
    return {
      status: "settled_ok",
      actualOutput: plan.guaranteedOutput,
      txHash: fakeHash,
    };
  }, [plan, writeContractAsync]);

  return {
    execute,
    isConfirming: !!txHash && waitLoading,
    address,
    txHash,
  };
}