# Personal Software Company OS

# Module System & Technical Architecture Specification v1.0

**Document 04 + Document 05**

**Version：v1.0**
**Status：Baseline Design Document**
**Purpose：定义个人软件资产如何形成、演进、复用，以及 PSCO OS 的工程技术实现方案**

> 2026-08-04 共识对齐注记：
> 本文继续作为 PSCO 的模块系统与技术架构基线，但当前工程阶段已经确认：
> - 默认技术基线继续为 `React + Go + PostgreSQL`；
> - `Rust Intelligence Layer` 不进入 `v0.1` 主工程范围，只保留为远期演进方向；
> - `Local First` 的当前解释是“数据所有权优先”，不等于必须切换到 `SQLite`；
> - `Module` 在 `v0.1` 中采用轻量契约字段与准入规则，不以前置完整 PMM 为阻断项；
> - `Repository` 在 `v0.1` 中先承担手动绑定与实现锚点角色；
> - `GitHub OAuth / 自动导入` 不作为当前阻断项。

---

# Document 04

# Personal Software Company Module System v1.0

---

# 1. 文档目的

Document 02 定义：

> 个人软件公司如何运行。

Document 03 定义：

> 系统管理哪些核心对象。

Document 04 进一步回答：

> 软件能力如何从一次项目开发中产生，并最终沉淀为个人长期资产？

---

# 2. 核心理念：从 Project 到 Capability

传统软件开发模式：

```text
Project

↓

Code

↓

Release

↓

Archive

```

问题：

项目结束后：

* 代码存在；
* 经验消失；
* 能力无法复用。

---

Personal Software Company 模式：

```text
Project

↓

Experience

↓

Reusable Module

↓

Capability

↓

Future Product

```

---

核心变化：

> 软件项目不再是终点，而是能力生产过程。

---

# 3. Module 的重新定义

## 3.1 Module 定义

Personal Software Company 中：

Module 不是简单代码包。

不是：

* npm package；
* Git repository；
* 文件夹。

---

Module 定义：

> 一个具有明确职责、稳定接口、独立生命周期，并经过真实项目验证的软件能力单元。

---

因此：

Module =

```text
Capability Boundary

+

Implementation

+

Interface

+

Knowledge

+

History

```

---

# 4. Module 的核心组成

一个完整 Module：

```text
Module

├── Interface
│
├── Implementation
│
├── Documentation
│
├── Tests
│
├── Examples
│
├── Decision History
│
├── Version History
│
└── Usage Context

```

---

## 4.1 Interface（接口）

决定：

模块如何被使用。

例如：

Auth Module：

```text
login()

logout()

refreshToken()

authorize()

```

---

接口是资产边界。

没有接口：

只有代码。

---

## 4.2 Implementation（实现）

具体技术：

例如：

Auth v1:

```text
JWT
```

Auth v2:

```text
OAuth
```

---

实现可以变化。

接口保持稳定。

---

## 4.3 Documentation（文档）

记录：

* 使用方式；
* 设计思想；
* 限制条件。

---

## 4.4 Tests（验证）

证明：

这个能力可靠。

---

## 4.5 Decision History（决策历史）

记录：

为什么这样设计。

例如：

```text
Decision:

选择 PostgreSQL


Context:

需要事务一致性


Alternative:

MongoDB


Reason:

关系模型更符合业务结构

```

---

这是未来 AI 理解你的重要上下文。

---

# 5. Module 与 Package 的区别

这是 PSCO 最重要的设计区别。

## Package

关注：

代码复用。

例如：

```text
lodash

react-router

```

---

## Module

关注：

能力复用。

例如：

```text
Authentication Capability

Payment Capability

Rental Domain Capability

```

---

Package 是：

技术资产。

Module 是：

软件公司资产。

---

# 6. Module 分类体系

不能按照语言分类。

错误：

```text
Go Module

React Module

Rust Module

```

因为技术会变化。

---

应该按照能力分类。

---

# 6.1 Foundation Modules

基础能力。

例如：

```text
Authentication

Authorization

Configuration

Logging

Error Handling

Deployment

```

---

# 6.2 Application Modules

应用能力。

例如：

```text
File Management

Notification

User Management

Search

```

---

# 6.3 Domain Modules

领域能力。

这是最高价值部分。

例如：

```text
Rental Domain

Trading Domain

Financial Analysis Domain

```

---

# 6.4 AI Modules

AI时代核心能力。

例如：

```text
LLM Gateway

RAG Pipeline

Agent Workflow

Evaluation System

```

---

# 6.5 Data Modules

数据能力。

例如：

```text
Data Ingestion

ETL Pipeline

Knowledge Extraction

Feature Pipeline

```

---

# 7. Module 生命周期

Module 不应该一开始就存在。

它应该自然形成。

---

## Stage 0：Prototype

项目内部代码。

例如：

```text
rento/src/auth
```

特点：

* 快速；
* 不稳定；
* 项目专用。

---

## Stage 1：Candidate

发现未来可能复用。

开始：

* 分离职责；
* 定义接口；
* 补充文档。

---

## Stage 2：Internal Module

个人内部资产。

特点：

* 可跨项目使用；
* 有版本。

例如：

```text
auth-module v1.0
```

---

## Stage 3：Stable Module

长期能力。

要求：

* 多项目验证；
* 稳定接口；
* 完善测试。

---

## Stage 4：Commercial Module

商业化能力。

例如：

未来：

出售 SaaS Starter Kit。

---

补充说明：

`Commercial` 阶段继续保留为长期演进方向，但不属于 `v0.1` 的执行目标。

---

生命周期：

```text
Prototype

↓

Candidate

↓

Internal

↓

Stable

↓

Commercial

```

---

# 8. Module Extraction Workflow

软件资产形成过程：

```text
Product Development


↓

发现重复能力


↓

Extract


↓

Define Interface


↓

Test


↓

Version


↓

Register


↓

Reuse

```

---

关键原则：

不是提前设计模块。

而是：

从真实需求中抽象。

---

# 9. Module Registry

PSCO OS 必须拥有：

个人模块注册中心。

类似：

企业内部 Package Registry。

---

记录：

```yaml
name:

auth-module


version:

3.0.0


category:

foundation


language:

go


used_by:

- Rento

- AI Platform


status:

stable

```

---

---

# Document 05

# Personal Software Company Technical Architecture v1.0

---

# 1. 架构目标

PSCO OS 不是高并发互联网 SaaS。

它首先是：

个人软件公司的内部基础设施。

因此设计目标：

## 第一优先级

长期维护。

## 第二优先级

开发效率。

## 第三优先级

性能扩展。

---

# 2. 核心技术原则

## Principle 1

简单可靠优先。

避免：

* Kubernetes；
* 复杂微服务；
* 过度分布式。

---

## Principle 2

领域模型优先。

代码只是实现。

---

## Principle 3

Local First。

个人资产首先属于个人。

---

当前阶段进一步解释为：

> 数据所有权、导出能力、备份能力与迁移能力优先。

这并不自动等于：

> `v0.1` 必须改用 `SQLite`。

---

## Principle 4

AI Ready。

但 AI 是增强层。

---

# 3. 总体架构

```text
                    React Client


                         |


                    API Layer


                         |


                    Go Core


                         |


                    PostgreSQL


                         |


                  Business Data


```

---

`v0.1` 当前正式架构只保留：

```text
React Client

↓

API Layer

↓

Go Core

↓

PostgreSQL
```

---

# 4. Frontend Architecture

## 技术选择

```text
React

+

TypeScript

+

Vite

+

React Router

+

Tailwind CSS

+

shadcn/ui

```

---

原因：

PSCO OS 本质：

复杂管理系统。

需要：

* Dashboard；
* 编辑器；
* 图关系；
* 数据展示；
* 搜索。

React生态成熟。

---

# 5. Backend Architecture

## 技术选择

Go。

---

原因：

PSCO核心：

不是计算。

而是：

领域管理。

Go优势：

* 简洁；
* 稳定；
* 长期维护成本低。

---

结构：

```text
backend/

cmd/

server/


internal/

venture/

product/

module/

decision/

experiment/


pkg/

database/

```

---

# 6. Database Architecture

## PostgreSQL

负责：

业务状态。

包括：

* Venture；
* Product；
* Module；
* Release；
* Decision。

---

例如：

Module 表：

```sql
modules

id

name

version

status

repository

```

---

# 7. Rust Intelligence Layer

Rust 不进入业务核心。

它负责：

高性能能力。

---

未来：

## 7.1 Code Index Engine

分析：

多个代码仓库。

---

## 7.2 Semantic Search Engine

寻找：

相关能力。

---

## 7.3 Dependency Analysis

分析：

模块关系。

---

架构：

```text
Go Core

↓

Intelligence API

↓

Rust Engine

```

---

补充说明：

该部分只保留为 **远期演进方向**。

在 `v0.1` 中：

- 不实现 Rust Intelligence Layer；
- 不为其增加前置工程复杂度；
- 不把自动扫描、依赖分析或语义搜索写成当前 MVP 的默认能力。

---

# 8. Repository Architecture

推荐：

Hybrid。

---

核心：

```text
personal-software-company-os/


├── registry

├── templates

├── documentation

└── projects

```

---

模块：

独立仓库。

例如：

```text
github/

auth-module

storage-module

ai-gateway-module

```

---

原因：

模块需要：

独立生命周期。

---

补充说明：

当前工程阶段已确认：

- PSCO 自身仓库可以维持一个清晰的主仓库结构；
- 被管理的模块继续按“独立生命周期”思路理解；
- `v0.1` 先支持手动 `Repository` 绑定，不以前置自动化仓库治理为目标。

---

# 9. Project 与 Module 关系

Product：

不拥有所有代码。

而是：

组合 Module。

---

例如：

Rento：

```yaml
modules:

auth:
 3.0.0

storage:
 2.0.0

notification:
 1.0.0

```

---

类似：

package-lock。

---

# 10. AI Integration Architecture

AI 不负责：

自动决定资产。

AI负责：

增强。

---

## AI Assistant Layer

能力：

### 1. Documentation Assistant

生成模块文档。

---

### 2. Development Assistant

理解：

项目上下文。

---

### 3. Migration Assistant

辅助升级模块。

---

### 4. Research Assistant

辅助机会探索。

---

架构：

```text
PSCO Context

↓

AI Context Provider

↓

LLM

```

---

# 11. MVP 技术范围

必须控制。

v0.1：

不做：

❌ 自动扫描代码
❌ 自动生成知识图谱
❌ AI自动推荐模块
❌ 自动理解全部代码

---

只做：

## Core Registry

```text
Module CRUD


Version Management


Project Binding


Decision Record

```

---

当前 `v0.1` 进一步收敛为：

- `Module CRUD`
- `Release Management`
- `Product / Repository / Module Binding`
- `Decision Record`
- `Dashboard` 最小聚合反馈

---

技术：

```text
React

+

Go

+

PostgreSQL

```

即可。

---

# 12. 从 Rento-miniX 开始验证

Rento-miniX 是第一个真实资产来源。

目标：

不是重构。

而是：

识别已有能力。

可能模块：

---

## Auth Module

身份认证能力。

---

## Storage Module

图片上传：

* 压缩；
* 预览；
* 存储。

---

## Deployment Module

部署体系：

* Caddy；
* systemd；
* PostgreSQL。

---

## SaaS Foundation Module

通用 SaaS 基础能力。

---

# 13. Document 04 + Document 05 最终结论

经过这两个文档：

PSCO OS 的技术核心已经明确：

---

## 软件资产形成模型：

```text
项目经验

↓

模块抽象

↓

模块复用

↓

能力形成

↓

未来产品

```

---

## 技术实现模型：

```text
React

↓

Go Domain Core

↓

PostgreSQL


+

Rust Intelligence Engine


+

AI Enhancement Layer

```

---

# 当前整体设计状态

目前四组核心文档已经形成：

```text
Document 00/01

为什么做、是什么


        ↓


Document 02/03

如何运行、管理什么


        ↓


Document 04/05

如何形成资产、如何实现


        ↓


Document 06/07

工作流与AI策略


        ↓


Document 08

UX Product Specification

```

---

下一阶段进入：

# Document 06 + Document 07

将解决两个关键问题：

1. **个人开发者每天如何使用 PSCO OS？**

   → Personal Software Company Workflow Engine v1.0

2. **AI 在这个系统中的正确位置是什么？**

   → Personal Software Company AI Strategy v1.0

这两个文档完成后，才适合进入最终的 UX 设计。
