import { useState } from 'react'
import { useQuery, useMutation } from '@tanstack/react-query'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from '@/components/ui/select'
import { fetchProductCandidates, fetchRepositoryCandidates, bindModuleToProduct, mapModuleToRepository } from '../data/module-registry-adapter'
import { toast } from 'sonner'
import type { ProductBinding, RepositoryMapping } from '../types'

interface ModuleBindingPanelProps {
  moduleId: string
  productBindings: ProductBinding[]
  repositoryMappings: RepositoryMapping[]
  onBindingSuccess: () => void
}

type PanelMode = 'closed' | 'product' | 'repository'

/**
 * ModuleBindingPanel — 绑定面板
 * §4.1 统一承接 BindModuleToProduct 与 MapModuleToRepository
 * §8.4 同一时刻只允许一个绑定面板处于打开态
 * §8.4 绑定成功后停留在 ModuleDetailPage，重新读取绑定结果
 * §8.4 绑定失败时错误停留在面板上下文，保留当前选择
 */
export function ModuleBindingPanel({
  moduleId,
  productBindings,
  repositoryMappings,
  onBindingSuccess,
}: ModuleBindingPanelProps) {
  // §8.4 面板状态：closed / open-idle / submitting / submit-success / submit-error
  const [panelMode, setPanelMode] = useState<PanelMode>('closed')
  const [selectedProductId, setSelectedProductId] = useState('')
  const [selectedRepositoryId, setSelectedRepositoryId] = useState('')
  const [submitError, setSubmitError] = useState<string | undefined>(undefined)

  // 候选读取
  const { data: productCandidates } = useQuery({
    queryKey: ['product-candidates'],
    queryFn: fetchProductCandidates,
    enabled: panelMode === 'product',
  })
  const { data: repositoryCandidates } = useQuery({
    queryKey: ['repository-candidates'],
    queryFn: fetchRepositoryCandidates,
    enabled: panelMode === 'repository',
  })

  const productMutation = useMutation({
    mutationFn: () => bindModuleToProduct({ moduleId, productId: selectedProductId }),
    onSuccess: () => {
      toast.success('产品绑定成功')
      setPanelMode('closed')
      setSelectedProductId('')
      setSubmitError(undefined)
      onBindingSuccess()
    },
    onError: (error: Error) => {
      // §8.4 绑定失败时错误停留在面板上下文
      setSubmitError(error.message)
    },
  })

  const repositoryMutation = useMutation({
    mutationFn: () => mapModuleToRepository({ moduleId, repositoryId: selectedRepositoryId }),
    onSuccess: () => {
      toast.success('仓库映射成功')
      setPanelMode('closed')
      setSelectedRepositoryId('')
      setSubmitError(undefined)
      onBindingSuccess()
    },
    onError: (error: Error) => {
      setSubmitError(error.message)
    },
  })

  const handleOpenPanel = (mode: PanelMode) => {
    // §8.4 同一时刻只允许一个绑定面板打开
    setSubmitError(undefined)
    setPanelMode(panelMode === mode ? 'closed' : mode)
  }

  const submitting = productMutation.isPending || repositoryMutation.isPending

  return (
    <Card>
      <CardHeader>
        <CardTitle>关联关系</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {/* 产品绑定区 */}
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm font-medium">产品绑定</span>
            <Button
              variant="outline"
              size="sm"
              onClick={() => handleOpenPanel('product')}
              disabled={submitting && panelMode !== 'product'}
            >
              {panelMode === 'product' ? '取消' : '绑定产品'}
            </Button>
          </div>
          <div className="space-y-1">
            {productBindings.length === 0 ? (
              <p className="text-xs text-muted-foreground">未绑定产品</p>
            ) : (
              productBindings.map((pb) => (
                <div key={pb.product_id} className="flex items-center gap-2">
                  <Badge variant="outline">{pb.product_name}</Badge>
                </div>
              ))
            )}
          </div>

          {/* 产品绑定面板 */}
          {panelMode === 'product' && (
            <div className="mt-2 space-y-2 rounded-md border p-3">
              <Select value={selectedProductId} onValueChange={setSelectedProductId}>
                <SelectTrigger>
                  <SelectValue placeholder="选择产品" />
                </SelectTrigger>
                <SelectContent>
                  {productCandidates?.map((pc) => (
                    <SelectItem key={pc.id} value={pc.id}>
                      {pc.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              {submitError && <p className="text-xs text-destructive">{submitError}</p>}
              <Button
                size="sm"
                disabled={!selectedProductId || submitting}
                onClick={() => productMutation.mutate()}
              >
                {submitting ? '提交中...' : '确认绑定'}
              </Button>
            </div>
          )}
        </div>

        {/* 仓库映射区 */}
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm font-medium">仓库映射</span>
            <Button
              variant="outline"
              size="sm"
              onClick={() => handleOpenPanel('repository')}
              disabled={submitting && panelMode !== 'repository'}
            >
              {panelMode === 'repository' ? '取消' : '映射仓库'}
            </Button>
          </div>
          <div className="space-y-1">
            {repositoryMappings.length === 0 ? (
              <p className="text-xs text-muted-foreground">未映射仓库</p>
            ) : (
              repositoryMappings.map((rm) => (
                <div key={rm.repository_id} className="flex items-center gap-2">
                  <Badge variant="outline">{rm.repository_name}</Badge>
                </div>
              ))
            )}
          </div>

          {/* 仓库映射面板 */}
          {panelMode === 'repository' && (
            <div className="mt-2 space-y-2 rounded-md border p-3">
              <Select value={selectedRepositoryId} onValueChange={setSelectedRepositoryId}>
                <SelectTrigger>
                  <SelectValue placeholder="选择仓库" />
                </SelectTrigger>
                <SelectContent>
                  {repositoryCandidates?.map((rc) => (
                    <SelectItem key={rc.id} value={rc.id}>
                      {rc.name}
                    </SelectItem>
                  ))}
                </SelectContent>
              </Select>
              {submitError && <p className="text-xs text-destructive">{submitError}</p>}
              <Button
                size="sm"
                disabled={!selectedRepositoryId || submitting}
                onClick={() => repositoryMutation.mutate()}
              >
                {submitting ? '提交中...' : '确认映射'}
              </Button>
            </div>
          )}
        </div>
      </CardContent>
    </Card>
  )
}
