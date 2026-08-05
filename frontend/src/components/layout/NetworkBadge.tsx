"use client";

import { useChainId } from "wagmi";
import { CHAIN_ID, CHAIN_NAME } from "@/lib/config";
import { StatusDot } from "@/components/ui/StatusDot";

/** 网络徽章:显示当前链与网络状态 */
export function NetworkBadge() {
  const chainId = useChainId();
  const isCorrect = chainId === CHAIN_ID;
  return (
    <span
      className="inline-flex items-center gap-2 rounded-lg border border-border bg-bg-elevated px-3 py-1.5 text-xs text-text-secondary"
      title={isCorrect ? `已连接 ${CHAIN_NAME}` : `请切换到 ${CHAIN_NAME}(chainId ${CHAIN_ID})`}
    >
      <StatusDot color={isCorrect ? "positive" : "shortfall"} />
      {CHAIN_NAME}
      {!isCorrect && <span className="text-shortfall">请切换网络</span>}
    </span>
  );
}