/**
 * wagmi 配置(文档 §4.1 / §16 开放问题 1、3)
 * P0:注入式钱包(MetaMask/Rabby)
 */

import { createConfig, http } from "wagmi";
import { injected } from "wagmi/connectors";
import { defineChain } from "viem";
import { CHAIN_ID, CHAIN_NAME, RPC_URL, EXPLORER_URL } from "./config";

/** Monad Testnet 链定义(chainId 可在 .env.local 覆盖) */
export const monadTestnet = defineChain({
  id: CHAIN_ID,
  name: CHAIN_NAME,
  nativeCurrency: { name: "MON", symbol: "MON", decimals: 18 },
  rpcUrls: {
    default: { http: [RPC_URL] },
    public: { http: [RPC_URL] },
  },
  blockExplorers: {
    default: { name: "Monad Explorer", url: EXPLORER_URL },
  },
  testnet: true,
});

export const wagmiConfig = createConfig({
  chains: [monadTestnet],
  connectors: [injected()],
  transports: {
    [monadTestnet.id]: http(RPC_URL),
  },
  ssr: true,
});

/** 供构建/类型使用的导出 */
export { CHAIN_ID };