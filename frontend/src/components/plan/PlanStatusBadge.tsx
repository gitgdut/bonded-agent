"use client";

import { STATUS_META } from "@/lib/plan-status";
import type { PlanStatus } from "@/lib/types";
import { StatusDot, type DotColor } from "@/components/ui/StatusDot";

const BADGE_STYLE: Record<DotColor, string> = {
  primary: "border-primary/40 bg-primary/10 text-primary",
  accent: "border-accent-tech/40 bg-accent-tech/10 text-accent-tech",
  positive: "border-positive/40 bg-positive/10 text-positive",
  shortfall: "border-shortfall/40 bg-shortfall/10 text-shortfall",
  failed: "border-failed/40 bg-failed/10 text-failed",
  neutral: "border-neutral/40 bg-neutral/10 text-neutral",
};

/** 计划状态徽章:胶囊形,同色系深底 + 亮字 + 状态灯(文档 §10.5) */
export function PlanStatusBadge({
  status,
  className = "",
}: {
  status: PlanStatus;
  className?: string;
}) {
  const meta = STATUS_META[status] ?? STATUS_META.loading;
  return (
    <span
      className={`inline-flex items-center gap-2 rounded-full border px-3 py-1 text-xs font-medium ${BADGE_STYLE[meta.color]} ${className}`}
    >
      <StatusDot
        color={meta.color}
        pulse={status === "executing" || status === "loading"}
      />
      {meta.label}
    </span>
  );
}