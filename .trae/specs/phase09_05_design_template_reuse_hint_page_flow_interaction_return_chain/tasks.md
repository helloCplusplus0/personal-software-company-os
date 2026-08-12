# Tasks

- [x] Task 1: 对齐 `phase09-05` 的上游冻结结论与页面交互基线
  - [x] SubTask 1.1: 对齐 `phase09-02` 的模板候选、handoff 与成功回流链
  - [x] SubTask 1.2: 对齐 `phase09-03` 的提示矩阵、CTA 与空态语义
  - [x] SubTask 1.3: 对齐 `phase09-04` 的合同、owner、`templateCandidateId` 漂移语义与 caller inventory
  - [x] SubTask 1.4: 参考 `phase08-05` 与 `phase04-06` 的页面流、交互流与返回链写法

- [x] Task 2: 冻结 `Weekly Review -> Product Create -> Product Detail` 的正式页面流
  - [x] SubTask 2.1: 明确 `Weekly Review` 是模板候选与提示的唯一主消费宿主
  - [x] SubTask 2.2: 明确 `Product Create` 只承接模板预填与解释延续，不重新承接候选选择主线
  - [x] SubTask 2.3: 明确 `Product Detail` 只承接模板来源摘要与 canonical binding CTA
  - [x] SubTask 2.4: 明确 `capability_gap_hint` 走 `Module Registry / Module Detail` 补齐支链，而不是进入 `Product Create` 主创建链

- [x] Task 3: 冻结模板候选、提示与 create 预填的关键交互流
  - [x] SubTask 3.1: 明确模板候选默认 active candidate、单选切换与 CTA 同步更新
  - [x] SubTask 3.2: 明确 `reuse_opportunity_hint / capability_gap_hint` 在 `Weekly Review` 与 `Product Create` 中的展示层级
  - [x] SubTask 3.3: 明确模板预填展示、可编辑语义与取消返回规则
  - [x] SubTask 3.4: 明确 `capability_gap_hint` 从 `Weekly Review` 与 `Product Create` 进入补齐页后的返回链与上下文保留规则

- [x] Task 4: 冻结返回链、空态、失败态与移动端降级
  - [x] SubTask 4.1: 明确 `templateSource` 与 `fromDashboard` 元数据并存时的用户可见返回链
  - [x] SubTask 4.2: 明确模板空态、提示空态、预填 unavailable 成功态与请求失败态
  - [x] SubTask 4.3: 明确移动浏览器下的单列降级、按钮布局与不可引入的第二套页面体系

- [x] Task 5: 完成 `phase09-05` 规格自检与一致性校验
  - [x] SubTask 5.1: 校验页面流与交互流足以直接进入实现
  - [x] SubTask 5.2: 校验模板与提示的正式消费位、返回链与空态语义已明确
  - [x] SubTask 5.3: 校验 `Weekly Review / Product Create / Product Detail` 没有重复长出第二套入口
  - [x] SubTask 5.4: 校验空态、失败态与回退路径可直接被浏览器验收用例消费
  - [x] SubTask 5.5: 校验 `capability_gap_hint` 的 CTA 页面流与 `phase09-03` canonical 路径完全一致，且返回链已机械冻结

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 2, Task 3
- Task 5 depends on Task 1, Task 2, Task 3, Task 4
