# phase09-11 模板复用与派生提示联调、浏览器验收与反回归验证报告

## 验收结论：通过

`phase09-08 / 09 / 10` 已交付的模板复用与派生提示能力，已在同一正式运行环境中重新拉通。当前 `phase09-11` 的工具链验证、模板相关 API smoke、`Weekly Review -> Product Create -> Product Detail` 浏览器闭环、以及 `Dashboard / Review / Product Detail / ReuseSummary` 最小反回归验证均通过，未发现阻断级运行时问题，当前实现链已具备进入后续收口的条件。

---

## 1. 验收环境与前置条件

- 前端服务：`http://127.0.0.1:5173`
- 后端服务：`http://127.0.0.1:8081`
- 运行口径：`React Web + ConnectRPC + /api` 正式主线
- 环境来源：复用用户手动开启的前后端服务；本轮未额外开启任何端口或旁路服务
- 健康检查：`GET http://127.0.0.1:8081/healthz` 返回 `200 application/json`

冻结结论：

- 本阶段验收未引入第二套前端入口、第二套 transport 主线或临时模板/提示旁路
- 本阶段继续只验证模板复用与派生提示当前已冻结能力，不扩写未来范围

---

## 2. 工具链验证

| 项目 | 结果 |
|------|------|
| `proto: make build && make gen && make lint && make breaking` | 通过 |
| `backend: go build ./...` | 通过 |
| `frontend: npx tsc -b --noEmit` | 通过 |
| `frontend: npm run build` | 通过 |

补充说明：

- 所有验证均直接在当前仓库工作树执行
- 本轮未因验收需要临时修改构建脚本、代理配置或运行端口

---

## 3. 模板相关 API smoke

### 3.1 模板候选读取

- 请求：`TemplateReuseService/ListTemplateCandidates`
- 输入：
  - `consumerSurface = TEMPLATE_CONSUMER_SURFACE_WEEKLY_REVIEW`
- 结果：
  - 成功返回 2 条模板候选
  - `defaultActiveCandidateId = 8eb612e43da13ded5dd0b1dbf1ba73a5`
  - 候选中包含 `Product A / Product B`

### 3.2 模板预填详情读取

- 请求：`TemplateReuseService/GetTemplateCandidatePrefill`
- 输入：
  - `templateCandidateId = 8eb612e43da13ded5dd0b1dbf1ba73a5`
  - `consumerSurface = TEMPLATE_CONSUMER_SURFACE_PRODUCT_CREATE`
- 结果：
  - `resolutionStatus = TEMPLATE_RESOLUTION_STATUS_RESOLVED`
  - 成功返回建议名称、建议描述、模块组合摘要
  - 成功返回 `capabilityGapHints`

### 3.3 派生提示读取

- 请求：`TemplateReuseService/GetDerivedInsightHints`
- 输入：
  - `templateCandidateId = 8eb612e43da13ded5dd0b1dbf1ba73a5`
  - `consumerSurface = TEMPLATE_CONSUMER_SURFACE_WEEKLY_REVIEW`
  - `reviewScopeKey = weekly-review`
- 结果：
  - `resolutionStatus = TEMPLATE_RESOLUTION_STATUS_RESOLVED`
  - 成功返回 `reuse_opportunity_hint`
  - 返回空缺口列表不被误判为空错误；当前环境下正式缺口提示由浏览器入口按 active candidate 正常呈现

冻结结论：

- 模板候选、模板预填详情、派生提示三类 API smoke 均命中当前正式合同与 transport 主线
- 未发现通过临时脚本私有语义才能成立的情况

---

## 4. 浏览器闭环验收

### 4.1 `Weekly Review` 入口断言

入口页面：`/reviews/weekly`

观察结果：

- 模板候选区存在，且出现成功态而非失败态
- active candidate 为单值，当前默认选中 `Product B`
- 派生提示区存在，且同时出现：
  - `复用机会：模块组合已被多个产品使用`
  - `能力缺口：补齐 Web Frontend`
- CTA 存在：
  - `基于模板创建产品`
  - `前往 Module Registry`

### 4.2 `templateCandidateId` 进入预填创建页

点击路径：

- `Weekly Review -> 基于模板创建产品`

导航结果：

- 页面进入 `/products/new?fromTemplateReuse=true&templateCandidateId=aafb9413b9c6f4165ef3201608c82897&templateSource=weekly-review`
- 已机械证明通过 `templateCandidateId` 进入正式预填创建页

### 4.3 `Product Create` 字段级预填断言

观察结果：

- 名称预填：`Product B (基于模板)`
- 描述预填：`基于模板「Product B」创建，包含 1 个模块：auth-service`
- 页面展示模板来源摘要：`Product B / 由 1 个模块组成的模板，源自 1 个产品`
- 页面展示解释性缺口提示：`能力模块：auth-service`
- 预填字段可编辑；本轮将名称改为：
  - `Product B Phase09-11 Acceptance 20260812`

冻结结论：

- 模板预填真实缩短了创建路径，而不是要求手工重新输入模板信息

### 4.4 创建成功并进入 `Product Detail`

提交结果：

- 创建成功后进入：
  - `/products/f75fb8b3-e0dd-4911-a291-7a5e27fe09b1?fromTemplateReuse=true&templateCandidateId=aafb9413b9c6f4165ef3201608c82897&templateSource=weekly-review`

结果页断言：

- 模板来源摘要可见：
  - `Product B`
  - `由 1 个模块组成的模板，源自 1 个产品`
- 候选模块组合摘要可见：
  - `基于模板「Product B」创建，包含 1 个模块：auth-service`
- canonical binding CTA 可见：
  - `为模板模块绑定仓库`
- `ReuseSummary` 区块成功返回空态：
  - `暂无复用反馈`
- 空态被视为成功 reread，不误判为失败

---

## 5. 派生提示动作承接证据

### 5.1 `Weekly Review -> Module Registry`

点击路径：

- `Weekly Review -> 前往 Module Registry`

导航结果：

- 进入：
  - `/modules?returnTo=weekly-review&returnCandidateId=aafb9413b9c6f4165ef3201608c82897&statusFilter=all`
- 说明 `capability_gap_hint` 已真实承接到既有 canonical `Module Registry` 页面

### 5.2 返回 `Weekly Review`

返回路径：

- `Module Registry -> 返回 Weekly Review`

结果：

- 成功返回 `/reviews/weekly?returnCandidateId=aafb9413b9c6f4165ef3201608c82897`
- active candidate 恢复为原候选
- 模板候选区、派生提示区与 CTA 继续保持可见

冻结结论：

- 派生提示已能真实支撑下一步动作，不只停留在静态说明文案

---

## 6. 最小反回归与 reread

### 6.1 Dashboard

观察结果：

- 从新建产品的 `Product Detail` 返回 `Dashboard` 后，页面正常加载
- 系统概览中 `产品` 计数由 `3` 变为 `4`
- `Current Focus / Asset Feedback / Recent Activity / Reuse Snapshot` 均正常可见
- 新产品 `Product B Phase09-11 Acceptance 20260812` 已出现在：
  - `Current Focus / Asset Feedback` 的缺口信号
  - `Recent Activity`

结论：

- `Dashboard` reread 成功

### 6.2 Review

观察结果：

- 从 `Dashboard` 再次进入 `Weekly Review` 后，页面正常加载
- `Overview / Template Candidates / Derived Hints / Recent Activity / Representative Signals / Reuse Snapshot` 均正常可见
- 新创建产品已出现在 Recent Activity 区块
- 模板候选与提示区未因当前轮创建而崩溃或漂移

结论：

- `Review` reread 成功

### 6.3 Product Detail

观察结果：

- 新建产品的 `Product Detail` 可正常打开
- 模板来源摘要、模块组合摘要与 canonical binding CTA 均存在
- `ReuseSummary` 局部空态成功呈现，未退化为整页错误

结论：

- `Product Detail` reread 成功

### 6.4 ReuseSummary

观察结果：

- `Dashboard` 中 `Reuse Snapshot` 正常呈现
- 新建产品 `Product Detail` 中 `复用摘要` 正常返回成功空态 `暂无复用反馈`

结论：

- `ReuseSummary` 在 dashboard 与 product-detail 两个正式挂接位都能成功 reread
- “无统计变化 / 成功空态”未被误判为失败

---

## 7. 阻断问题核对

| 问题类型 | 结果 |
|----------|------|
| 模板候选读取失败 | 未发现 |
| active candidate 分裂或丢失 | 未发现 |
| 派生提示缺失或 CTA 失效 | 未发现 |
| `templateCandidateId` 未进入预填创建页 | 未发现 |
| 创建成功后未回流到 `Product Detail` | 未发现 |
| `Product Detail` 模板来源摘要缺失 | 未发现 |
| 候选模块组合摘要缺失 | 未发现 |
| canonical binding CTA 缺失 | 未发现 |
| `Dashboard / Review / ReuseSummary` reread 失败 | 未发现 |
| 将成功空态误判为失败 | 未发现 |

---

## 8. 非目标边界

本阶段继续明确不做：

- `dry-run`
- `AI Context Enhancement`
- `Venture`

冻结结论：

- 本轮验收只验证模板复用与派生提示当前已冻结能力
- 未将未来能力写成当前既成事实

---

## 9. 最终判断

`phase09-11` 已满足以下通过条件：

1. 工具链验证通过
2. 模板候选、模板预填详情、派生提示三类 API smoke 通过
3. `Weekly Review -> Product Create 预填 -> Product Detail` 浏览器闭环通过
4. 模板预填真实缩短创建路径的证据已留档
5. 派生提示真实支撑下一步动作的证据已留档
6. `Dashboard / Review / Product Detail / ReuseSummary` 最小反回归与 reread 通过
7. 成功空态未被误判为失败
8. 本阶段边界未漂移

因此，`phase09-11` 当前可以按“通过验收”处理，`phase09` 模板复用与派生提示实现链已具备进入后续收口的条件。
