import { useState, useEffect } from 'react'
import { useQuery, useMutation } from '@tanstack/react-query'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Skeleton } from '@/components/ui/skeleton'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { fetchRepositoryModuleCandidates, mapModuleToRepository } from '../data/repository-binding-adapter'
import { toast } from 'sonner'
import type { MappedModuleSummary } from '../types'

interface RepositoryModuleMappingPanelProps {
  repositoryId: string
  mappedModules: MappedModuleSummary[]
  /** phase04-06 互斥展开：由父页面控制面板开闭 */
  open: boolean
  onOpenChange: (open: boolean) => void
  /** phase04-06 来源上下文：从 Module Detail 带入的 moduleId，用于预填候选选择 */
  prefillModuleId?: string
  onBindingSuccess: () => void
}

/**
 * RepositoryModuleMappingPanel — Repository Detail 中的 Module 映射面板
 *
 * phase04-05 组件树冻结：是该页面中唯一直接承接 MapModuleToRepository 写入的组件
 * phase04-06 状态模型：
 * - 候选读取状态：closed / pending / ready / empty / error
 * - 写入状态：idle / submitting / submit-success / submit-error
 * - 绑定成功后停留在 RepositoryBindingDetailPage 并重新读取详情结果（reread）
 * - 绑定失败时停留在面板上下文，保留当前已选候选 Module
 * - 候选 Module 为空时展示明确的无可绑定候选空状态提示
 * - 同一时刻只允许一个绑定面板处于打开态（互斥展开，由父页面控制）
 *
 * phase04-06 来源上下文预填：
 * - 从 Module Detail 携带 moduleId / moduleName / fromModuleDetail 进入时，
 *   打开面板并预填候选 Module 选择
 */
export function RepositoryModuleMappingPanel({
  repositoryId,
  mappedModules,
  open,
  onOpenChange,
  prefillModuleId,
  onBindingSuccess,
}: RepositoryModuleMappingPanelProps) {
  const [selectedModuleId, setSelectedModuleId] = useState('')
  const [submitError, setSubmitError] = useState<string | undefined>(undefined)

  // 候选读取 — phase04-06 候选读取独立于详情读取
  const { data: candidates, isLoading: candidatesLoading, isError: candidatesError } = useQuery({
    queryKey: ['repository-module-candidates', repositoryId],
    queryFn: () => fetchRepositoryModuleCandidates(repositoryId),
    enabled: open,
  })

  // phase04-06 来源上下文预填：从 Module Detail 带入 moduleId 时自动打开面板并预选
  useEffect(() => {
    if (prefillModuleId && !open) {
      onOpenChange(true)
    }
  }, [prefillModuleId, open, onOpenChange])

  // 候选列表加载完成后，若存在 prefillModuleId 且在候选列表中，则预选
  useEffect(() => {
    if (prefillModuleId && candidates && candidates.length > 0) {
      const matched = candidates.find((c) => c.module_id === prefillModuleId)
      if (matched && !selectedModuleId) {
        setSelectedModuleId(matched.module_id)
      }
    }
  }, [prefillModuleId, candidates, selectedModuleId])

  const mutation = useMutation({
    mutationFn: () => mapModuleToRepository({ repositoryId, moduleId: selectedModuleId }),
    onSuccess: () => {
      // phase04-06 绑定成功后停留在 RepositoryBindingDetailPage 并重新读取详情结果
      toast.success('模块映射成功')
      onOpenChange(false)
      setSelectedModuleId('')
      setSubmitError(undefined)
      onBindingSuccess()
    },
    onError: (error: Error) => {
      // phase04-06 绑定失败时错误停留在面板上下文，保留当前已选候选
      setSubmitError(error.message)
    },
  })

  const handleToggle = () => {
    setSubmitError(undefined)
    onOpenChange(!open)
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between gap-2">
          <span>已映射模块</span>
          <Button
            variant="outline"
            size="sm"
            onClick={handleToggle}
            disabled={mutation.isPending}
          >
            {open ? '取消' : '映射模块'}
          </Button>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        {/* 已映射 Module 列表 — 只读摘要 */}
        <div className="space-y-1">
          {mappedModules.length === 0 ? (
            <p className="text-xs text-muted-foreground">未映射模块</p>
          ) : (
            mappedModules.map((mm) => (
              <div key={mm.module_id} className="flex items-center gap-2">
                <Badge variant="outline">{mm.module_name}</Badge>
                <Badge variant={mm.module_status === 'active' ? 'default' : 'secondary'}>
                  {mm.module_status}
                </Badge>
              </div>
            ))
          )}
        </div>

        {/* 映射面板 — phase04-06 候选读取状态机 */}
        {open && (
          <div className="space-y-2 rounded-md border p-3">
            {candidatesLoading ? (
              <div className="space-y-2">
                <Skeleton className="h-8 w-full" />
                <p className="text-xs text-muted-foreground">加载候选模块...</p>
              </div>
            ) : candidatesError ? (
              <p className="text-xs text-destructive">候选模块读取失败，请重试</p>
            ) : candidates && candidates.length === 0 ? (
              // empty — phase04-06 候选为空不得误报为接口错误
              <p className="text-xs text-muted-foreground">无可映射候选模块（所有 active 模块已映射或不存在）</p>
            ) : (
              <>
                <Select value={selectedModuleId} onValueChange={setSelectedModuleId}>
                  <SelectTrigger>
                    <SelectValue placeholder="选择模块" />
                  </SelectTrigger>
                  <SelectContent>
                    {candidates?.map((mc) => (
                      <SelectItem key={mc.module_id} value={mc.module_id}>
                        {mc.module_name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                {submitError && <p className="text-xs text-destructive">{submitError}</p>}
                <Button
                  size="sm"
                  disabled={!selectedModuleId || mutation.isPending}
                  onClick={() => mutation.mutate()}
                >
                  {mutation.isPending ? '提交中...' : '确认映射'}
                </Button>
              </>
            )}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
