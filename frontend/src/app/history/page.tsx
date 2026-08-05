"use client";

import Link from "next/link";
import { useMemo, useState } from "react";
import { useAccount } from "wagmi";
import { usePlanEvents } from "@/hooks/usePlanEvents";
import { PlanStatusBadge } from "@/components/plan/PlanStatusBadge";
import { EmptyState } from "@/components/ui/EmptyState";
import { Skeleton } from "@/components/ui/Skeleton";
import { Button } from "@/components/ui/Button";
import { TxLink } from "@/components/ui/TxLink";
import { HISTORY, PLAN } from "@/lib/copy";
import { describeSwap, formatDateTimeShort, formatAmount } from "@/lib/format";
import { derivePlanStatus } from "@/lib/plan-status";
import { INPUT_TOKEN, OUTPUT_TOKEN } from "@/lib/config";
import type { PlanStatus } from "@/lib/types";

type Filter = (typeof HISTORY.filters)[number];

const FILTER_TO_STATUS: Record<Filter, PlanStatus | "all"> = {
  全部: "all",
  待执行: "open",
  履约: "settled_ok",
  赔付: "settled_shortfall",
  失败: "failed",
  过期: "expired",
};

/** 历史页:按状态筛选的担保计划列表(文档 §6.3) */
export default function HistoryPage() {
  const { address } = useAccount();
  const { plans, isLoading } = usePlanEvents(address);
  const [filter, setFilter] = useState<Filter>("全部");

  const rows = useMemo(() => {
    const target = FILTER_TO_STATUS[filter];
    return plans
      .map((p) => ({ ...p, derivedStatus: derivePlanStatus(p) }))
      .filter((p) => target === "all" || p.derivedStatus === target)
      .sort((a, b) => (b.updatedAt ?? b.deadline) - (a.updatedAt ?? a.deadline));
  }, [plans, filter]);

  return (
    <div className="space-y-6">
      <div className="pt-4">
        <h1 className="text-2xl font-bold text-text-primary sm:text-[28px]">
          {HISTORY.title}
        </h1>
      </div>

      {/* 筛选条 */}
      <div className="flex flex-wrap gap-2">
        {HISTORY.filters.map((f) => (
          <button
            key={f}
            onClick={() => setFilter(f)}
            className={`rounded-full border px-3.5 py-1.5 text-xs transition-colors ${
              filter === f
                ? "border-primary bg-primary/15 text-primary"
                : "border-border bg-bg-elevated text-text-secondary hover:border-primary/50"
            }`}
          >
            {f}
          </button>
        ))}
      </div>

      {/* 列表 */}
      {isLoading ? (
        <div className="space-y-3">
          {[0, 1, 2].map((i) => (
            <Skeleton key={i} className="h-24 w-full" />
          ))}
        </div>
      ) : rows.length === 0 ? (
        <EmptyState
          title={HISTORY.empty}
          action={
            <Link href="/">
              <Button>{HISTORY.goRequest}</Button>
            </Link>
          }
        />
      ) : (
        <ul className="space-y-3">
          {rows.map((plan) => {
            const txHash = plan.txHashes[plan.txHashes.length - 1];
            return (
              <li key={plan.planId} className="glass-card px-4 py-3.5">
                <div className="flex flex-wrap items-center justify-between gap-3">
                  <div className="min-w-0">
                    <div className="flex items-center gap-2">
                      <Link
                        href={`/plans/${encodeURIComponent(plan.planId)}`}
                        className="truncate font-mono text-sm font-medium text-text-primary hover:text-primary"
                      >
                        #{plan.planId}
                      </Link>
                      <PlanStatusBadge status={plan.derivedStatus} />
                    </div>
                    <p className="mt-1 truncate text-sm text-text-secondary">
                      {describeSwap(
                        plan.inputAmount,
                        plan.expectedOutput,
                        INPUT_TOKEN,
                        OUTPUT_TOKEN,
                      )}
                    </p>
                    <p className="mt-0.5 text-xs text-text-muted">
                      {formatDateTimeShort(plan.updatedAt ?? plan.deadline)}
                      {plan.isMock ? " · 演示数据" : ""}
                    </p>
                  </div>

                  <div className="flex shrink-0 flex-col items-end gap-1.5">
                    {plan.actualOutput && (
                      <span className="num text-xs text-text-secondary">
                        到账 {formatAmount(plan.actualOutput)} {OUTPUT_TOKEN}
                      </span>
                    )}
                    {plan.shortfallPaid && (
                      <span className="num text-xs text-shortfall">
                        补足 {formatAmount(plan.shortfallPaid)} {OUTPUT_TOKEN}
                      </span>
                    )}
                    {plan.actualOutput && (
                      <span className="num text-xs text-text-secondary">
                        {INPUT_TOKEN} 支出 {formatAmount(plan.inputAmount)}
                      </span>
                    )}
                    {txHash && <TxLink txHash={txHash} label={PLAN.viewOnExplorer} />}
                  </div>
                </div>
              </li>
            );
          })}
        </ul>
      )}
    </div>
  );
}