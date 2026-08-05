# 合约浏览器验证指南

Monadscan 的 API 被 Cloudflare 保护，无法通过 forge verify-contract 自动验证。需手动在浏览器上操作。

## 已准备好的文件

扁平化源码在 `contracts/flats/` 目录：

| 合约 | 扁平化文件 |
|------|-----------|
| MockUSDC | `flats/MockUSDC.flattened.sol` |
| MockDex | `flats/MockDex.flattened.sol` |
| BondedExecutor | `flats/BondedExecutor.flattened.sol` |

## 手动验证步骤

1. 打开 https://testnet.monadscan.com/
2. 搜索合约地址（见下方）
3. 点击 "Contract" 标签 → "Verify & Publish"
4. 填写：
   - **Compiler**: Solidity 0.8.28
   - **Optimization**: No
   - **License**: MIT
   - 粘贴对应 `flats/*.flattened.sol` 内容
5. 提交

### 三个合约地址

```
MockUSDC:       0xD5B1b42929188280631ef2502c78AA61e1A56e0a
MockDex:        0x57dB71757893eE968197B80faB991F16086ec55e
BondedExecutor: 0x5F9C61C6ff535245f8A8112F9BAcf82b39Ee25D0
```

### MockDex 构造参数

```
_tUSDC: 0xD5B1b42929188280631ef2502c78AA61e1A56e0a
_initialRate: 100000000000000000000
```

### BondedExecutor 构造参数

```
_tUSDC: 0xD5B1b42929188280631ef2502c78AA61e1A56e0a
```

> ⚠️ Monadscan 目前有 Cloudflare 验证，可能需要多次尝试。如果实在无法通过，部署记录（deploy-monad-testnet.md）中的 TxHash 可作为链上证据。
