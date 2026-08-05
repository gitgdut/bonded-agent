"use client";

import type { Plan } from "@/lib/types";
import { PLAN } from "@/lib/copy";
import { isSettled } from "@/lib/plan-status";
import { TxLink } from "@/components/ui/TxLink";

/** 结算解释卡片(文档 §8.3):一句话解释 + 区块链接 */
export function SettlementExplanation({ plan }: { plan: Plan }) {
  if (!isSettled(plan.status)) return null;

  const explainFn =
    PLAN.explain[plan.status as keyof typeof PLAN.explain] ?? PLAN.explain.expired;
  const text = explainFn(plan);
  const txHash = plan.txHashes[plan.txHashes.length - 1];

  return (
    <div className="glass-card space-y-3 p-5">
      <h3 className="text-sm font-semibold text-text-primary">结算说明</h3>
      <p className="text-sm leading-relaxed text-text-secondary">{text}</p>
      {txHash && <TxLink txHash={txHash} label={PLAN.viewOnExplorer} />}
    </div>
  );
}