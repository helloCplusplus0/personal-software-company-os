# phase14-03 Checklist

## 三表 DDL 级设计草案

- [x] `standards` 主表：全局作用域无 repository 锚点（裁决⑤）；name 全局唯一；status 受控枚举 draft/active/retired；directory_tree JSONB NOT NULL（无树状态不存在）
- [x] `standard_revisions`：只追加不更新；change_summary 必填；ON DELETE CASCADE + standard_id 索引
- [x] `standard_bindings`：target_type 4 枚举可扩展（CON-02 不设限）；target_id 无 DB 外键（多态必然，存在性 service 层校验）；role 2 枚举可扩展（组合合法性按 phase14-02 八格矩阵）；唯一约束四元组；brief 反查索引 (target_type, target_id) + standard_id 正查索引
- [x] DDL 风格对齐 0010（注释块头 / 字段级行注释 / gen_random_uuid / TEXT+CHECK / TIMESTAMPTZ / 内联 UNIQUE / idx_ 前缀）

## DirectoryTreeNode jsonb 序列化规范

- [x] Go 结构 6 字段与 baseline §3.3 节点矩阵一一对应；json tag 与 omitempty 语义冻结（role/summary/ref 空省略、缺失视为空串）
- [x] children nil/[] 等价规则冻结（file 与无子 directory 均允许两种形态）
- [x] 根节点表示决策冻结：单根 directory、name="."、role 必空、summary 可选、不参与同层唯一性
- [x] proto 字段对齐表冻结（phase14-04 逐字引用：NODE_TYPE_DIRECTORY=1 / NODE_TYPE_FILE=2）
- [x] 示例标注准确：真实迁移产物全挂根下（phase13-11 样本 entry_ref 实证均为根级路径），拆树分支为通用规则前向演示

## 8 条树校验规则形式化

- [x] R1-R8 逐条有判定逻辑 + 稳定错误码 + 节点路径定位（自根起 / 连接）
- [x] 与 baseline §3.3 八条规则一一对应；三处收紧性细化（R2 file 计数 / R5 第 6 层限 file / R8 路径自洽性）显式列入 MODIFIED 声明，只收紧不放宽
- [x] 校验入口冻结：Create / Update（含整树替换与状态变更）service 层 validate
- [x] 字段级细化冻结：role 1-32（软约定值域不硬卡）、summary ≤2000、name `^[A-Za-z0-9._-]{1,64}$`、序列化 ≤65536 字节
- [x] 每条规则可直接写单元测试（含非法用例与合法边界用例）

## 迁移映射冻结

- [x] 源 1（canonical_root_file_bindings）字段级映射：file_name→name、role 原值保留、required 显式退役（结构性吸收，同 markdown_resolvable 模式）、ref 派生 `/`+file_name
- [x] 源 2（governance_global_asset_bindings）字段级映射：kind 并入 summary 前缀格式（`"[" + kind + "] " + structured_summary`，对齐 phase14-02 冻结口径）、role 原值保留
- [x] entry_ref 规范化三态规则：https:// 原值挂根下 / 裸名 `/` 前缀挂根下 / 含路径拆树（directory 段生成或复用 + 末段 file）
- [x] 0010 无 external_url 列实证留档；ref 的 URL 能力为前向能力承接其语义定位
- [x] 同名冲突合并规则：迁移顺序先源 1 后源 2；源 2 非空值增强既有节点（role/summary 覆盖），空值保留源 1
- [x] 裁决⑧多画像合并口径沿用（当前仅 1 实例；未来多画像取最新 updated_at，差异记首条 revision）
- [x] 零丢失验收范围显式冻结：5 字段映射 + 源 1 的 file_name/role；required 与 markdown_resolvable 为显式退役项不在范围

## 一致性核对

- [x] 与 shared_baseline §3.2/§3.3/§3.4 单值一致（细化全部只收紧不放宽）
- [x] 与 phase14-02 零丢失映射 / 显式退役口径逐字一致
- [x] 与 0010 DDL 风格及旧两表字段实证一致（含 AGENTS.md 双表并存断言经 phase13-11 验收报告 L34/L35 实证）
- [x] dev_plan L29-32 范围与 DoD 三条逐条满足
- [x] 零代码改动、零三件套正文改动、零根级改动（git status 验证）

## 复核与收口

- [x] 子代理独立复核通过（发现项已修复回填 spec）
- [x] tasks.md 全部勾选并附执行记录
- [x] 变更未提交，待用户最终确认后手动提交
