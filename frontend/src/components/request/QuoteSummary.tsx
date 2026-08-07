"use client";

import type { Quote } from "@/lib/types";
import { AmountToken } from "@/components/ui/AmountToken";
import { Button } from "@/components/ui/Button";
import { QUOTE, HOME } from "@/lib/copy";
import { DEFAULT_DEADLINE_MINUTES, OUTPUT_TOKEN } from "@/lib/config";
import { formatAmount } from "@/lib/format";

/** 报价预览(文档 §6.1.1):预期输出 + 保证输出 + 锁定保证金 + 有效期 */
export function QuoteSummary({
  quote,
  onCreatePlan,
  creating,
}: {
  quote: Quote;
  onCreatePlan: () => void;
  creating?: boolean;
}) {
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

      <dl className="grid grid-cols-2 gap-4">
        <div>
          <dt className="text-xs text-text-muted">{QUOTE.expectedOutput}</dt>
          <dd className="mt-1">
            <AmountToken value={quote.expectedOutput} token={OUTPUT_TOKEN} />
          </dd>
        </div>
        <div>
          <dt className="text-xs text-text-muted">
            {QUOTE.guaranteedOutput}
            {quote.operator && (
              <span className="ml-1 text-text-muted">
                ({(quote.operator.guaranteedRatio * 100).toFixed(0)}%)
              </span>
            )}
          </dt>
          <dd className="mt-1">
            <AmountToken value={quote.expectedOutput} token={OUTPUT_TOKEN} glow />
          </dd>
        </div>
        <div>
          <dt className="text-xs text-text-muted">{QUOTE.bondLocked}</dt>
          <dd className="mt-1 text-sm text-text-primary">
            约 {formatAmount((BigInt(quote.expectedOutput) * 12n) / 100n)} {OUTPUT_TOKEN}
          </dd>
        </div>
        <div>
          <dt className="text-xs text-text-muted">{QUOTE.validity}</dt>
          <dd className="mt-1 text-sm text-text-primary">
            {QUOTE.minutes(DEFAULT_DEADLINE_MINUTES)}
          </dd>
        </div>
        <div>
          <dt className="text-xs text-text-muted">
            服务费 ({((quote.operator?.serviceFeeBps ?? 30) / 100).toFixed(2)}%)
          </dt>
          <dd className="mt-1 text-sm text-text-primary">
            <AmountToken value={((BigInt(quote.expectedOutput) * BigInt(quote.operator?.serviceFeeBps ?? 30)) / 10000n).toString()} token={OUTPUT_TOKEN} />
          </dd>
        </div>
      </dl>

      <div className="data-flow-line" aria-hidden />

      <Button onClick={onCreatePlan} disabled={creating} className="w-full">
        {creating ? HOME.creatingPlan : HOME.createPlan}
      </Button>
      {quote.isMock && (
        <p className="text-center text-[11px] text-text-muted">演示报价,非真实汇率</p>
      )}
    </div>
  );
}