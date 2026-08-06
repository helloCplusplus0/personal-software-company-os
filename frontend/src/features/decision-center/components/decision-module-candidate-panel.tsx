/**
 * DecisionModuleCandidatePanel — 候选 Module 读取与目标关联面板
 *
 * §5.10 候选读取为 Decision Detail 的附属读取：
 * - 候选来源为当前已存在的 modules
 * - 候选范围同时覆盖 active 与 archived 的 Module
 * - 排序采用 status(active 优先) -> module_name 升序
 * - 已建立 Decision -> Module 关联的目标不得再次出现在候选中
 * - 无可关联候选时返回空列表，不误报为接口错误
 *
 * phase03-13 spec §"候选读取与最小目标关联"：
 * - 通过 useQuery 读取候选列表，区分 pending / ready / empty / error
 * - 通过 useMutation 触发 LinkDecisionToTarget
 * - 关联成功后停留在当前 DecisionDetailPage
 * - onSuccess 中失效当前详情、候选列表与决策列表相关查询
 *
 * phase03-05 组件树冻结：
 * - DecisionModuleCandidatePanel 承接候选读取与目标选择
 * - DecisionLinkActions 内联于此组件
 */
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import { Plus, AlertCircle } from 'lucide-react'
import { toast } from 'sonner'
import { fetchDecisionModuleCandidates, linkDecisionToTarget } from '../data/decision-center-adapter'
import type { DecisionModuleCandidate } from '../types'

interface DecisionModuleCandidatePanelProps {
  decisionId: string
}

/** 候选 Module 状态标签 */
const STATUS_LABEL: Record<string, string> = {
  active: 'Active',
  archived: 'Archived',
}

export function DecisionModuleCandidatePanel({ decisionId }: DecisionModuleCandidatePanelProps) {
  const queryClient = useQueryClient()

  // §5.10 候选读取
  const { data: candidates, isLoading, isError, error, refetch } = useQuery({
    queryKey: ['decision-module-candidates', decisionId],
    queryFn: () => fetchDecisionModuleCandidates(decisionId),
    enabled: Boolean(decisionId),
  })

  // §5.10 LinkDecisionToTarget 写入
  const linkMutation = useMutation({
    mutationFn: (moduleId: string) =>
      linkDecisionToTarget({
        decisionId,
        target_type: 'module',
        module_id: moduleId,
      }),
    onSuccess: () => {
      // phase03-13 spec：失效当前详情、候选列表与决策列表
      queryClient.invalidateQueries({ queryKey: ['decision-detail', decisionId] })
      queryClient.invalidateQueries({ queryKey: ['decision-module-candidates', decisionId] })
      queryClient.invalidateQueries({ queryKey: ['decision-list'] })
      toast.success('关联成功')
    },
    onError: (err: Error) => {
      toast.error('关联失败：' + err.message)
    },
  })

  /** 渲染候选列表项 */
  const renderCandidate = (c: DecisionModuleCandidate) => (
    <div
      key={c.module_id}
      className="flex items-center justify-between rounded-md border p-3"
    >
      <div className="flex items-center gap-2">
        <span className="text-sm font-medium">{c.module_name}</span>
        <Badge variant="outline" className="text-xs">
          {STATUS_LABEL[c.status] ?? c.status}
        </Badge>
      </div>
      <Button
        type="button"
        size="sm"
        variant="outline"
        disabled={linkMutation.isPending}
        onClick={() => linkMutation.mutate(c.module_id)}
      >
        <Plus className="mr-1 h-3 w-3" />
        关联
      </Button>
    </div>
  )

  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-base">关联模块</CardTitle>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          // pending 状态
          <div className="space-y-2">
            {Array.from({ length: 2 }).map((_, i) => (
              <Skeleton key={i} className="h-12 w-full" />
            ))}
          </div>
        ) : isError ? (
          // error 状态
          <div className="rounded-md border border-destructive/50 bg-destructive/10 p-3">
            <div className="flex items-center gap-2 text-destructive">
              <AlertCircle className="h-4 w-4" />
              <p className="text-sm">候选读取失败：{(error as Error).message}</p>
            </div>
            <Button type="button" variant="outline" size="sm" className="mt-2" onClick={() => refetch()}>
              重试
            </Button>
          </div>
        ) : !candidates || candidates.length === 0 ? (
          // empty 状态 — 不误报为接口错误
          <p className="text-sm text-muted-foreground py-2">
            暂无可关联的模块候选
          </p>
        ) : (
          // ready 状态
          <div className="space-y-2">
            {candidates.map(renderCandidate)}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
