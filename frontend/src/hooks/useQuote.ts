"use client";

import { useQuery, useMutation } from "@tanstack/react-query";
import { useRouter } from "next/navigation";
import { parseEther } from "viem";
import {
  fetchQuote,
  createPlan,
  fetchPlan,
  DEFAULT_DEADLINE_MINUTES,
} from "@/lib/api";
import type { CreatePlanRequest, Plan, Quote } from "@/lib/types";
import { derivePlanStatus } from "@/lib/plan-status";

/** 报价查询:传入用户输入的 MON 数量(字符串,如 "1"),内部转为 wei 再请求 */
export function useQuote(inputAmount?: string, enabled = false) {
  return useQuery<Quote>({
    queryKey: ["quote", inputAmount],
    queryFn: () => fetchQuote(parseEther(inputAmount as `${number}`).toString()),
    enabled: enabled && !!inputAmount,
    staleTime: 30_000,
    retry: 1,
  });
}

/** 创建担保计划并跳转到计划页 */
export function useCreatePlan() {
  const router = useRouter();
  return useMutation({
    mutationFn: (req: CreatePlanRequest) => createPlan(req),
    onSuccess: (res) => {
      router.push(`/plans/${encodeURIComponent(res.planId)}`);
    },
  });
}

/** 计划详情:读取后端/链上数据并派生状态 */
export function usePlan(planId?: string) {
  const query = useQuery<Plan>({
    queryKey: ["plan", planId],
    queryFn: () => fetchPlan(planId as string),
    enabled: !!planId,
    staleTime: 15_000,
    retry: 1,
  });

  return {
    ...query,
    plan: query.data,
    status: query.data ? derivePlanStatus(query.data) : ("loading" as const),
  };
}

/** 开计划默认时长 */
export { DEFAULT_DEADLINE_MINUTES };