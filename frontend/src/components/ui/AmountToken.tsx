"use client";

import { formatAmount } from "@/lib/format";

interface AmountTokenProps {
  value?: string; // wei
  token: string;
  decimals?: number;
  skeleton?: boolean;
  glow?: boolean;
  className?: string;
}

/** 代币金额:等宽数字 + 代币名;加载时显示骨架 */
export function AmountToken({
  value,
  token,
  decimals = 6,
  skeleton = false,
  glow = false,
  className = "",
}: AmountTokenProps) {
  if (skeleton || value === undefined) {
    return (
      <span className="skeleton inline-block h-5 w-20 align-middle" aria-hidden />
    );
  }
  return (
    <span className={`num ${glow ? "text-glow-bond" : ""} font-semibold ${className}`}>
      {formatAmount(value, decimals)}{" "}
      <span className="text-sm font-normal text-text-secondary">{token}</span>
    </span>
  );
}