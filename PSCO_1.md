# Personal Software Company OS

# Operating Model & Domain Model Specification v1.0

**Document 02 + Document 03**

**Version：v1.0**
**Status：Baseline Design Document**
**Purpose：定义个人软件公司如何运行，以及 PSCO OS 内部管理的核心对象模型**

> 2026-08-04 共识对齐注记：
> 本文继续承载 PSCO 的运行模型与长期领域模型基线，但当前工程阶段已经确认：
> - `v0.1` 的执行子集优先聚焦 `Product / Module / Release / Decision / Repository`，`Venture` 为可选实体；
> - `Capability` 在 `v0.1` 中作为派生视图或结果层，不作为核心 CRUD 实体；
> - `Feature / Opportunity / Experiment` 继续保留在长期理论模型中，但不进入 `v0.1` 主执行范围；
> - `Decision` 需要采用低摩擦、结构化的最小模板；
> - `Repository` 在 `v0.1` 中先承担手动绑定与实现载体角色，不以前置自动化集成为目标。

---

# Document 02

# Personal Software Company Operating Model v1.0

---

# 1. 文档目的

Document 00/01 解决：

> 为什么需要 Personal Software Company OS？

本文件进一步回答：

> 如果一个个人开发者真的像运营一家软件公司一样工作，那么这家公司应该如何运行？

---

传统公司的运行模式：

```text
战略部门
    ↓
产品部门
    ↓
研发部门
    ↓
运营部门
    ↓
资产沉淀
```

PSCO OS 的目标：

不是复制传统公司组织结构。

而是：

将一个人的创造过程工程化。

---

# 2. Personal Software Company 的基本运行模型

个人软件公司由四个核心循环组成：

```text
                    Personal Software Company


                          


        Innovation Loop

              ↓

        Product Loop

              ↓

        Engineering Loop

              ↓

        Capability Loop

              ↓

        Innovation Loop

```

这四个循环构成长期复利飞轮。

---

# 3. Innovation Loop（创新循环）

## 目的

解决：

> 做什么？

---

传统个人开发：

```text
突然想到一个想法

↓

马上开发

↓

投入大量时间

↓

发现没有需求
```

---

PSCO：

```text
Idea

↓

Opportunity

↓

Hypothesis

↓

Experiment

↓

Decision

```

---

## 核心对象

### Opportunity（机会）

不是功能。

不是项目。

而是：

> 一个值得探索的问题或商业机会。

例如：

错误：

```
我要开发一个租房APP
```

正确：

```
大量个人房东缺少低成本数字化管理工具
```

---

## Opportunity 必须包含：

### Problem

存在什么问题？

---

### Target User

谁遇到问题？

---

### Hypothesis

为什么认为可以解决？

---

### Validation Method

如何验证？

---

例如：

```yaml
Opportunity:

problem:
房东管理租约效率低

user:
个人房东

hypothesis:
提供轻量化SaaS可以降低管理成本

validation:
访谈20名房东

```

---

# 4. Venture Loop（事业探索循环）

## 目的

解决：

> 是否值得长期投入？

---

Opportunity 经过验证后：

进入：

Venture。

---

Venture 定义：

> 一个长期探索方向，而不是一个具体产品。

例如：

---

Opportunity：

```
房东管理困难
```

↓

Venture：

```
Rental Software Ecosystem
```

↓

Product：

```
Rento-miniX
```

---

一个 Venture 可以产生多个产品。

---

例如：

未来：

```
AI知识管理 Venture

    ├── Personal Knowledge OS

    ├── Research Assistant

    └── Decision Platform

```

---

# 5. Product Loop（产品循环）

## 目的

解决：

> 创造什么价值？

---

模型：

```text
Venture

↓

Product

↓

Feature

↓

User Value

```

---

## Product

不是代码仓库。

而是：

> 面向用户提供持续价值的商业实体。

---

Product 必须明确：

---

## Value Proposition

为什么用户需要？

---

## Customer

谁使用？

---

## Business Model

如何产生价值？

---

## Product Lifecycle

阶段：

```text
Idea

↓

Prototype

↓

MVP

↓

Growth

↓

Stable

↓

Retirement

```

---

# 6. Feature Loop（功能循环）

## 目的

连接用户需求和工程实现。

---

Feature 不应该是：

“开发一个按钮”。

而应该：

```text
User Problem

↓

Feature

↓

Capability Requirement

↓

Module Composition

```

---

例如：

用户需求：

```
房东需要查看租约状态
```

↓

Feature：

```
Lease Dashboard
```

↓

需要能力：

```
Document Module

Notification Module

Rental Domain Module

```

---

# 7. Engineering Loop（工程循环）

## 目的

解决：

> 如何持续生产软件？

---

传统：

```text
Feature

↓

Coding

↓

Deploy

```

结束。

---

PSCO：

```text
Feature

↓

Module Composition

↓

Development

↓

Release

↓

Module Evaluation

↓

Capability Growth

```

---

核心区别：

每一次开发：

都是能力建设。

---

# 8. Capability Loop（能力循环）

这是 PSCO 最核心的差异。

---

定义：

Capability：

> 个人软件公司已经拥有，可以持续创造价值的能力。

---

例如：

不是：

```
写过一个登录页面
```

而是：

```
拥有身份认证能力
```

---

能力形成：

```text
Project Experience

↓

Module

↓

Multiple Usage

↓

Capability

```

---

# 9. Release Loop（版本循环）

软件能力必须持续演进。

---

模型：

```text
Release

↓

Feedback

↓

Improvement

↓

New Capability

```

---

例如：

Auth Module：

```text
v1

简单登录


v2

权限系统


v3

OAuth


v4

Multi-tenant Identity

```

---

最终：

模块变成长期资产。

---

# 10. PSCO Operating Cycle 总模型

最终：

```text
        Opportunity


             ↓


        Venture


             ↓


        Product


             ↓


        Feature


             ↓


        Module


             ↓


        Release


             ↓


        Capability


             ↓


        New Opportunity

```

---

这就是：

个人软件公司的生产循环。

---

---

# Document 03

# Personal Software Company Domain Model v1.0

---

# 1. 文档目的

Document 02 描述：

> 公司如何运行。

Document 03 描述：

> 系统里面到底管理什么。

即：

PSCO OS 的核心数据模型。

---

# 2. Domain Model 总览

PSCO OS 核心实体：

```text
Opportunity

↓

Venture

↓

Product

↓

Feature

↓

Module

↓

Release


辅助实体：

Decision

Experiment

Capability

Repository


---

补充说明：

上面的对象总览是 **长期完整模型**。

为了进入工程落地，当前 `v0.1` 的执行子集收敛为：

```text
Product

↓

Module

↓

Release


辅助实体：

Decision

Repository

Venture（可选）
```

---

# 3. Opportunity（机会）

## 定义

机会是：

> 尚未验证的问题、需求或商业假设。

---

属性：

```text
Opportunity

id

title

problem

target_user

hypothesis

validation_status

created_time

```

---

生命周期：

```text
Idea

↓

Research

↓

Validation

↓

Accepted

↓

Rejected

```

---

# 4. Venture（长期探索方向）

## 定义

Venture 是：

> 围绕某个长期方向建立的探索体系。

---

例如：

不是：

```
Rento
```

而是：

```
Personal Real Estate Software
```

---

属性：

```text
Venture

id

name

mission

vision

strategy

status

```

---

状态：

```text
Exploring

↓

Building

↓

Operating

↓

Archived

```

---

# 5. Product（产品）

## 定义

Product：

> 面向真实用户持续创造价值的软件产品。

---

属性：

```text
Product

id

venture_id

name

description

customer

value_proposition

business_model

stage

```

---

生命周期：

```text
Prototype

↓

MVP

↓

Growth

↓

Stable

↓

Retired

```

---

# 6. Feature（产品能力需求）

## 定义

Feature：

> 用户价值到软件能力之间的连接点。

---

属性：

```text
Feature

id

product_id

name

description

priority

status

```

---

Feature 不直接拥有代码。

它关联：

Module。

---

补充说明：

`Feature` 仍然是 PSCO 长期模型中的关键桥接对象。

但在当前 `v0.1` 中，它不进入主执行范围，优先通过 `Product -> Module` 绑定与 `Decision` 留痕承接最小闭环。

---

关系：

```text
Feature

uses

Module

```

---

# 7. Module（软件资产）

## 定义

Module：

> 可独立演化、验证、复用的软件能力单元。

---

属性：

```text
Module

id

name

category

interface

repository

version

status

```

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

补充说明：

`Commercial` 继续作为长期演进阶段保留。

但在 `v0.1` 中，推荐默认只使用：

```text
Prototype

↓

Candidate

↓

Stable
```

---

这是 PSCO 最核心资产。

---

# 8. Release（版本）

## 定义

Release：

> 软件能力的一次正式演进。

---

属性：

```text
Release

id

module_id

version

change_log

date

```

---

例如：

```text
Auth Module

v1.0

↓

v2.0

↓

v3.0

```

---

# 9. Capability（能力）

## 定义

Capability：

> 多个模块和经验形成的长期个人能力。

---

注意：

Capability 不是代码。

例如：

Module：

```
Payment Module
```

Capability：

```
Building SaaS Monetization Systems
```

---

属性：

```text
Capability

id

name

related_modules

experience_level

```

---

补充说明：

`Capability` 在长期模型中成立，但当前工程阶段已经确认：

> 在 `v0.1` 中，`Capability` 优先作为派生结果层，而不是要求用户手动维护的重实体。

也就是说，第一版重点是先沉淀模块、版本、使用关系与决策，再由系统聚合形成能力视图。

---

# 10. Decision（决策记录）

## 定义

Decision：

> 在特定上下文中的关键选择。

---

例如：

为什么选择 PostgreSQL？

---

属性：

```text
Decision

id

title

context

problem

alternatives

choice

reason

impact

status

```

---

Decision 的作用：

不是保存所有知识。

而是保存：

影响未来的判断。

---

# 11. Experiment（实验）

## 定义

Experiment：

> 用于验证假设的最小行动。

---

属性：

```text
Experiment

id

hypothesis

method

result

decision

```

---

例如：

验证：

```
用户是否愿意支付租赁管理费用？
```

---

# 12. Repository（代码仓库）

## 定义

Repository：

> Module 或 Product 的代码实现载体。

---

属性：

```text
Repository

id

url

technology

owner

binding_scope

```

---

注意：

Repository 不是核心。

只是：

Module 的实现。

---

补充说明：

在 `v0.1` 中，`Repository` 的首要职责是承接：

- Product 与代码实现之间的绑定；
- Module 与代码实现之间的手动映射；
- 后续导入与复用判断的现实锚点。

自动导入、自动同步与更复杂的多仓库治理不属于当前阻断项。

---

# 13. Entity Relationship 总图

```text
                  Opportunity

                       |

                       ↓

                  Venture

                       |

                       ↓

                  Product

                       |

                       ↓

                  Feature

                       |

                       ↓

                  Module

                       |

                       ↓

                  Release



Module

   |

   ↓

Capability



Opportunity

   |

   ↓

Experiment



Decision

   |

   ↓

所有核心对象

```

---

# 14. PSCO OS 核心数据哲学

这里需要明确：

PSCO OS 不管理：

“我做了多少任务”。

它管理：

“我形成了多少能力”。

---

传统软件管理：

```text
Task

↓

Done

```

---

PSCO：

```text
Problem

↓

Solution

↓

Capability

↓

Future Value

```

---

# 15. Document 02 + Document 03 最终结论

经过这两个文档定义：

Personal Software Company OS 已经明确：

## 它模拟的是：

一家软件公司的完整运行机制。

---

运行层：

```text
Opportunity

↓

Venture

↓

Product

↓

Feature

↓

Module

↓

Capability
```

---

数据层：

```text
Opportunity Entity

Venture Entity

Product Entity

Feature Entity

Module Entity

Release Entity

Decision Entity

Experiment Entity

```

---

下一阶段：

# Document 04 + Document 05

将继续解决：

1. **软件能力如何真正沉淀？**

   → Personal Software Company Module System v1.0

2. **这个系统如何工程实现？**

   → Personal Software Company Architecture v1.0

这两个文档会决定 PSCO OS 是否从“理念体系”进入“可开发的软件系统”。
