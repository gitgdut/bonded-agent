"use client";

import type { ButtonHTMLAttributes } from "react";

type Variant = "primary" | "secondary";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: Variant;
}

/** 统一按钮:primary 紫→靛渐变 / secondary 描边(样式见 globals.css) */
export function Button({ variant = "primary", className = "", ...props }: ButtonProps) {
  const cls = variant === "primary" ? "btn-primary" : "btn-secondary";
  return (
    <button
      className={`${cls} inline-flex items-center justify-center gap-2 px-5 py-2.5 text-sm font-medium transition-colors ${className}`}
      {...props}
    />
  );
}