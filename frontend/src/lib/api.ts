/**
 * 后端接口封装(唯一 HTTP 出口)
 * 对应文档 §12.1 的三个接口:
 *   GET  /quote       ?inputAmount=...(wei)
 *   POST /plans
 *   GET  /plans/:id
 *
 * 对接说明:
 *   - 后端就绪后,在 .env.local 设置 NEXT_PUBLIC_USE_MOCK=0 与 NEXT_PUBLIC_API_BASE_URL
 *   - 所有请求自动携带统一错误处理,错误结构见 ApiError
 *   - 演示数据(mock)均带 isMock 标记,UI 会显示"演示"角标
 */

import {
  API_BASE_URL,
  USE_MOCK,
  DEFAULT_DEADLINE_MINUTES,
  MOCK_API_DELAY_MS,
  OUTPUT_TOKEN,
} from "./config";
import type {
  ApiError,
  CreatePlanRequest,
  CreatePlanResponse,
  OperatorStats,
  Plan,
  Quote,
} from "./types";
import { parseEther, formatEther } from "viem";

/* ---------------- 基础请求工具 ---------------- */

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(`${API_BASE_URL}${path}`, {
    ...init,
    headers: { "Content-Type": "application/json", ...init?.headers },
  });

  if (!res.ok) {
    let error: ApiError = { code: "UNKNOWN", message: `请求失败(${res.status})` };
    try {
      const body = await res.json();
      error = body as ApiError;
    } catch {
      /* 忽略非 JSON 错误体 */
    }
    throw error;
  }

  return (await res.json()) as T;
}

function delay(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/* ---------------- Mock 数据(演示模式) ---------------- */

const MOCK_RATE = 0.95; // 1 MON -> 0.95 tUSDC(模拟汇率)

function mockQuote(inputAmountWei: string): Quote {
  const expected = (BigInt(inputAmountWei) * 95n) / 100n;
  return {
    inputAmount: inputAmountWei,
    outputToken: OUTPUT_TOKEN,
    expectedOutput: expected.toString(),
    simulatedRate: String(MOCK_RATE),
    timestamp: Date.now(),
    protocol: "simple-amm",
    isMock: true,
  };
}

function mockCreatePlan(req: CreatePlanRequest): CreatePlanResponse {
  const now = Date.now();
  const input = BigInt(req.inputAmount);
  const expected = BigInt(req.expectedOutput);
  const guaranteed = (expected * 99n) / 100n; // 保证输出 = 预期 * 0.99
  const maxComp = (input * 10n) / 100n; // 最大赔付 = 输入 * 0.1
  const failComp = (input * 2n) / 100n; // 失败补偿 = 输入 * 0.02
  const planId = `plan-${now.toString(36)}${Math.floor(Math.random() * 1e4).toString(36)}`;
  return {
    planId,
    guaranteedOutput: guaranteed.toString(),
    maxCompensation: maxComp.toString(),
    failureCompensation: failComp.toString(),
    deadline: now + req.deadlineMinutes * 60_000,
    target: "0x0000000000000000000000000000000000000000",
    calldataHash: `0x${"ab".repeat(32)}`,
    txHash: `0x${"12".repeat(32)}`,
    isMock: true,
  };
}

function mockPlan(planId: string): Plan {
  const now = Date.now();
  const input = "1000000000000000000"; // 1 MON
  const expected = "950000000000000000"; // 0.95 tUSDC
  return {
    planId,
    status: "open",
    user: "0x0000000000000000000000000000000000000000",
    inputAmount: input,
    expectedOutput: expected,
    guaranteedOutput: "940500000000000000",
    maxCompensation: "100000000000000000",
    failureCompensation: "20000000000000000",
    deadline: now + 15 * 60_000,
    txHashes: [`0x${"12".repeat(32)}`],
    isMock: true,
  };
}

/** 简易 mock 历史(演示用,真实模式由链上事件重建) */
function mockHistory(): Plan[] {
  const base = mockPlan("demo-ok");
  const short = mockPlan("demo-short");
  const failed = mockPlan("demo-failed");
  return [
    { ...base, planId: "demo-ok", status: "settled_ok", actualOutput: "950000000000000000", bondReleased: true, updatedAt: Date.now() - 3600_000 },
    { ...short, planId: "demo-short", status: "settled_shortfall", actualOutput: "930000000000000000", shortfallPaid: "10500000000000000", bondReleased: true, updatedAt: Date.now() - 7200_000 },
    { ...failed, planId: "demo-failed", status: "failed", refunded: true, compensation: "20000000000000000", updatedAt: Date.now() - 86_400_000 },
  ];
}

/* ---------------- 业务接口 ---------------- */

/** GET /quote?inputAmount=... */
export async function fetchQuote(inputAmountWei: string): Promise<Quote> {
  if (USE_MOCK) {
    await delay(MOCK_API_DELAY_MS);
    return mockQuote(inputAmountWei);
  }
  const params = new URLSearchParams({ inputAmount: inputAmountWei });
  return request<Quote>(`/quote?${params.toString()}`);
}

/** POST /plans */
export async function createPlan(req: CreatePlanRequest): Promise<CreatePlanResponse> {
  if (USE_MOCK) {
    await delay(MOCK_API_DELAY_MS);
    return mockCreatePlan(req);
  }
  return request<CreatePlanResponse>("/plans", {
    method: "POST",
    body: JSON.stringify(req),
  });
}

/** GET /plans/:id */
export async function fetchPlan(planId: string): Promise<Plan> {
  if (USE_MOCK) {
    await delay(MOCK_API_DELAY_MS);
    return mockPlan(planId);
  }
  return request<Plan>(`/plans/${encodeURIComponent(planId)}`);
}

/** 历史页:真实模式由链上事件重建(见 hooks/usePlanEvents),mock 模式返回演示数据 */
export async function fetchHistoryPlans(): Promise<Plan[]> {
  if (USE_MOCK) {
    await delay(MOCK_API_DELAY_MS);
    return mockHistory();
  }
  // 真实模式:前端订阅链上事件后本地重建,不依赖后端列表接口
  return [];
}

/** GET /operators — 运营方列表及声誉数据 */
function mockOperators(): OperatorStats[] {
  return [
    {
      address: "0xB2F400E688a21f79F11c7e5c016eEaE436CA9E4C",
      name: "Bodier",
      totalPlans: 3,
      successPlans: 3,
      shortfallPlans: 0,
      failedPlans: 0,
      totalVolume: "3000000000000000000",
      reputationScore: 90,
      successRate: 100,
      serviceFeeBps: 30,
      guaranteedRatio: 0.90,
      isDefault: false,
    },
    {
      address: "0xDc17fDad88A81Ce3C28F83b0a1706f535723EfbD",
      name: "SwiftSwap",
      totalPlans: 5,
      successPlans: 4,
      shortfallPlans: 1,
      failedPlans: 0,
      totalVolume: "5000000000000000000",
      reputationScore: 82,
      successRate: 80,
      serviceFeeBps: 25,
      guaranteedRatio: 0.92,
      isDefault: false,
    },
    {
      address: "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B",
      name: "SecureTrade",
      totalPlans: 2,
      successPlans: 1,
      shortfallPlans: 0,
      failedPlans: 1,
      totalVolume: "2000000000000000000",
      reputationScore: 55,
      successRate: 50,
      serviceFeeBps: 40,
      guaranteedRatio: 0.95,
      isDefault: false,
    },
  ];
}

export async function fetchOperators(): Promise<OperatorStats[]> {
  if (USE_MOCK) {
    await delay(MOCK_API_DELAY_MS);
    return mockOperators();
  }
  return request<OperatorStats[]>("/operators");
}

/** 金额工具:用于构造请求体 */
export { parseEther, formatEther };

/** 默认计划时长 */
export { DEFAULT_DEADLINE_MINUTES };