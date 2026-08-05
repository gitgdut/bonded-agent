"use client";

export type ToastType = "success" | "error" | "info";

const TOAST_STYLE: Record<ToastType, string> = {
  success: "border-positive/50 bg-positive/10 text-positive",
  error: "border-failed/50 bg-failed/10 text-failed",
  info: "border-primary/50 bg-primary/10 text-text-primary",
};

/** 轻提示(成功/失败/信息) */
export function Toast({
  type,
  message,
  className = "",
}: {
  type: ToastType;
  message: string;
  className?: string;
}) {
  if (!message) return null;
  return (
    <div
      role="status"
      className={`rounded-lg border px-4 py-3 text-sm backdrop-blur ${TOAST_STYLE[type]} ${className}`}
    >
      {message}
    </div>
  );
}