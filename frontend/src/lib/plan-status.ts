/**
 * 计划状态机(文档 §8)
 * 派生顺序:先看结算事件;若无执行事件,再按 deadline 判断 open/expired
 */

import type { Plan, PlanStatus } from "./types";

/** 根据计划数据派生当前状态 */
export function derivePlanStatus(plan: Plan, now = Date.now()): PlanStatus {
  // 简化:以后端/链上给出的 status 为准;本地补充 deadline 判断
  if (plan.status === "open" && plan.deadline <= now) return "expired";
  return plan.status;
}

/** 状态 -> 徽章语义(颜色、文案) */
export const STATUS_META: Record<
  PlanStatus,
  { label: string; color: "positive" | "shortfall" | "failed" | "neutral" | "primary" | "accent" }
> = {
  loading: { label: "读取中", color: "neutral" },
  open: { label: "待执行", color: "primary" },
  expired: { label: "已过期", color: "neutral" },
  executing: { label: "执行中", color: "accent" },
  settled_ok: { label: "已履约", color: "positive" },
  settled_shortfall: { label: "已赔付", color: "shortfall" },
  failed: { label: "已失败", color: "failed" },
  rejected: { label: "已取消", color: "neutral" },
};

/** 是否允许执行 */
export function canExecute(status: PlanStatus): boolean {
  return status === "open";
}

/** 是否已结算(展示解释卡片) */
export function isSettled(status: PlanStatus): boolean {
  return (
    status === "settled_ok" ||
    status === "settled_shortfall" ||
    status === "failed"
  );
}
