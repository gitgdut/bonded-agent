"use client";

import { useEffect, useState } from "react";
import type { Plan, PlanStatus } from "@/lib/types";
import { PLAN, planHeading } from "@/lib/copy";
import { formatAmount, describeSwap } from "@/lib/format";
import { canExecute } from "@/lib/plan-status";
import { INPUT_TOKEN, OUTPUT_TOKEN } from "@/lib/config";
import { AmountToken } from "@/components/ui/AmountToken";
import { Button } from "@/components/ui/Button";
import { PlanStatusBadge } from "./PlanStatusBadge";
import { StatusDot } from "@/components/ui/StatusDot";

/**
 * 承诺卡片(文档 §7,核心组件)
 * A 标题 / B 交易摘要 / C 保证条款(不可折叠) / D 风险边界 / E 操作
 */
export function PromiseCard({
  plan,
  status,
  executing,
  wrongWallet,
  onExecute,
  onGaslessExecute,
}: {
  plan: Plan;
  status: PlanStatus;
  executing?: boolean;
  wrongWallet?: boolean;
  onExecute?: () => void;
  onGaslessExecute?: () => void;
}) {
  const [now, setNow] = useState<number>();

  // 倒计时每秒刷新(等宽数字防跳动)
  useEffect(() => {
    const timer = setInterval(() => setNow(Date.now()), 1000);
    return () => clearInterval(timer);
  }, []);

  const heading = planHeading(plan.planId, plan.deadline);

  const actionDisabled = !canExecute(status) || executing || wrongWallet;

  let actionLabel: string = PLAN.execute;
  if (status === "expired") actionLabel = PLAN.expired;
  else if (executing) actionLabel = PLAN.executing;
  else if (wrongWallet) actionLabel = PLAN.wrongWallet;
  else if (!canExecute(status)) actionLabel = PLAN.settled;

  return (
    <div className="glass-card overflow-hidden">
      {/* 顶部状态条 */}
      <div className="border-b border-border/60 bg-bg-elevated/60 px-5 py-3">
        <div className="flex flex-wrap items-center justify-between gap-2">
          <h2 className="font-mono text-sm font-semibold text-text-primary">
            {heading.title}
          </h2>
          <PlanStatusBadge status={status} />
        </div>
      </div>

      <div className="space-y-5 px-5 py-5">
        {/* B 交易摘要 */}
        <section aria-label={PLAN.summary}>
          <h3 className="mb-2 text-xs font-semibold uppercase tracking-wider text-text-muted">
            {PLAN.summary}
          </h3>
          <p className="text-lg font-semibold text-text-primary">
            {describeSwap(
              plan.inputAmount,
              plan.expectedOutput,
              INPUT_TOKEN,
              OUTPUT_TOKEN,
            )}
          </p>
        </section>

        <div className="data-flow-line" aria-hidden />

        {/* C 保证条款(最高信息层级,不可折叠) */}
        <section aria-label={PLAN.guarantees}>
          <h3 className="mb-3 text-xs font-semibold uppercase tracking-wider text-text-muted">
            {PLAN.guarantees}
          </h3>
          <dl className="space-y-3">
            <div className="flex items-center justify-between rounded-lg border border-bond/30 bg-bond/5 px-4 py-3">
              <dt className="text-sm text-text-secondary">{PLAN.fieldGuaranteedOutput}</dt>
              <dd>
                <AmountToken value={plan.guaranteedOutput} token={OUTPUT_TOKEN} glow />
              </dd>
            </div>
            <div className="flex items-center justify-between px-1 text-sm">
              <dt className="text-text-secondary">{PLAN.fieldBond}</dt>
              <dd className="num text-text-primary">
                {formatAmount(plan.maxCompensation)} {OUTPUT_TOKEN}
              </dd>
            </div>
            <div className="flex items-center justify-between px-1 text-sm">
              <dt className="text-text-secondary">{PLAN.fieldMaxCompensation}</dt>
              <dd className="num text-text-primary">
                {formatAmount(plan.maxCompensation)} {OUTPUT_TOKEN}
              </dd>
            </div>
            <div className="flex items-center justify-between px-1 text-sm">
              <dt className="text-text-secondary">{PLAN.fieldFailureCompensation}</dt>
              <dd className="num text-text-primary">
                {formatAmount(plan.failureCompensation)} {OUTPUT_TOKEN}
              </dd>
            </div>
            <div className="flex items-center justify-between px-1 text-sm">
              <dt className="text-text-secondary">{PLAN.fieldDeadline}</dt>
              <dd className="num text-text-primary">
                {heading.deadlineText}
                <span className="ml-2 rounded bg-primary-soft px-1.5 py-0.5 text-xs text-primary">
                  {now ? heading.countdown(now) : "—"}
                </span>
              </dd>
            </div>
          </dl>
        </section>

        {/* D 风险边界(固定模板) */}
        <section aria-label={PLAN.riskBoundary}>
          <h3 className="mb-2 text-xs font-semibold uppercase tracking-wider text-text-muted">
            {PLAN.riskBoundary}
          </h3>
          <ul className="space-y-1.5 text-xs leading-relaxed text-text-secondary">
            {PLAN.riskItems.map((item) => (
              <li key={item} className="flex gap-2">
                <span className="mt-1 text-primary">▸</span>
                <span>{item}</span>
              </li>
            ))}
          </ul>
        </section>

        {/* E 操作 */}
        <section className="border-t border-border/60 pt-4">
          <Button
            className="w-full"
            disabled={actionDisabled}
            onClick={onExecute}
          >
            {actionLabel}
          </Button>
          {onGaslessExecute && status === "open" && !executing && (
            <Button
              className="w-full mt-2"
              variant="secondary"
              disabled={wrongWallet}
              onClick={onGaslessExecute}
            >
              免 Gas 执行 (EIP-712)
            </Button>
          )}
          {wrongWallet && !executing && (
            <p className="mt-2 flex items-center justify-center gap-1.5 text-xs text-shortfall">
              <StatusDot color="shortfall" />
              {PLAN.wrongWallet}
            </p>
          )}
          {plan.isMock && (
            <p className="mt-2 text-center text-[11px] text-text-muted">
              演示计划,执行不产生真实链上交易
            </p>
          )}
        </section>
      </div>
    </div>
  );
}
