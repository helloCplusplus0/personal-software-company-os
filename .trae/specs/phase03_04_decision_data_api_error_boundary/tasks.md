# Tasks

- [x] Task 1: 冻结 `Decision Center` 当前阶段的数据读写范围。将当前主线真正需要的读写对象与动作写成单值结论。
  - [x] SubTask 1.1: 明确当前阶段写入范围只承接 `RecordDecision` 与 `LinkDecisionToTarget`
  - [x] SubTask 1.2: 明确当前阶段读取范围只承接列表读取、详情读取、`Decision -> Module` 候选读取与创建后的最小回流读取
  - [x] SubTask 1.3: 明确读取前提只包含 `decisions`、`decision_links` 与 `modules`
  - [x] SubTask 1.4: 验证当前阶段未引入 `Decision -> Product / Repository` 的写入主线

- [x] Task 2: 冻结最小接口承接前提。将当前页面主线与接口承接关系写清。
  - [x] SubTask 2.1: 明确 `Decision Center / List` 承接列表读取接口与进入详情入口
  - [x] SubTask 2.2: 明确 `Decision Create` 承接 `RecordDecision` 创建写入接口
  - [x] SubTask 2.3: 明确 `Decision Detail` 承接详情读取、候选读取与 `LinkDecisionToTarget` 关联写入接口
  - [x] SubTask 2.4: 验证未把当前动作拆成第二套并列业务工作台

- [x] Task 3: 冻结读动作与写动作的最小接口分组。将接口分组控制在当前阶段最小可执行范围内。
  - [x] SubTask 3.1: 明确 `DecisionListRead`、`DecisionDetailRead`、`DecisionModuleCandidateRead` 的最小读接口分组
  - [x] SubTask 3.2: 明确 `DecisionWrite` 与 `DecisionLinkWrite` 的最小写接口分组
  - [x] SubTask 3.3: 验证未提前冻结完整聚合查询、趋势分析查询或跨页面反馈接口
  - [x] SubTask 3.4: 验证未引入审批、投票、批量变更或自动化写入接口

- [x] Task 4: 冻结 `RecordDecision` 的请求校验与失败语义归属。将创建动作的关键异常路径写成单值结论。
  - [x] SubTask 4.1: 明确必填字段缺失返回明确校验失败语义
  - [x] SubTask 4.2: 明确 `status` 越界或去空白后为空返回明确校验失败语义
  - [x] SubTask 4.3: 明确这些错误归属落在 `DecisionWrite`
  - [x] SubTask 4.4: 验证未把创建校验错误错误映射为资源不存在或重复冲突

- [x] Task 5: 冻结 `Decision -> Module` 候选读取与 `LinkDecisionToTarget` 的错误语义前提。将读取错误与关联写入错误分开写清。
  - [x] SubTask 5.1: 明确候选读取空结果返回空列表语义，不映射为资源不存在
  - [x] SubTask 5.2: 明确 `Decision` 本身不存在时返回资源不存在语义
  - [x] SubTask 5.3: 明确目标类型越界返回明确校验失败语义
  - [x] SubTask 5.4: 明确目标 `Decision` 或 `Module` 不存在时返回资源不存在语义
  - [x] SubTask 5.5: 明确重复关联返回重复冲突语义
  - [x] SubTask 5.6: 明确这些错误归属落在 `DecisionDetailRead` 或 `DecisionLinkWrite` 对应边界

- [x] Task 6: 完成关键异常路径与非目标接口校验。确认本次 `phase03-04` 规格可直接作为后续接口、合同与验收设计的上游。
  - [x] SubTask 6.1: 验证 `RecordDecision` 与 `LinkDecisionToTarget` 的关键异常路径已进入阶段规划
  - [x] SubTask 6.2: 验证未把关键异常路径后移到联调阶段再补定义
  - [x] SubTask 6.3: 验证未提前冻结 `pending_decision_signals`、`Dashboard` 或聚合分析接口
  - [x] SubTask 6.4: 验证本次规格与 `Contract First`、`phase03` 三件套及 `phase03-01 ~ 03` 规格保持一致

# Task Dependencies

- `Task 2` depends on `Task 1`
- `Task 3` depends on `Task 1` and `Task 2`
- `Task 4` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 5` depends on `Task 1`, `Task 2`, and `Task 3`
- `Task 6` depends on `Task 1`, `Task 2`, `Task 3`, `Task 4`, and `Task 5`
