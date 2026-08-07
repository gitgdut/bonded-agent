/**
 * Bonded Agent 前后端契约类型
 * 对应《Bonded-Agent-frontend-design.md》§12.1
 * 金额字段统一使用字符串(wei,18 位小数),避免精度丢失
 */

/** 计划状态机(文档 §8.1) */
export type PlanStatus =
  | "loading" // 链上数据读取中
  | "open" // PlanOpened 已发出,未过期,无执行事件
  | "expired" // 当前区块时间 > deadline 且无执行事件
  | "executing" // 用户已提交交易,等待回执
  | "settled_ok" // PlanExecuted 且无 ShortfallPaid
  | "settled_shortfall" // PlanExecuted 且有 ShortfallPaid
  | "failed" // PlanFailed
  | "rejected"; // 本地状态:用户主动放弃(链上无记录)

/** GET /quote 响应 */
export interface Quote {
  inputAmount: string; // wei,如 "1000000000000000000"
  outputToken: string; // 目标代币,如 "tUSDC"
  expectedOutput: string; // wei,预期输出
  simulatedRate: string; // 模拟汇率(每 1 个输入代币可换得的输出)
  timestamp: number; // 报价生成时间(ms)
  protocol: string; // 报价来源协议 (Moss-style discover pipeline)
  isMock?: boolean; // 演示数据标记(仅 mock 模式)
  operator?: OperatorStats; // 选中的运营方信息 (多运营方支持)
}

/** POST /plans 请求体 */
export interface CreatePlanRequest {
  inputAmount: string; // wei
  expectedOutput: string; // wei
  deadlineMinutes: number;
  userAddress: string; // 用户钱包地址，由后端写入 plan.user
}

/** POST /plans 响应 */
export interface CreatePlanResponse {
  planId: string;
  guaranteedOutput: string; // wei,保证输出
  maxCompensation: string; // wei,最大赔付
  failureCompensation: string; // wei,失败补偿
  deadline: number; // 截止时间戳(ms)
  target: string; // 目标合约地址
  calldataHash: string;
  txHash: string; // 开计划交易哈希
  isMock?: boolean;
}

/** GET /plans/:id 响应(计划详情) */
export interface Plan {
  planId: string;
  status: PlanStatus;
  user: string; // 计划所属钱包地址
  inputAmount: string; // wei
  expectedOutput: string; // wei,预期输出
  guaranteedOutput: string; // wei,保证输出
  maxCompensation: string; // wei,最大赔付
  failureCompensation: string; // wei,失败补偿
  deadline: number; // 截止时间戳(ms)
  txHashes: string[]; // 相关交易哈希列表
  actualOutput?: string; // wei,结算后实际到账
  shortfallPaid?: string; // wei,自动补足差额
  compensation?: string; // wei,失败补偿实付
  refunded?: boolean; // 失败时是否退还本金
  bondReleased?: boolean; // 保证金是否已释放
  updatedAt?: number; // 最近更新时间(ms)
  isMock?: boolean; // 演示数据标记
}

/** 链上事件(历史页/计划页重建用) */
export interface PlanEvent {
  type:
    | "PlanOpened"
    | "PlanExecuted"
    | "ShortfallPaid"
    | "PlanFailed"
    | "BondReleased";
  planId: string;
  user?: string;
  txHash: string;
  blockNumber: bigint;
  data: Record<string, unknown>;
}

/** GET /operators 响应 — 运营方声誉数据 */
export interface OperatorStats {
  address: string;
  name?: string;
  totalPlans: number;
  successPlans: number;
  shortfallPlans: number;
  failedPlans: number;
  totalVolume: string;
  reputationScore: number;
  successRate: number;
  serviceFeeBps: number;
  guaranteedRatio: number;
  isDefault?: boolean;
}

/** 后端错误结构(约定) */
export interface ApiError {
  code: string;
  message: string;
}