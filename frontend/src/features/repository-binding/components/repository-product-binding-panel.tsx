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
import { fetchRepositoryProductCandidates, bindRepositoryToProduct } from '../data/repository-binding-adapter'
import { toast } from 'sonner'
import type { BoundProductSummary } from '../types'

interface RepositoryProductBindingPanelProps {
  repositoryId: string
  boundProducts: BoundProductSummary[]
  /** phase04-06 互斥展开：由父页面控制面板开闭 */
  open: boolean
  onOpenChange: (open: boolean) => void
  /** phase04-06 来源上下文：从 Product Detail 带入的 productId，用于预填候选选择 */
  prefillProductId?: string
  onBindingSuccess: () => void
}

/**
 * RepositoryProductBindingPanel — Repository Detail 中的 Product 绑定面板
 *
 * phase04-05 组件树冻结：是该页面中唯一直接承接 BindRepositoryToProduct 写入的组件
 * phase04-06 状态模型：
 * - 候选读取状态：closed / pending / ready / empty / error
 * - 写入状态：idle / submitting / submit-success / submit-error
 * - 绑定成功后停留在 RepositoryBindingDetailPage 并重新读取详情结果（reread）
 * - 绑定失败时停留在面板上下文，保留当前已选候选 Product
 * - 候选 Product 为空时展示明确的无可绑定候选空状态提示
 * - 同一时刻只允许一个绑定面板处于打开态（互斥展开，由父页面控制）
 *
 * phase04-06 来源上下文预填：
 * - 从 Product Detail 携带 productId / productName / fromProductDetail 进入时，
 *   打开面板并预填候选 Product 选择
 */
export function RepositoryProductBindingPanel({
  repositoryId,
  boundProducts,
  open,
  onOpenChange,
  prefillProductId,
  onBindingSuccess,
}: RepositoryProductBindingPanelProps) {
  const [selectedProductId, setSelectedProductId] = useState('')
  const [submitError, setSubmitError] = useState<string | undefined>(undefined)

  // 候选读取 — phase04-06 候选读取独立于详情读取
  const { data: candidates, isLoading: candidatesLoading, isError: candidatesError } = useQuery({
    queryKey: ['repository-product-candidates', repositoryId],
    queryFn: () => fetchRepositoryProductCandidates(repositoryId),
    enabled: open,
  })

  // phase04-06 来源上下文预填：从 Product Detail 带入 productId 时自动打开面板并预选
  useEffect(() => {
    if (prefillProductId && !open) {
      onOpenChange(true)
    }
  }, [prefillProductId, open, onOpenChange])

  // 候选列表加载完成后，若存在 prefillProductId 且在候选列表中，则预选
  useEffect(() => {
    if (prefillProductId && candidates && candidates.length > 0) {
      const matched = candidates.find((c) => c.product_id === prefillProductId)
      if (matched && !selectedProductId) {
        setSelectedProductId(matched.product_id)
      }
    }
  }, [prefillProductId, candidates, selectedProductId])

  const mutation = useMutation({
    mutationFn: () => bindRepositoryToProduct({ repositoryId, productId: selectedProductId }),
    onSuccess: () => {
      // phase04-06 绑定成功后停留在 RepositoryBindingDetailPage 并重新读取详情结果
      toast.success('产品绑定成功')
      onOpenChange(false)
      setSelectedProductId('')
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
          <span>已绑定产品</span>
          <Button
            variant="outline"
            size="sm"
            onClick={handleToggle}
            disabled={mutation.isPending}
          >
            {open ? '取消' : '绑定产品'}
          </Button>
        </CardTitle>
      </CardHeader>
      <CardContent className="space-y-3">
        {/* 已绑定 Product 列表 — 只读摘要 */}
        <div className="space-y-1">
          {boundProducts.length === 0 ? (
            <p className="text-xs text-muted-foreground">未绑定产品</p>
          ) : (
            boundProducts.map((bp) => (
              <div key={bp.product_id} className="flex items-center gap-2">
                <Badge variant="outline">{bp.product_name}</Badge>
                <Badge variant={bp.product_status === 'active' ? 'default' : 'secondary'}>
                  {bp.product_status}
                </Badge>
              </div>
            ))
          )}
        </div>

        {/* 绑定面板 — phase04-06 候选读取状态机 */}
        {open && (
          <div className="space-y-2 rounded-md border p-3">
            {candidatesLoading ? (
              <div className="space-y-2">
                <Skeleton className="h-8 w-full" />
                <p className="text-xs text-muted-foreground">加载候选产品...</p>
              </div>
            ) : candidatesError ? (
              <p className="text-xs text-destructive">候选产品读取失败，请重试</p>
            ) : candidates && candidates.length === 0 ? (
              // empty — phase04-06 候选为空不得误报为接口错误
              <p className="text-xs text-muted-foreground">无可绑定候选产品（所有 active 产品已绑定或不存在）</p>
            ) : (
              <>
                <Select value={selectedProductId} onValueChange={setSelectedProductId}>
                  <SelectTrigger>
                    <SelectValue placeholder="选择产品" />
                  </SelectTrigger>
                  <SelectContent>
                    {candidates?.map((pc) => (
                      <SelectItem key={pc.product_id} value={pc.product_id}>
                        {pc.product_name}
                      </SelectItem>
                    ))}
                  </SelectContent>
                </Select>
                {submitError && <p className="text-xs text-destructive">{submitError}</p>}
                <Button
                  size="sm"
                  disabled={!selectedProductId || mutation.isPending}
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
