"use client";

import { useState } from "react";
import {
  useAccount,
  useConnect,
  useConnectors,
  useDisconnect,
} from "wagmi";
import { Button } from "@/components/ui/Button";
import { shortAddress } from "@/lib/format";
import { COPY } from "@/lib/copy";

/**
 * Wallet connection button.
 * Handles wagmi v3's "reconnecting" state that can get stuck after page reload.
 */
export function WalletButton() {
  const { address, isConnected, status } = useAccount();
  const { connectAsync } = useConnect();
  const connectors = useConnectors();
  const { disconnectAsync } = useDisconnect();
  const [connecting, setConnecting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Show connected UI only when fully resolved
  if (isConnected && address) {
    return (
      <div className="flex items-center gap-2">
        <span className="rounded-lg border border-primary/40 bg-primary-soft px-3 py-1.5 font-mono text-xs text-text-primary">
          {shortAddress(address)}
        </span>
        <button
          onClick={() => disconnectAsync()}
          className="text-xs text-text-muted transition-colors hover:text-text-primary"
        >
          断开
        </button>
      </div>
    );
  }

  // Find the injected connector (MetaMask / Rabby)
  const injectedConnector =
    connectors.find(
      (c) =>
        c.type === "injected" ||
        c.id === "injected" ||
        c.name.toLowerCase().includes("meta") ||
        c.name.toLowerCase().includes("rabby"),
    ) ?? connectors[0];

  const handleConnect = async () => {
    setError(null);
    if (!injectedConnector) {
      setError("未检测到钱包插件，请安装 MetaMask 或 Rabby");
      return;
    }
    try {
      setConnecting(true);

      // If wagmi is stuck in a stale "reconnecting" state, disconnect first
      if (status === "reconnecting") {
        try {
          await disconnectAsync();
        } catch {
          // Ignore — connector might already be disconnected
        }
        // Small delay to let wagmi flush state
        await new Promise((r) => setTimeout(r, 300));
      }

      await connectAsync({ connector: injectedConnector });
    } catch (e: unknown) {
      const msg = e instanceof Error ? e.message : String(e);

      // User rejected — not an error to report
      if (
        msg.includes("rejected") ||
        msg.includes("denied") ||
        msg.includes("cancelled")
      ) {
        return;
      }

      // wagmi cookie persistence: connector thinks it's already connected
      if (msg.includes("already connected") || msg.includes("AlreadyConnected")) {
        try {
          await disconnectAsync();
          await new Promise((r) => setTimeout(r, 300));
          await connectAsync({ connector: injectedConnector });
          return;
        } catch (retryErr: unknown) {
          const retryMsg =
            retryErr instanceof Error ? retryErr.message : String(retryErr);
          if (
            retryMsg.includes("rejected") ||
            retryMsg.includes("denied") ||
            retryMsg.includes("cancelled")
          ) {
            return;
          }
          setError(retryMsg || "连接失败，请重试");
          return;
        }
      }

      setError(msg || "连接失败，请重试");
    } finally {
      setConnecting(false);
    }
  };

  // Only show spinner when user explicitly clicked connect.
  // Don't block the UI when wagmi is stuck in its own reconnecting state.
  const showSpinner = connecting;

  return (
    <div className="flex flex-col items-end gap-1">
      <Button
        onClick={handleConnect}
        disabled={!injectedConnector || connecting}
      >
        {showSpinner ? "连接中…" : COPY.connectWallet}
      </Button>
      {error && <p className="text-xs text-failed">{error}</p>}
      {!connectors.length && !error && (
        <p className="text-xs text-text-muted">请安装 MetaMask 或 Rabby 钱包</p>
      )}
    </div>
  );
}
