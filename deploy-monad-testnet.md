# Bonded Agent — Monad Testnet 部署记录

## 部署信息

| 项目 | 值 |
|------|-----|
| 网络 | Monad Testnet (Chain ID: 10143) |
| RPC | https://testnet-rpc.monad.xyz |
| 浏览器 | https://testnet.monadscan.com |
| 部署者 | 0xdc1746793D71256Be16B70188297ab5aB056208F |
| 部署日期 | 2026-08-05 |

## 合约地址

| 合约 | 地址 | 部署 TxHash |
|------|------|------------|
| MockUSDC | `0xD5B1b42929188280631ef2502c78AA61e1A56e0a` | `0x1df324881d2d7a0ff740b2c592be519a5b165cab540c9aac846b13258e61924d` |
| MockDex | `0x57dB71757893eE968197B80faB991F16086ec55e` | `0x001f173ee1b7e57187100916babe5543bd0d30a8da704e571855cd31195e538a` |
| BondedExecutor | `0x5F9C61C6ff535245f8A8112F9BAcf82b39Ee25D0` | `0x0d1d32bbc7b1da5e7ffccc6cdefabe69309af5d9f9dd3c46580a8c1358f157cc` |

## 初始化交易

| 操作 | TxHash |
|------|--------|
| Mint 10000 tUSDC → deployer | `0x3b8b97d2029577c8a140c2b263003d1f01ed133fcf1d7c998fc7fdab74ea9447` |
| Transfer 5000 tUSDC → MockDex | `0x27fd6fd0defd01fc256034f602ced9476b6b7ae3281ff3163fdf4615df6a64a5` |

## 验证交易

### 场景 1 — 正常执行（actual ≥ guaranteed）

| 项目 | 值 |
|------|-----|
| Plan ID | `0x3c7c18b3e636a0322b80cfab54ac036e56d22933848ebc6fd0aaa2aeec61ab25` |
| 输入 | 1 MON |
| 汇率 | 100 tUSDC/MON |
| 期望输出 | 100 tUSDC |
| 保证输出 | 90 tUSDC |
| 创建 Plan Tx | `0x1a648a631dcb2dc797a0803d8e67d6a03cb0c21ffdc4a615cb22f270dbbc38dd` |
| 执行 Tx | `0x94086a8d67c8e0db61f630986cbcc4f762290808442047fd21d3d67c3cd0fd30` |
| 结果 | ✅ 实际 100 ≥ 保证 90，bond 全额释放，0 赔付 |

### 场景 2 — 赔付执行（actual < guaranteed）

| 项目 | 值 |
|------|-----|
| Plan ID | `0xb9bc036b11ef677b9a6844b0243c0e79323ec0813dfb288253c4f07c9269c9e3` |
| 输入 | 1 MON |
| 创建时汇率 | 100 tUSDC/MON |
| 执行时汇率 | 50 tUSDC/MON |
| 期望输出 | 100 tUSDC |
| 保证输出 | 90 tUSDC |
| 创建 Plan Tx | `0x71dca33fd2e8e2a22e22cdd34f65d82e3ca80903786422003def5ffb22e54a47` |
| 执行 Tx | `0xe1a4684058b4915170d4c1c19801ec4b0ef1b1e2fdf77dd9b380d7045052a1c4` |
| 结果 | 🔴 实际 50 < 保证 90，差额 40，赔付 20（maxCompensation 上限），用户收到 70 tUSDC |

## Go Agent 环境变量

```bash
export MONAD_RPC=https://testnet-rpc.monad.xyz
export USDC_ADDR=0xD5B1b42929188280631ef2502c78AA61e1A56e0a
export DEX_ADDR=0x57dB71757893eE968197B80faB991F16086ec55e
export EXECUTOR_ADDR=0x5F9C61C6ff535245f8A8112F9BAcf82b39Ee25D0
export GUARANTEE_RATIO=0.90
export MAX_COMPENSATION=20000000000000000000
export FAILURE_COMPENSATION=5000000000000000000
```
