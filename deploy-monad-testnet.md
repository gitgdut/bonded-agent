# Bonded Agent — Monad Testnet 部署记录

## 部署信息

| 项目 | 值 |
|------|-----|
| 网络 | Monad Testnet (Chain ID: 10143) |
| RPC | https://testnet-rpc.monad.xyz |
| 浏览器 | https://testnet.monadscan.com |
| 部署者 | 0xdc1746793D71256Be16B70188297ab5aB056208F |
| 部署日期 | 2026-08-08 |

## 合约地址

| 合约 | 地址 | 部署 TxHash |
|------|------|------------|
| tUSDC (MockUSDC) | `0xD5B1b42929188280631ef2502c78AA61e1A56e0a` | `0x1df324881d2d7a0ff740b2c592be519a5b165cab540c9aac846b13258e61924d` |
| MockDex | `0x8FA4112739F1fAd68EEB82f75360800354992803` | `forge script DeployMockDex --broadcast --via-ir --skip-simulation` |
| BondedExecutor | `0x195da9B8600717f758F1f4300b9A812CB2CAb768` | `forge script DeployBondedExecutor --broadcast --via-ir` |
| ERC-8004 Identity | `0xe18Fd55C1935554499b0E8c74Db967ccFD46F3c1` | — |
| ERC-8004 Reputation | `0x8A2d492418Ba6b765fE719321ADF2e6f56d9AeA7` | — |

## 初始化交易

| 操作 | TxHash |
|------|--------|
| Mint 10000 tUSDC → deployer | `0x3b8b97d2029577c8a140c2b263003d1f01ed133fcf1d7c998fc7fdab74ea9447` |
| Fund MockDex (2000 tUSDC) | `0x642f4e6e0258c05a06b6d21678f27d3e064d02e18af7f86b70047c44c4d78432` |

## MockDex 说明

MockDex 是带 `setRate(uint256)` 后门的测试 DEX，用于演示不同汇率场景：
- `rate()` — 查询当前汇率
- `setRate(uint256)` — 修改汇率（onlyOwner）
- `swap(uint256 minOutput)` — 用 MON 兑换 tUSDC
- `getAmountOut(uint256)` — 查询预期输出（与 SimpleAMMPair 接口兼容）

## Go Agent 环境变量

```bash
MONAD_RPC=https://testnet-rpc.monad.xyz
USDC_ADDR=0xD5B1b42929188280631ef2502c78AA61e1A56e0a
DEX_ADDR=0x8FA4112739F1fAd68EEB82f75360800354992803
EXECUTOR_ADDR=0x195da9B8600717f758F1f4300b9A812CB2CAb768
GUARANTEE_RATIO=0.90
MAX_COMPENSATION=20000000000000000000
FAILURE_COMPENSATION=5000000000000000000
```
