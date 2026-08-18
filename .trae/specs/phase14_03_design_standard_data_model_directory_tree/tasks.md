# Tasks

- [x] Task 1: 建立 phase14-03 数据模型与目录树设计 spec 工件
  - [x] SubTask 1.1: 冻结三表 DDL 级设计草案（`standards` / `standard_revisions` / `standard_bindings`；风格对齐 0010 迁移：注释块头 + 字段级行注释 + `gen_random_uuid()` / `TEXT+CHECK` / `TIMESTAMPTZ` / 内联 UNIQUE / `idx_` 前缀索引）
    - 执行记录：DDL 草案冻结（spec §ADDED-1）；关键决策留档——`directory_tree JSONB NOT NULL`（无树状态不存在，创建即单根空树）、`target_id` 无 DB 外键（多态必然，存在性由 service 层校验）、revisions 只追加不更新、bindings 唯一约束四元组 + 双索引（brief 反查 + standard 正查）。
  - [x] SubTask 1.2: 冻结 `DirectoryTreeNode` jsonb 序列化规范（Go 结构 + json tag + omitempty 语义 + children nil/[] 等价 + 根节点表示决策 + proto 字段对齐表 + 示例）
    - 执行记录：规范冻结（spec §ADDED-2）；根节点表示决策为单根 `name="."`（仓库根语义）；示例标注已修正——phase13-11 固定样本 entry_ref 实证均为根级路径，真实迁移产物全挂根下，拆树分支为通用规则前向演示。
  - [x] SubTask 1.3: 形式化 8 条树校验规则（每条：判定逻辑 + 稳定错误码 + 节点路径定位；字段级细化 role 1-32 / summary ≤2000）
    - 执行记录：R1-R8 表格化冻结（spec §ADDED-3）；校验入口为 Create/Update（含整树替换与状态变更）service 层 validate；role 约定值域为软约定（硬校验仅长度）；name 中 `.` / `..` 不赋予特殊路径语义。
  - [x] SubTask 1.4: 冻结旧画像两表 → 树节点字段级迁移映射（源 1 canonical `file_name/role/required`；源 2 global assets `name/kind/entry_ref/role/structured_summary`；required 显式退役、kind 并入 summary 前缀、entry_ref 规范化三态规则、同名冲突合并、零丢失验收范围）
    - 执行记录：字段级映射冻结（spec §ADDED-4）；迁移顺序固定先源 1 后源 2；冲突合并规则为源 2 非空值增强既有节点；裁决⑧多画像合并口径沿用 phase14-02；AGENTS.md 双表并存断言经 phase13-11 验收报告 L34/L35 实证。

- [x] Task 2: 一致性核对
  - [x] SubTask 2.1: 与 shared_baseline §3.2/§3.3/§3.4 单值一致（字段矩阵、节点矩阵、8 条规则、绑定矩阵）
    - 执行记录：逐项核对通过——三表字段与 §3.2 一致（revisions 的 id/standard_id/created_at 为 DDL 级必要补充）；节点 6 字段与 §3.3 一致；R1-R8 与 §3.3 八条规则一一对应（R2 按 file 计数、R5 第 6 层限 file、R8 路径自洽性三处收紧性细化已显式列入 MODIFIED 声明）；bindings 枚举/唯一约束/扩展规则与 §3.4 一致。
  - [x] SubTask 2.2: 与 phase14-02 零丢失映射与显式退役口径一致
    - 执行记录：5 字段映射（name→节点名 / entry_ref→ref / role→role / structured_summary→summary / kind 并入摘要语义）逐字对齐；kind 前缀格式 `"[" + kind + "] " + structured_summary` 为 phase14-02 冻结口径的单值化落地；required 退役理由与 markdown_resolvable 同模式（结构性吸收 + 不进零丢失验收），语义一致。
  - [x] SubTask 2.3: 与 0010 迁移 DDL 风格及旧两表字段实证一致
    - 执行记录：0010 全文读取实证——旧两表字段为 `file_name/role/required` 与 `name/kind/entry_ref/role/structured_summary`，无 `external_url` 列（spec 断言成立）；DDL 草案风格（注释块头 / 字段级行注释 / CHECK 枚举 / idx_ 前缀）对齐 0010。
  - [x] SubTask 2.4: dev_plan L29-32 范围与 DoD 逐条满足
    - 执行记录：范围三件（三表字段级设计 / DirectoryTreeNode jsonb 规范 / 8 条校验规则完整落地）全部覆盖；DoD 三条满足——字段矩阵与校验规则可直接进入 0011 迁移与 proto 实现（DDL 草案 + proto 对齐表）、双清单合一落为树节点单一承载（两源合并一棵树）、ref 承接 entry_ref 语义无丢失（规范化三态规则 + URL 前向能力承接 external_url 语义定位）。
  - [x] SubTask 2.5: 零代码改动、零三件套正文改动、零根级改动（git status 验证）
    - 执行记录：git status 验证本子任务仅新增 phase14_03 目录，无其他变更。

- [x] Task 3: 独立复核与收口
  - [x] SubTask 3.1: 子代理独立复核（DDL 草案风格与约束完备性、8 条规则可测试性、迁移映射闭环、示例标注准确性、与上游零漂移）
    - 执行记录（2026-08-18 独立复核代理）：复核证据链覆盖 A-F 六项（DDL 草案 L31-76 / 8 条规则 L137-161 / 迁移映射 L162-203 / jsonb 规范 L91-112 / MODIFIED 声明 L205-212 与上游 0010、phase14-02、baseline §3.3、dev_plan L29-32、phase13-11 验收报告 L34-35 逐一对上），无 FAIL 项。主代理闭环特别核查点：①entry_ref 根级路径断言与 phase13-11 验收报告 L34-L35 一致（复核代理确认）；②源表 created_at 去向本 spec 未显式声明——已修复（迁移映射节补充时间戳列口径，引用 phase14-02 显式不迁移清单：created_at+updated_at 不迁移、目标时间戳由迁移写入时刻生成、首条 revision 承接合并来源）；③非根 directory 的 role 非空为设计意图（baseline"directory 可空"= 允许为空而非必空，spec 与 baseline 一致无需修改）。最终结论：PASS。
  - [x] SubTask 3.2: 复核通过后勾选 checklist 收口；变更保持未提交，待用户最终确认后手动提交
    - 执行记录：checklist.md 全部勾选；变更未提交（phase14_03 目录 untracked），待用户确认后手动提交。

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 2
