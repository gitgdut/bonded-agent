/**
 * 链上事件辅助(文档 §12.3)
 * 事件订阅的公共常量与过滤逻辑,供 hooks/usePlanEvents 与历史页使用
 */

import type { PlanEvent } from "./types";

/** 计划生命周期事件名(与合约 ABI 对应) */
export const PLAN_EVENT_NAMES = [
  "PlanOpened",
  "PlanExecuted",
  "ShortfallPaid",
  "PlanFailed",
  "BondReleased",
] as const;

export type PlanEventName = (typeof PLAN_EVENT_NAMES)[number];

/**
 * 将链上 raw log 归一化为 PlanEvent
 * TODO(合约发布后):按真实 ABI 的事件参数解析 topics/data
 */
export function normalizeEvent(raw: {
  eventName: string;
  args: Record<string, unknown>;
  transactionHash: string;
  blockNumber: bigint;
}): PlanEvent {
  return {
    type: raw.eventName as PlanEventName,
    planId: String(raw.args.planId ?? raw.args.nonce ?? ""),
    user: raw.args.user ? String(raw.args.user) : undefined,
    txHash: raw.transactionHash,
    blockNumber: raw.blockNumber,
    data: raw.args,
  };
}

/** 按用户地址过滤事件 */
export function filterEventsByUser(events: PlanEvent[], user?: string): PlanEvent[] {
  if (!user) return [];
  return events.filter((e) => !e.user || e.user.toLowerCase() === user.toLowerCase());
}

/** 事件列表按 planId 去重合并(保留最新事件) */
export function mergeEventsByPlan(events: PlanEvent[]): Map<string, PlanEvent[]> {
  const map = new Map<string, PlanEvent[]>();
  for (const e of events) {
    const list = map.get(e.planId) ?? [];
    list.push(e);
    map.set(e.planId, list);
  }
  return map;
}