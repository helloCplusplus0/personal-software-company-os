# PSCO v0.1 Entity Boundary

## 1. 核心实体

以下对象进入 `v0.1` 主执行范围：

- `Product`
- `Module`
- `Release`
- `Decision`
- `Repository`
- `Venture`（可选）

## 2. 派生层

以下对象在 `v0.1` 中不作为重实体维护：

- `Capability`

其展示应来自：

- 模块数量
- 模块稳定状态
- 版本演进
- 复用关系
- 决策沉淀

## 3. 后移对象

以下对象保留在总体理论模型中，但不进入 `v0.1` 主执行范围：

- `Opportunity`
- `Feature`
- `Experiment`

## 4. 当前执行闭环

`v0.1` 当前优先验证的最小闭环为：

`Product -> Module -> Release -> Decision -> Repository Binding -> Feedback`

## 5. 当前非目标

当前不以以下事项作为 `v0.1` 阻断项：

- GitHub OAuth 自动导入
- Rust Intelligence Layer
- 完整 PMM 正式规范
- 完整 PCP 正式协议
- 独立 AI Assistant 主导航
- 自动扫描代码
- 自动知识图谱
