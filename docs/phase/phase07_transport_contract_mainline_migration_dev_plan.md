# phase07_transport_contract_mainline_migration_dev_plan

## 1. 文档定位

本文档定义 `phase07_transport_contract_mainline_migration` 的执行顺序、子任务范围、DoD 与明确不做。

`phase07` 是 `phase06` 收口之后、`mvp0.3` 业务阶段之前的前置基础 phase。它不是传输层试点，也不是只冻结“以后用 ConnectRPC”，而是要完成 `/plan -> /spec -> 实现 -> 验收 -> 收口` 全链路，并最终交付 `phase01 ~ phase06` canonical 业务接口的正式传输主线切换。

相较于 `phase06`，本阶段在任务拆分上显式吸取以下经验：

- 结构性隐患一旦进入 `audit -> escalate-phase`，就不能继续留在“边做业务边顺手收敛”
- 传输层主线切换如果只做新接口默认规则，而不清旧主线，最终一定会演化成双真相长期并存
- Connect 生成链、Go handler 挂载与前端客户端切换必须作为同一个交付相位处理，不能拆成互相等待的散装任务

## 2. 本阶段目标

在保持 `.proto` 为唯一长期合同源、`chi` 保留装配职责的前提下，交付：

- `buf` 生成链升级
- `ConnectRPC` 业务传输主线
- `phase01 ~ phase06` canonical 业务接口一次性迁移
- 全链路回归验收与旧 JSON 业务主线退场

使仓库在进入 `mvp0.3` 业务主线前，完成“业务接口真相直接收敛到 `.proto + ConnectRPC`”的结构切换。

## 3. 子任务清单

### 第一组：冻结类子任务

### phase07-01 冻结本阶段迁移范围、canonical 业务接口清单与非业务端点边界

范围：

- 冻结 `phase01 ~ phase06` 必须在本阶段迁移的 canonical 业务接口清单
- 冻结到 `service / RPC / 当前外部访问路径 / 页面或动作 owner` 级别的迁移总表
- 冻结不纳入迁移的非业务端点范围
- 冻结 phase 收口时“业务主线已切换”的判定标准

DoD：

- 已明确哪些接口必须迁移，哪些端点继续保留在 `chi + net/http`
- 已明确每个已交付业务模块下的 RPC、当前入口路径与所属页面/动作承接位
- 已明确 phase 收口后不再允许保留手写 JSON 业务主线
- 不把“只迁部分模块、其余继续兼容”解释为本阶段完成

### phase07-02 冻结 `chi + ConnectRPC + buf` 的正式组合方式

范围：

- 冻结 `chi` 的唯一职责
- 冻结 Connect handler 的挂载方式
- 冻结浏览器侧统一 `/api` 前缀、Vite dev proxy、Caddy 与本地启动链的访问路径承接方式
- 冻结 `buf.gen.yaml` 的插件与产物落点
- 冻结前端客户端生成方式

DoD：

- `chi` 不再承担业务合同定义职责
- `connectrpc/go` 与 `bufbuild/es` 的生成链要求明确
- 外部访问路径不长出第二套并列基址
- Go / TS 产物落点单值化

### phase07-03 冻结迁移过程兼容策略与 phase 收口退场标准

范围：

- 冻结迁移过程是否允许短时并存 JSON adapter
- 冻结临时并存的使用前提与禁止事项
- 冻结当前真实 legacy / compat 业务入口的退场清单
- 冻结 phase 收口前的退场标准

DoD：

- 已明确“兼容只允许存在于迁移过程”
- 不允许把临时 adapter 写成长期正式接口
- 已明确当前 legacy / compat 业务入口 inventory，至少覆盖：
  - `/api/candidates/products`
  - `/api/candidates/repositories`
  - `/api/modules/{moduleId}/bindings/products`
  - `/api/modules/{moduleId}/bindings/repositories`
- 每个 legacy / compat 入口都已明确：
  - 当前调用方
  - 对应替代 RPC / Connect path
  - 允许并存的最晚时点
  - 退场时的删除证据与回归证据
- phase 收口时旧 JSON 业务主线必须退场

### 第二组：实现设计产出类子任务

### phase07-04 产出 Go Connect handler、service implementation 与 chi mount 设计

范围：

- 产出 generated handler 与 service implementation 的组合方式
- 产出 `chi` middleware、Connect interceptor 与错误处理链的承接位
- 产出 procedure path、route group 与 router 结构调整方案
- 产出 domain error -> proto error code -> Connect error 的单值映射方案

DoD：

- 后端传输结构足以直接进入实现
- `chi` 与 Connect 各自职责不再模糊
- request id、logging、recovery、auth / context 注入等横切逻辑的唯一承接位明确
- 错误语义不会在迁移后长出第二套长期映射
- 不引入第二套路由组织模式

### phase07-05 产出前端生成客户端、切片承接位与 query/application 迁移设计

范围：

- 产出前端生成客户端承接位
- 产出 query / application 层的调用切换策略
- 产出前端正式写动作的 mutation owner inventory 与收口策略
- 产出旧 API adapter 的回收顺序

DoD：

- 前端不会因切传输层而长出第二套调用组织
- query 与 mutation 边界继续保持清晰
- 已明确当前组件 / 页面级 mutation 清单，并对每一项给出：
  - 回收到 `application` owner
  - 暂时保留在切片内固定承接位
  - 或明确标记为 phase07 允许存在的短时过渡位及退场条件
- canonical 写动作至少覆盖：
  - create product / repository / module / decision
  - create release
  - bind module to product
  - map module to repository
  - bind repository to product
  - link decision to target
  - trigger export
  - trigger backup verify
- 设计结果足以直接指导现有 adapter 回收

### phase07-06 产出 phase01 ~ phase06 业务接口迁移矩阵与回归验收设计

范围：

- 产出 `service / RPC / 当前入口路径 / 页面动作 owner` 级别的业务接口迁移顺序
- 产出跨模块回归清单
- 产出 fixture、联调、验收与退场证据矩阵
- 产出 legacy endpoint retirement inventory 与 frontend mutation owner inventory 的验收映射
- 产出 Vite、本地启动脚本、验收脚本与 CI 中 proto 生成链的迁移清单

DoD：

- 每个已交付业务模块、每个 canonical RPC 都有对应迁移 owner 与回归项
- 已明确 `/api` 单一访问前缀在 dev、验收与部署链路中的承接位
- 已明确现有验收脚本、reset 脚本与构建脚本的迁移要求
- 已明确 legacy / compat 入口退场所需的 endpoint 级证据
- 已明确前端正式写动作 owner 收口完成态与允许保留的过渡项清单
- 回归验收足以证明主线等价迁移成立
- 已明确 phase 收口所需的最终证据

### 第三组：规格、实现与验收子任务

### phase07-07 产出首份传输主线完全切换正式规格文档

范围：

- 基于前置冻结与设计产出 `phase07` 对应的 `/spec`
- 作为后续实现与下一阶段的直接上游规格来源

DoD：

- 文档完整覆盖迁移范围、生成链、后端、前端、验收与退场标准

### phase07-08 落实 buf 生成链与 ConnectRPC 正式合同产物主线

范围：

- 调整 `proto/buf.gen.yaml`
- 调整生成脚本、验收脚本、启动链与文档
- 生成 Go Connect 与前端客户端产物

DoD：

- `make gen` / `buf generate` 可稳定产出新主线所需产物
- 本地开发、验收与 CI 对生成链的调用入口保持一致
- 生成链与根级规则保持一致

### phase07-09 落实 Go 后端业务传输主线切换

范围：

- 将 `phase01 ~ phase06` 业务接口切到 Connect handler / service implementation
- 调整 `router.go` 与相关 transport owner
- 落实 `/api` 前缀下的 Connect procedure path 挂载与 `chi` middleware 继承
- 落实 Connect interceptor 与错误语义映射
- 按 legacy endpoint retirement inventory 清退旧 compat 业务入口
- 清退旧手写 JSON 业务 handler 主线

DoD：

- canonical 业务接口已完成正式切换
- `chi` 只保留 shell 与非业务端点职责
- 外部访问路径对浏览器侧保持单一 `/api` 基址
- 已按 inventory 清退或关闭所有列入退场范围的 legacy / compat 业务入口
- 不再新增或保留手写 JSON 业务主线

状态同步：

- `2026-08-11`：`phase07-09` 已完成源码优先独立复核、阻断修复与运行时验收
- 当前后端结论：canonical Connect transport 已切到正式主线；L1/L2 候选 compat 路由已退场；L3/L4 绑定 compat 薄壳仍保留到 `phase07-10`
- 当前证据入口：`.trae/specs/phase07_09_cut_go_backend_transport_mainline/checklist.md`
- 下一步：进入 `phase07-10`，继续完成前端客户端、adapter 与 mutation owner 收口

### phase07-10 落实前端客户端与业务切片调用切换

范围：

- 将前端业务调用切到生成客户端或等价 Connect transport adapter
- 回收旧 API adapter
- 调整开发环境代理与传输层承接位
- 按 mutation owner inventory 回收或冻结正式写动作承接位
- 保持既有 query / application 边界

DoD：

- 前端业务主线已完成正式切换
- 页面与切片行为保持等价
- 不引入第二套前端 API 基址或第二套 transport owner
- 所有 canonical 写动作都已有单一正式 owner；若仍存在过渡 mutation，已被显式列入允许清单并附退场条件
- 不保留第二套长期 fetch / JSON 主线

### phase07-11 完成 phase01 ~ phase06 联调、回归与退场验收

范围：

- 覆盖 Module / Decision / Product / Repository / Dashboard / Onboarding / Export / Backup / Reuse Summary
- 覆盖本地启动链、验收脚本与 CI 生成链
- 验证旧业务主线退场
- 验证 legacy endpoint retirement inventory 与 frontend mutation owner inventory 已按计划收口
- 留存正式验收记录

DoD：

- 回归通过
- 开发、验收与部署链路均已通过单一 `/api` + Connect 主线运行
- legacy / compat 业务入口 inventory 已逐项核销，不存在未声明残留
- canonical 写动作 owner 已逐项核销，不存在未声明的页面 / 组件级长期正式 mutation 主线
- phase 收口时不再保留旧手写 JSON 业务主线
- 已形成可供后续 `mvp0.3` 业务 phase 直接承接的正式结论

### phase07-12 完成根级同步与 `mvp0.3` 业务阶段进入条件回写

范围：

- 回写 `AGENTS.md`
- 回写 `plan.md`
- 回写 `docs/README.md`
- 回写 `architecture_map.md`

DoD：

- 根级入口已反映 `phase07` 的前置基础地位
- 后续进入 `mvp0.3` 业务阶段的条件明确
- 不提前猜测未建立阶段名称

## 4. 本阶段明确不做

- 在本阶段直接实现 `Operating Review Loop`
- 在本阶段直接实现 `Template Reuse`
- 在本阶段直接实现 `Derived Intelligence Deepening`
- 在本阶段直接执行真实项目 `dry-run`
- 在本阶段引入新的长期业务实体
- 在本阶段引入重型 gRPC 基础设施、微服务或第二套路由框架

## 5. 本阶段 Done 标准

只有当以下条件同时满足时，`phase07` 才算完成：

1. `phase07-01 ~ phase07-07` 的冻结与规格上游已经成立
2. `phase07-08 ~ phase07-10` 的生成链、后端与前端主线切换已经落地
3. `phase07-11` 已证明 `phase01 ~ phase06` canonical 业务接口切换成立
4. `phase07-12` 已完成根级同步，且未把后续业务 phase 写成既成事实
