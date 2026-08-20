/**
 * ProgressCurrentPhaseCard — 当前 phase 派生卡
 *
 * phase15-05 §"DP-1 裁决"：
 * - 数据源唯一经 use-repository-progress-read（GetProjectBrief.progress 投影，
 *   与 agent 消费主路径同源）；组件内零派生逻辑——无任何从事件集合计算
 *   当前 phase / 最新任务的代码路径
 * - 空值两情形（从未开始 / 全部完结，current_phase_key 空串同型零值）
 *   统一文案"暂无进行中 phase"，不做区分（Obs-1 闭环）
 * - 完结态承接：phase_completed 后当前卡转空态统一文案，完结事实经时间轴
 *   最新 phase_completed 事件行组合确认（非当前卡派生）
 */
import { useRepositoryProgressRead } from '../data/use-repository-progress-read'

interface ProgressCurrentPhaseCardProps {
  repositoryId: string
}

export function ProgressCurrentPhaseCard({ repositoryId }: ProgressCurrentPhaseCardProps) {
  const { data, isLoading, isError } = useRepositoryProgressRead(repositoryId)

  if (isLoading) {
    return <p className="text-xs text-muted-foreground">加载中...</p>
  }
  if (isError) {
    return <p className="text-xs text-destructive">进度摘要读取失败</p>
  }

  // 空态：current_phase_key 空串（含从未开始 / 全部完结两情形，统一文案不区分）
  if (!data || !data.current_phase_key) {
    return (
      <div className="space-y-1">
        <p className="text-xs text-muted-foreground">暂无进行中 phase</p>
        {!data?.latest_task_completed ? (
          <p className="text-xs text-muted-foreground">暂无任务完成记录</p>
        ) : null}
      </div>
    )
  }

  return (
    <div className="space-y-1">
      {/* 进行中态：current_phase_key（code/mono 样式）+ current_phase_label */}
      <div className="flex flex-wrap items-center gap-1.5">
        <code className="text-xs">{data.current_phase_key}</code>
        <span className="text-xs">{data.current_phase_label}</span>
      </div>
      {/* 最新完成任务行：task_key + title + occurred_at（浏览器本地时区展示，DP-3） */}
      {data.latest_task_completed ? (
        <p className="min-w-0 truncate text-xs text-muted-foreground">
          {data.latest_task_completed.task_key} {data.latest_task_completed.title}
          {data.latest_task_completed.occurred_at
            ? ` ${data.latest_task_completed.occurred_at.toLocaleString()}`
            : ''}
        </p>
      ) : (
        <p className="text-xs text-muted-foreground">暂无任务完成记录</p>
      )}
    </div>
  )
}
