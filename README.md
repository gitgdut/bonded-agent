# Bonded Agent

> 🏗️ **Monad Playground Hackathon** | Team Backchannel

[English](#english) | [中文](#中文)

---

## English

### What is Bonded Agent?

**Bonded Agent** adds financial accountability to AI Agent transactions. When an Agent proposes an on-chain swap (e.g., "swap 1 MON for tUSDC"), the Agent operator locks collateral (tUSDC) as a guarantee. If the actual execution result falls below the promised minimum, the smart contract **automatically compensates the user** from the locked bond — no admin, no LLM judgment, no manual dispute resolution.

### The Problem

Simulation shows what *might* happen under current conditions — it is not a promise. By the time a user signs, on-chain state and prices may have shifted. Today, if an Agent's prediction doesn't hold, the operator bears zero financial responsibility.

### How It Works

```
User: "Swap 1 MON → tUSDC"
          ↓
Agent simulates, posts a Plan with guaranteed output + locked bond
          ↓
User reviews the guarantee, signs
          ↓
BondedExecutor executes swap, measures actual output
          ↓
  actual ≥ guaranteed? → Bond released back to operator
  actual < guaranteed? → Shortfall auto-paid to user from bond
  swap fails?          → User refunded + failure compensation
```

### P0 Scope (Hackathon Deliverable)

| Component | Description |
|-----------|-------------|
| **Swap pair** | MON → tUSDC only |
| **Operator** | Single team-controlled operator |
| **Network** | Monad Testnet (Chain ID: 10143) |
| **Contracts** | MockUSDC, MockDex, BondedExecutor |
| **Guarantee** | Purely on-chain, measurable by token balance delta |

### Contracts

- **MockUSDC.sol** — Test output token & bond asset (deployer-mintable)
- **MockDex.sol** — Simulated DEX with configurable rate and `minOutput` support
- **BondedExecutor.sol** — Core: plan creation, bond locking, swap execution, balance measurement, auto-compensation

### What We Explicitly Do NOT Do

- ❌ General natural-language promises
- ❌ LLM or human judgment for dispute resolution
- ❌ Insurance pools, cross-chain, governance, or token issuance
- ❌ Merge with AnnoPilot during this hackathon

### Tech Stack

| Layer | Technology |
|-------|------------|
| Smart Contracts | Solidity (Foundry) |
| Agent Backend | Go + go-ethereum |
| Frontend | React / Next.js |

---

## 中文

### Bonded Agent 是什么？

**Bonded Agent** 为 AI Agent 的链上交易引入金融责任制。当 Agent 提议一笔 Swap（如"把 1 MON 换成 tUSDC"），Agent 运营方需锁定 tUSDC 作为保证金。如果实际执行结果低于承诺的最低标准，智能合约会**自动从保证金中赔付用户**——无需管理员、无需 LLM 判断、无需人工仲裁。

### 要解决的问题

模拟只能说明交易在某个时刻某种状态下*可能*产生什么结果，它不是保证。用户签名前，链上状态和价格可能已经变化。目前，即使 Agent 的预测没有兑现，运营方也几乎不承担任何经济责任。

### 工作流程

```
用户：把 1 MON 换成 tUSDC
          ↓
Agent 模拟交易，发布包含保证输出 + 锁定保证金的 Plan
          ↓
用户查看保证条款，签名
          ↓
BondedExecutor 执行 Swap，测量实际输出
          ↓
  实际 ≥ 保证？ → 保证金退还运营方
  实际 < 保证？ → 差额自动从保证金赔付用户
  Swap 失败？   → 退还用户输入 + 失败补偿
```

### P0 范围（黑客松交付）

| 组件 | 说明 |
|------|------|
| **交易对** | 仅 MON → tUSDC |
| **运营方** | 单一团队控制 |
| **网络** | Monad Testnet（Chain ID: 10143） |
| **合约** | MockUSDC、MockDex、BondedExecutor |
| **赔付判定** | 纯链上，通过代币余额变化测量 |

### 合约说明

- **MockUSDC.sol** — 测试输出代币与保证金资产（仅部署者可铸造）
- **MockDex.sol** — 模拟 DEX，可配置汇率，支持 `minOutput`
- **BondedExecutor.sol** — 核心：创建担保计划、锁定保证金、执行 Swap、测量余额、自动赔付

### 明确不做的事

- ❌ 通用自然语言承诺
- ❌ 由 LLM 或管理员判断是否违约
- ❌ 保险池、跨链、治理、发币
- ❌ 本次黑客松中与 AnnoPilot 合并

### 技术栈

| 层 | 技术 |
|----|------|
| 智能合约 | Solidity（Foundry） |
| Agent 后端 | Go + go-ethereum |
| 前端 | React / Next.js |

---

## Monad Testnet Deployment

| Contract | Address |
|----------|---------|
| MockUSDC | [`0xD5B1b42929188280631ef2502c78AA61e1A56e0a`](https://testnet.monadscan.com/address/0xD5B1b42929188280631ef2502c78AA61e1A56e0a) |
| MockDex | [`0x57dB71757893eE968197B80faB991F16086ec55e`](https://testnet.monadscan.com/address/0x57dB71757893eE968197B80faB991F16086ec55e) |
| BondedExecutor | [`0x5F9C61C6ff535245f8A8112F9BAcf82b39Ee25D0`](https://testnet.monadscan.com/address/0x5F9C61C6ff535245f8A8112F9BAcf82b39Ee25D0) |

**Verified scenarios (see [deploy-monad-testnet.md](deploy-monad-testnet.md)):**

| Scenario | Result | Tx |
|----------|--------|-----|
| Normal (actual ≥ guaranteed) | ✅ Bond released, 0 compensation | `0x94086a...fd30` |
| Shortfall (actual < guaranteed) | 🔴 20 tUSDC auto-paid | `0xe1a468...a1c4` |

## Team

**Backchannel** — Monad Playground Hackathon

| Role | Responsibility |
|------|---------------|
| Researcher / PM | Product scope, promise validity, testing, pitch |
| Marketing / Ops | Positioning, copy, demos, evidence collection |
| Dev: Contracts + Agent Backend | Solidity (Foundry), Go + go-ethereum agent service |
| Dev: Frontend | React / Next.js, wallet integration, promise card UI |
| Design | Information hierarchy, promise card, tx status visuals |
