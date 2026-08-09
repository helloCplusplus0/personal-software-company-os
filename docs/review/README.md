# Review Index

## 1. 职责

本目录承载专家评审与交叉汇总文档，作为 PSCO 方案讨论、评审与决策参考的留档区。

评审文档不直接承担当前 workflow 的正式规则，但作为最终共识与后续规格的上游输入参考。

## 2. 当前文档

### 2.1 第一轮专家评审

- [PSCO_Evaluation-GPT54.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO_Evaluation-GPT54.md)
- [PSCO_Review_deepseek-v4-flash.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO_Review_deepseek-v4-flash.md)
- [PSCO-Design-Review-GLM-52.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-Design-Review-GLM-52.md)
- [PSCO-evaluation-deepseek-v4-pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-evaluation-deepseek-v4-pro.md)
- [PSCO-review-qwen37-pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-review-qwen37-pro.md)

### 2.2 第二轮交叉汇总

- [PSCO-summarize-feedback-dsv4flash.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-summarize-feedback-dsv4flash.md)
- [PSCO-summarize-feedback-GLM52.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-summarize-feedback-GLM52.md)
- [PSCO-summarize-feedback-GPT54.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-summarize-feedback-GPT54.md)

### 2.3 阶段方向参考

- [PSCO-mvp02-GLM52.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-GLM52.md) — mvp0.1 验收完结后，由 GLM-5.2 以 GPT54 共识标准为参照输出的 mvp0.2 推进方向计划，作为下一阶段正式 phase 规划的上游参考
- [PSCO-next-phase-mvp02-GPT54.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-next-phase-mvp02-GPT54.md) — 基于 `PSCO_0.md ~ PSCO_4.md` 与 `phase01 ~ phase05` 已验收现实，由 GPT54 视角输出的下一阶段推进方向评审文档
- [PSCO-mvp02-deepseekv4flash.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-deepseekv4flash.md) — 基于 `PSCO_0.md ~ PSCO_4.md` 与 `phase01 ~ phase05` 已验收现实，由 DeepSeek-V4-Flash 视角输出的下一阶段推进方向标准计划文档；在 Onboarding / Review / Derived Intelligence 三线基础上，额外强调闭合 `mvp_spec §7.3/§7.4` 导出 / 备份规格合规负债
- [PSCO-mvp02-summarize-feedback-deepseekv4flash.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-summarize-feedback-deepseekv4flash.md) — 对 `docs/review/` 下五份 MVP0.2 方向评审（GPT54 / GLM-5.2 / DS-Pro / DS-Flash / Qwen3.7-Pro）的交叉总结；明确支持"不扩实体 / 导出闭合 / dry-run 独立 / Decision 复用后移"，反对 Camp 2"复用感知优先"的主线排序，主张 Operating Loop Foundation（Onboarding → Review → Derived）并把复用/复利内容纳入第三阶段
- [PSCO-mvp02-deepseekv4pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-deepseekv4pro.md) — 结合其 2026-08-04 独立方案评审回看，对 mvp0.2 三线方向与导出 / 备份负债闭合给出更强工程论证
- [PSCO-mvp02-qwen37pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-qwen37pro.md) — 以复用感知、模板级复用、Local First 导出与真实项目 dry-run 为重点的 mvp0.2 方向评审
- [PSCO-mvp02-summarize-feedback-GPT54.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-summarize-feedback-GPT54.md) — 对五份 mvp0.2 方向文档的交叉比较与立场归档，明确哪些观点支持、部分支持或反对，作为后续正式 `/plan` 的参考仲裁入口
- [PSCO-mvp02-summarize-feedback-deepseekv4pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-summarize-feedback-deepseekv4pro.md) — 对五份 mvp0.2 方向评审的系统性汇总与仲裁；识别"经营闭环优先"与"复利感知优先"两条路线，主张二者融合：阶段一 Onboarding+数据主权+复用感知基础，阶段二 Review Loop+模板复用+派生智能深化+dry-run
- [PSCO-mvp02-summarize-feedback-qwen37pro.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-summarize-feedback-qwen37pro.md) — 对五份 mvp0.2 方向评审的总结与仲裁建议，明确支持/反对的观点及理由，作为后续正式 `/plan` 的参考依据
- [PSCO-mvp02-summarize-feedback-GLM52.md](file:///home/dell/Projects/personal-software-company-os/docs/review/PSCO-mvp02-summarize-feedback-GLM52.md) — 作为"复利感知优先"路线提出者（`PSCO-mvp02-GLM52.md` 作者）的立场二次校准；诚实修正原始文档三项短板（预冻结 phase 命名、未纳入 Onboarding、Decision 整体后移过激），坚持 `module_reuse_summary` 与 `capability_summary` 必须在第一阶段完整落地，支持 DS-Pro 提出的融合方案并做一项关键强化

### 2.4 迁移跳转

- [Personal Software Company OS v2.0.md](file:///home/dell/Projects/personal-software-company-os/docs/review/Personal%20Software%20Company%20OS%20v2.0.md) -> `TECH_STACK_BASELINE.md`

## 3. 规则

- 评审与交叉汇总文档归入 `docs/review/`，不进入 `docs/audit/`
- `docs/audit/` 只承接内部审计工作流（`audit_issue` / `audit_analysis`）
- 评审文档不直接承担当前阶段正式规则
