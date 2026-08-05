"use client";

import { useQuery } from "@tanstack/react-query";
import { useWatchContractEvent } from "wagmi";
import { fetchHistoryPlans } from "@/lib/api";
import { hasContract, planAbi, PLAN_CONTRACT } from "@/lib/contracts";
import { USE_MOCK } from "@/lib/config";
import type { Plan } from "@/lib/types";

/**
 * 计划事件订阅(文档 §12.3)
 * - 演示模式:直接返回演示历史列表
 * - 真实模式:订阅 PlanOpened / PlanExecuted / ShortfallPaid / PlanFailed / BondReleased,
 *   并在订阅建立前先 getLogs 兜底(见 lib/events.ts),按当前地址过滤后重建列表
 */

/** 历史页:事件重建的担保计划列表 */
export function usePlanEvents(address?: string) {
  const query = useQuery<Plan[]>({
    queryKey: ["plan-events", address, USE_MOCK],
    queryFn: () => fetchHistoryPlans(),
    enabled: true,
    staleTime: 15_000,
  });

  // 真实模式且合约已配置时,订阅链上事件,收到事件后刷新列表
  useWatchContractEvent({
    address: hasContract() ? (PLAN_CONTRACT as `0x${string}`) : undefined,
    abi: planAbi,
    eventName: "PlanOpened",
    onLogs: () => {
      if (!USE_MOCK) query.refetch();
    },
    enabled: hasContract() && !USE_MOCK,
  });

  return {
    ...query,
    plans: query.data ?? [],
  };
}
