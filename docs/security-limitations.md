# 安全限制 & 已知问题

> Bonded Agent v0.1 — Monad Playground Hackathon

## 安全模型

### 信任假设

| 假设 | 说明 |
|------|------|
| 合约代码正确 | 依赖 Solidity 编译器和 Foundry 测试覆盖 |
| MockDex 汇率可信 | 部署者手动设置汇率，无去中心化预言机 |
| 运营方有偿付能力 | 创建 Plan 前需预留足够 tUSDC 保证金 |
| 用户信任 Plan 参数 | 用户需自行验证 expectedOutput / guaranteedOutput 是否合理 |

### 赔付机制

```
actualOutput < guaranteedOutput → 从保证金赔付差额（上限 maxCompensation）
swap 完全失败                  → 退还 MON + 付 failureCompensation
```

**关键属性**：赔付是纯合约逻辑，不依赖管理员、AI 判断或链下仲裁。

---

## 已知限制

### 1. 无滑点保护

calldata 固定为 `swap(uint256(0))`——minOutput=0。不提供传统 DEX 的滑点保护。

**原因**：BondedExecutor 用保证金赔付替代滑点保护。用户不需要设 minOutput，因为合约保证最低输出。

**风险**：如果 MockDex 汇率剧烈波动，实际输出可能远低于预期，但赔付上限由 maxCompensation 限制。

### 2. 单运营方模型

当前仅支持单一运营方创建 Plan。无多运营方竞争、无保证金池机制。

### 3. 单交易对

仅支持 MON → tUSDC。不支持其他代币对或复杂路由。

### 4. Plan 不可修改

Plan 创建后无法更新参数（金额、保证比例、过期时间）。需取消并重建。

### 5. 无重入保护

`executePlan` 遵循 CEI 模式（Checks → Effects → Interactions），在外部调用前设置 `executed = true`，理论上安全。但未使用 OpenZeppelin `ReentrancyGuard`。

### 6. 保证金计价单一

保证金和输出均为同一种代币（tUSDC）。不支持 ETH 或其他资产作为保证金。

### 7. calldata 预哈希

Plan 创建时固定 calldataHash，执行时必须完全匹配。这防止了 calldata 篡改，但也意味着：
- 无法动态调整 swap 参数
- MockDex 接口变更需重新部署 BondedExecutor

### 8. 无 Gas 退款

用户执行 Plan 时支付 Gas，无论结果如何都不退还 Gas。

### 9. 单次执行

每个 Plan 只能执行一次。不支持分批执行或多次 swap。

---

## 已知攻击面 & 缓解

| 攻击 | 风险 | 缓解 |
|------|------|------|
| **重入** | 低 | CEI 模式 + executed 标志前置 |
| **calldata 篡改** | 无 | calldataHash 校验 |
| **nonce 重用** | 无 | usedNonces 映射 |
| **Plan 重放** | 无 | executed 标志 |
| **过期 Plan 执行** | 无 | deadline 检查 |
| **非授权用户执行** | 无 | msg.sender == plan.user |
| **保证金不足** | 无 | transferFrom 前置检查 |
| **运营方跑路** | 中 | 保证金已锁定在合约中，用户始终可执行 |
| **运营方操纵汇率** | 中 | 赔付机制提供经济惩罚（保证金损失） |

---

## 测试覆盖

10 个 Foundry 测试全部通过：

```
✓ NormalFulfillment        — 正常执行，bond 释放
✓ ShortfallExactCompensation — 差额 ≤ maxCompensation，精确赔付
✓ ShortfallMaxCompensation   — 差额 > maxCompensation，按上限赔付
✓ SwapFailure_CallRevert     — Swap 回滚，退 MON + 失败补偿
✓ CalldataTampered           — calldata 不匹配，拒绝
✓ NonceReuseRejected         — nonce 重用，拒绝
✓ ReplayRejected             — Plan 重放，拒绝
✓ PlanExpired                — 过期 Plan，拒绝
✓ WrongUserRejected          — 非指定用户，拒绝
✓ InsufficientBondRejected   — 保证金不足，拒绝
```

---

## 生产化建议（非本次 Hackathon 范围）

- [ ] 接入去中心化预言机（Chainlink）获取汇率
- [ ] 引入 `maxSlippage` 参数作为第二层保护
- [ ] 多运营方 + 保证金竞价
- [ ] 支持任意 ERC-20 交易对
- [ ] 正式审计（Trail of Bits / OpenZeppelin）
- [ ] 事件索引服务（The Graph / Goldsky）
- [ ] 前端显示赔付概率和风险评分
