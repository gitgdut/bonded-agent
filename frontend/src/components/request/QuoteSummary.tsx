"use client";

import type { Quote } from "@/lib/types";
import { AmountToken } from "@/components/ui/AmountToken";
import { Button } from "@/components/ui/Button";
import { QUOTE, HOME } from "@/lib/copy";
import { DEFAULT_DEADLINE_MINUTES, OUTPUT_TOKEN, MAX_COMPENSATION } from "@/lib/config";
import { formatAmount } from "@/lib/format";

/** 报价预览（文档 §6.1.1）：预期 → 服务费 → 净到账 → 保证到账 → 风险敞口 → 保证金 → 有效期 */
export function QuoteSummary({
  quote,
  onCreatePlan,
  creating,
}: {
  quote: Quote;
  onCreatePlan: () => void;
  creating?: boolean;
}) {
  const feeBps = BigInt(quote.operator?.serviceFeeBps ?? 30);
  const ratio = quote.operator?.guaranteedRatio ?? 0.9;
  const expected = BigInt(quote.expectedOutput);

  const feeAmount = (expected * feeBps) / 10000n;
  const netAfterFee = expected - feeAmount;
  // guaranteed = netAfterFee * ratio (e.g. 0.90 → 9000 bps)
  const ratioBps = BigInt(Math.round(ratio * 10000));
  const guaranteed = (netAfterFee * ratioBps) / 10000n;
  const riskCovered = netAfterFee > guaranteed ? netAfterFee - guaranteed : 0n;

  const bondDisplay = MAX_COMPENSATION; // 20 tUSDC = deposit = max(maxComp, failureComp)

  return (
    <div className="glass-card space-y-4 p-5">
      <div className="flex items-center justify-between">
        <h2 className="text-sm font-semibold text-text-secondary">报价预览</h2>
        <div className="flex items-center gap-2">
          <span className="rounded-full border border-accent-warm/40 bg-accent-warm/10 px-2 py-0.5 text-[11px] text-accent-warm">
            {quote.protocol ?? "simple-amm"}
          </span>
          <span className="rounded-full border border-accent-tech/40 bg-accent-tech/10 px-2 py-0.5 text-[11px] text-accent-tech">
            {HOME.simulatedRateTag}
          </span>
        </div>
      </div>

      {/* Vertical flow */}
      <div className="space-y-0">
        {/* 1. 预期输出 */}
        <FlowRow>
          <FlowLabel>{QUOTE.expectedOutput}</FlowLabel>
          <AmountToken value={quote.expectedOutput} token={OUTPUT_TOKEN} />
        </FlowRow>

        {/* 2. 服务费 */}
        <FlowRow muted>
          <FlowLabel>
            服务费（{(Number(feeBps) / 100).toFixed(2)}%）
          </FlowLabel>
          <span className="num text-sm text-text-muted">
            -{formatAmount(feeAmount.toString())} {OUTPUT_TOKEN}
          </span>
        </FlowRow>

        {/* connector line */}
        <div className="ml-4 h-4 border-l border-white/10" />

        {/* 3. 净到账 */}
        <FlowRow>
          <FlowLabel>净到账</FlowLabel>
          <AmountToken value={netAfterFee.toString()} token={OUTPUT_TOKEN} />
        </FlowRow>

        {/* connector line */}
        <div className="ml-4 h-4 border-l border-white/10" />

        {/* 4. 保证到账 — THIS WAS THE BUG FIX */}
        <FlowRow highlight>
          <FlowLabel>
            {QUOTE.guaranteedOutput}
            {quote.operator && (
              <span className="ml-1 text-text-muted">
                ({(ratio * 100).toFixed(0)}%)
              </span>
            )}
          </FlowLabel>
          <AmountToken value={guaranteed.toString()} token={OUTPUT_TOKEN} glow />
        </FlowRow>

        {/* connector line */}
        <div className="ml-4 h-4 border-l border-white/10" />

        {/* 5. 风险敞口 */}
        <FlowRow muted>
          <FlowLabel>风险敞口（无担保部分）</FlowLabel>
          <span className="num text-sm text-accent-warm">
            {formatAmount(riskCovered.toString())} {OUTPUT_TOKEN}
          </span>
        </FlowRow>
      </div>

      <div className="flex items-center gap-4 border-t border-white/10 pt-3">
        {/* 6. 保证金 */}
        <div className="flex-1">
          <dt className="text-xs text-text-muted">{QUOTE.bondLocked}</dt>
          <dd className="num mt-0.5 text-sm font-semibold text-text-primary">
            {formatAmount(bondDisplay)} {OUTPUT_TOKEN}
          </dd>
        </div>
        {/* 7. 有效期 */}
        <div className="flex-1 text-right">
          <dt className="text-xs text-text-muted">{QUOTE.validity}</dt>
          <dd className="num mt-0.5 text-sm font-semibold text-text-primary">
            {QUOTE.minutes(DEFAULT_DEADLINE_MINUTES)}
          </dd>
        </div>
      </div>

      <Button onClick={onCreatePlan} disabled={creating} className="w-full">
        {creating ? HOME.creatingPlan : HOME.createPlan}
      </Button>
      {quote.isMock && (
        <p className="text-center text-[11px] text-text-muted">演示报价，非真实汇率</p>
      )}
    </div>
  );
}

/* ── Vertical flow helpers ───────────────────────────────── */

function FlowRow({
  children,
  muted,
  highlight,
}: {
  children: React.ReactNode;
  muted?: boolean;
  highlight?: boolean;
}) {
  return (
    <div
      className={`flex items-center justify-between rounded-md px-3 py-2 ${
        highlight
          ? "border border-accent-bond/30 bg-accent-bond/8"
          : muted
            ? ""
            : ""
      }`}
    >
      {children}
    </div>
  );
}

function FlowLabel({
  children,
}: {
  children: React.ReactNode;
}) {
  return <span className="text-xs text-text-secondary">{children}</span>;
}
