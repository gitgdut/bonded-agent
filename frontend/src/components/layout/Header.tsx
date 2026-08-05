"use client";

import Link from "next/link";
import { NetworkBadge } from "./NetworkBadge";
import { WalletButton } from "./WalletButton";
import { COPY, NAV } from "@/lib/copy";

/** 顶部导航:品牌 + 首页/历史入口 + 网络徽章 + 钱包 */
export function Header() {
  return (
    <header className="sticky top-0 z-20 border-b border-border/60 bg-bg/85 backdrop-blur-md">
      <div className="mx-auto flex h-14 max-w-3xl items-center justify-between gap-3 px-4">
        <Link href="/" className="flex items-center gap-2">
          <span className="flex h-6 w-6 items-center justify-center rounded-md bg-primary/20 text-primary">
            ◆
          </span>
          <span className="bg-gradient-to-r from-primary to-accent-tech bg-clip-text text-sm font-bold text-transparent">
            {COPY.appName}
          </span>
        </Link>

        <div className="flex items-center gap-3">
          <Link
            href="/"
            className="text-sm text-text-secondary transition-colors hover:text-primary"
          >
            {NAV.home}
          </Link>
          <Link
            href="/history"
            className="text-sm text-text-secondary transition-colors hover:text-primary"
          >
            {NAV.history}
          </Link>
          <NetworkBadge />
          <WalletButton />
        </div>
      </div>
    </header>
  );
}