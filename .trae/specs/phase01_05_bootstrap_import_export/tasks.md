# Tasks

- [x] Task 1: 冻结首轮冷启动路径。将零数据用户从空状态进入第一版可用状态的最小路径写成单值结论。
  - [x] SubTask 1.1: 明确首轮必须允许手动创建 `Product / Module / Repository / Decision`
  - [x] SubTask 1.2: 明确首轮必须允许完成基础绑定关系
  - [x] SubTask 1.3: 验证首轮路径不依赖额外前置配置或外部系统接入

- [x] Task 2: 冻结冷启动的阻断边界。明确当前阶段哪些自动化和外部集成不进入首轮路径。
  - [x] SubTask 2.1: 明确 `GitHub OAuth / 自动导入` 不进入 `v0.1` 首轮
  - [x] SubTask 2.2: 明确自动扫描、自动同步和外部集成不作为冷启动前提
  - [x] SubTask 2.3: 验证首轮默认路径仍以手动创建和手动绑定为主

- [x] Task 3: 冻结导入路径与手动录入边界。将当前阶段支持与不支持的导入口径写清。
  - [x] SubTask 3.1: 明确当前阶段必须先保证 `Product / Module / Decision / Repository` 的手动录入闭环
  - [x] SubTask 3.2: 明确 `v0.1` 当前不冻结任何正式导入能力，首轮资产统一以手动录入为主
  - [x] SubTask 3.3: 明确未来可补充导入说明，但不能写成当前默认能力
  - [x] SubTask 3.4: 验证没有因为未来导入设想而削弱当前手动录入闭环

- [x] Task 4: 冻结最小导出要求。将 `Local First` 语义下用户带走核心资产的最低要求写成可执行前提。
  - [x] SubTask 4.1: 明确导出至少覆盖 `Product / Module / Release / Repository / Decision` 及基础绑定关系，语义为"面向用户带走核心资产数据"
  - [x] SubTask 4.2: 明确导出结果必须足以支持用户理解和迁移核心资产
  - [x] SubTask 4.3: 验证导出策略没有被留成“后续再说”

- [x] Task 5: 冻结最小备份要求。将当前实例的基础备份能力写成单值结论。
  - [x] SubTask 5.1: 明确 `v0.1` 至少需要一种基础备份路径，语义为"面向当前实例保留与恢复"
  - [x] SubTask 5.2: 明确备份要求与数据所有权、迁移能力和恢复预期一致
  - [x] SubTask 5.3: 明确当前阶段不要求自动连续备份、多端同步或复杂灾备体系
  - [x] SubTask 5.4: 验证导出与备份不依赖 GitHub 或第三方平台作为唯一前提

- [x] Task 6: 完成规格校验。检查本次 `phase01-05` 规格是否具备进入后续子任务的条件。
  - [x] SubTask 6.1: 验证已明确首轮如何从零建立 `Product / Module / Decision / Repository` 基础资产
  - [x] SubTask 6.2: 验证已明确当前阶段是否支持导入、支持什么导入、哪些仍采用手动录入
  - [x] SubTask 6.3: 验证已明确 `Local First` 语义下的最小导出 / 备份要求
  - [x] SubTask 6.4: 验证本次规格与 `phase01` 三件套和最终共识保持一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`