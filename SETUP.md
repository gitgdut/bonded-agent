# Bonded Agent 启动指南

> 写给队友：如何拉取最新代码并在本地跑起来

---

## 1. 环境要求

| 依赖 | 版本 | 检查命令 |
|------|------|----------|
| **Go** | ≥ 1.21（go.mod 写 1.26.2，1.21+ 即可） | `go version` |
| **Node.js** | ≥ 18（推荐 22 LTS） | `node -v` |
| **pnpm** / npm | 任意（项目用 npm） | `npm -v` |
| **Foundry** | 最新 stable（仅合约编译需要） | `forge --version` |

> Foundry 安装：`curl -L https://foundry.paradigm.xyz | bash` 然后 `foundryup`

---

## 2. 拉取代码

```bash
git clone git@github.com:gitgdut/bonded-agent.git
cd bonded-agent
git checkout feat/http-api
```

确认在正确的分支和 commit：

```bash
git log --oneline -1
# 应该看到: ee50f2e feat: BondedExecutor V2 — 合约审计修复 + coverageFloor + E2E
```

---

## 3. 项目结构

```
bonded-agent/
├── agent/                  # Go 后端（API + Agent 逻辑）
│   ├── main.go
│   ├── internal/           # 核心逻辑
│   ├── contracts/          # abigen 生成的 Go 绑定
│   ├── .env                # 后端环境变量（需自行创建）
│   └── go.mod
├── frontend/               # Next.js 前端
│   ├── src/
│   │   ├── app/            # App Router 页面
│   │   ├── components/     # React 组件
│   │   ├── hooks/          # 自定义 hooks
│   │   └── lib/            # 配置、类型、工具
│   ├── .env.local          # 前端环境变量（需自行创建）
│   └── package.json
├── contracts/              # Solidity 合约（Foundry）
│   ├── src/BondedExecutor.sol
│   ├── script/DeployBondedExecutor.s.sol
│   └── foundry.toml
└── README.md
```

---

## 4. 配置环境变量

### 4.1 后端：`agent/.env`

在 `agent/` 目录下创建 `.env` 文件：

```bash
# Monad Testnet RPC
MONAD_RPC=https://testnet-rpc.monad.xyz

# 合约地址（Monad Testnet，已部署）
USDC_ADDR=0xD5B1b42929188280631ef2502c78AA61e1A56e0a
DEX_ADDR=0x3F2F80002024C53D444537884708570706bE68c5
EXECUTOR_ADDR=0x195da9B8600717f758F1f4300b9A812CB2CAb768

# 担保参数
GUARANTEE_RATIO=0.90
MAX_COMPENSATION=20000000000000000000
FAILURE_COMPENSATION=5000000000000000000

# ERC-8004 身份合约（Monad Testnet）
ERC8004_IDENTITY_ADDR=0xe18Fd55C1935554499b0E8c74Db967ccFD46F3c1
ERC8004_REPUTATION_ADDR=0x8A2d492418Ba6b765fE719321ADF2e6f56d9AeA7

# Operator 私钥（需替换为你自己的测试网私钥）
OPERATOR_PRIVATE_KEY=<你的测试网私钥>
```

> ⚠️ **私钥安全**：`.env` 已在 `.gitignore` 中，不会被提交。私钥仅用于测试网，不要使用主网私钥。

### 4.2 前端：`frontend/.env.local`

在 `frontend/` 目录下创建 `.env.local` 文件：

```bash
# Monad Testnet
NEXT_PUBLIC_CHAIN_ID=10143
NEXT_PUBLIC_CHAIN_NAME=Monad Testnet
NEXT_PUBLIC_RPC_URL=https://testnet-rpc.monad.xyz
NEXT_PUBLIC_EXPLORER_URL=https://testnet.monadscan.com

# BondedExecutor 合约地址（Monad Testnet）
NEXT_PUBLIC_PLAN_CONTRACT=0x195da9B8600717f758F1f4300b9A812CB2CAb768

# Go Agent HTTP API — 通过 Next.js rewrite 代理，避免跨域
NEXT_PUBLIC_API_BASE_URL=/api
NEXT_PUBLIC_USE_MOCK=0
```

### 4.3 合约地址速查表

| 合约 | 地址（Monad Testnet） |
|------|----------------------|
| tUSDC（MockUSDC） | `0xD5B1b42929188280631ef2502c78AA61e1A56e0a` |
| DEX（SimpleAMMPair） | `0x3F2F80002024C53D444537884708570706bE68c5` |
| BondedExecutor V2 | `0x195da9B8600717f758F1f4300b9A812CB2CAb768` |
| ERC-8004 Identity | `0xe18Fd55C1935554499b0E8c74Db967ccFD46F3c1` |
| ERC-8004 Reputation | `0x8A2d492418Ba6b765fE719321ADF2e6f56d9AeA7` |

---

## 5. 安装依赖 & 编译

### 5.1 Go 后端

```bash
cd agent
go mod download
go build -o agent .
```

编译成功会在 `agent/` 目录下生成 `agent` 可执行文件。

### 5.2 前端

```bash
cd frontend
npm install
```

### 5.3 合约（可选，仅当你需要修改/重新部署合约时）

```bash
cd contracts
forge build --via-ir
```

---

## 6. 启动服务

需要**两个终端窗口**：

### 终端 1：Go 后端（端口 8787）

```bash
cd agent
./agent serve
```

看到以下输出表示启动成功：

```
✓ Connected to chain 10143 as operator 0x...
API listening on :8787
```

### 终端 2：Next.js 前端（端口 3000）

```bash
cd frontend
npx next dev --webpack -p 3000
```

> ⚠️ 必须加 `--webpack` 参数，当前项目未迁移到 Turbopack。

看到 `Ready in ...s` 表示启动成功。

### 访问

浏览器打开 **http://localhost:3000**

---

## 7. 功能验证

### 7.1 后端 API 检查

```bash
# 健康检查（如果实现了的话）
curl http://localhost:8787/operators

# 报价测试（替换数值测试）
curl -X POST http://localhost:8787/quote \
  -H "Content-Type: application/json" \
  -d '{"inputAmount":"1000000000000000000","inputToken":"0x0000000000000000000000000000000000000000","outputToken":"0xD5B1b42929188280631ef2502c78AA61e1A56e0a","operator":"0xdc1746793D71256Be16B70188297ab5aB056208F"}'
```

### 7.2 前端流程验证

1. 打开 http://localhost:3000
2. 连接钱包（MetaMask / Rabby），切换到 **Monad Testnet**（Chain ID 10143）
3. 输入金额 → 点击"获取报价"
4. 查看报价预览（预期输出、服务费、保证输出、保证金）
5. 点击"创建担保计划" → 签名交易
6. 等待 Plan 创建确认
7. 点击"执行" → 签名（EIP-712 离线签名 或 直接 executePlan）

---

## 8. 常见问题

### Q: `go build` 报依赖错误？

```bash
cd agent
go mod tidy
go build -o agent .
```

### Q: 前端启动报 "Couldn't find any pages or app directory"？

确认你在 `frontend/` 目录下执行的命令：
```bash
cd /path/to/bonded-agent/frontend
npx next dev --webpack -p 3000
```

### Q: 前端请求 API 报 404？

确认 Go 后端已启动且监听 `:8787`。Next.js 的 `rewrites` 会将 `/api/*` 代理到 `localhost:8787/*`。

### Q: 交易一直 pending？

检查：
1. Monad Testnet RPC 是否可用：`curl https://testnet-rpc.monad.xyz -X POST -H "Content-Type: application/json" -d '{"jsonrpc":"2.0","method":"eth_chainId","params":[],"id":1}'`
2. 钱包是否切到 Monad Testnet（Chain ID 10143）
3. 账户是否有足够的 MON gas 费

### Q: executePlanWithSignature 报 InvalidSignature？

EIP-712 domain version 已固定为 `"2"`。不要将前端 `useExecutePlan.ts` 中的 `version: "2"` 改为其他值。

### Q: 我想用自己的测试网私钥做 Operator？

修改 `agent/.env` 中的 `OPERATOR_PRIVATE_KEY`，重启后端即可。不需要重新部署合约——BondedExecutor 支持多个 Operator 共用同一个合约。

---

## 9. 分支策略

```
master       ← 稳定可部署
develop      ← 开发集成
feat/xxx     ← 新功能分支（从 develop 切，PR 回 develop）
feat/http-api ← 当前开发分支（HTTP API + 合约审计修复）
fix/xxx      ← 修 bug（从 master 切，cherry-pick 回 develop）
```

日常开发不要直接推 master。PR 至少一人 review 后合并。

---

## 10. 团队联系方式

- **Dev1（合约 + 后端）**：合约问题、API 问题、部署问题
- **Dev2（前端）**：UI 问题、钱包连接问题、报价展示问题

有问题直接群里喊 😄
