"use client";

import { EXPLORER_URL, USE_MOCK } from "@/lib/config";
import { COPY } from "@/lib/copy";

/**
 * 区块浏览器交易链接。
 * 演示模式(mock)下生成的交易哈希在链上不存在,点击只会打开无效页面,
 * 因此渲染为普通提示文字而不是链接。
 */
export function TxLink({
  txHash,
  label,
  className = "",
}: {
  txHash: string;
  label?: string;
  className?: string;
}) {
  if (!txHash) return null;
  if (USE_MOCK) {
    return (
      <span
        className={`text-sm text-text-muted ${className}`}
        title="演示模式生成的交易哈希,链上不存在"
      >
        {label ?? "在区块浏览器中查看"} · {COPY.demoTxHint}
      </span>
    );
  }
  return (
    <a
      href={`${EXPLORER_URL}/tx/${txHash}`}
      target="_blank"
      rel="noopener noreferrer"
      className={`text-sm text-primary underline-offset-4 hover:underline ${className}`}
    >
      {label ?? "在区块浏览器中查看"} ↗
    </a>
  );
}