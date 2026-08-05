# Bonded Agent — 前端接口文档

> 给 Dev 2（React/Next.js 前端）的合约调用说明

## 网络

| 项目 | 值 |
|------|-----|
| 网络 | Monad Testnet |
| Chain ID | 10143 |
| RPC | https://testnet-rpc.monad.xyz |
| 浏览器 | https://testnet.monadscan.com |

## 合约地址

| 合约 | 地址 | ABI 文件 |
|------|------|----------|
| MockUSDC | `0xD5B1b42929188280631ef2502c78AA61e1A56e0a` | `agent/abi/MockUSDC.json` |
| MockDex | `0x57dB71757893eE968197B80faB991F16086ec55e` | `agent/abi/MockDex.json` |
| BondedExecutor | `0x5F9C61C6ff535245f8A8112F9BAcf82b39Ee25D0` | `agent/abi/BondedExecutor.json` |

合约 ABI 文件路径：`agent/abi/*.json`，可直接导入 ethers.js / viem / wagmi。

## Plan 结构体

```solidity
struct Plan {
    address user;              // 谁来执行
    address operator;          // 谁创建的（运营方）
    uint256 inputAmount;       // 交易的 MON 数量（wei）
    uint256 expectedOutput;    // 模拟预期输出（信息性）
    uint256 guaranteedOutput;  // 保证最低输出（不足则赔付）
    uint256 maxCompensation;   // 最大赔付额
    uint256 failureCompensation; // Swap 失败时的赔付
    address target;            // MockDex 地址
    bytes32 calldataHash;      // calldata 哈希（防篡改）
    uint256 deadline;          // Unix 时间戳，超时失效
    uint256 nonce;             // 唯一性
    bool executed;             // 是否已执行
}
```

## 前端需要调用的合约方法

### 1. 查询汇率（只读，免费）

**MockDex.rate()** → uint256

```js
const rate = await dexContract.rate();
// rate / 1e18 = tUSDC per MON, e.g. 100e18 = 100 tUSDC/MON
```

### 2. 模拟 Swap（只读，免费）

链下计算即可：`output = monAmount * rate / 1e18`

### 3. 查询 Plan（只读，免费）

**BondedExecutor.plans(bytes32 planId)** → Plan struct

```js
const plan = await executorContract.plans(planId);
// plan.executed → 是否已执行
// plan.deadline → 过期时间
```

### 4. 创建 Plan（写，需要 tUSDC 授权）

**步骤 A**：用户先授权 tUSDC 给 BondedExecutor

**MockUSDC.approve(BondedExecutor地址, amount)**

**步骤 B**：调用 BondedExecutor.openPlan

```solidity
function openPlan(
    bytes32 planId,      // keccak256(abi.encode(user, operator, nonce))
    Plan calldata plan,  // Plan 结构体
    bytes calldata data  // swap(0) 的 calldata
)
```

**calldata 编码**：`swap(uint256)` 的 ABI 编码，minOutput=0

```js
const iface = new ethers.Interface(["function swap(uint256 minOutput)"]);
const calldata = iface.encodeFunctionData("swap", [0n]);
```

**planId 计算**：

```solidity
planId = keccak256(abi.encode(user, operator, nonce))
```

```js
const planId = ethers.keccak256(
  ethers.solidityPacked(
    ["address", "address", "uint256"],
    [user, operator, nonce]
  )
);
```

Nonce 用 `Date.now()` 或计数器即可。

### 5. 执行 Plan（写，需要 MON）

**BondedExecutor.executePlan(bytes32 planId, bytes calldata)**

```js
const calldata = iface.encodeFunctionData("swap", [0n]);
const tx = await executorContract.executePlan(planId, calldata, {
  value: plan.inputAmount  // 必须等于 plan.inputAmount
});
```

## 事件监听

### 创建成功

```solidity
event PlanOpened(
    bytes32 indexed planId,
    address indexed user,
    address indexed operator,
    uint256 guaranteedOutput,
    uint256 maxCompensation,
    uint256 deadline
);
```

### 执行成功（正常）

```solidity
event PlanExecuted(bytes32 indexed planId, uint256 actualOutput, uint256 paidToUser);
// paidToUser = 0 表示正常，无赔付
```

### 赔付

```solidity
event ShortfallPaid(bytes32 indexed planId, uint256 guaranteed, uint256 actual, uint256 shortfall);
```

### Swap 失败

```solidity
event PlanFailed(bytes32 indexed planId, uint256 refundedMON, uint256 compensationPaid);
```

### 保证金释放

```solidity
event BondReleased(bytes32 indexed planId, address indexed operator, uint256 amount);
```

## 前端 demo 流程

```
1. 显示当前汇率："1 MON = 100 tUSDC"
2. 用户输入 MON 数量 → 计算预期输出
3. 显示保证条款："保证至少 {guaranteed} tUSDC（{ratio}%）"
4. 用户确认 → 调 approve → 调 openPlan
5. 显示 Plan 卡片（planId、金额、保证、倒计时）
6. 用户点"执行" → 调 executePlan（带 value）
7. 监听 PlanExecuted / ShortfallPaid 事件
8. 显示结果：收到多少 tUSDC、有无赔付
```

## 已知限制

- calldata 固定为 `swap(0)`（minOutput=0），前端不需要构造 calldata
- 目前仅支持 MON → tUSDC 一个交易对
- 单运营方、单用户模型
- Plan 过期后无法执行，需创建新 Plan
