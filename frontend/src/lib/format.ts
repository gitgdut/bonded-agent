/**
 * 格式化工具(文档 §14.1)
 * 金额统一以 wei 字符串输入,输出可读字符串
 */

import { formatEther } from "viem";

/** wei -> 可读金额,保留最多 6 位小数,去除末尾 0 */
export function formatAmount(wei: string | bigint | number, decimals = 6): string {
  let value: string;
  try {
    value = formatEther(BigInt(wei));
  } catch {
    return "0";
  }
  const [int, frac] = value.split(".");
  if (!frac) return int;
  const trimmed = frac.slice(0, decimals).replace(/0+$/, "");
  return trimmed ? `${int}.${trimmed}` : int;
}

/** 时间戳(ms) -> "YYYY-MM-DD HH:mm" */
export function formatDateTime(ts: number): string {
  const d = new Date(ts);
  const pad = (n: number) => String(n).padStart(2, "0");
  return `${d.getFullYear()}-${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}`;
}

/** 时间戳(ms) -> "MM-DD HH:mm:ss" */
export function formatDateTimeShort(ts: number): string {
  const d = new Date(ts);
  const pad = (n: number) => String(n).padStart(2, "0");
  return `${pad(d.getMonth() + 1)}-${pad(d.getDate())} ${pad(d.getHours())}:${pad(d.getMinutes())}:${pad(d.getSeconds())}`;
}

/** 倒计时文本(ms),过期返回 "已过期" */
export function formatCountdown(targetTs: number, now: number): string {
  const diff = targetTs - now;
  if (diff <= 0) return "已过期";
  const s = Math.floor(diff / 1000);
  const h = Math.floor(s / 3600);
  const m = Math.floor((s % 3600) / 60);
  const sec = s % 60;
  const pad = (n: number) => String(n).padStart(2, "0");
  if (h > 0) return `${h}:${pad(m)}:${pad(sec)}`;
  return `${pad(m)}:${pad(sec)}`;
}

/** 地址缩写 0x1234…abcd */
export function shortAddress(address: string, chars = 4): string {
  if (!address) return "";
  if (address.length <= 2 + chars * 2) return address;
  return `${address.slice(0, 2 + chars)}…${address.slice(-chars)}`;
}

/** 交易哈希缩写 */
export function shortTxHash(hash: string): string {
  return shortAddress(hash, 6);
}

/** 金额描述: "1 MON -> 0.95 tUSDC" */
export function describeSwap(inputWei: string, outputWei: string, inToken: string, outToken: string): string {
  return `${formatAmount(inputWei)} ${inToken} → ${formatAmount(outputWei)} ${outToken}`;
}

/** 校验输入金额(正数、精度 <= 18),非法返回错误信息,合法返回 null */
export function validateInputAmount(raw: string): string | null {
  const value = raw.trim();
  if (!value) return "请输入金额";
  const regex = /^\d+(\.\d+)?$/;
  if (!regex.test(value)) return "金额格式不正确";
  const parts = value.split(".");
  if (parts[1] && parts[1].length > 18) return "精度最多 18 位小数";
  if (BigInt(parts[0]) === 0n && (!parts[1] || parts[1].replace(/0+$/, "") === "")) {
    return "金额必须大于 0";
  }
  return null;
}