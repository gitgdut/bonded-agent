"use client";

import { useParams } from "next/navigation";
import Link from "next/link";
import { useState } from "react";
import { useAccount } from "wagmi";
import { usePlan } from "@/hooks/useQuote";
import { useExecutePlan, type ExecuteResult } from "@/hooks/useExecutePlan";
import { PromiseCard } from "@/components/plan/PromiseCard";
import { SettlementExplanation } from "@/components/plan/SettlementExplanation";
import { ErrorBanner } from "@/components/ui/ErrorBanner";
import { Skeleton } from "@/components/ui/Skeleton";
import { EmptyState } from "@/components/ui/EmptyState";
import { Button } from "@/components/ui/Button";
import { PLAN } from "@/lib/copy";
import { USE_MOCK } from "@/lib/config";
import type { Plan } from "@/lib/types";

/** 计划页:承诺卡片 + 执行 + 结算(文档 §6.2) */
export default function PlanPage() {
  const params = useParams<{ id: string }>();
  const planId = params?.id;

  const { plan, status, isLoading, error } = usePlan(planId);
  const { address } = useAccount();
  const { execute, isConfirming } = useExecutePlan(plan);

  const [execResult, setExecResult] = useState<ExecuteResult>();
  const [execError, setExecError] = useState<unknown>();

  // 钱包校验:演示模式跳过;真实模式校验计划归属
  const wrongWallet =
    !USE_MOCK &&
    !!plan &&
    !!address &&
    !!plan.user &&
    plan.user.toLowerCase() !== address.toLowerCase();

  // 执行成功后合并结算结果,驱动结算区渲染
  const displayPlan: Plan | undefined = plan
    ? {
        ...plan,
        status: (execResult?.status ?? status) as Plan["status"],
        actualOutput: execResult?.actualOutput ?? plan.actualOutput,
        shortfallPaid: execResult?.shortfallPaid ?? plan.shortfallPaid,
        txHashes: execResult?.txHash
          ? [...plan.txHashes, execResult.txHash]
          : plan.txHashes,
      }
    : undefined;

  const handleExecute = async () => {
    setExecError(undefined);
    try {
      const result = await execute();
      setExecResult(result);
    } catch (e) {
      setExecError(e);
    }
  };

  if (isLoading) {
    return (
      <div className="space-y-4">
        <Skeleton className="h-8 w-2/3" />
        <Skeleton className="h-72 w-full" />
        <Skeleton className="h-40 w-full" />
      </div>
    );
  }

  if (error || !plan || !displayPlan) {
    return (
      <EmptyState
        title={PLAN.notFound}
        action={
          <Link href="/">
            <Button>{PLAN.reRequest}</Button>
          </Link>
        }
      />
    );
  }

  return (
    <div className="space-y-6">
      <PromiseCard
        plan={displayPlan}
        status={displayPlan.status}
        executing={isConfirming}
        wrongWallet={wrongWallet}
        onExecute={handleExecute}
      />
      <SettlementExplanation plan={displayPlan} />
      <ErrorBanner error={execError} />
    </div>
  );
}