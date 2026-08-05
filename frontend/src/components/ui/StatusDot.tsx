"use client";

type DotColor = "primary" | "accent" | "positive" | "shortfall" | "failed" | "neutral";

const DOT_STYLE: Record<DotColor, string> = {
  primary: "bg-primary",
  accent: "bg-accent-tech",
  positive: "bg-positive",
  shortfall: "bg-shortfall",
  failed: "bg-failed",
  neutral: "bg-neutral",
};

/** 链上状态灯(6px 圆点,可呼吸闪烁) */
export function StatusDot({ color, pulse = false }: { color: DotColor; pulse?: boolean }) {
  return (
    <span
      className={`status-dot ${DOT_STYLE[color]} ${pulse ? "status-dot-pulse" : ""}`}
      aria-hidden
    />
  );
}

export type { DotColor };