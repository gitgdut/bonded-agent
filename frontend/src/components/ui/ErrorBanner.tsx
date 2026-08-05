"use client";

import { errorMessage } from "@/lib/copy";
import { Button } from "./Button";

/** 错误提示条:展示失败原因 + 可操作重试 */
export function ErrorBanner({
  error,
  onRetry,
  className = "",
}: {
  error: unknown;
  onRetry?: () => void;
  className?: string;
}) {
  if (!error) return null;
  return (
    <div
      role="alert"
      className={`flex items-start justify-between gap-4 rounded-lg border border-failed/40 bg-failed/10 px-4 py-3 text-sm text-failed ${className}`}
    >
      <span>⚠ {errorMessage(error)}</span>
      {onRetry && (
        <Button variant="secondary" onClick={onRetry} className="shrink-0 !px-3 !py-1 text-xs">
          重试
        </Button>
      )}
    </div>
  );
}