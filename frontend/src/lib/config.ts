/**
 * 全局配置(环境变量驱动,便于部署与后端对接)
 * 复制 .env.example 为 .env.local 并修改
 */

/** 后端服务地址(运营方 API) */
/** 后端服务地址。为空时走 Next.js rewrite 代理（/quote → localhost:8787）。 */
export const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_BASE_URL ?? "http://localhost:8787";
export const USE_API_PROXY = API_BASE_URL === "";

/**
 * 是否使用演示数据。
 * 后端未就绪时置 "1" 可完整体验流程;对接后置 "0"。
 * 所有演示数据都会带 isMock 标记,并在 UI 上标注"演示"。
 */
export const USE_MOCK = (process.env.NEXT_PUBLIC_USE_MOCK ?? "1") === "1";

/** Monad Testnet(默认 10143,可在 .env.local 覆盖) */
export const CHAIN_ID = Number(process.env.NEXT_PUBLIC_CHAIN_ID ?? 10143);
export const CHAIN_NAME =
  process.env.NEXT_PUBLIC_CHAIN_NAME ?? "Monad Testnet";
export const RPC_URL =
  process.env.NEXT_PUBLIC_RPC_URL ?? "https://testnet-rpc.monad.xyz";
export const EXPLORER_URL =
  process.env.NEXT_PUBLIC_EXPLORER_URL ?? "https://testnet.monadexplorer.com";

/** 担保合约地址(由部署方提供) */
export const PLAN_CONTRACT_ADDRESS =
  process.env.NEXT_PUBLIC_PLAN_CONTRACT ?? "";

/** 交易对 */
export const INPUT_TOKEN = "MON";
export const OUTPUT_TOKEN = "tUSDC";

/** 默认值 */
export const DEFAULT_INPUT_AMOUNT = "1";
export const DEFAULT_DEADLINE_MINUTES = 15;

/** mock 请求模拟延迟(ms) */
export const MOCK_API_DELAY_MS = 600;