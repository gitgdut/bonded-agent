# Bonded Agent

> 🏗️ **Monad Playground Hackathon** | Team Backchannel

---

## Quick Start / 快速开始

### Prerequisites / 环境要求

| Dependency | Version | Check |
|------------|---------|-------|
| Go | ≥ 1.21 | `go version` |
| Node.js | ≥ 18 | `node -v` |
| Foundry | latest | `forge --version` |

### Clone & Checkout / 拉取代码

```bash
git clone git@github.com:zzxwjm/backchannel-monad-playground.git
cd backchannel-monad-playground
git checkout feat/http-api
```

### Configure / 配置

**agent/.env**（backend / 后端）：

```bash
MONAD_RPC=https://testnet-rpc.monad.xyz
USDC_ADDR=0xD5B1b42929188280631ef2502c78AA61e1A56e0a
DEX_ADDR=0x8FA4112739F1fAd68EEB82f75360800354992803
EXECUTOR_ADDR=0x195da9B8600717f758F1f4300b9A812CB2CAb768
GUARANTEE_RATIO=0.90
MAX_COMPENSATION=20000000000000000000
FAILURE_COMPENSATION=5000000000000000000
ERC8004_IDENTITY_ADDR=0xe18Fd55C1935554499b0E8c74Db967ccFD46F3c1
ERC8004_REPUTATION_ADDR=0x8A2d492418Ba6b765fE719321ADF2e6f56d9AeA7
OPERATOR_PRIVATE_KEY=<your testnet key>
```

**frontend/.env.local**（frontend / 前端）：

```bash
NEXT_PUBLIC_CHAIN_ID=10143
NEXT_PUBLIC_CHAIN_NAME=Monad Testnet
NEXT_PUBLIC_RPC_URL=https://testnet-rpc.monad.xyz
NEXT_PUBLIC_EXPLORER_URL=https://testnet.monadscan.com
NEXT_PUBLIC_PLAN_CONTRACT=0x195da9B8600717f758F1f4300b9A812CB2CAb768
NEXT_PUBLIC_API_BASE_URL=/api
NEXT_PUBLIC_USE_MOCK=0
```

### Build / 编译

```bash
# backend / 后端
cd agent && go mod download && go build -o agent .

# frontend / 前端
cd frontend && npm install
```

### Run / 启动

Two terminals / 两个终端：

```bash
# Terminal 1: API server (port 8787)
cd agent && ./agent serve

# Terminal 2: Frontend (port 3000)
cd frontend && npx next dev --webpack -p 3000
```

Open / 打开 **http://localhost:3000**

> Full troubleshooting guide / 完整排障指南：[SETUP.md](SETUP.md)

---

## What & Why / 是什么 & 为什么

### English

**Bonded Agent** adds financial accountability to AI Agent transactions. When an Agent proposes an on-chain swap (e.g., "swap 1 MON for tUSDC"), the Agent operator locks collateral (tUSDC) as a guarantee. If the actual execution result falls below the promised minimum, the smart contract **automatically compensates the user** from the locked bond — no admin, no LLM judgment, no manual dispute resolution.

**The Problem**: Simulation shows what *might* happen under current conditions — it is not a promise. By the time a user signs, on-chain state and prices may have shifted. Today, if an Agent's prediction doesn't hold, the operator bears zero financial responsibility.

### 中文

**Bonded Agent** 为 AI Agent 的链上交易引入金融责任制。当 Agent 提议一笔 Swap（如"把 1 MON 换成 tUSDC"），Agent 运营方需锁定 tUSDC 作为保证金。如果实际执行结果低于承诺的最低标准，智能合约会**自动从保证金中赔付用户**——无需管理员、无需 LLM 判断、无需人工仲裁。

**核心问题**：模拟只能说明交易在某个时刻*可能*产生什么结果，它不是保证。用户签名前，链上状态和价格可能已经变化。目前，即使 Agent 的预测没有兑现，运营方也几乎不承担任何经济责任。

---

## How It Works / 工作流程

```
User / 用户: "Swap 1 MON → tUSDC"
          ↓
Agent simulates, posts a Plan with guaranteed output + locked bond
Agent 模拟交易，发布包含保证输出 + 锁定保证金的 Plan
          ↓
User reviews the guarantee, signs
用户查看保证条款，签名
          ↓
BondedExecutor executes swap, measures actual output
BondedExecutor 执行 Swap，测量实际输出
          ↓
  actual ≥ guaranteed? → Bond released back to operator / 保证金退还运营方
  actual < guaranteed? → Shortfall auto-paid to user from bond / 差额自动赔付
  swap fails?          → User refunded + failure compensation / 退还用户 + 失败补偿
```

### The Guarantee Math / 保证输出怎么算

```
expectedOutput   → 预期输出（链上模拟结果）
  - fee (0.30%)  → 扣除服务费
  = netAfterFee  → 净到账
  × 90%          → 乘以保证比率
  = guaranteedOutput → 保证输出（最低到账）
  - maxComp      → 扣除最大赔付
  = coverageFloor → 执行底线（低于此值不执行）
```

> **Key insight / 关键设计**：guaranteedOutput 低于 expectedOutput 约 10%，这个缓冲吸收了滑点和 MEV 波动，让运营方不会因为正常的市场波动而频繁被罚没。

### Three Settlement Paths / 三种结算路径

| Path / 路径 | Trigger / 触发条件 | Outcome / 结果 |
|-------------|-------------------|----------------|
| **Success** | actualOutput ≥ guaranteedOutput | Bond fully released to operator / 保证金全额退还 |
| **Shortfall** | coverageFloor ≤ actualOutput < guaranteedOutput | User paid the difference from bond; remainder returned to operator / 差额从保证金赔付用户，剩余退还运营方 |
| **Failed** | actualOutput < coverageFloor **or** swap reverted | User's input refunded + failure compensation (5 tUSDC); operator's MON refunded via pull / 退还用户输入 + 失败补偿；运营方 MON 通过 pull 取回 |

---

## Using the App / 操作指南

### User Flow / 用户流程

1. **Connect wallet** / 连接钱包 → switch to Monad Testnet (Chain 10143)
2. **Enter amount** / 输入金额 → click "获取报价" (Get Quote)
3. **Review quote** / 查看报价：
   - Expected output / 预期输出
   - Service fee (0.30%) / 服务费
   - **Guaranteed output** (highlighted) / **保证输出**（高亮）
   - Risk uncovered / 风险敞口
   - Bond locked (20 tUSDC) / 锁定保证金
   - Validity (15 min) / 有效期
4. **Create Plan** / 创建担保计划 → sign transaction
5. **Execute** / 执行：
   - **On-chain** (executePlan) — user pays gas, submits tx directly / 用户付 gas，直接提交
   - **Off-chain signature** (executePlanWithSignature) — user signs EIP-712 for free, operator submits tx / 用户免费签名，运营方代提交

### Verification / 验证

```bash
# Check backend operators / 检查后端运营方列表
curl http://localhost:8787/operators

# Get a quote / 获取报价
curl -X POST http://localhost:8787/quote \
  -H "Content-Type: application/json" \
  -d '{"inputAmount":"1000000000000000000","inputToken":"0x0000000000000000000000000000000000000000","outputToken":"0xD5B1b42929188280631ef2502c78AA61e1A56e0a","operator":"0xdc1746793D71256Be16B70188297ab5aB056208F"}'
```

---

## Architecture / 技术架构

```
┌──────────────┐     ┌──────────────┐     ┌──────────────────┐
│   Frontend   │────▶│  Go Agent    │────▶│  BondedExecutor  │
│  Next.js 16  │     │  :8787       │     │  (Monad Testnet) │
│  wagmi/viem  │     │  go-ethereum │     │  Solidity        │
└──────────────┘     └──────────────┘     └──────────────────┘
       │                     │                      │
       │  /api/* rewrite     │  RPC calls           │  on-chain events
       │  → :8787            │  contract calls       │  balanceOf checks
       ▼                     ▼                      ▼
  Wallet (MetaMask)    Monad Testnet RPC        Smart Contract State
```

### Project Structure / 项目结构

```
bonded-agent/
├── agent/                    # Go backend / 后端
│   ├── main.go               # Entry point / 入口
│   ├── internal/
│   │   ├── agent.go          # Core agent logic / 核心 Agent 逻辑
│   │   ├── server.go         # HTTP API handlers
│   │   └── reputation.go     # On-chain reputation engine / 链上声誉引擎
│   ├── contracts/            # go-ethereum bindings (abigen) / Go 合约绑定
│   └── .env                  # Backend config / 后端配置
├── frontend/                 # Next.js frontend / 前端
│   ├── src/
│   │   ├── app/              # App Router pages / 页面
│   │   ├── components/       # React components / 组件
│   │   │   ├── request/      # Quote + create plan flow / 报价 + 创建流程
│   │   │   ├── operator/     # Operator picker + card / 运营方选择器
│   │   │   └── ui/           # Shared UI primitives / 通用 UI 组件
│   │   ├── hooks/            # wagmi + React Query hooks
│   │   └── lib/              # Config, types, formatters / 配置、类型、格式化
│   └── .env.local            # Frontend config / 前端配置
└── contracts/                # Solidity (Foundry) / 智能合约
    ├── src/
    │   ├── BondedExecutor.sol  # Core contract / 核心合约
    │   ├── MockUSDC.sol        # Test token / 测试代币
    │   └── SimpleAMMPair.sol   # Mock DEX / 模拟 DEX
    ├── script/
    │   └── DeployBondedExecutor.s.sol
    └── foundry.toml
```

### Tech Stack / 技术栈

| Layer / 层 | Technology / 技术 | Notes / 备注 |
|------------|-------------------|-------------|
| Smart Contracts | Solidity (Foundry) | `--via-ir` pipeline, 0.8.28 |
| Agent Backend | Go + go-ethereum | abigen bindings, EIP-712 signing |
| Frontend | React 19 + Next.js 16.3 | webpack, wagmi v3, viem v2 |
| Network | Monad Testnet | Chain ID 10143 |

---

## Contracts / 合约详情

### Deployed Addresses / 部署地址 (Monad Testnet)

| Contract / 合约 | Address / 地址 |
|-----------------|---------------|
| tUSDC (MockUSDC) | [`0xD5B1b42929188280631ef2502c78AA61e1A56e0a`](https://testnet.monadscan.com/address/0xD5B1b42929188280631ef2502c78AA61e1A56e0a) |
| DEX (MockDex) | [`0x8FA4112739F1fAd68EEB82f75360800354992803`](https://testnet.monadscan.com/address/0x8FA4112739F1fAd68EEB82f75360800354992803) |
| BondedExecutor V2 | [`0x195da9B8600717f758F1f4300b9A812CB2CAb768`](https://testnet.monadscan.com/address/0x195da9B8600717f758F1f4300b9A812CB2CAb768) |
| ERC-8004 Identity | [`0xe18Fd55C1935554499b0E8c74Db967ccFD46F3c1`](https://testnet.monadscan.com/address/0xe18Fd55C1935554499b0E8c74Db967ccFD46F3c1) |
| ERC-8004 Reputation | [`0x8A2d492418Ba6b765fE719321ADF2e6f56d9AeA7`](https://testnet.monadscan.com/address/0x8A2d492418Ba6b765fE719321ADF2e6f56d9AeA7) |

### Key Contract Design / 核心合约设计

**BondedExecutor.sol** — the core guarantee & settlement contract / 核心担保结算合约：

- **Plan struct** — 12 fields including operator, nonce, coverageFloor, bondDeposit, failureCompensation / 12 个字段：operator、nonce、coverageFloor、bondDeposit、failureCompensation
- **planId** = `keccak256(operator, nonce)` — per-operator nonce prevents replay / 每个运营方独立 nonce，防重放
- **bondDeposit** = `max(maxCompensation, failureCompensation)` — deposited in tUSDC at plan creation / 创建 Plan 时以 tUSDC 存入
- **Try/catch** execution — inner swap call measured against coverageFloor, outer routes to success/shortfall/failure / try/catch 原子执行：内层 swap 调用检测输出，外层路由至成功/差额/失败
- **Pull-mode MON refunds** — failed plans refund MON to `pendingRefunds[operator]` rather than pushing (safe against reentrancy) / 失败时 MON 写入 `pendingRefunds` 映射，运营方主动取回（防重入）
- **Malicious calldata guard** — post-exec `balanceOf(this) >= totalLockedBonds` invariant check / 执行后检查 `balanceOf(this) >= totalLockedBonds`，防止恶意 calldata 窃取资金
- **EIP-712 domain version `"2"`** — offline signature authorization for gasless user execution / 离线签名授权，用户无需付 gas

### P0 Scope / P0 范围

| Component / 组件 | Description / 说明 |
|-----------|-------------|
| Swap pair / 交易对 | MON → tUSDC only |
| Operator / 运营方 | Team-controlled (multi-operator ready) |
| Network / 网络 | Monad Testnet (Chain ID: 10143) |
| Contracts / 合约 | MockUSDC, SimpleAMMPair, BondedExecutor |
| Guarantee / 赔付 | Purely on-chain, token balance delta / 纯链上，代币余额变化 |

### What We Do NOT Do / 明确不做

- ❌ Natural-language promises / 通用自然语言承诺
- ❌ LLM or human judgment for disputes / LLM 或人工仲裁
- ❌ Insurance pools, cross-chain, governance, token issuance / 保险池、跨链、治理、发币
- ❌ Merge with AnnoPilot / 与 AnnoPilot 合并

---

## Team & Workflow / 团队 & 协作

### Team / 团队

**Backchannel** — Monad Playground Hackathon

| Role / 角色 | Responsibility / 职责 |
|-------------|----------------------|
| Researcher / PM | Product scope, testing, pitch / 产品范围、测试、路演 |
| Marketing / Ops | Positioning, copy, demos / 定位、文案、演示 |
| Dev: Contracts + Backend | Solidity (Foundry), Go agent service |
| Dev: Frontend | React / Next.js, wallet integration, UI |
| Design | Information hierarchy, promise card, tx status |

### Branch Strategy / 分支策略

```
master       ← stable, deployable / 稳定可部署
develop      ← integration / 开发集成
feat/xxx     ← feature branch (cut from develop, PR back)
feat/http-api ← current / 当前分支
fix/xxx      ← hotfix (cut from master, cherry-pick to develop)
```

> 日常开发不要直接推 master。PR 至少一人 review 后合并。
> Never push directly to master. At least one review before merge.

### FAQ / 常见问题

| Problem / 问题 | Solution / 解决 |
|---------------|----------------|
| Frontend can't find pages | Run from `frontend/` directory / 在 frontend/ 目录下执行 |
| API returns 404 | Ensure Go backend is on `:8787` / 确认 Go 后端在 8787 端口 |
| Tx pending forever | Check RPC, chain ID, gas balance / 检查 RPC、链 ID、gas 余额 |
| InvalidSignature error | EIP-712 version must be `"2"`, don't change it / 不要改 version |
| `go build` errors | Run `go mod tidy` first / 先跑 `go mod tidy` |
