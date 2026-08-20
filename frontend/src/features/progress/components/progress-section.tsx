/**
 * ProgressSection — Repository detail 挂载的"项目进度"区外壳与编排
 *
 * 2026-08-19 用户 UI 反馈裁决（修订 phase15-05 组件树冻结条目）：
 * - 进度区是维护工作台（表单录入 + 删除 + 过滤），视觉重量须与
 *   "已绑定产品 / 相关决策"等工作台卡片对齐——容器 Card 化
 *   （card-title text-base font-semibold），不再沿用
 *   StandardReadonlySummary 的裸 section 轻量只读摘要模式（定位错位）
 * - 页面级第三全宽区块定位不变（Standard 摘要之后，不进 grid 工作台区）；
 *   Standard 摘要区保持 phase14 先例不动（用户裁决：不同步统一）
 * - 内部子区域保持紧凑密度（text-xs），表单与时间轴以 border-t pt-3 分隔
 */
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { ProgressCurrentPhaseCard } from './progress-current-phase-card'
import { ProgressEventForm } from './progress-event-form'
import { ProgressTimelineList } from './progress-timeline-list'

interface ProgressSectionProps {
  repositoryId: string
}

export function ProgressSection({ repositoryId }: ProgressSectionProps) {
  return (
    <Card className="min-w-0">
      <CardHeader>
        <CardTitle>项目进度</CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        <ProgressCurrentPhaseCard repositoryId={repositoryId} />
        <div className="border-t pt-3">
          <ProgressEventForm repositoryId={repositoryId} />
        </div>
        <div className="border-t pt-3">
          <ProgressTimelineList repositoryId={repositoryId} />
        </div>
      </CardContent>
    </Card>
  )
}
