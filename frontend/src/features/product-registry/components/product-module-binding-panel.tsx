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
import { fetchProductModuleCandidates, bindModuleToProduct } from '../data/product-registry-adapter'
import { toast } from 'sonner'
import type { BoundModuleSummary } from '../types'

interface ProductModuleBindingPanelProps {
  productId: string
  boundModules: BoundModuleSummary[]
  /** phase04-06 来源上下文：从 Module Detail 带入的 moduleId，用于预填候选选择 */
  prefillModuleId?: string
  onBindingSuccess: () => void
}

/**
 * ProductModuleBindingPanel — Product Detail 中的 Module 绑定面板
 *
 * phase04-05 组件树冻结：是 Product Detail 中唯一直接承接 BindModuleToProduct 写入的组件
 * phase04-06 状态模型：
 * - 候选读取状态：closed / pending / ready / empty / error
 * - 写入状态：idle / submitting / submit-success / submit-error
 * - 绑定成功后停留在 ProductDetailPage 并重新读取详情结果（reread）
 * - 绑定失败时停留在面板上下文，保留当前已选候选 Module
 * - 候选 Module 为空时展示明确的无可绑定候选空状态提示
 *
 * phase04-06 来源上下文预填：
 * - 从 Module Detail 携带 moduleId / moduleName / fromModuleDetail 进入时，
 *   打开面板并预填候选 Module 选择
 */
export function ProductModuleBindingPanel({
  productId,
  boundModules,
  prefillModuleId,
  onBindingSuccess,
}: ProductModuleBindingPanelProps) {
  // phase04-06 面板状态：closed / open
  const [panelOpen, setPanelOpen] = useState(false)
  const [selectedModuleId, setSelectedModuleId] = useState('')
  const [submitError, setSubmitError] = useState<string | undefined>(undefined)

  // 候选读取 — phase04-06 候选读取独立于详情读取
  const { data: candidates, isLoading: candidatesLoading, isError: candidatesError } = useQuery({
    queryKey: ['product-module-candidates', productId],
    queryFn: () => fetchProductModuleCandidates(productId),
    enabled: panelOpen,
  })

  // phase04-06 来源上下文预填：从 Module Detail 带入 moduleId 时自动打开面板并预选
  useEffect(() => {
    if (prefillModuleId && !panelOpen) {
      setPanelOpen(true)
    }
  }, [prefillModuleId, panelOpen])

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
    mutationFn: () => bindModuleToProduct({ productId, moduleId: selectedModuleId }),
    onSuccess: () => {
      // phase04-06 绑定成功后停留在 ProductDetailPage 并重新读取详情结果
      toast.success('模块绑定成功')
      setPanelOpen(false)
      setSelectedModuleId('')
      setSubmitError(undefined)
      onBindingSuccess()
    },
    onError: (error: Error) => {
      // phase04-06 绑定失败时错误停留在面板上下文，保留当前已选候选
      setSubmitError(error.message)
    },
  })

  const handleTogglePanel = () => {
    setSubmitError(undefined)
    setPanelOpen(!panelOpen)
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between gap-2">
          <span>已绑定模块</span>
          <Button
            variant="outline"
            size="sm"
            onClick={handleTogglePanel}
            disabled={mutation.isPending}
          >
            {panelOpen ? '取消' : '绑定模块'}
          </Button>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        {/* 已绑定 Module 列表 — 只读摘要 */}
        <div className="space-y-1">
          {boundModules.length === 0 ? (
            <p className="text-xs text-muted-foreground">未绑定模块</p>
          ) : (
            boundModules.map((bm) => (
              <div key={bm.module_id} className="flex items-center gap-2">
                <Badge variant="outline">{bm.module_name}</Badge>
                <Badge variant={bm.module_status === 'active' ? 'default' : 'secondary'}>
                  {bm.module_status}
                </Badge>
              </div>
            ))
          )}
        </div>

        {/* 绑定面板 — phase04-06 候选读取状态机 */}
        {panelOpen && (
          <div className="space-y-2 rounded-md border p-3">
            {candidatesLoading ? (
              // pending
              <div className="space-y-2">
                <Skeleton className="h-8 w-full" />
                <p className="text-xs text-muted-foreground">加载候选模块...</p>
              </div>
            ) : candidatesError ? (
              // error
              <p className="text-xs text-destructive">候选模块读取失败，请重试</p>
            ) : candidates && candidates.length === 0 ? (
              // empty — phase04-06 候选为空不得误报为接口错误
              <p className="text-xs text-muted-foreground">无可绑定候选模块（所有 active 模块已绑定或不存在）</p>
            ) : (
              // ready
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
                  {mutation.isPending ? '提交中...' : '确认绑定'}
                </Button>
              </>
            )}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
