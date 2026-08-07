"use client";

import { useParams } from "next/navigation";
import Link from "next/link";
import { useMemo, useState } from "react";
import { useAccount } from "wagmi";
import { usePlan } from "@/hooks/useQuote";
import { useExecutePlan } from "@/hooks/useExecutePlan";
import { PromiseCard } from "@/components/plan/PromiseCard";
import { SettlementExplanation } from "@/components/plan/SettlementExplanation";
import { ErrorBanner } from "@/components/ui/ErrorBanner";
import { Skeleton } from "@/components/ui/Skeleton";
import { EmptyState } from "@/components/ui/EmptyState";
import { Button } from "@/components/ui/Button";
import { PLAN } from "@/lib/copy";
import { USE_MOCK } from "@/lib/config";
import type { Plan } from "@/lib/types";

/** Plan page: promise card + execute + settlement (doc §6.2) */
export default function PlanPage() {
  const params = useParams<{ id: string }>();
  const planId = params?.id;

  const { plan, status, isLoading, error } = usePlan(planId);
  const { address } = useAccount();
  const { executeWithSignature, isConfirming, result: execResult } = useExecutePlan(plan);

  const [execError, setExecError] = useState<unknown>();

  // Wallet check: skip in demo mode; in real mode verify plan ownership
  const wrongWallet =
    !USE_MOCK &&
    !!plan &&
    !!address &&
    !!plan.user &&
    plan.user.toLowerCase() !== address.toLowerCase();

  // Merge the reactive execution result with backend plan data.
  // Only trust execResult for final statuses (not "executing").
  // For "executing", show the backend-derived status as fallback.
  const displayPlan: Plan | undefined = useMemo(() => {
    if (!plan) return undefined;

    const finalStatus = execResult?.status;
    const isFinal =
      finalStatus === "settled_ok" ||
      finalStatus === "settled_shortfall" ||
      finalStatus === "failed";

    return {
      ...plan,
      // Only override status with final on-chain results; otherwise trust backend
      status: (isFinal ? finalStatus : status) as Plan["status"],
      actualOutput: execResult?.actualOutput ?? plan.actualOutput,
      shortfallPaid: execResult?.shortfallPaid ?? plan.shortfallPaid,
      compensation: execResult?.compensation ?? plan.compensation,
      txHashes: execResult?.txHash
        ? [...plan.txHashes, execResult.txHash]
        : plan.txHashes,
    };
  }, [plan, execResult, status]);

  const handleExecute = async () => {
    setExecError(undefined);
    try {
      await executeWithSignature();
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
        executing={isConfirming || execResult?.status === "executing"}
        wrongWallet={wrongWallet}
        onExecute={handleExecute}
      />
      <SettlementExplanation plan={displayPlan} />
      <ErrorBanner error={execError} />
    </div>
  );
}
