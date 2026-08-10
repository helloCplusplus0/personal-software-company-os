# Phase06-08 导出、备份与恢复前提的后端模块边界设计 Spec

## Why

`phase06-03` 已冻结 `Export / Backup / restore prerequisites / backup verified` 的正式语义与覆盖矩阵，`phase06-05` 已冻结 `.proto` 为唯一合同源与 `backup_snapshot` 必须保持可重新读取的约束。但如果不在进入实现前把导出 / 备份相关后端模块边界、接口分组、数据装配 owner、`export_snapshot / backup_snapshot` 的读模型承接位、`backup verified` 校验链与既有服务 / 脚本 / 数据库的边界关系设计到可直接落地的程度，后续 `phase06-10` 合同实现与 `phase06-15` 后端实现仍会在模块拆分粒度、跨模块数据读取方式与校验链 owner 归属上各自猜测。

## What Changes

- 产出 `Export` 与 `Backup` 两个独立后端模块的边界与接口分组
- 产出数据装配责任边界：覆盖矩阵中每类数据的读取 owner
- 产出 `export_snapshot / backup_snapshot` 的后端读模型承接位
- 产出 `backup verified` 最小后端校验链与每步 owner 归属
- 产出与既有 canonical 模块、`platform` 装配点、数据库脚本的边界关系
- 产出 `.proto` 合同文件分组与接口 owner
- 产出 `backup_snapshot` 在 `BackupWrite` owner 内的读时 `read/verify` 子路径设计
- 不提前冻结脚本名、存储介质与最终目录细节

## Impact

- Affected specs:
  - `phase06_onboarding_sovereignty_reuse_foundation`
  - `phase06-03` 导出 / 备份 / 恢复前提语义
  - `phase06-05` 合同 / 传输 / 源码约束
  - `phase06_onboarding_sovereignty_reuse_foundation_shared_baseline`
- Affected code:
  - 新增 `backend/internal/export/` 模块（handler / service / candidate / types.go / errors.go / response.go）
  - 新增 `backend/internal/backup/` 模块（handler / service / candidate / repository / types.go / errors.go / response.go）
  - 修改 `backend/internal/platform/router.go` 新增 `buildExport / mountExport / buildBackup / mountBackup` 装配函数
  - 新增 `proto/psco/export/v1/export.proto`
  - 新增 `proto/psco/backup/v1/backup.proto`

---

## 1. 当前后端模块结构概览

### 1.1 既有模块装配模式

当前后端在 `backend/internal/platform/router.go` 中以"组合根"模式装配所有业务模块。每个模块遵循四层结构：

```
backend/internal/{module}/
  handler/           — HTTP handler（query / command 分组）
  service/           — 业务逻辑（query / command service）
  candidate/         — 跨模块只读读取器（由消费方模块拥有）
  repository/        — 数据库存储层（store）
  types.go           — DTO 与域类型
  errors.go          — 错误定义
  response.go        — 响应 envelope 辅助
```

`platform` 包负责把 handler / service / repository / candidate 装配到 chi 子路由，避免业务模块根包与子包之间的导入循环。

### 1.2 既有模块清单

| 模块 | 路径 | 接口分组 | 装配函数 |
| --- | --- | --- | --- |
| Module Registry | `backend/internal/moduleregistry/` | 读组 + 写组 | `mountModuleRegistry` |
| Decision Center | `backend/internal/decisioncenter/` | 读组 + 写组 | `mountDecisionCenter` |
| Product Registry | `backend/internal/productregistry/` | 读组 + 写组 | `mountProductRegistry` |
| Repository Binding | `backend/internal/repositorybinding/` | 读组 + 写组 | `mountRepositoryBinding` |
| Dashboard | `backend/internal/dashboard/` | 仅读组 | `mountDashboard` |

### 1.3 跨模块读取模式

当前跨模块读取统一遵循"消费方模块拥有 candidate reader"模式：

```
// Dashboard 的 candidate reader 由 Dashboard 模块拥有，直接使用 pool 查询 canonical 表
overviewReaders := dashboardcandidate.NewOverviewReaders(pool)
querySvc := dashboardservice.NewQueryService(overviewReaders, ...)
```

- candidate reader 由消费方模块的 `candidate/` 子包定义和拥有
- service 层不直接写跨模块 SQL，通过 candidate reader 接口隔离
- candidate reader 直接使用 `*pgxpool.Pool` 查询 canonical 表

---

## 2. Export 模块边界设计

### 2.1 模块定位

`Export` 模块是 `phase06-08` 新增的独立后端模块，承接"面向用户带走核心资产数据"的导出能力。

### 2.2 模块结构

```
backend/internal/export/
  handler/
    query_handler.go        — ExportRead（读取 export_snapshot：资产范围、最近创建摘要、空态/预览态）
    command_handler.go      — ExportWrite（触发导出、装配数据、返回结果）
  service/
    query_service.go        — export_snapshot 读取逻辑（资产范围、最近创建摘要、空态/预览态）
    command_service.go      — 导出装配逻辑（组装核心资产 + 关系数据）
  candidate/
    asset_reader.go         — 跨模块只读读取器（读取 9 类 canonical 表）
  types.go                  — Export DTO（从 .proto 单向派生）
  errors.go                 — 导出错误定义（装配失败 / 产物生成失败 / 不可读取）
  response.go               — 响应 envelope 辅助
```

### 2.3 接口分组

| 接口 | HTTP 方法 | 路由 | Owner | 职责 |
| --- | --- | --- | --- | --- |
| `ExportRead` | GET | `/api/dashboard/export` | export.query_handler → query_service | 读取 `export_snapshot`：资产范围、最近一次导出创建摘要；无历史结果时返回预览态 snapshot |
| `ExportWrite` | POST | `/api/dashboard/export` | export.command_handler → command_service | 装配核心资产 + 关系数据，生成导出结果 |

### 2.4 职责边界

- Export 模块 SHALL 只读 canonical 表，不写入任何 canonical 数据
- Export 模块 SHALL 通过自有 `candidate/asset_reader.go` 读取跨模块数据，不依赖 canonical 模块的 service
- Export 模块 SHALL 承接 `export_snapshot` 的正式读模型，但当前阶段不提前冻结其底层元数据承接介质
- Export 模块 SHALL NOT 承接备份、恢复或校验职责

### 2.5 `export_snapshot` 读模型

- `ExportRead` 是 `export_snapshot` 的正式读取 owner
- `export_snapshot` 至少承接：
  - 当前可导出的资产范围
  - 最近一次导出创建时间
  - 最近一次导出创建结果摘要
- 若当前尚无历史导出结果，`ExportRead` 必须返回基于当前 canonical 数据装配得到的预览态 `export_snapshot`
- 当前阶段不要求冻结 `export_snapshot` 的底层元数据落盘方式；可以使用轻量元数据来源或等价承接位，但不得让读模型悬空

---

## 3. Backup 模块边界设计

### 3.1 模块定位

`Backup` 模块是 `phase06-08` 新增的独立后端模块，承接"面向当前实例保留与恢复前提校验"的备份能力。

### 3.2 模块结构

```
backend/internal/backup/
  handler/
    query_handler.go        — `BackupWrite` 的 `read/verify` 子路径（读取 backup_snapshot、校验 manifest / 覆盖矩阵 / schema）
    command_handler.go      — BackupWrite（触发备份、创建产物 + manifest）
  service/
    query_service.go        — `BackupWrite` 的读取与校验逻辑（backup verified 校验链）
    command_service.go      — 备份创建逻辑（装配资产 + 生成 manifest + 记录 schema/version）
  candidate/
    asset_reader.go         — 跨模块只读读取器（读取 9 类 canonical 表，同 Export 覆盖范围）
  repository/
    backup_store.go         — 备份产物 / manifest 元数据持久化（如需要）
  types.go                  — Backup DTO（从 .proto 单向派生）
  errors.go                 — 备份错误定义（产物失败 / manifest 缺失 / 覆盖不完整 / schema 缺失）
  response.go               — 响应 envelope 辅助
```

### 3.3 接口分组

| 接口 | HTTP 方法 | 路由 | Owner | 职责 |
| --- | --- | --- | --- | --- |
| `BackupWrite.command` | POST | `/api/dashboard/backup` | backup.command_handler → command_service | 装配核心资产 + manifest + schema/version，创建备份产物 |
| `BackupWrite.read_verify` | GET | `/api/dashboard/backup` | backup.query_handler → query_service | 作为 `BackupWrite` 的读时 `read/verify` 子路径，读取 `backup_snapshot` 并校验 manifest / 覆盖矩阵 / schema/version |

### 3.4 职责边界

- Backup 模块 SHALL 只读 canonical 表，不写入任何 canonical 数据
- Backup 模块 SHALL 通过自有 `candidate/asset_reader.go` 读取跨模块数据，不依赖 canonical 模块的 service
- Backup 模块 SHALL 在 `BackupWrite` owner 内包含正式读时 `read/verify` 子路径，不得以 `BackupWrite` 响应附带代替 snapshot 读取侧
- Backup 模块 SHALL NOT 承接导出职责
- Backup 模块 SHALL NOT 在当前阶段执行真正 restore 写回

### 3.5 `backup_snapshot` 读取侧独立性

- `backup_snapshot` 的读取侧必须由 `BackupWrite` owner 内的正式 `read/verify` 子路径承接
- `BackupWrite` 的响应只返回备份创建结果摘要（产物标识、创建时间、是否成功）
- `backup_snapshot` 的完整内容（manifest + 覆盖矩阵 + schema/version）必须通过该 `read/verify` 子路径独立读取
- phase06-05 已冻结："不得把 `BackupWrite` 响应中附带了一次 `manifest` 解释为已满足可重新读取要求"

---

## 4. 数据装配责任边界与覆盖矩阵

### 4.1 Export 覆盖矩阵与装配 owner

| 资产数据 | 来源表 | 装配 Owner | 装配方式 |
| --- | --- | --- | --- |
| `products` | `products` | `export.candidate.AssetReader` | 只读 SQL 查询 |
| `modules` | `modules` | `export.candidate.AssetReader` | 只读 SQL 查询 |
| `releases` | `releases` | `export.candidate.AssetReader` | 只读 SQL 查询 |
| `repositories` | `repositories` | `export.candidate.AssetReader` | 只读 SQL 查询 |
| `decisions` | `decisions` | `export.candidate.AssetReader` | 只读 SQL 查询 |
| `decision_links` | `decision_links` | `export.candidate.AssetReader` | 只读 SQL 查询 |
| `product_modules` | `product_modules` | `export.candidate.AssetReader` | 只读 SQL 查询 |
| `product_repositories` | `product_repositories` | `export.candidate.AssetReader` | 只读 SQL 查询 |
| `module_repositories` | `module_repositories` | `export.candidate.AssetReader` | 只读 SQL 查询 |

### 4.2 Backup 覆盖矩阵与装配 owner

| 资产数据 | 来源表 | 装配 Owner | 装配方式 |
| --- | --- | --- | --- |
| 上述 9 类核心资产 | 同 Export | `backup.candidate.AssetReader` | 只读 SQL 查询（同覆盖范围） |
| `manifest` | 备份元数据 | `backup.service.CommandService` | 服务层组装 |
| 备份创建时间 | 系统时间 | `backup.service.CommandService` | 服务层生成 |
| `schema / version` | 当前实例 schema 版本 | `backup.service.CommandService` | 服务层读取 |

### 4.3 装配约束

- `export.candidate.AssetReader` 和 `backup.candidate.AssetReader` 各自独立拥有，分别读取相同覆盖范围的 canonical 表
- 两个 reader 遵循"消费方模块拥有 candidate reader"模式，不依赖 canonical 模块的 service
- candidate reader 直接使用 `*pgxpool.Pool` 查询 canonical 表，service 层不直接写跨模块 SQL
- 覆盖矩阵中的 9 类数据必须全部装配，不得只导出主实体而缺失关系数据

---

## 5. `backup verified` 最小后端校验链与 owner 归属

### 5.1 校验链定义

| 步骤 | 校验内容 | Owner 接口 | Owner 模块 | 成功条件 |
| --- | --- | --- | --- | --- |
| 1. 备份产物生成 | 装配资产 + 生成 manifest + 写入产物 | `BackupWrite.command` | backup.command_service | 产物已生成且可标识 |
| 2. snapshot 可重新读取 | 读取 backup_snapshot | `BackupWrite.read_verify` | backup.query_service | snapshot 可读取且可解析 |
| 3. manifest 可解析 | 解析 manifest 内容 | `BackupWrite.read_verify` | backup.query_service | manifest 结构完整且可解析 |
| 4. 覆盖矩阵可核对 | 校验 manifest 中的覆盖矩阵 | `BackupWrite.read_verify` | backup.query_service | 9 类核心资产全部在 manifest 中列出 |
| 5. schema/version 可校验 | 校验 manifest 中的 schema/version | `BackupWrite.read_verify` | backup.query_service | schema/version 字段存在且可校验 |

### 5.2 校验链 owner 归属

- 步骤 1（"备份产物生成成功"）由 `BackupWrite` 承接
- 步骤 2-5（"恢复前提已校验"）由 `BackupWrite` owner 内的 `read/verify` 子路径承接
- `backup verified` 的正式读取判定必须留在 `BackupWrite` owner 语义内，不得再长出第二套独立 canonical owner
- `BackupWrite` 的响应不得包含 `backup verified` 判定结果（因为步骤 2-5 尚未执行）

### 5.3 校验链约束

- 只有步骤 1-5 全部成功，当前阶段才允许判定 `backup verified` 成立
- 仅步骤 1 成功（文件写出成功）不等于 `backup verified`
- `BackupWrite` 的 `read/verify` 子路径必须能独立执行步骤 2-5，不依赖 `BackupWrite.command` 的响应数据
- 步骤 2-5 的校验失败必须分别归类（manifest 缺失 / 覆盖不完整 / schema 缺失），不得统一收敛为笼统"校验失败"

---

## 6. 与既有服务、脚本、数据库的边界关系

### 6.1 与 canonical 模块的边界

| 边界 | 约束 |
| --- | --- |
| Export / Backup → canonical 表 | 只读，通过各自 candidate reader 查询，不修改 canonical 数据 |
| Export / Backup → canonical service | 不依赖，不注入 canonical 模块的 QueryService / CommandService |
| canonical 模块 → Export / Backup | 无反向依赖，canonical 模块不感知 Export / Backup 的存在 |

### 6.2 与 platform 装配点的边界

| 装配函数 | 职责 | 装配顺序约束 |
| --- | --- | --- |
| `buildExport(pool)` | 构造 export 的 query / command service，返回供 mount 使用 | 必须在 canonical 模块建表后调用 |
| `mountExport(r, querySvc, commandSvc)` | 注册 `/api/dashboard/export` 路由 | 无特殊顺序约束 |
| `buildBackup(pool)` | 构造 backup 的 query / command service，返回供 mount 使用 | 必须在 canonical 模块建表后调用 |
| `mountBackup(r, querySvc, commandSvc)` | 注册 `/api/dashboard/backup` 路由 | 无特殊顺序约束 |

> 约束：Export / Backup 模块的 candidate reader 直接使用 `*pgxpool.Pool`，不需要 canonical 模块的 service 注入。装配顺序只需保证 canonical 表已建表（由迁移脚本保证），不需要 canonical service 已构造。

### 6.3 与数据库脚本的边界

| 脚本类型 | 边界关系 |
| --- | --- |
| 既有 migrations | 不修改，Export 不引入新表；Backup 是否引入备份元数据表由后续实现决定，当前不冻结 |
| 既有 seeds | 不修改，不依赖特定 seed 数据 |
| 既有 reset / fixture 脚本 | 不修改，但验收 fixture 需要覆盖 Export / Backup 场景 |

### 6.4 不冻结的细节

- 不提前冻结脚本名（如备份脚本、恢复脚本的具体命名）
- 不提前冻结存储介质（文件系统路径、内存、数据库表等）
- 不提前冻结最终目录细节（备份产物存放路径、导出文件格式）
- 不提前冻结备份产物的文件格式（JSON、tar、zip 等）

---

## 7. `.proto` 合同分组与接口 owner

### 7.1 合同文件分组

| .proto 文件 | 接口 | 消息类型 | Owner 模块 |
| --- | --- | --- | --- |
| `proto/psco/export/v1/export.proto` | `ExportRead` / `ExportWrite` | ExportReadRequest / ExportReadResponse / ExportSnapshot / ExportWriteRequest / ExportWriteResponse / ExportResult / ExportCoverage | export |
| `proto/psco/backup/v1/backup.proto` | `BackupWrite` / `BackupWrite.read_verify` | BackupWriteRequest / BackupWriteResponse / ReadBackupSnapshotRequest / ReadBackupSnapshotResponse / BackupSnapshot / BackupManifest / BackupCoverage / SchemaVersion | backup |

### 7.2 合同约束

- 新增 `.proto` 必须复用既有 `proto/Makefile` 与 `proto/buf.yaml / buf.gen.yaml` 入口
- 不得新增第二套 `buf.yaml`、`buf.gen.yaml` 或并列 proto 根目录
- `make build / gen / lint / breaking` 必须对新增 .proto 保持可运行
- HTTP DTO（`types.go`）必须从 `.proto` 单向派生，不得新增 `.proto` 中不存在的业务字段
- `ExportRead` 的响应消息必须包含 `export_snapshot`（资产范围 + 创建时间 + 创建结果摘要或空态/预览态），且与 `.proto` 单值一致
- `BackupWrite.read_verify` 的响应消息必须包含 manifest + 覆盖矩阵 + schema/version，且与 `.proto` 单值一致

### 7.3 `backup_snapshot` 合同约束

- `BackupWrite` 响应消息只包含创建结果摘要（产物标识、创建时间、成功状态）
- `BackupWrite.read_verify` 响应消息包含完整 `backup_snapshot`（manifest + 覆盖矩阵 + schema/version + 校验结果）
- `backup_snapshot` 的字段必须从 `.proto` 单向承接到 HTTP DTO 与前端消费模型
- 不得在 HTTP 层或前端类型中补出与 `.proto` 不一致的第二套字段语义

---

## 8. 模块间依赖关系图

```
                    ┌─────────────────────────────────┐
                    │       platform/router.go         │
                    │  (组合根：装配所有模块到 chi)     │
                    └──────┬──────────┬───────────────┘
                           │          │
           ┌───────────────┤          │
           │               │          │
           ▼               ▼          ▼
    ┌─────────────┐ ┌───────────┐ ┌───────────┐
    │   Export    │ │  Backup   │ │ Canonical │
    │   Module    │ │  Module   │ │  Modules  │
    │             │ │           │ │ (MR/DC/   │
    │ handler/    │ │ handler/  │ │  PR/RB)   │
    │ service/    │ │ service/  │ │           │
    │ candidate/  │ │ candidate/│ │           │
    │   └ reader ─┼─┤   └ reader├─┤ canonical │
    │     (只读)  │ │     (只读)│ │  tables   │
    └─────────────┘ └───────────┘ └───────────┘
```

- Export / Backup 模块通过 candidate reader 只读 canonical 表
- Export / Backup 模块之间无直接依赖
- Export / Backup 模块不依赖 canonical 模块的 service 层
- canonical 模块不感知 Export / Backup 模块的存在

---

## ADDED Requirements

### Requirement: Export 与 Backup 后端模块独立冻结

系统 SHALL 将 `Export` 与 `Backup` 冻结为两个独立的后端模块，各自拥有独立的 handler / service / candidate 层，不合并为同一模块。

#### Scenario: 模块独立性

- **WHEN** 接手者实现 phase06 导出 / 备份后端
- **THEN** 必须新增 `backend/internal/export/` 与 `backend/internal/backup/` 两个独立模块
- **AND** 两个模块各自拥有独立的 `handler/`、`service/`、`candidate/` 子包
- **AND** 两个模块之间不得存在直接代码依赖
- **AND** 不得把 Export / Backup 合并为同一个模块或同一个 service

### Requirement: Export 模块接口分组与 owner 冻结

系统 SHALL 将 `ExportRead` 与 `ExportWrite` 冻结为 Export 模块的正式接口，并把 `ExportRead` 收为 `export_snapshot` 的正式读取 owner。

#### Scenario: ExportRead owner

- **WHEN** 接手者实现 `ExportRead`
- **THEN** 接口必须由 `export.handler.QueryHandler` → `export.service.QueryService` 承接
- **AND** HTTP 路由必须为 `GET /api/dashboard/export`
- **AND** 返回内容必须承接 `export_snapshot`
- **AND** `export_snapshot` 至少必须包含当前可导出的资产范围、最近一次导出创建时间、最近一次导出创建结果摘要
- **AND** 若当前尚无历史导出结果，必须返回预览态或空态 `export_snapshot`，而不是把读模型留空

#### Scenario: ExportWrite owner

- **WHEN** 接手者实现 `ExportWrite`
- **THEN** 接口必须由 `export.handler.CommandHandler` → `export.service.CommandService` 承接
- **AND** HTTP 路由必须为 `POST /api/dashboard/export`
- **AND** 装配逻辑必须通过 `export.candidate.AssetReader` 读取 9 类 canonical 表
- **AND** 当前阶段不要求冻结独立 export history 表或最终元数据介质
- **AND** 但 `ExportWrite` 产出的创建时间与结果摘要必须能被 `ExportRead` 重新承接为 `export_snapshot`

### Requirement: Backup 模块接口分组与 owner 冻结

系统 SHALL 将 `BackupWrite` 冻结为 Backup 模块的唯一正式 owner，并在该 owner 内同时承接 command 子路径与读时 `read/verify` 子路径。

#### Scenario: BackupWrite owner

- **WHEN** 接手者实现 `BackupWrite`
- **THEN** 接口必须由 `backup.handler.CommandHandler` → `backup.service.CommandService` 承接
- **AND** HTTP 路由必须为 `POST /api/dashboard/backup`
- **AND** 装配逻辑必须通过 `backup.candidate.AssetReader` 读取 9 类 canonical 表
- **AND** 响应只返回创建结果摘要，不包含完整 `backup_snapshot`

#### Scenario: `BackupWrite.read_verify` 子路径

- **WHEN** 接手者实现 `BackupWrite` 的读时 `read/verify` 子路径
- **THEN** 接口必须由 `backup.handler.QueryHandler` → `backup.service.QueryService` 承接
- **AND** HTTP 路由必须为 `GET /api/dashboard/backup`
- **AND** 该子路径仍属于 `BackupWrite` 的正式 owner 语义
- **AND** 必须独立读取 `backup_snapshot`，不依赖 `BackupWrite.command` 的响应数据
- **AND** 返回内容必须包含 manifest + 覆盖矩阵 + schema/version + 校验结果

### Requirement: 数据装配 candidate reader owner 冻结

系统 SHALL 将导出 / 备份的跨模块数据读取冻结为各自模块的 candidate reader，遵循"消费方模块拥有 candidate reader"模式。

#### Scenario: Export candidate reader

- **WHEN** 接手者实现 Export 数据装配
- **THEN** 必须在 `export/candidate/` 子包中定义 `AssetReader`
- **AND** `AssetReader` 直接使用 `*pgxpool.Pool` 查询 canonical 表
- **AND** service 层通过 `AssetReader` 接口读取，不直接写跨模块 SQL
- **AND** 不得注入 canonical 模块的 QueryService / CommandService

#### Scenario: Backup candidate reader

- **WHEN** 接手者实现 Backup 数据装配
- **THEN** 必须在 `backup/candidate/` 子包中定义 `AssetReader`
- **AND** `AssetReader` 覆盖范围与 Export 相同（9 类 canonical 表）
- **AND** Backup 的 manifest / 创建时间 / schema/version 由 `backup.service.CommandService` 组装，不在 candidate reader 中生成

#### Scenario: 覆盖矩阵完整性

- **WHEN** 接手者装配导出 / 备份数据
- **THEN** 9 类核心资产（products / modules / releases / repositories / decisions / decision_links / product_modules / product_repositories / module_repositories）必须全部读取
- **AND** 不得只读取主实体而缺失关系数据

### Requirement: `backup verified` 校验链与 owner 归属冻结

系统 SHALL 将 `backup verified` 的最小后端校验链冻结为 5 个步骤，前 1 步由 `BackupWrite.command` 承接，后 4 步由 `BackupWrite.read_verify` 承接。

#### Scenario: 备份产物生成成功

- **WHEN** 系统执行 `BackupWrite`
- **THEN** `backup.service.CommandService` 必须装配核心资产 + 生成 manifest + 记录 schema/version
- **AND** 产物生成成功后返回创建结果摘要
- **AND** 此步骤成功不等同于 `backup verified`

#### Scenario: 恢复前提已校验

- **WHEN** 系统执行 `BackupWrite.read_verify` 校验
- **THEN** `backup.service.QueryService` 必须独立读取 `backup_snapshot`
- **AND** 必须校验 manifest 可解析
- **AND** 必须校验覆盖矩阵中 9 类核心资产全部列出
- **AND** 必须校验 schema/version 字段存在且可校验
- **AND** 只有 4 项校验全部通过才允许判定 `backup verified`

#### Scenario: 校验失败分类

- **WHEN** `BackupWrite.read_verify` 校验失败
- **THEN** 必须能区分以下失败类型：manifest 缺失或不可解析 / 覆盖矩阵缺失或不完整 / schema/version 缺失或不可校验
- **AND** 不得把所有校验失败统一收敛为笼统"校验失败"

### Requirement: `backup_snapshot` 读取侧独立性冻结

系统 SHALL 将 `backup_snapshot` 的读取侧冻结为 `BackupWrite` owner 内的正式 `read/verify` 子路径承接，不得以 `BackupWrite` 响应附带代替。

#### Scenario: 读取侧独立

- **WHEN** 接手者实现 `backup_snapshot` 读取
- **THEN** `BackupWrite.read_verify` 必须能独立读取 snapshot（不依赖 `BackupWrite.command` 响应中的数据）
- **AND** `BackupWrite` 响应不得包含完整 `backup_snapshot`（只返回创建结果摘要）
- **AND** `BackupWrite.read_verify` 的 `.proto` 合同、HTTP DTO 与前端消费模型必须保持单值一致

### Requirement: platform 装配点冻结

系统 SHALL 在 `platform/router.go` 中新增 `buildExport / mountExport / buildBackup / mountBackup` 四个装配函数。

#### Scenario: Export 装配

- **WHEN** 接手者在 platform 中装配 Export 模块
- **THEN** 必须通过 `buildExport(pool)` 构造 query / command service
- **THEN** 必须通过 `mountExport(r, querySvc, commandSvc)` 注册路由
- **AND** 装配顺序只需保证 canonical 表已建表，不需要 canonical service 已构造

#### Scenario: Backup 装配

- **WHEN** 接手者在 platform 中装配 Backup 模块
- **THEN** 必须通过 `buildBackup(pool)` 构造 query / command service
- **THEN** 必须通过 `mountBackup(r, querySvc, commandSvc)` 注册路由
- **AND** 装配顺序只需保证 canonical 表已建表，不需要 canonical service 已构造

### Requirement: `.proto` 合同分组冻结

系统 SHALL 将导出 / 备份的 `.proto` 合同冻结为两个独立文件，各自承接对应模块的接口消息。

#### Scenario: Export proto

- **WHEN** 接手者定义 Export 合同
- **THEN** 必须新增 `proto/psco/export/v1/export.proto`
- **AND** 必须包含 `ExportRead` 与 `ExportWrite` 的请求 / 响应消息
- **AND** 必须包含 `ExportSnapshot` 或与其单值等价的正式读模型消息
- **AND** 必须复用既有 `proto/Makefile` 与 `proto/buf.yaml / buf.gen.yaml` 入口

#### Scenario: Backup proto

- **WHEN** 接手者定义 Backup 合同
- **THEN** 必须新增 `proto/psco/backup/v1/backup.proto`
- **AND** 必须包含 `BackupWrite` 的写入请求 / 响应消息，以及 `BackupWrite.read_verify` 子路径对应的读取请求 / 响应消息
- **AND** 必须包含 `BackupManifest` / `BackupCoverage` / `SchemaVersion` 消息类型
- **AND** `BackupWrite.read_verify` 响应必须包含完整 `backup_snapshot`（manifest + 覆盖矩阵 + schema/version + 校验结果）

### Requirement: 与既有 canonical 模块边界冻结

系统 SHALL 保证 Export / Backup 模块不修改 canonical 数据、不依赖 canonical service。

#### Scenario: 只读边界

- **WHEN** 接手者实现 Export / Backup
- **THEN** 两个模块对 canonical 表只允许只读查询
- **AND** 不得注入 canonical 模块的 QueryService / CommandService
- **AND** canonical 模块不感知 Export / Backup 的存在

#### Scenario: 不冻结脚本名与存储介质

- **WHEN** 接手者设计备份存储方案
- **THEN** 当前阶段不提前冻结脚本名、存储介质与最终目录细节
- **AND** 但 `backup verified` 校验链的 owner 归属必须已冻结
- **AND** 后续实现选择存储方案时不得改变校验链 owner 归属

## MODIFIED Requirements

### Requirement: platform 装配层模块清单

`platform/router.go` 在 `phase06-08` 中 SHALL 在既有 5 个模块（Module Registry / Decision Center / Product Registry / Repository Binding / Dashboard）之外，新增 Export 与 Backup 两个模块的装配函数。

#### Scenario: 装配函数清单

- **WHEN** 接手者在 platform 中注册模块
- **THEN** 必须新增 `buildExport / mountExport / buildBackup / mountBackup` 四个函数
- **AND** 既有 5 个模块的装配函数不得改变
- **AND** Export / Backup 的路由必须注册到 `/api` 子路由器下

## REMOVED Requirements

### Requirement: 导出 / 备份可由 canonical 模块的 service 直接承接

**Reason**: 让 canonical 模块的 service 直接承接导出 / 备份职责会破坏模块边界，导致 canonical 模块承担非自身职责的数据装配与产物生成，且无法满足"消费方模块拥有 candidate reader"的跨模块读取约束。

**Migration**: phase06 后续实现统一改为：新增独立的 `export` 与 `backup` 模块，各自拥有 candidate reader 读取 canonical 表，service 层只承接导出 / 备份的业务编排。
