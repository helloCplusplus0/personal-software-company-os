# Tasks

- [x] Task 1: 冻结 `phase06` 新增接口的唯一合同源与传输边界。把 `.proto`、`chi + HTTP JSON`、HTTP envelope 和 breaking 规则写成单值结论。
  - [x] SubTask 1.1: 明确 `Onboarding / Export / Backup / Reuse Summary` 新增或扩展接口必须先在 `.proto` 中冻结字段、枚举、response envelope 与错误语义
  - [x] SubTask 1.2: 明确 `chi + HTTP JSON` 只承担传输适配职责，不承担第二套合同定义职责
  - [x] SubTask 1.3: 明确 handler 必须把 path / query / body 显式组装为 Proto request 或与之单向对齐的 DTO
  - [x] SubTask 1.4: 明确已冻结为 envelope 的读取接口不得退回为裸对象、裸数组或临时 map
  - [x] SubTask 1.5: 明确 breaking 变更、`reserved` 与 `main#proto` 基准的最小兼容性规则
  - [x] SubTask 1.6: 明确 breaking check 失败必须被显式处理为阻断，不得通过 `continue-on-error` / `allow-failure` 或等价机制降级为 warning

- [x] Task 2: 冻结 `Export / Backup / Reuse Summary` 的 DTO 一致性执行口径。把 `.proto`、HTTP DTO 与前端消费模型之间的映射关系压成验收前提。
  - [x] SubTask 2.1: 明确后端 `types.go` 只承接与 `.proto` 对齐的过渡 JSON DTO
  - [x] SubTask 2.2: 明确前端 `types.ts` 与 `api-adapter.ts` 只能做单向承接或显式字段转换
  - [x] SubTask 2.3: 明确 `Export / Backup` 的覆盖矩阵、`manifest`、`schema / version` 与错误语义不得在 HTTP 层再长第二套口径
  - [x] SubTask 2.4: 明确 `ReuseSummaryRead` 的 `module_reuse_summary / capability_summary` 字段与空态 / 错误语义必须从 `.proto` 单向承接
  - [x] SubTask 2.5: 明确 `OnboardingRead`（含 `first_run_state` 状态枚举与跃迁语义）的 DTO 一致性必须从 `.proto` 单向承接
  - [x] SubTask 2.6: 明确 `backup_snapshot` 读取侧必须由当前阶段 `Backup` 能力中的正式读时承接位负责，不得只由 `BackupWrite` 响应附带，且其 DTO 一致性从 `.proto` 单向承接

- [x] Task 3: 冻结前端 `application / query / mutation / shared` 四条约束在当前阶段的执行口径。避免 phase06 新增写路径继续散落在页面和面板中。
  - [x] SubTask 3.1: 明确 `query` owner 只承接读取、query key、只读解包与 `queryOptions` 级别配置
  - [x] SubTask 3.2: 明确正式 `useMutation`、失效刷新、成功回流、错误归一化统一收敛到切片内固定 `application` 承接位
  - [x] SubTask 3.3: 明确页面、表单、面板组件只保留字段收集、提交事件与局部状态展示职责
  - [x] SubTask 3.4: 明确既有 page-level / panel-level `useMutation` 可作为过渡现实保留，但 `phase06` 新增写路径不得复制这种模式
  - [x] SubTask 3.5: 明确 `shared` 晋升必须延后到跨切片稳定复用后

- [x] Task 4: 冻结 `buf` 工具链入口与阶段验收口径。让合同一致性和源码边界从实现前就进入门禁，而不是验收末尾再补。
  - [x] SubTask 4.1: 明确 `proto/Makefile` 与 `proto/buf.yaml / buf.gen.yaml` 是唯一正式工具链入口
  - [x] SubTask 4.2: 明确当前阶段不得新增第二套 proto 根目录、私有 buf 配置或旁路生成脚本
  - [x] SubTask 4.3: 明确验收时至少覆盖 `buf build / lint / generate / breaking`
  - [x] SubTask 4.4: 明确验收时必须同时检查合同一致性与前端写路径承接位边界
  - [x] SubTask 4.5: 明确验收时必须显式验证 `OnboardingRead`（含 `first_run_state`）的 `.proto -> HTTP DTO -> 前端消费模型` 单值一致，不得只验证 `Export / Backup / Reuse Summary` 而遗漏 `Onboarding`
  - [x] SubTask 4.6: 明确验收时必须显式验证 `backup_snapshot` 读取侧的 `.proto -> HTTP DTO -> 前端消费模型` 单值一致，不得只验证 `BackupWrite` 写入响应而遗漏读取侧

- [x] Task 5: 完成规格一致性校验。验证本次 `phase06-05` 规格与 `phase06` shared baseline、architecture plan、project_rules 以及 phase06-03 / 04 已冻结语义保持一致。
  - [x] SubTask 5.1: 验证本次规格没有重写 `phase06-03` 的 `Export / Backup` 正式语义，只补合同与 DTO 约束
  - [x] SubTask 5.2: 验证本次规格没有重写 `phase06-04` 的 `ReuseSummaryRead` 读模型语义，只补合同、owner 与 DTO 一致性约束
  - [x] SubTask 5.3: 验证本次规格与 `project_rules.md` 中 `.proto` 唯一合同源、HTTP DTO 单向派生、mutation 固定承接位等规则一致
  - [x] SubTask 5.4: 验证本次规格足以支撑后续 `phase06-10` 合同实现、前端写路径回收与阶段验收

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1` and `Task 2`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
