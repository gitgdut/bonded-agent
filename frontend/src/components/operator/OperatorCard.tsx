"use client";

import type { OperatorStats } from "@/lib/types";
import { shortAddress } from "@/lib/format";

/** 声誉评分颜色映射 */
function scoreColor(score: number): string {
  if (score >= 80) return "text-green-400";
  if (score >= 60) return "text-yellow-400";
  return "text-red-400";
}

function scoreBg(score: number): string {
  if (score >= 80) return "bg-green-400/10 border-green-400/30";
  if (score >= 60) return "bg-yellow-400/10 border-yellow-400/30";
  return "bg-red-400/10 border-red-400/30";
}

/** 运营方卡片：头像、名称、评分、成功率、费率 */
export function OperatorCard({
  operator,
  selected,
  onClick,
}: {
  operator: OperatorStats;
  selected?: boolean;
  onClick?: () => void;
}) {
  const name = operator.name ?? shortAddress(operator.address);
  const feePercent = ((operator.serviceFeeBps ?? 30) / 100).toFixed(2);

  const handleClick = () => {
    onClick?.();
  };

  return (
    <div
      role="button"
      tabIndex={0}
      onClick={handleClick}
      onKeyDown={(e) => {
        if (e.key === "Enter" || e.key === " ") handleClick();
      }}
      style={selected ? {
        border: "2px solid #8b5cf6",
        boxShadow: "0 0 16px rgba(139, 92, 246, 0.35), 0 8px 24px rgba(0, 0, 0, 0.45)",
        background: "rgba(139, 92, 246, 0.08)",
      } : undefined}
      className={`glass-card w-full p-4 text-left transition-all cursor-pointer select-none ${
        selected
          ? ""
          : "hover:ring-1 hover:ring-white/10"
      }`}
    >
      <div className="flex items-center gap-3">
        {/* Avatar */}
        <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-full bg-accent-tech/20 text-sm font-bold text-accent-tech">
          {name.charAt(0).toUpperCase()}
        </div>

        {/* Name + address */}
        <div className="min-w-0 flex-1">
          <div className="flex items-center gap-2">
            <span className="text-sm font-semibold text-text-primary">
              {name}
            </span>
            {operator.erc8004AgentId && (
              <span className="rounded bg-purple-400/15 px-1.5 py-0.5 text-[10px] text-purple-400">
                ERC-8004 #{operator.erc8004AgentId}
              </span>
            )}
            {operator.isDefault && (
              <span className="rounded bg-accent-tech/15 px-1.5 py-0.5 text-[10px] text-accent-tech">
                Default
              </span>
            )}
          </div>
          <p className="truncate text-[11px] text-text-muted">
            {operator.address}
          </p>
        </div>

        {/* Score */}
        <div
          className={`flex h-10 w-10 shrink-0 items-center justify-center rounded-full border text-sm font-bold ${scoreColor(operator.reputationScore)} ${scoreBg(operator.reputationScore)}`}
        >
          {Math.round(operator.reputationScore)}
        </div>
      </div>

      {/* Stats row */}
      <div className="mt-3 flex gap-4 text-[11px]">
        <span className="text-text-muted">
          成功率{" "}
          <span className="text-text-primary">
            {operator.successRate.toFixed(0)}%
          </span>
        </span>
        <span className="text-text-muted">
          费率{" "}
          <span className="text-text-primary">{feePercent}%</span>
        </span>
        <span className="text-text-muted">
          保证率{" "}
          <span className="text-text-primary">
            {((operator.guaranteedRatio ?? 0.9) * 100).toFixed(0)}%
          </span>
        </span>
        <span className="text-text-muted">
          交易{" "}
          <span className="text-text-primary">{operator.totalPlans}</span>
        </span>
      </div>
    </div>
  );
}
