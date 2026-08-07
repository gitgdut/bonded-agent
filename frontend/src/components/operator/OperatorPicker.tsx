"use client";

import type { OperatorStats } from "@/lib/types";
import { OperatorCard } from "./OperatorCard";

/** 运营方选择器：展示卡片列表，用户选择后高亮 */
export function OperatorPicker({
  operators,
  selected,
  onSelect,
  loading,
  error,
}: {
  operators: OperatorStats[];
  selected?: string;
  onSelect: (operator: OperatorStats) => void;
  loading?: boolean;
  error?: Error | null;
}) {
  if (loading) {
    return (
      <div className="space-y-2">
        <h2 className="text-sm font-semibold text-text-secondary">
          选择运营方
        </h2>
        <div className="glass-card flex items-center justify-center p-8">
          <div className="h-4 w-4 animate-spin rounded-full border-2 border-accent-tech border-t-transparent" />
          <span className="ml-3 text-sm text-text-muted">加载运营方列表中…</span>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="space-y-2">
        <h2 className="text-sm font-semibold text-text-secondary">
          选择运营方
        </h2>
        <div className="glass-card flex items-center justify-center p-8 border-red-400/30">
          <span className="text-sm text-red-400">
            加载失败：{error.message}
          </span>
        </div>
      </div>
    );
  }

  if (!operators.length) {
    return null;
  }

  return (
    <div className="space-y-2">
      <h2 className="text-sm font-semibold text-text-secondary">
        选择运营方
      </h2>
      <div className="grid gap-3 sm:grid-cols-3">
        {operators.map((op) => (
          <OperatorCard
            key={op.address}
            operator={op}
            selected={selected === op.address}
            onClick={() => onSelect(op)}
          />
        ))}
      </div>
    </div>
  );
}
