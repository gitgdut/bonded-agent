# Bonded Agent 前端设计方案(P0)

> 版本:v1.1 | 日期:2026-08-04 | 状态:规划稿,未实现 | 变更:§10 视觉规范更新为紫黑科技金融风
> 适用范围:Monad Playground 黑客松项目 Bonded Agent 的 P0 前端(Next.js + TypeScript + viem/wagmi + Tailwind CSS)
> 关联文档:《Bonded Agent 技术实现方案(P0 版)》

---

## 1. 文档目的

本文档是前端部分的详细设计规格,供实现者直接编码使用,涵盖:

- 信息架构、路由与页面流程
- 承诺卡片(核心组件)的区块规格
- 计划状态机与用户可见文案
- 组件体系与关键技术决策(wagmi 配置、事件订阅、错误处理)
- 文案合规、视觉规范、可访问性
- 测试与验收标准、里程碑任务拆解

## 2. 设计目标与原则

### 2.1 设计目标

| 目标 | 说明 | 度量方式 |
| --- | --- | --- |
| 签名前可理解 | 用户能在点击"执行"前,用普通语言复述保证条款 | 可用性测试三问(见 §14.4) |
| 结算后可解释 | 结算后展示区块链接 + 一句话解释 | 历史页每条可自解释 |
| 真实性 | 所有状态只来自链上事件与合约读取,禁止硬编码结果 | 代码审查 + e2e |
| 不误导 | 模拟汇率、测试网、资产性质有明确标注 | 文案合规检查(§9) |

### 2.2 设计原则

1. **承诺卡片优先**:保证条款是最高信息层级,位于第一屏,签名前不可折叠关键风险。
2. **普通语言优先**:技术术语(bond/liability/expiry/shortfall)一律映射为用户语言。
3. **真实事件驱动**:UI 状态绑定合约事件,运营方服务仅提供报价与开计划。
4. **渐进披露**:默认只展示用户需要决策的信息,详细条款按需展开。
5. **防误导**:不使用"保险、保证盈利、100% 安全"等表述;不夸大模拟部分。

## 3. 范围

### 3.1 本期(P0)包含

- 单一交易类型:测试网原生 MON → tUSDC(金额可输入,默认 1 MON)
- 单一运营方(团队控制),前端不展示运营方管理能力
- 请求 → 报价 → 生成担保计划 → 承诺卡片 → 执行 → 结算全流程
- 计划页、历史页、钱包连接、网络与事件订阅
- 状态覆盖:待执行、过期、执行中、履约、赔付、失败、拒绝

### 3.2 本期(P0)不包含

- 真实 Moss 接入、EIP-712 离线签名、真实 DEX、服务费、多运营方、声誉、ERC-8004、多操作
- 运营方管理台(开计划/保证金管理)由运营方服务与脚本完成,前端只读展示
- 后端数据库:历史记录从链上事件重建


## 4. 技术栈与工程结构

### 4.1 技术栈

| 层 | 选型 | 说明 |
| --- | --- | --- |
| 框架 | Next.js(App Router)+ TypeScript | 静态托管友好,路由清晰 |
| 链交互 | wagmi + viem | 连接注入钱包、读/写合约、事件订阅 |
| 数据缓存 | React Query | 报价与链上数据缓存、轮询 |
| 样式 | Tailwind CSS | 主题 token 见 §10 |
| 状态管理 | 组件局部状态 + React Query | 不引入全局状态库 |

### 4.2 前端目录结构

```text
web/
├─ app/
│  ├─ layout.tsx            # 全局布局(Header/Footer/Providers)
│  ├─ page.tsx              # 首页:连接钱包 + 请求表单
│  ├─ plans/[id]/page.tsx   # 计划页:承诺卡片 + 执行 + 状态
│  ├─ history/page.tsx      # 历史页:计划列表
│  └─ globals.css           # Tailwind 入口与主题 token
├─ components/
│  ├─ layout/               # Header、Footer、NetworkBadge、WalletButton
│  ├─ request/              # RequestForm、QuoteSummary
│  ├─ plan/                 # PromiseCard、PlanStatusBadge、SettlementExplanation
│  └─ ui/                   # ErrorBanner、Toast、AmountToken、Skeleton、EmptyState、TxLink
├─ lib/
│  ├─ wagmi.ts              # 链、transport、连接器配置
│  ├─ contracts.ts          # 合约地址/ABI 读取与环境变量
│  ├─ events.ts             # 事件订阅与计划状态重建
│  ├─ status.ts             # 状态机派生逻辑
│  ├─ copy.ts               # 全部用户可见文案(合规集中管理)
│  └─ format.ts             # 金额、时间、地址格式化
├─ hooks/
│  ├─ usePlan.ts            # 读取 + 派生单个计划状态
│  ├─ usePlanEvents.ts      # 订阅计划事件
│  ├─ useQuote.ts           # GET /quote
│  ├─ useCreatePlan.ts      # POST /plans
│  └─ useExecutePlan.ts     # executePlan 写交易
└─ .env.local               # 环境变量(见 §7.2)
```

## 5. 信息架构与路由

### 5.1 路由表

| 路由 | 页面 | 进入条件 | 主要动作 |
| --- | --- | --- | --- |
| `/` | 首页 | 无 | 连接钱包、输入请求、获取报价、生成计划 |
| `/plans/[id]` | 计划页 | 任意 | 查看承诺卡片、执行、查看结算状态 |
| `/history` | 历史页 | 已连接钱包 | 按状态筛选计划列表、打开链接 |

### 5.2 页面流转

```text
首页(/)
   │ 用户输入 1 MON → 点击"获取担保报价"
   ▼
报价预览(首页内联)
   │ 点击"生成担保计划"(POST /plans)
   ▼
计划页(/plans/[id]) ← 承诺卡片(签名前完整可见)
   │ 点击"同意并执行 Swap"
   ▼
执行中 → 结算(履约/赔付/失败)/ 过期 / 拒绝
   ▼
历史页(/history) 汇总全部计划与链接
```

### 5.3 全局布局

- `Header`:左侧品牌"Bonded Agent";右侧依次为历史入口、网络徽章(Monad Testnet)、钱包按钮。
- 主内容区:最大宽度 `max-w-3xl` 居中,承诺卡片场景下不引入侧栏干扰。
- `Footer`:三行固定披露——"运行于 Monad Testnet"、"当前市场为模拟(MockDex),非真实 DEX"、"测试资产,无真实价值"。


## 6. 页面详细设计

### 6.1 首页 `/`

#### 6.1.1 区块顺序(自上而下)

1. **标题区**:H1"有担保的 Swap",副标题一句话:"Agent 运营方为交易结果锁定保证金,结果不足自动赔付"。
2. **钱包区**(未连接时):引导卡片 + "连接钱包"按钮;已连接时显示地址与 MON 余额。
3. **请求表单**:金额输入(默认 `1`,单位 MON)+ 目标代币(固定展示 tUSDC,下拉禁用)+ 提交按钮"获取担保报价"。
4. **报价预览**(获取报价后出现):
   - 预期输出:约 X tUSDC(标注"模拟汇率")
   - 保证输出:至少 Y tUSDC
   - 运营方锁定保证金:Z tUSDC
   - 有效期预览:预计 N 分钟
   - 按钮"生成担保计划"(loading 态防重复提交)
5. **错误区**:`ErrorBanner` 展示失败原因与可操作提示。

#### 6.1.2 交互与校验

- 金额输入:仅正数、精度 ≤ 18 位;输入非法时按钮禁用并提示。
- 未连接钱包:点击"获取担保报价"先触发连接引导,不发起请求。
- 生成计划成功:跳转 `/plans/[id]`;失败(保证金不足/RPC 错误/超时)回到报价预览并提示重试或重新报价。
- 报价有有效期:计划生成失败提示"报价已过期,请重新获取"。

### 6.2 计划页 `/plans/[id]`

#### 6.2.1 区块顺序(自上而下)

1. **标题区**:担保交易计划 #planId + `PlanStatusBadge`。
2. **承诺卡片**(见 §7,核心组件)。
3. **结算区**(执行/结算后出现):`SettlementExplanation` + 交易链接。
4. **风险说明区**:折叠面板,默认展开前两项,其余可展开(核心条款不允许折叠)。

#### 6.2.2 页面行为

- 进入页面:并行读取链上 `plans(nonce)` 与本地 `/plans/:id`(运营方元数据),读取中显示骨架屏。
- 未找到计划(planId 无效):`EmptyState` + "返回首页重新发起请求"。
- 钱包未连接或当前地址 ≠ 计划用户:执行按钮禁用,提示"请使用计划所属的钱包地址"。
- 执行中:按钮防重复提交,`waitForTransactionReceipt` 期间轮询区块与事件。
- 结算后:按钮区替换为解释卡片,不再允许操作。

### 6.3 历史页 `/history`

#### 6.3.1 区块顺序

1. **标题区**:我的担保计划。
2. **筛选条**:全部 / 待执行 / 履约 / 赔付 / 失败 / 过期。
3. **列表**:每行一条计划卡片:
   - 左侧:planId、时间、金额方向(1 MON → X tUSDC)。
   - 中间:状态徽章 + 一句解释文案。
   - 右侧:区块浏览器链接图标。
4. **空态**:无记录时提示"还没有担保计划,去发起一笔 Swap"。

#### 6.3.2 数据来源

- 订阅/拉取 `PlanOpened`、`PlanExecuted`、`ShortfallPaid`、`PlanFailed`、`BondReleased` 事件,过滤 `user == 当前地址`。
- 按 `planId` 去重合并为列表;状态由 `lib/status.ts` 统一派生。

## 7. 承诺卡片详细设计(核心组件)

### 7.1 布局规格(自上而下,不可折叠关键区块)

| 区块 | 内容 | 视觉权重 | 是否可折叠 |
| --- | --- | --- | --- |
| A 标题 | 担保交易计划 #planId + 状态徽章 | 中 | 否 |
| B 交易摘要 | 1 MON → 预期约 X tUSDC(附"模拟汇率"角标) | 中 | 否 |
| C 保证条款 | 保证输出 / 锁定保证金 / 最大赔付 / 失败补偿 / 有效期截止 | 最高 | 否 |
| D 风险边界 | 只保证最低到账、赔付上限、模拟市场声明 | 高 | 否(前三项) |
| E 操作 | "同意并执行 Swap"按钮 + 辅助说明 | 高 | 否 |

### 7.2 保证条款区字段规格

| 字段 | 链上字段 | 展示格式 | 备注 |
| --- | --- | --- | --- |
| 保证输出 | `guaranteedOutput` | 至少 X tUSDC | 主数字,最大字号 |
| 运营方已锁定保证金 | 事件/读取 | Y tUSDC | 注明"由运营方锁定" |
| 最大赔付 | `maxCompensation` | Z tUSDC | 注明"赔付上限" |
| 失败补偿 | `failureCompensation` | W tUSDC | 注明"交易失败时另付" |
| 有效期截止 | `deadline` | 本地时间 + 倒计时 | 以区块时间校准 |

### 7.3 风险与边界区文案(固定模板)

- "本计划只保证这一笔 Swap 的最低到账,不保证盈利。"
- "仅当实际到账低于保证值时自动补足,赔付上限为锁定的保证金。"
- "当前市场为模拟(MockDex),汇率由团队配置,非真实 DEX;真实市场接入为后续版本。"
- "测试网与测试资产,无真实价值。"

### 7.4 执行按钮规格

| 状态 | 按钮文案 | 可用性 |
| --- | --- | --- |
| 待执行(有效期内) | 同意并执行 Swap | 可用 |
| 已过期 | 计划已过期 | 禁用 |
| 执行中(已提交) | 等待交易确认… | 禁用 + 防重复 |
| 已结算 | 本计划已结算 | 禁用 |
| 钱包未连接/非计划用户 | 请连接计划所属钱包 | 禁用 + 提示 |


## 8. 计划状态机

### 8.1 状态定义与派生规则

```ts
type PlanStatus =
  | 'loading'      // 链上数据读取中
  | 'open'         // PlanOpened 已发出,未过期,无执行事件
  | 'expired'      // 当前区块时间 > deadline 且无执行事件
  | 'executing'    // 用户已提交交易,等待回执
  | 'settled_ok'   // PlanExecuted 且无 ShortfallPaid
  | 'settled_shortfall' // PlanExecuted 且有 ShortfallPaid
  | 'failed'       // PlanFailed
  | 'rejected';    // 本地状态:用户主动放弃(链上无记录)
```

派生顺序:先读事件;若均无,再按 deadline 判断 open/expired;`executing` 由本地交易状态临时置入;`rejected` 仅本地,刷新后按链上事实回退。

### 8.2 事件 → 状态映射

| 事件 | 状态变化 | 附带数据 |
| --- | --- | --- |
| `PlanOpened` | → open | planId、保证金锁定 |
| `PlanExecuted` | → settled_ok / settled_shortfall | actualOutput、netOutput |
| `ShortfallPaid` | 叠加标记 | shortfall 金额 |
| `PlanFailed` | → failed | refunded、compensation |
| `BondReleased` | 结算详情补充 | 释放金额 |

### 8.3 用户可见解释模板(结算后)

| 状态 | 解释文案 |
| --- | --- |
| settled_ok | "执行成功,到账 X tUSDC,达到保证值;运营方保证金已释放。" |
| settled_shortfall | "实际到账 X tUSDC,低于保证值,已自动补足差额 Y tUSDC。" |
| failed | "交易失败,已退还 1 MON,并支付失败补偿 W tUSDC。" |
| expired | "计划已过期,未执行,无任何资金变动。" |

每条解释下方附"在区块浏览器中查看"链接。

## 9. 文案与合规

### 9.1 术语映射(普通语言)

| 技术术语 | 用户语言 |
| --- | --- |
| bond | 运营方锁定的保证金 |
| liability / maxCompensation | 最大赔付 |
| expiry / deadline | 有效期截止 |
| shortfall | 差额自动补足 |
| calldata / target | 目标合约与调用数据(技术说明页) |

### 9.2 禁用表述

- 保险、承保、理赔(保险语义)
- 保证盈利、稳赚、无风险
- 100% 安全、生产环境安全
- 官方、受信任 DEX(模拟市场)

### 9.3 必须标注

- 模拟汇率(MockDex,团队可配置)
- Monad Testnet
- 测试资产,无真实价值

所有文案集中在 `lib/copy.ts`,实现期禁止在组件内散写。

## 10. 视觉规范(紫黑科技金融风,建议,实现时可微调)

### 10.1 设计方向

- 整体基调:**紫黑深色主题**——近黑带紫调底色 + 高饱和紫罗兰主色,营造"链上资产守护"的科技感与安全感。
- 三个关键词:**紫黑、科技、金融**——紫黑定基调,科技表现链上实时与数据可视化,金融表现金额、保证金与状态语义。
- P0 固定深色主题,不跟随系统 `prefers-color-scheme`,保证视觉一致性;深色底上仍保证对比度 ≥ 4.5:1(WCAG AA)。

### 10.2 主题 token(Tailwind CSS 变量)

| Token | 值 | 用途 |
| --- | --- | --- |
| `--color-bg` | 紫黑 #0b0612 | 页面底色(近黑带紫) |
| `--color-bg-elevated` | #150e22 | 卡片/面板底 |
| `--color-bg-elevated-2` | #1d1430 | 悬浮/选中底 |
| `--color-border` | #2c1f45 | 卡片描边、分隔线 |
| `--color-primary` | #8b5cf6 | 主按钮、链接、品牌色 |
| `--color-primary-strong` | #7c3aed | 主按钮 hover、强调描边 |
| `--color-primary-soft` | rgba(139,92,246,.16) | 选中底、标签底、发光蒙层 |
| `--color-accent-tech` | #22d3ee | 科技点缀:数据流、扫描线、网络状态灯 |
| `--color-bond` | 金 #f59e0b | 保证金/保证条款/金额强调(金融金) |
| `--color-positive` | 绿 #34d399 | 履约/成功状态 |
| `--color-shortfall` | 橙 #fb923c | 赔付状态 |
| `--color-failed` | 红 #f87171 | 失败状态 |
| `--color-neutral` | 灰 #9ca3af | 过期/拒绝/辅助信息 |
| `--color-text-primary` | #f4efff | 主文字(紫调白) |
| `--color-text-secondary` | #a9a0bd | 次级文字 |
| `--color-text-muted` | #6e6580 | 辅助/占位文字 |

- 状态语义固化:绿=履约/成功,橙=赔付,红=失败,金=保证金/保证输出,灰=过期/拒绝。
- 涨跌约定采用加密市场惯例:**绿涨/红跌**;金额方向与汇率波动统一按此配色,不采用 A 股"红涨绿跌"。

### 10.3 科技与金融元素

**科技感**

- 页面背景:紫黑底 + 极淡 24px 网格线(`--color-primary-soft`,透明度 ≤ 6%),顶部可叠加一道紫→靛渐变光晕。
- 品牌与标题:可选极细紫色描边/光晕(`text-shadow`),H1 配紫→靛渐变下划线。
- 链上状态灯:卡片与徽章前置 6px 圆形状态灯,待执行/执行中/结算分别以紫、青(呼吸闪烁)、绿呈现,体现"实时上链"。
- 数据流点缀:报价预览与计划页加一条横向青色流动细线(`--color-accent-tech`),寓意链上事件流。
- 玻璃拟态:卡片用 `rgba(21,14,34,.72)` + `backdrop-blur(12px)` + 1px 紫色描边,深色上保持层次。

**金融感**

- 金额数字一律使用等宽数字字体(JetBrains Mono / IBM Plex Mono),`font-variant-numeric: tabular-nums`,避免金额与倒计时跳动。
- 保证金/保证输出用金色 `--color-bond` 加粗呈现,可配图形化徽标(如锁形图标);文案仍守合规,不出现"保险"语义。
- 报价预览与历史行可内嵌迷你走势图(sparkline/柱状),示意"模拟汇率"波动,并标注"模拟"。
- 顶部可选行情走马灯(Ticker):MON / tUSDC 模拟汇率 + 网络状态,置于 Header 下,仅展示、不做交易入口。
- 结算解释卡片底部附"在区块浏览器中查看"链接,带外链图标与 hover 紫色下划线。

### 10.4 字号与层级

- 页面 H1:24–28px,粗体,`--color-text-primary`
- 卡片标题:18px,粗体
- 保证输出主数字:28–32px,粗体,金色 `--color-bond` + 轻微发光(`text-shadow 0 0 12px rgba(245,158,11,.35)`)
- 等宽数字:16px 起,`tabular-nums`
- 辅助说明:14px,常规,`--color-text-muted`
- 字体栈:正文 Inter/系统字体;数字与代码 JetBrains Mono;标题可选 Space Grotesk

### 10.5 组件样式

| 组件 | 样式 |
| --- | --- |
| 卡片 | 圆角 12px,玻璃底 + 1px `--color-border` 描边,阴影 `0 8px 24px rgba(0,0,0,.45)` 叠加 `0 0 24px rgba(139,92,246,.08)` 紫辉光 |
| 主按钮 | 紫→靛渐变(`#8b5cf6 → #6366f1`),圆角 8px,hover 上浮 + 外发光;禁用态降透明度并去掉辉光 |
| 次按钮/链接 | 透明底 + 1px 描边,hover 描边转主色、文字转紫 |
| 输入框 | 深底 `--color-bg-elevated` + 描边,focus 时 1px 主色描边 + 2px 主色光环 |
| 状态徽章 | 胶囊形,同色系深底 + 亮字(如履约 `rgba(52,211,153,.14)` 底 + `#34d399` 字),前置状态灯,对比度 ≥ 4.5:1 |
| 骨架屏 | 紫灰渐变 shimmer,圆角与卡片一致,不引起布局跳动 |

### 10.6 动效与反馈

- 页面/卡片进入:淡入 + 上移 8px,200ms,`prefers-reduced-motion` 下关闭。
- 卡片 hover:边框由 `--color-border` 渐亮至主色 50%,辉光增强。
- 状态变化:徽章状态灯脉冲 2 次;倒计时每秒刷新(等宽数字防跳动)。
- 交易提交:执行中按钮转 loading 图标 + 青色流动线,示意链上确认中。
- 所有动效 ≤ 300ms,不阻塞操作,不产生布局位移。

### 10.7 深色主题注意

- 正文不用低对比紫(如 `#a855f7` 直接做正文);正文一律用 `--color-text-primary/secondary/muted`。
- 状态色在深底上用 400 级亮色,不用 800 级深色,保证对比度 ≥ 4.5:1。
- 玻璃拟态保留不透明度兜底:不支持 `backdrop-filter` 的浏览器回退到 `--color-bg-elevated`。
- 金色强调只用于金额/保证金,不做大面积背景,避免喧宾夺主。

## 11. 组件体系

### 11.1 组件清单与职责

| 组件 | 职责 | 关键 props |
| --- | --- | --- |
| `Layout` | 页面骨架、Provider 注入 | children |
| `Header` | 品牌、导航、钱包区 | — |
| `WalletButton` | 连接/断开/显示地址 | — |
| `NetworkBadge` | 显示当前链与网络状态 | chainId |
| `RequestForm` | 金额输入与校验 | onQuote |
| `QuoteSummary` | 报价预览 | quote |
| `PromiseCard` | 承诺卡片(§7) | plan、status |
| `PlanStatusBadge` | 状态徽章 | status |
| `SettlementExplanation` | 结算解释 + 交易链接 | settlement |
| `TxLink` | 拼接 explorer URL | txHash |
| `AmountToken` | 代币金额格式化 + 加载占位 | value、decimals |
| `ErrorBanner` | 错误提示与操作指引 | error、onRetry |
| `Toast` | 轻提示(成功/失败) | message、type |
| `Skeleton` | 加载骨架 | — |
| `EmptyState` | 空态引导 | title、action |

## 12. 数据与事件流

### 12.1 报价与开计划(前后端契约)

`GET /quote`

```text
请求:?inputAmount=1000000000000000000
响应:{ inputAmount, outputToken, expectedOutput, simulatedRate, timestamp }
```

`POST /plans`

```text
请求:{ inputAmount, expectedOutput, deadlineMinutes }
响应:{ planId, guaranteedOutput, maxCompensation, failureCompensation, deadline, target, calldataHash, txHash }
```

`GET /plans/:id`

```text
响应:{ planId, status, user, inputAmount, expectedOutput, guaranteedOutput,
       maxCompensation, failureCompensation, deadline, txHashes[] }
```

### 12.2 执行流程(前端视角)

1. `useExecutePlan` 调用 `executePlan(planId, calldata)`,原生 MON 走 `value: inputAmount`。
2. 钱包签名 → `waitForTransactionReceipt`。
3. 回执确认后,事件订阅同步刷新状态到结算态。
4. 任何一步失败:按错误码映射为可读文案(§12.4)并允许重试(非幂等错误除外)。

### 12.3 事件订阅

- 使用 `useWatchContractEvent` 订阅 `PlanOpened` / `PlanExecuted` / `ShortfallPaid` / `PlanFailed` / `BondReleased`。
- 计划页按 `planId` 过滤;历史页按当前地址过滤。
- 订阅建立前先拉一次历史事件(`getLogs`)兜底,避免刷新丢状态。

### 12.4 错误码映射

| 错误 | 用户文案 | 处理 |
| --- | --- | --- |
| `PlanExpired` | 计划已过期,请重新生成 | 禁用按钮,提供"重新发起" |
| `AlreadyExecuted` | 该计划已执行过 | 刷新为结算态 |
| `InvalidCalldataHash` | 调用数据不匹配 | 提示重新获取报价 |
| `NotPlanUser` | 该计划不属于当前钱包 | 提示切换钱包 |
| `InsufficientBond` | 运营方保证金不足 | 提示稍后重试 |
| 用户拒绝签名 | 已取消,无资金变动 | 回到待执行态 |
| RPC/网络错误 | 网络异常,请重试 | 重试按钮 |


## 13. 响应式、可访问性与安全

- **响应式**:桌面双列(条款区 + 操作区),移动端单列;承诺卡片移动端优先可读,保证条款区不换行截断。
- **可访问性**:对比度 ≥ 4.5:1(WCAG AA);按钮与链接键盘可达;状态变化用 ARIA `aria-live` 播报;倒计时与加载不引起布局跳动。
- **性能**:报价结果与链上数据经 React Query 缓存;列表分页/按需拉取。
- **安全**:客户端不持有私钥;不打印敏感错误(环境变量、签名原文);所有展示数据来自链上或运营方只读接口。

## 14. 测试与验收

### 14.1 单元测试

- `lib/status.ts`:事件 → 状态派生,含 deadline 边界(恰好在过期前/后)。
- `lib/copy.ts`:文案模板参数填充;禁用词扫描。
- `lib/format.ts`:金额格式化(18 位小数)、时间与倒计时。
- 错误码映射:每个错误码返回预期文案。

### 14.2 组件测试

- 承诺卡片:各状态渲染、过期禁用、非计划用户禁用、加载骨架。
- `RequestForm`:非法输入拦截、未连接钱包引导。
- `WalletButton` / `NetworkBadge`:未连接、错误链、断线状态。
- 防重复提交:双击按钮只产生一次交易。

### 14.3 端到端(e2e)

正常案例:连接钱包 → 请求 → 报价 → 生成计划 → 承诺卡片 → 执行 → 履约结算 → 看到解释与链接。
赔付案例:运营方调低 MockDex 汇率 → 同上流程 → 结算为赔付态,展示差额补足。
分支案例:过期、用户拒绝签名、网络错误、备用 RPC 切换、全新浏览器环境。

### 14.4 可用性测试(≥3 次)

- 三问:谁提供保证金 / 保证了什么 / 为什么可能赔付。
- 观察:签名前是否读完卡片;是否误把"预期输出"当"保证输出";是否理解模拟市场。
- 输出:记录原话与困惑点,驱动承诺卡片与文案修改。

### 14.5 验收标准

- 新用户连接钱包后,签名前能复述保证条款。
- 所有展示状态来自真实链上事件,无硬编码结果。
- 模拟汇率与测试网有明确标注;无禁用表述。
- 两套链上案例(履约/赔付)可完整走通并展示链接。

## 15. 里程碑任务拆解

### M3:脚手架、布局、请求与报价

- [ ] Next.js + Tailwind 脚手架,Provider(wagmi/React Query)接入
- [ ] `lib/wagmi.ts` 链配置与连接器
- [ ] Header/Footer/NetworkBadge/WalletButton
- [ ] RequestForm + 校验
- [ ] useQuote + QuoteSummary
- [ ] useCreatePlan + 跳转计划页
- [ ] ErrorBanner / Toast 基础设施

### M4:承诺卡片、执行、状态机与历史

- [ ] PromiseCard(§7 全部区块)
- [ ] lib/status.ts 状态机 + 事件订阅(usePlanEvents)
- [ ] useExecutePlan + 回执追踪 + 防重复
- [ ] SettlementExplanation + TxLink
- [ ] 历史页:事件重建列表 + 筛选
- [ ] 错误码映射与用户文案
- [ ] 响应式与骨架屏

### M6:打磨、合规与验证

- [ ] 文案合规复查(禁用词扫描、术语映射)
- [ ] 可用性测试 ≥3 次并按反馈修改
- [ ] 视觉一致性检查(桌面/移动)
- [ ] 全新浏览器端到端清单
- [ ] 备用 RPC 与恢复演练

## 16. 开放问题(实现前确认)

| # | 问题 | 影响 | 默认处理 |
| --- | --- | --- | --- |
| 1 | Monad Testnet chainId 与 explorer URL 模板 | 链配置、链接 | 实现时从 RPC/文档确认,环境变量注入 |
| 2 | `GET /quote` 的汇率来源(链上读取 or 服务端配置) | 报价一致性 | 优先读 MockDex 链上汇率 |
| 3 | 支持的钱包连接器范围 | 连接体验 | 注入式钱包(MetaMask/Rabby),按环境补 WalletConnect |
| 4 | deadline 默认时长 | 报价与计划有效期 | 默认 15 分钟,运营方服务可配置 |
| 5 | POST /plans 是否返回已上链计划的事件 | 跳转时机 | 返回后按 planId 轮询/订阅 `PlanOpened` |

---

*本文档为规划稿,未执行任何实现。实现时以《Bonded Agent 技术实现方案(P0 版)》为准。*

