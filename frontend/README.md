# Bonded Agent · 前端(web)

Monad Playground 黑客松项目 Bonded Agent 的 P0 前端。
基于 Next.js 16(App Router)+ TypeScript + Tailwind CSS 4 + wagmi/viem + React Query,紫黑科技金融风。

## 快速开始

```bash
npm install
cp .env.example .env.local   # 按需修改
npm run dev                  # http://localhost:3000
```

## 目录结构

```text
src/
├─ app/
│  ├─ layout.tsx        # 根布局(Header/Footer/Providers)
│  ├─ page.tsx          # 首页:连接钱包 + 请求表单 + 报价预览
│  ├─ plans/[id]/page.tsx  # 计划页:承诺卡片 + 执行 + 结算
│  └─ history/page.tsx  # 历史页:事件重建列表 + 筛选
├─ components/
│  ├─ layout/           # Header、Footer、NetworkBadge、WalletButton
│  ├─ request/          # RequestForm、QuoteSummary
│  ├─ plan/             # PromiseCard、PlanStatusBadge、SettlementExplanation
│  └─ ui/               # Button、ErrorBanner、Toast、AmountToken、Skeleton、EmptyState、TxLink、StatusDot
├─ hooks/               # useQuote、useCreatePlan、usePlan、usePlanEvents、useExecutePlan
└─ lib/
   ├─ config.ts         # 环境变量与链配置
   ├─ types.ts          # 前后端契约类型(金额统一 wei 字符串)
   ├─ api.ts            # 后端 HTTP 出口(唯一请求封装,含演示数据模式)
   ├─ format.ts         # 金额/时间/地址格式化
   ├─ plan-status.ts    # 计划状态机派生
   ├─ copy.ts           # 全部用户可见文案(合规集中管理)
   ├─ wagmi.ts          # wagmi 配置(Monad Testnet + injected)
   ├─ contracts.ts      # 合约地址/ABI(待部署后配置)
   └─ events.ts         # 链上事件常量与归一化
```

## 后端对接

所有后端请求集中在 `src/lib/api.ts`,对应设计文档 §12.1 的三个接口:

- `GET /quote?inputAmount=<wei>`
- `POST /plans`(body:`{ inputAmount, expectedOutput, deadlineMinutes }`)
- `GET /plans/:id`

对接步骤:

1. 在 `.env.local` 设置 `NEXT_PUBLIC_API_BASE_URL=<后端地址>`。
2. 后端就绪后设置 `NEXT_PUBLIC_USE_MOCK=0`(演示模式默认开启,便于无后端体验)。
3. 合约部署后设置 `NEXT_PUBLIC_PLAN_CONTRACT=<地址>` 并在 `src/lib/contracts.ts` 补齐 ABI。

金额字段统一使用 wei 字符串(18 位小数),前端展示层负责格式化,避免精度丢失。

## 环境变量

见 `.env.example`。默认 Monad Testnet(chainId 10143),可在 `.env.local` 覆盖。

## 设计文档

界面规范与页面规格见《Bonded-Agent-frontend-design.md》。
