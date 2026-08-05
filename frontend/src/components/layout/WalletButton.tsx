"use client";

import { useAccount, useConnect, useDisconnect } from "wagmi";
import { Button } from "@/components/ui/Button";
import { shortAddress } from "@/lib/format";
import { COPY } from "@/lib/copy";

/** 钱包连接按钮:未连接显示连接,已连接显示地址 + 断开 */
export function WalletButton() {
  const { address, isConnected } = useAccount();
  const { connect, connectors } = useConnect();
  const { disconnect } = useDisconnect();

  if (isConnected && address) {
    return (
      <div className="flex items-center gap-2">
        <span className="rounded-lg border border-primary/40 bg-primary-soft px-3 py-1.5 font-mono text-xs text-text-primary">
          {shortAddress(address)}
        </span>
        <button
          onClick={() => disconnect()}
          className="text-xs text-text-muted transition-colors hover:text-text-primary"
        >
          断开
        </button>
      </div>
    );
  }

  return (
    <Button
      onClick={() => connect({ connector: connectors[0] })}
      disabled={!connectors.length}
    >
      {COPY.connectWallet}
    </Button>
  );
}