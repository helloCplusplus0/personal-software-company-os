# Tasks

- [x] Task 1: 冻结 `Decision` 最小结构化模板。将当前阶段唯一允许的记录字段写成单值结论，避免表单、DTO 与合同层继续漂移。
  - [x] SubTask 1.1: 明确最小字段集合为 `title / context / problem / alternatives / choice / reason / impact / status`
  - [x] SubTask 1.2: 明确当前阶段不引入复杂审批、投票或自动化字段

- [x] Task 2: 冻结字段级 `required / optional` 与创建校验前提。把必填字段、空字符串规则与非法输入判定写成可直接实现的最小规则。
  - [x] SubTask 2.1: 明确 `title / context / problem / choice / reason / status` 为必填
  - [x] SubTask 2.2: 明确 `alternatives / impact` 为可选
  - [x] SubTask 2.3: 明确必填字段去首尾空白后不得为空字符串
  - [x] SubTask 2.4: 明确非法 `status` 与非法 `alternatives` 条目必须返回明确校验失败

- [x] Task 3: 冻结 `alternatives` 的最小结构。确保前后端和合同层都采用同一解释，不各自发明嵌套模型。
  - [x] SubTask 3.1: 明确 `alternatives` 为按顺序保留的文本条目集合
  - [x] SubTask 3.2: 明确允许空集合
  - [x] SubTask 3.3: 明确单个条目去首尾空白后不得为空字符串
  - [x] SubTask 3.4: 明确不允许扩写为嵌套对象结构

- [x] Task 4: 冻结 `Decision.status` 的最小枚举与状态语义。把每个状态的当前阶段含义写成前后端、页面与合同都可复用的单值结论。
  - [x] SubTask 4.1: 明确最小状态集合为 `proposed / active / superseded / archived`
  - [x] SubTask 4.2: 明确 `proposed` 的状态语义
  - [x] SubTask 4.3: 明确 `active` 的状态语义
  - [x] SubTask 4.4: 明确 `superseded` 的状态语义
  - [x] SubTask 4.5: 明确 `archived` 的状态语义
  - [x] SubTask 4.6: 明确不引入额外新状态

- [x] Task 5: 冻结 `Decision List` 与 `Decision Detail` 的最小展示模型。将列表与详情至少展示哪些字段、关键字段计算口径、空值语义与最小来源上下文写成可直接进入后续页面设计的读模型。
  - [x] SubTask 5.1: 明确列表页至少展示 `title / status / created_at / link_count / linked_module_summary`
  - [x] SubTask 5.2: 明确 `link_count` 仅统计 `decision_links` 中已建立的 `Decision -> Module` 有效关联数，不混入 `Product / Repository`
  - [x] SubTask 5.3: 明确 `linked_module_summary` 按 `module_name` 升序取前 `3` 个名称，超出 `3` 个时末尾附加 `+N`
  - [x] SubTask 5.4: 明确无关联时 `link_count` 返回 `0`，`linked_module_summary` 返回空字符串，不返回 `null`
  - [x] SubTask 5.5: 明确详情页至少展示 `title / context / problem / alternatives / choice / reason / impact / status / created_at`
  - [x] SubTask 5.6: 明确详情页还必须展示已关联目标结果
  - [x] SubTask 5.7: 明确详情页必须展示最小来源上下文，并区分从 `Module Detail` 带上下文进入与从列表直接进入两种来源
  - [x] SubTask 5.8: 明确来源上下文的具体字段结构与入口上下文冻结由 `phase03-03` 承接，本阶段不提前定义
  - [x] SubTask 5.9: 明确列表页不扩写为完整结构化模板全文展示

- [x] Task 6: 完成规格校验。确认本次 `phase03-02` 规格可以直接作为后续页面、接口与合同设计的上游。
  - [x] SubTask 6.1: 验证 `Decision` 模板字段已经单值化
  - [x] SubTask 6.2: 验证 `required / optional`、`alternatives` 与创建校验规则已经明确
  - [x] SubTask 6.3: 验证 `status` 枚举与状态语义已经明确
  - [x] SubTask 6.4: 验证列表与详情的最小读模型已经明确，包含 `link_count / linked_module_summary` 计算口径、空值语义与最小来源上下文
  - [x] SubTask 6.5: 验证未引入超出 `v0.1` 的复杂审批、投票或自动化字段

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1`
- `Task 4` depends on `Task 1`
- `Task 5` depends on `Task 1`, `Task 2`, `Task 3`, and `Task 4`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
