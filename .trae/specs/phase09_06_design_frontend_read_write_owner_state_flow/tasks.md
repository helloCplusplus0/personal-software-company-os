# Tasks

- [x] Task 1: 对齐 `phase09-06` 的上游冻结结论、真实代码入口与前端 owner 基线
  - [x] SubTask 1.1: 对齐 `phase09-04` 的合同、caller inventory 与前端 owner 边界
  - [x] SubTask 1.2: 对齐 `phase09-05` 的页面流、返回链与空态/失败态要求
  - [x] SubTask 1.3: 对齐 `phase08-06`、`phase06-07`、`phase07-05` 中关于 read/application 分层的既有口径
  - [x] SubTask 1.4: 复核当前真实前端承接位：`useWeeklyReviewRead`、`useReviewAction`、`useCreateDraftProduct`、`ProductCreatePage`、`ProductDetailPage`

- [x] Task 2: 冻结 `template-reuse` 读层切片落点与 `Weekly Review` 的正式消费方式
  - [x] SubTask 2.1: 明确模板候选、模板预填、模板来源复读与派生提示的 read layer 落点
  - [x] SubTask 2.1.1: 明确模板相关 read owner 的物理落点只能在 `template-reuse/data/`，不得再落到 `product-registry/data/`
  - [x] SubTask 2.2: 明确 `Weekly Review` 页面继续只消费 `useWeeklyReviewRead`
  - [x] SubTask 2.3: 明确新增模板消费位与既有 `reuseSnapshot / representativeSignals` 的边界矩阵

- [x] Task 3: 冻结 `Product Create` 的 handoff application owner、表单状态边界与成功回流分工
  - [x] SubTask 3.1: 明确 `use-product-create-template-handoff` 或等价单一 application owner 的职责
  - [x] SubTask 3.2: 明确 `useCreateDraftProduct` 继续是唯一正式 create mutation owner
  - [x] SubTask 3.3: 明确模板预填不能引入第二套 create form state 主线
  - [x] SubTask 3.3.1: 明确当前 `ProductCreateForm` 本地字段状态必须升级为唯一正式 create form state owner，而不是与临时 store 并存
  - [x] SubTask 3.3.2: 明确从 `Module Registry / Module Detail` 返回时只能通过同一条 form state 主线恢复草稿
  - [x] SubTask 3.4: 明确成功回流、unavailable 成功态、请求失败态与错误反馈的分层承接

- [x] Task 4: 冻结 `Product Detail` 的模板来源复读、caller-owner 映射与必须回收的临时编排点
  - [x] SubTask 4.1: 明确 `Product Detail` 模板来源复读只能通过单一 read owner 消费
  - [x] SubTask 4.2: 明确 `Product Detail` 继续只导向 canonical binding path，不长第二套详情写路径
  - [x] SubTask 4.3: 产出 `Weekly Review / Product Create / Product Detail` 的 caller-owner 一对一映射表
  - [x] SubTask 4.4: 识别并禁止继续扩写页面级临时编排模式

- [x] Task 5: 完成 `phase09-06` 规格自检与一致性校验
  - [x] SubTask 5.1: 校验 `query` 与 `application` 边界明确
  - [x] SubTask 5.2: 校验 `Product Create` canonical mutation owner 不被模板逻辑侵入替换
  - [x] SubTask 5.3: 校验 caller 不会跨页面漂移成第二套 owner
  - [x] SubTask 5.4: 校验 `Product Detail` 模板来源复读不会长出第二套详情写路径
  - [x] SubTask 5.5: 校验前端不会因为模板预填而新增第二套 create form state 主线
  - [x] SubTask 5.6: 校验从缺口补齐页返回后的草稿恢复机制已冻结为单值
  - [x] SubTask 5.7: 校验模板只读 owner 的物理切片落点不存在 `template-reuse` / `product-registry` 双轨

# Task Dependencies

- Task 2 depends on Task 1
- Task 3 depends on Task 1, Task 2
- Task 4 depends on Task 1, Task 2, Task 3
- Task 5 depends on Task 1, Task 2, Task 3, Task 4
