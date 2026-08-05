"use client";

import { useMemo, useState } from "react";
import { useAccount } from "wagmi";
import { parseEther } from "viem";
import { Button } from "@/components/ui/Button";
import { validateInputAmount } from "@/lib/format";
import { HOME, COPY } from "@/lib/copy";
import { DEFAULT_INPUT_AMOUNT, OUTPUT_TOKEN } from "@/lib/config";

/** 请求表单:金额输入 + 获取报价(文档 §6.1) */
export function RequestForm({
  onQuote,
  loading,
}: {
  onQuote: (amountWei: string) => void;
  loading?: boolean;
}) {
  const { isConnected } = useAccount();
  const [raw, setRaw] = useState(DEFAULT_INPUT_AMOUNT);

  const error = useMemo(() => validateInputAmount(raw), [raw]);

  const handleSubmit = () => {
    if (!isConnected) {
      alert(HOME.walletPrompt);
      return;
    }
    const err = validateInputAmount(raw);
    if (err) return;
    onQuote(parseEther(raw.trim() as `${number}`).toString());
  };

  return (
    <div className="glass-card space-y-4 p-5">
      <div className="grid gap-4 sm:grid-cols-2">
        <label className="block">
          <span className="mb-1.5 block text-sm text-text-secondary">{HOME.inputLabel}</span>
          <input
            type="text"
            inputMode="decimal"
            value={raw}
            onChange={(e) => setRaw(e.target.value)}
            placeholder="1"
            aria-invalid={!!error}
            className="w-full rounded-lg border border-border bg-bg-elevated px-3 py-2.5 font-mono text-text-primary outline-none transition-colors placeholder:text-text-muted focus:border-primary focus:ring-2 focus:ring-primary/40"
          />
          {error && <span className="mt-1 block text-xs text-failed">{error}</span>}
        </label>

        <label className="block">
          <span className="mb-1.5 block text-sm text-text-secondary">{HOME.outputLabel}</span>
          <div className="flex w-full items-center justify-between rounded-lg border border-border bg-bg-elevated px-3 py-2.5 text-text-primary">
            {OUTPUT_TOKEN}
            <span className="text-xs text-text-muted">固定</span>
          </div>
        </label>
      </div>

      <Button onClick={handleSubmit} disabled={!!error || loading} className="w-full">
        {loading ? HOME.generating : HOME.getQuote}
      </Button>

      {!isConnected && (
        <p className="text-center text-xs text-text-muted">{HOME.walletPrompt}</p>
      )}
    </div>
  );
}

export { COPY };