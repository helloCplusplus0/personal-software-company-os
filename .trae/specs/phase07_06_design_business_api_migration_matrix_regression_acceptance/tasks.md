# Tasks

- [x] Task 1: 冻结 phase07-06 的输入边界与当前源码事实，建立单一矩阵的上游基线。
  - [x] SubTask 1.1: 复用 phase07-01 的 34 条 canonical RPC 总表，明确它是本次迁移矩阵的唯一业务接口上游 → `design.md` §1.1（唯一上游基线表）
  - [x] SubTask 1.2: 复用 phase07-03 的 4 条 legacy / compat endpoint inventory，明确它是本次退场证据矩阵的唯一上游 → `design.md` §1.1（4 条 legacy inventory）
  - [x] SubTask 1.3: 复用 phase07-05 的 frontend read / mutation owner 设计，明确 11 项正式写动作 owner 与 route / candidate read caller 必须进入验收映射 → `design.md` §1.1（11 项 mutation owner）
  - [x] SubTask 1.4: 盘点当前真实脚本与工具链入口 → `design.md` §1.2（工具链事实表，含 CI 缺口显式建账）

- [x] Task 2: 产出 canonical 业务接口迁移矩阵，覆盖 34 条 RPC 的迁移顺序、owner 与回归项。
  - [x] SubTask 2.1: 为每条 RPC 定义统一列结构：`service / RPC / 当前入口路径 / 方法 / 当前 owner / 目标 Connect path / 迁移 owner / 波次 / 回归项 / 最终证据` → `design.md` §2.1-2.10（34 条逐条完整矩阵）
  - [x] SubTask 2.2: 明确 9 个业务模块的波次划分与依赖关系 → `design.md` §2（Wave 1-4 波次表 + §7 执行顺序图）
  - [x] SubTask 2.3: 对 ModuleRegistryService 下 4 条仍带 compat 语义的 RPC 明确 transport inventory 与正式业务 owner 的区分规则 → `design.md` §2.10（4 条 compat RPC 单独分组，标注"transport inventory 不等于 canonical business owner"）

- [x] Task 3: 产出跨模块回归清单与 frontend owner 验收映射。
  - [x] SubTask 3.1: 为 9 模块产出模块内回归项 → `design.md` §3.1（9 模块 17 项回归）
  - [x] SubTask 3.2: 产出跨模块回归项（CR1-CR9）→ `design.md` §3.2（9 条跨模块联动 + §3.3 6 条 route 级回归）
  - [x] SubTask 3.3: 将 11 项 frontend 正式写动作 owner 升级为验收映射表 → `design.md` §5.1（11 项逐条映射，含 fixture + 触发位 + 回流检查项 + 最晚核销时点）+ §5.2（4 组 candidate read + mutation 联合验收）

- [x] Task 4: 产出 fixture、联调、退场证据与最终收口证据矩阵。
  - [x] SubTask 4.1: 复用现有 reset / acceptance 脚本，建立 fixture / 环境入口到回归项的映射 → `design.md` §4.1（5 个脚本逐项映射，含默认恢复 + 联调步骤 + 期望结果）
  - [x] SubTask 4.2: 为 4 条 legacy / compat endpoint 逐项绑定路由删除证据、handler / adapter 删除证据与替代 Connect 回归证据 → `design.md` §4.2（4 条逐项退场证据矩阵）
  - [x] SubTask 4.3: 产出 phase07-11 最终证据包结构 → `design.md` §4.3（8 份证据文件清单）+ §4.4（8 条阻断条件）

- [x] Task 5: 产出 Vite、本地脚本、proto 生成链与 CI 缺口的迁移清单。
  - [x] SubTask 5.1: 明确 /api 单一基址在 vite.config.ts、本地启动、验收脚本与部署链路中的承接要求 → `design.md` §6.1（Vite/Frontend 3 项迁移清单）
  - [x] SubTask 5.2: 明确 proto/Makefile 与 buf.gen.yaml 的当前事实、phase07 目标状态与验证命令 → `design.md` §6.2（Proto 生成链 5 项迁移清单）
  - [x] SubTask 5.3: 明确当前仓库没有现成 .github/workflows/* 的事实，并把 CI proto 生成链缺口纳入迁移清单与收口证据 → `design.md` §6.4（CI 缺口显式建账 + 替代方案）

- [x] Task 6: 完成规格一致性校验，确保本次设计可直接指导 phase07-07 ~ phase07-11。
  - [x] SubTask 6.1: 校验与 phase07-01 的 34 RPC / 9 模块 / 单一 /api 基线一致 → `design.md` §8 一致性声明
  - [x] SubTask 6.2: 校验与 phase07-03 的 legacy 退场窗口与 endpoint 级核销规则一致 → `design.md` §8 一致性声明
  - [x] SubTask 6.3: 校验与 phase07-05 的 query / application / mutation owner 收口规则一致 → `design.md` §8 一致性声明
  - [x] SubTask 6.4: 校验与 phase06-16 的 fixture / 联调 / 验收入口口径一致，不新增第二套验收事实源 → `design.md` §8 一致性声明

# Task Dependencies

- Task 2 depends on Task 1 ✅
- Task 3 depends on Task 1, Task 2 ✅
- Task 4 depends on Task 1, Task 2, Task 3 ✅
- Task 5 depends on Task 1 ✅
- Task 6 depends on Task 2, Task 3, Task 4, Task 5 ✅