# Personal Software Company OS v2.0

# 个人开发者长期软件资产架构与技术栈标准

> 当前状态说明：本文档已从评审资料提升为根级正式技术基线文档。
> 从现在开始，PSCO 项目及其后续项目的技术栈选择，统一以本文档为准；不得再以 `AGENTS-OLD.md` 或临时偏好替代。

**版本：v2.0**
**状态：长期基准架构（Baseline Architecture）**
**适用对象：个人开发者 / 一人软件公司 / 长生命周期软件资产建设**

---

# 0. 文档定位

本文档定义未来 5～10 年个人开发者的软件工程标准。

目标：

不是追求：

* 最新技术
* 最大规模
* 企业级复杂度

而是建立：

> 一个个人开发者可以长期维护、低成本运行、持续积累的软件生产体系。

核心指标：

| 指标     | 目标    |
| ------ | ----- |
| 开发效率   | 最高    |
| AI协作效率 | 最高    |
| 运行成本   | 最低    |
| 维护复杂度  | 最低    |
| 资产复用率  | 最高    |
| 生命周期   | 10年以上 |

---

# 1. 核心战略原则

## 1.1 软件资产化，而不是项目化

传统：

```
需求
 ↓
开发
 ↓
上线
 ↓
维护
 ↓
废弃
```

个人软件公司：

```
基础能力资产

        |
        |
  ----------------

  产品A

  产品B

  产品C

```

每个项目必须沉淀：

* UI组件
* 数据模型
* API规范
* 认证模块
* 权限系统
* 部署方案
* 业务经验

---

# 1.2 架构第一原则

## Modular Monolith First

默认：

> 模块化单体优先。

不默认：

* 微服务
* Kubernetes
* 服务网格
* 消息中间件

原因：

个人开发者最大风险：

不是性能不足。

而是：

> 系统复杂度超过个人维护能力。

---

# 2. 最终架构决策：双路线模型

经过讨论，确定：

个人开发者不采用单一技术栈。

采用：

# Product Track

和

# Durable System Track

---

# 3. Product Track（产品快速迭代路线）

## 定位

适合：

* Rento
* SaaS
* 管理系统
* AI应用
* Web工具
* MVP产品

核心目标：

> 快速创造价值。

---

## 标准架构

```
              React

                |

              Hono

                |

           PostgreSQL

```

---

# 技术栈

## Frontend

```
React
+
Vite
+
TypeScript
```

---

## Router

```
TanStack Router
```

理由：

* 完整类型安全
* AI代码生成友好
* 参数自动推导

---

## 数据请求

```
TanStack Query
```

负责：

* API缓存
* 请求状态
* 服务端数据同步

---

## 客户端状态

```
Zustand
```

负责：

* UI状态
* 临时状态
* 用户偏好

---

## UI

```
Tailwind CSS

+

shadcn/ui
```

原则：

组件代码属于自己。

避免：

大型UI框架锁定。

---

# Backend

```
Hono
+
TypeScript
```

定位：

完整业务后端。

负责：

* API
* 用户逻辑
* 权限
* 数据访问
* 工作流

---

# Database

```
PostgreSQL
```

负责：

OLTP：

* 用户
* 房源
* 合同
* 订单
* 状态

---

# ORM

新项目：

```
Drizzle ORM
```

原因：

* SQL优先
* 轻量
* AI友好
* 无复杂生成流程

已有项目：

```
Prisma继续使用
```

原则：

> 不为了架构纯洁而重写业务资产。

---

# API规范

## Schema First

禁止：

```
Backend定义字段

↓

Frontend猜字段
```

采用：

```
Schema

↓

Backend

↓

Frontend

```

技术：

```
Zod

+

Hono RPC

```

效果：

前后端：

端到端类型安全。

---

# 4. Durable System Track（长期运行系统路线）

## 定位

适合：

* 世界模型
* 数据基础设施
* 交易系统
* 实时系统
* 大规模分析
* 核心计算系统

目标：

> 低资源、长期稳定运行。

---

# 标准架构

```
                 React

                    |

              Go Backend

                    |

             PostgreSQL

                    |

              Rust Engine

```

---

# Backend

## Go

Go承担：

完整后端职责：

* HTTP API
* 用户系统
* 权限
* 业务逻辑
* 数据访问
* 服务管理
* 后台任务

原因：

* 单二进制
* 内存低
* 部署简单
* 长期稳定

### Go 路线默认后端栈

在 `Durable System Track` 中，Go 后端默认采用以下最小栈：

* `net/http`
* `chi`
* `pgx / pgxpool`
* `log/slog`

职责划分：

* `chi`：HTTP 路由、子路由挂载、中间件编排
* `net/http`：基础传输抽象
* `pgx / pgxpool`：PostgreSQL 驱动与连接池
* `log/slog`：结构化日志

原则：

> `chi` 是轻量传输层与路由装配工具，不是合同定义工具。

> Go 路线默认保持 SQL-first 与标准库优先，不为“框架完整性”额外叠加第二套服务框架。

---

# Compute Layer

## Rust

Rust不是普通后端。

定位：

> 高性能计算引擎。

负责：

* 文档解析
* 搜索引擎
* 图计算
* 向量计算
* 数据处理
* 实时计算

---

# 5. Hono 与 Go 的最终关系

## 不采用：

```
React

↓

Hono

↓

Go

↓

Rust
```

原因：

这会导致：

* Node运行时
* Go运行时

同时存在。

增加：

* 部署复杂度
* 内存占用
* 调试成本

---

最终规则：

## Hono路线：

```
React
+
Hono
+
Postgres
```

## Go路线：

```
React
+
Go
+
Postgres
+
Rust
```

二者二选一。

---

# 6. 多语言协作标准

问题：

TypeScript、Go、Rust无法共享类型。

解决：

不是强行统一语言。

而是：

## Contract First

架构：

```
             Schema


               |


    -----------------------


    TS        Go        Rust

```

---

推荐：

## Protocol Buffers

用于：

* API
* RPC
* 服务通信

例如：

```
proto

 |
 |
 +--- TypeScript

 +--- Go

 +--- Rust

```

### Contract First 的落地规则

在 `Durable System Track` 中：

* `.proto` 是唯一长期合同源
* `buf lint / breaking` 是合同演进门禁
* HTTP JSON 可以作为浏览器友好的传输形式存在
* 但 JSON 传输层不得形成与 `.proto` 并列的第二套 canonical contract
* 若 Go 路线需要在 HTTP 之上承接业务合同，默认优先采用 `ConnectRPC` 作为 `.proto` 对齐的正式业务传输层
* `chi` 继续承担顶层路由、子路由挂载、中间件编排与非业务端点承载；`healthz / readyz / metrics / debug` 一类基础设施端点不必强行纳入 `.proto`

这意味着：

* 路由、中间件、HTTP handler 可以由 `chi + net/http` 承接
* 但对外字段、枚举、错误语义、response envelope 必须与 `.proto` 单值一致
* 若存在 JSON request / response DTO，它们只能从 `.proto` 单向派生或显式映射，不得自行演化出第二套语义
* 业务接口不应再把手写 JSON DTO 当作长期事实源；若存在存量 JSON 业务端点，应视为兼容适配层，而不是第二套 canonical API

因此，Go 路线的正确关系是：

```
.proto
  ↓
ConnectRPC / Go / TypeScript 生成物或显式映射
  ↓
chi + net/http transport shell
```

而不是：

```
chi JSON contract
  ↘
   .proto contract
```

---

# 7. 移动端标准

## 默认：

```
React Native

+

Expo
```

原因：

共享：

* TypeScript
* API Client
* Schema
* 业务模型

架构：

```
React Web


      |

 Shared Packages


      |

React Native

```

---

原则：

不是：

```
Web代码=Mobile代码
```

而是：

```
业务能力共享
```

---

# 8. 数据资产架构

核心原则：

业务数据和长期资产分离。

架构：

```
PostgreSQL

      |

 Export

      |

 Parquet

      |

 DuckDB

```

---

## PostgreSQL

事务：

```
用户

订单

合同

状态

权限
```

---

## Parquet

长期资产：

```
历史数据

事件

日志

分析数据

训练数据
```

---

## DuckDB

分析：

```
统计

研究

模型准备

探索
```

---

# 9. 后台任务设计

默认：

不要引入消息队列。

采用：

```
Application

     |

Scheduler

     |

Task

```

工具：

* cron
* systemd timer
* 内置job

---

只有达到：

* 高并发
* 大规模异步
* 多服务通信

才引入：

* NATS
* Kafka
* Redis Queue

---

# 10. Runtime标准

## Product Track

生产：

```
Node.js LTS
```

原因：

* 生态成熟
* 稳定

---

开发：

```
Bun
```

用途：

* 包管理
* 脚本执行
* 开发加速

注意：

Bun不是必须生产运行环境。

---

## Durable Track

生产：

```
Go Binary

+

Rust Binary
```

直接运行：

```
systemd
```

---

# 11. 部署标准

核心：

# Single Server First

默认：

```
Linux VPS


Caddy


Application


PostgreSQL


systemd

```

---

不默认：

Docker。

原因：

个人长期运行：

Docker增加：

* 镜像管理
* 构建复杂度
* 存储消耗

---

部署：

```
GitHub

↓

Build

↓

Release

↓

Server

↓

systemd restart

```

---

# 12. 项目结构标准

推荐：

```
software-company/


apps/

 ├── web

 ├── mobile

 ├── api



packages/

 ├── ui

 ├── schema

 ├── auth

 ├── storage

 ├── utils



database/


docs/


```

---

# 13. 技术选择决策树

以后任何项目：

## Step 1

是不是业务产品？

例如：

Rento

SaaS

管理系统

是：

```
React

+

Hono

+

Postgres

```

---

## Step 2

是否长期7×24运行？

是否资源敏感？

是：

```
React

+

Go

+

Postgres
```

---

## Step 3

是否存在计算瓶颈？

是：

增加：

```
Rust Engine
```

---

# 14. 明确禁止事项

## 不因为流行引入：

❌ Kubernetes

❌ 微服务

❌ Docker全流程

❌ GraphQL

❌ Kafka

❌ Redis缓存层

❌ Elasticsearch

除非：

真实业务证明需要。

---

# 15. 最终技术栈表

| 领域     | 标准                       |
| ------ | ------------------------ |
| Web    | React + Vite             |
| Mobile | React Native + Expo      |
| 语言     | TypeScript               |
| 快速后端   | Hono                     |
| 长期后端   | Go                       |
| 计算引擎   | Rust                     |
| 数据库    | PostgreSQL               |
| ORM    | Drizzle                  |
| 已有ORM  | Prisma                   |
| Schema | Zod                      |
| RPC    | Hono RPC / Protobuf      |
| 状态管理   | TanStack Query + Zustand |
| 路由     | TanStack Router          |
| UI     | Tailwind + shadcn/ui     |
| 分析数据库  | DuckDB                   |
| 数据格式   | Parquet                  |
| 部署     | Caddy + systemd          |
| 运行方式   | Single Server First      |

---

# 16. 最终架构理念

最终形成：

```
              价值创造


                 |

        -----------------

        |               |

 Product Track     Durable Track


 React             React

 Hono              Go

 Postgres          Postgres


                   |

                Rust


```

核心思想：

> **Hono用于快速创造产品，Go用于长期运行系统，Rust用于构建核心能力。**

> **不要提前复杂化，也不要依赖未来重构。通过明确边界，让不同技术路线长期共存。**

---

# 17. 最终一句话原则

> **个人软件公司不是选择一种“最强语言”，而是建立一套可以持续十年积累的软件生产系统。**

这份 v2.0 应作为后续 Rento-miniX、AI World Model、交易系统等项目的技术基准文档。
