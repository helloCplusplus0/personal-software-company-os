# Personal Software Company OS

# Product UX Specification v1.0

**Document 08**

**Version：v1.0**
**Status：Baseline UX Design**
**Purpose：定义 Personal Software Company OS 作为真实软件产品时，用户如何理解、使用和获得价值**

> 2026-08-04 共识对齐注记：
> 本文继续承载 PSCO 的长期 UX 基线，但当前工程阶段已经确认：
> - `v0.1` 的正式目标是“资产登记 + 决策留痕 + 基础复用反馈”；
> - `Decision` 保留在 MVP；
> - `Capability` 在 `v0.1` 中以派生结果层呈现，而不是要求用户主动维护；
> - `AI Assistant` 不作为 `v0.1` 一级主导航；
> - `Knowledge / Experiment / Feature` 保留在长期模型与未来 UX 中，但不进入 `v0.1` 主执行范围；
> - 当前 IA 与页面范围需要以最终共识文档为准。

---

# 1. UX Design Goal（产品体验目标）

前面的文档解决了：

* 为什么存在；
* 如何运行；
* 管理什么；
* 如何形成资产；
* 如何实现技术架构；
* AI 如何参与。

Document 08 解决最后一个问题：

> 一个个人开发者每天打开 Personal Software Company OS，到底做什么？

---

核心原则：

> PSCO OS 不是让用户“管理更多东西”，而是让用户“更高效地经营自己的软件公司”。

---

因此 UX 不能设计成：

❌ Notion 页面集合
❌ Jira 任务列表
❌ GitHub 仓库浏览器
❌ AI Chat 界面

---

而应该设计成：

> 一个个人软件公司的控制中心（Personal Software Company Command Center）。

---

# 2. 用户画像（Primary User）

## Persona：

个人开发者 / 独立创业者 / 小型软件公司创始人

特点：

* 同时探索多个方向；
* 独立完成产品开发；
* 使用 AI 辅助开发；
* 长期积累代码和经验；
* 希望形成商业化能力。

---

典型场景：

### 场景 1

“我突然想到一个产品方向，我如何判断是否值得投入？”

---

### 场景 2

“我要开发一个新产品，有哪些已有能力可以复用？”

---

### 场景 3

“几个月后，我为什么当初这样设计？”

---

### 场景 4

“我已经做了很多项目，但是感觉没有形成积累。”

---

PSCO UX 必须解决这些问题。

---

# 3. Product Mental Model（用户心智模型）

用户打开系统，不应该想到：

“我要填写数据”。

应该想到：

“我正在运营我的软件公司。”

---

首页应该类似：

## Personal Company Dashboard

```text
+------------------------------------------------+

 Personal Software Company


 Active Ventures

 ┌───────────────────────┐
 │ Personal AI Platform   │
 │ Status: Building      │
 └───────────────────────┘


 Products

 ┌───────────────────────┐
 │ Rento-miniX            │
 │ Stage: MVP             │
 └───────────────────────┘


 Capabilities

 42 Modules
 12 Stable Capabilities


 Decisions

 5 Pending Decisions


 Experiments

 3 Running


+------------------------------------------------+

```

---

# 4. Information Architecture（信息架构）

整体导航：

```text
Personal Software Company


├── Dashboard

├── Products

├── Modules

├── Decisions

├── Ventures (Optional)

└── Settings

```

---

补充说明：

- 上面是当前 `v0.1` 推荐导航，而不是长期全量导航；
- `Experiments / Knowledge / AI Assistant` 保留为后续能力方向，不进入第一版主导航；
- `Project / Repository Binding` 优先作为 `Product` 与 `Module` 内的主功能块承接。

---

# 5. Dashboard UX（公司驾驶舱）

Dashboard 不是统计页面。

不是：

```
项目数量
代码行数
提交次数
```

---

这些指标没有价值。

---

Dashboard 展示：

## 5.1 Current Focus

当前最重要事情。

例如：

```
Current Venture:

Personal Software Company OS


Current Goal:

Validate MVP architecture


Next Decision:

Module registry design

```

---

## 5.2 Capability Growth

显示：

个人能力增长。

---

在 `v0.1` 中，这里的 `Capability Growth` 应理解为：

> 基于模块数量、复用关系、稳定状态、版本演进与决策沉淀生成的派生反馈。

而不是要求用户手动创建和维护能力实体。

例如：

```
Software Capability


Authentication

██████████ Stable


AI Integration

██████░░░ Candidate


Data Pipeline

████░░░░ Experimental

```

---

## 5.3 Asset Evolution

不是代码量。

而是：

能力资产。

例如：

```
Modules:


2026

+ Auth Module

+ AI Gateway

+ Data Pipeline


Total:

37 reusable capabilities

```

---

# 6. Venture Management UX

## Venture 页面

对应：

长期战略方向。

---

补充说明：

`Venture` 在 `v0.1` 中保留，但作为 **可选实体**，不强制成为所有用户的第一层输入入口。

---

例如：

```
Venture:

Personal Software Company OS


Mission:

Build a personal software production system


Products:

├── PSCO OS
├── Module Registry
└── AI Assistant


Status:

Building

```

---

用户关注：

不是任务。

而是：

方向是否正确。

---

# 7. Product Management UX

Product 页面：

连接商业和工程。

---

结构：

```
Product


Overview

Value Proposition


Users


Features


Modules


Releases


Metrics

```

---

例如：

Rento-miniX：

```
Product:

Rento-miniX


Value:

Help individual landlords manage rental business


Features:

✓ Tenant Management

✓ Contract Management

✓ Payment Tracking


Modules:

Auth v3

Storage v2

Notification v1

```

---

这里体现：

产品不是代码仓库。

---

# 8. Feature Management UX

Feature 是用户价值入口。

---

补充说明：

`Feature` 继续保留在长期 UX 模型中。

但它不属于 `v0.1` 的主执行范围，第一版优先通过：

`Product -> Module -> Decision`

承接最小价值闭环。

---

Feature 页面：

```
Feature:

Lease Reminder


Problem:

Landlords forget renewal dates


Solution:

Automatic notification


Required Capability:

Notification Module


Status:

Released

```

---

Feature 自动连接：

产品：

↓

模块：

↓

版本。

---

# 9. Module Library UX（核心页面）

这是 PSCO 最重要页面。

---

不是代码列表。

而是：

个人能力库。

---

设计：

```
Module Library


Search Capability...


Categories:


Foundation

Application

Domain

AI

Data


--------------------------------

Authentication Module


Version:

3.2


Status:

Stable


Used By:

Rento-miniX

AI Platform


Capability:

Identity Management


```

---

用户看到的是：

“我拥有了什么能力”。

---

# 10. Module Detail UX

一个 Module：

```
Authentication Module


Capability:

Identity Management


Interface:

login()
logout()


Implementation:

Go


Version:

3.2


Used Projects:

Rento

AI Platform


Decision History:


Why PostgreSQL?

Why JWT?


Evolution:


v1 Basic Auth

v2 OAuth

v3 Multi Tenant

```

---

这才体现：

软件资产。

---

# 11. Decision Center UX

这是区别于普通工具的重要地方。

---

页面：

```
Important Decisions


Decision:


Use PostgreSQL


Context:

Rental domain requires relational consistency


Alternatives:

MongoDB


Reason:

Better transaction model


Impact:

Foundation of SaaS architecture

```

---

未来 AI 最需要这里。

---

# 12. Experiment Center UX

对应：

创新验证。

---

补充说明：

`Experiment Center` 属于后续版本方向，不进入 `v0.1` MVP 范围。

---

页面：

```
Experiment


Hypothesis:

Users need rental automation


Method:

Interview 20 landlords


Result:

Positive


Decision:

Build MVP

```

---

避免：

凭感觉开发。

---

# 13. AI Assistant UX

这里非常重要。

AI 不应该是：

ChatGPT窗口。

---

应该是：

Context-aware Assistant。

---

例如：

用户进入 Module 页面：

AI知道：

```
Current Module:

Storage


Related Products:

Rento


Past Decisions:

Image compression strategy


```

---

AI可以：

> “Rento 使用的图片处理方案可能适用于当前项目。”

---

但是：

不是：

“自动修改”。

---

补充说明：

`AI Assistant` 的长期方向成立，但在 `v0.1` 中：

- 不作为一级主导航；
- 优先以页面内的 context-aware enhancement 出现；
- 例如在 `Module`、`Product`、`Decision` 页面提供低干扰提示。

---

# 14. New Product Creation UX

这是未来最有价值场景。

用户：

创建新产品。

---

流程：

```
Create Product


↓

Define Value


↓

Select Venture


↓

AI Suggest Existing Capability


↓

Compose Modules


↓

Generate Project Template


```

---

例如：

用户：

我要做 SaaS。

系统：

已有：

```
Auth Module

Billing Module

Notification Module

Deployment Module

```

---

生成：

Starter Architecture。

---

# 15. Daily Usage Flow

真实一天：

---

## 开始工作

打开 Dashboard。

看到：

当前 Venture。

---

## 开发

进入 Product。

查看 Modules 与相关 Decision。

---

## 遇到设计问题

记录 Decision。

---

## 发现可复用代码

创建 Module。

---

## 完成版本

发布 Release。

---

## 周末复盘

查看：

Capability Growth。

---

# 16. MVP UX Scope（非常重要）

不要一开始实现全部。

v0.1：

只需要：

---

## 1. Dashboard

查看当前状态。

---

## 2. Module Registry

管理软件资产。

---

## 3. Product Registry

记录产品和模块关系。

---

## 4. Decision Records

保存关键决策。

---

## 5. Project Binding

项目绑定模块。

---

当前进一步明确为：

## 5. Project / Repository Binding

承接：

- Product 与 Repository 的绑定；
- Product 与 Module 的绑定；
- Module 与 Repository 的实现映射。

---

技术：

React

*

Go

*

PostgreSQL

即可。

---

# 17. 不应该实现的功能（初期）

明确禁止：

## ❌ 自动扫描所有代码

原因：

复杂且价值不确定。

---

## ❌ 自动生成知识图谱

原因：

容易产生大量无价值关系。

---

## ❌ AI自动判断最佳方案

原因：

责任边界错误。

---

## ❌ 类似 Notion 编辑器

原因：

偏离核心。

---

# 18. 最终 UX 定义

Personal Software Company OS 用户体验不是：

> “帮我管理代码。”

而是：

> “让我知道我拥有什么能力，我正在构建什么，我过去为什么这样选择，以及下一次如何更快创造价值。”

---

# 19. 完整系统闭环

最终：

```
                Human


                 ↓


        Opportunity Discovery


                 ↓


             Product


                 ↓


            Development


                 ↓


             Module


                 ↓


           Capability


                 ↓


        Personal Software Company


                 ↓


              AI Amplification

```

---

# 20. Document 00-08 总结

至此 Personal Software Company OS v1.0 完成完整设计：

| 文档          | 定义     |
| ----------- | ------ |
| Document 00 | 战略愿景   |
| Document 01 | 核心哲学   |
| Document 02 | 公司运行模型 |
| Document 03 | 领域模型   |
| Document 04 | 软件资产模型 |
| Document 05 | 技术架构   |
| Document 06 | 工作流    |
| Document 07 | AI边界   |
| Document 08 | 产品UX   |

---

## 最终产品定位

一句话：

> Personal Software Company OS 是一个帮助个人开发者运营自己的软件公司的 AI 原生生产操作系统，通过机会管理、产品开发、模块积累和能力复利，使个人开发者获得长期的软件生产能力。

---

至此，方案已经从最初的“AI代码资产记忆系统”完成了一次根本转型：

**从工具 → 系统 → 个人软件公司的基础设施。**

下一步真正进入工程阶段时，不应该继续扩大概念，而应该开始设计：

**Personal Software Company OS v0.1 MVP Specification**

也就是：

> 第一版到底做哪些页面、哪些数据库表、哪些 API、哪些模块，如何在 4-8 周内跑出第一个可用版本。
