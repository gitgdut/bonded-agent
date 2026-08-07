/**
 * 全部用户可见文案(合规集中管理,文档 §9)
 * 实现期禁止在组件内散写文案
 */

import type { Plan } from "./types";
import { formatAmount, formatDateTime, formatCountdown } from "./format";
import { INPUT_TOKEN, OUTPUT_TOKEN } from "./config";

/* ---------------- 全局 ---------------- */

export const COPY = {
  appName: "Bonded Agent",
  tagline: "Agent 运营方为交易结果锁定保证金,结果不足自动赔付",
  footer: [
    "运行于 Monad Testnet",
    "Moss-style Protocol Pipeline: discover→load→action→simulate",
    "协议: simple-amm (恒定积 AMM)  |  可扩展至 PancakeSwap / Kuru",
    "测试资产,无真实价值",
  ] as const,
  connectWallet: "连接钱包",
  demoBadge: "演示数据",
  demoTxHint: "演示数据 · 链上无真实交易",
} as const;

/* ---------------- 顶部导航 ---------------- */

export const NAV = {
  home: "首页",
  history: "历史",
} as const;

/* ---------------- 首页 ---------------- */

export const HOME = {
  title: "有担保的 Swap",
  inputLabel: `输入金额(${INPUT_TOKEN})`,
  outputLabel: `目标代币`,
  getQuote: "获取担保报价",
  generating: "获取报价中…",
  createPlan: "生成担保计划",
  creatingPlan: "生成中…",
  walletPrompt: "请先连接钱包后再获取报价",
  simulatedRateTag: "模拟汇率",
} as const;

/* ---------------- 报价预览 ---------------- */

export const QUOTE = {
  expectedOutput: "预期输出",
  guaranteedOutput: "保证输出",
  bondLocked: "运营方锁定保证金",
  validity: "有效期预览",
  rate: "模拟汇率",
  serviceFee: "服务费 (0.3%)",
  minutes: (n: number) => `约 ${n} 分钟`,
} as const;

/* ---------------- 承诺卡片(文档 §7) ---------------- */

export const PLAN = {
  title: "担保交易计划",
  summary: "交易摘要",
  guarantees: "保证条款",
  riskBoundary: "风险边界",
  fieldGuaranteedOutput: "保证输出",
  fieldBond: "运营方已锁定保证金",
  fieldMaxCompensation: "最大赔付",
  fieldFailureCompensation: "失败补偿",
  fieldDeadline: "有效期截止",
  execute: "同意并执行 Swap",
  executing: "等待交易确认…",
  expired: "计划已过期",
  settled: "本计划已结算",
  wrongWallet: "请连接计划所属钱包",
  reRequest: "返回首页重新发起请求",
  notFound: "未找到该计划",
  riskItems: [
    "本计划只保证这一笔 Swap 的最低到账,不保证盈利。",
    "仅当实际到账低于保证值时自动补足,赔付上限为锁定的保证金。",
    "当前市场为模拟(MockDex),汇率由团队配置,非真实 DEX;真实市场接入为后续版本。",
    "测试网与测试资产,无真实价值。",
  ] as const,
  // 结算解释模板(文档 §8.3)
  explain: {
    settled_ok: (p: Plan) =>
      `执行成功,到账 ${formatAmount(p.actualOutput ?? p.guaranteedOutput)} ${OUTPUT_TOKEN},达到保证值;运营方保证金已释放。`,
    settled_shortfall: (p: Plan) =>
      `实际到账 ${formatAmount(p.actualOutput ?? "0")} ${OUTPUT_TOKEN},低于保证值,已自动补足差额 ${formatAmount(p.shortfallPaid ?? "0")} ${OUTPUT_TOKEN}。`,
    failed: (p: Plan) =>
      `交易失败,已退还 ${formatAmount(p.inputAmount)} ${INPUT_TOKEN}${p.compensation ? `,并支付失败补偿 ${formatAmount(p.compensation)} ${OUTPUT_TOKEN}` : ""}。`,
    expired: () => "计划已过期,未执行,无任何资金变动。",
  } as const,
  viewOnExplorer: "在区块浏览器中查看",
} as const;

/** 计划页头部描述 */
export function planHeading(planId: string, deadline: number) {
  return {
    title: `担保交易计划 #${planId}`,
    deadlineText: `有效期截止 ${formatDateTime(deadline)}`,
    countdown: (now: number) => formatCountdown(deadline, now),
  };
}

/* ---------------- 历史页 ---------------- */

export const HISTORY = {
  title: "我的担保计划",
  filters: ["全部", "待执行", "履约", "赔付", "失败", "过期"] as const,
  empty: "还没有担保计划,去发起一笔 Swap",
  goRequest: "发起担保 Swap",
} as const;

/* ---------------- 运营方选择器 ---------------- */

export const OPERATOR = {
  pickerTitle: "选择运营方",
  loading: "加载运营方列表中…",
  noOperators: "暂无可用运营方",
  score: "声誉评分",
  successRate: "成功率",
  fee: "费率",
  guaranteeRatio: "保证比率",
  totalPlans: "交易数",
} as const;

/* ---------------- 错误映射(文档 §12.4) ---------------- */

export const ERROR_COPY: Record<string, string> = {
  PlanExpired: "计划已过期,请重新生成",
  AlreadyExecuted: "该计划已执行过",
  InvalidCalldataHash: "调用数据不匹配,请重新获取报价",
  NotPlanUser: "该计划不属于当前钱包,请切换钱包",
  InsufficientBond: "运营方保证金不足,请稍后重试",
  UserRejected: "已取消,无资金变动",
  NETWORK: "网络异常,请重试",
  NOT_FOUND: "未找到该计划",
};

export function errorMessage(err: unknown): string {
  if (err && typeof err === "object" && "code" in err) {
    const code = (err as { code: string }).code;
    if (ERROR_COPY[code]) return ERROR_COPY[code];
  }
  if (err instanceof Error && err.message) return err.message;
  return "发生未知错误,请重试";
}