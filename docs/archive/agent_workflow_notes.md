# Agent Workflow Notes

## 1. 职责

本文档只回答一件事：PSCO 当前如何稳步推进。

它不重复承载项目边界、技术基线和目录落点；这些分别以 `PSCO-summarize-feedback.md`、`project_rules.md`、`architecture_map.md` 为准。

## 2. 当前默认推进链

PSCO 当前默认遵循：

`/plan -> 复核 -> /spec -> 实现 -> 验收 -> 收口`

在正式进入实现前，必须先完成规格冻结。

## 3. 三类 workflow

### 3.1 phase

适用于结构化阶段任务：

1. 在 `plan.md` 确认是否需要新阶段
2. 进入该阶段 `/plan`
3. 产出 `architecture_plan / dev_plan / shared_baseline`
4. 复核通过后，再进入 `/spec` 与实现

### 3.2 fix

适用于局部问题与小型增强：

1. 记录 issue
2. 产出 analysis
3. 进入 `/spec`
4. 实现并验收

### 3.3 audit

适用于方案仲裁、边界冲突、路线优选：

1. 记录 issue
2. 产出 analysis
3. 在 `keep-as-is / enter-fix / enter-improvement / escalate-phase` 中选唯一去向
4. 若进入改善，再继续 `/spec` 与实现

## 4. 稳定推进的控制点

- 任何阶段都先看上游真相源，再动手
- 任何新文档都先确定分类目录，再决定文件名
- 任何结构化任务结束后，都回到共识文档和规则文档做一致性校验
- 任何未冻结的选择，都必须显式标注“待确认”，不能靠猜

## 5. 接手顺序

1. `AGENTS.md`
2. `project_rules.md`
3. `plan.md`
4. `architecture_map.md`
5. `PSCO-summarize-feedback.md`
6. `docs/README.md`
