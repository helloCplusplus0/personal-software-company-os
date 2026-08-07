import { Link } from '@tanstack/react-router'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Badge } from '@/components/ui/badge'
import { Button } from '@/components/ui/button'
import { ArrowRight } from 'lucide-react'
import type { ProductBinding, RepositoryMapping } from '../types'

interface ModuleBindingPanelProps {
  moduleId: string
  moduleName: string
  productBindings: ProductBinding[]
  repositoryMappings: RepositoryMapping[]
}

/**
 * ModuleBindingPanel — 兼容入口组件（phase04-13 收敛）
 *
 * phase04-05 / phase04-13 已冻结：
 * - Module Detail 中的绑定面板从直接写入承接位回落为只读摘要展示与兼容跳转入口
 * - 不再在 Module Detail 内直接提交 BindModuleToProduct 或 MapModuleToRepository
 * - 正式绑定写入统一迁移到 ProductDetailPage 与 RepositoryBindingDetailPage 承接
 *
 * 兼容跳转入口（phase04-13 spec "Module Detail 发起 Product/Repository 绑定动作"）：
 * - 目标未确定 → 跳转到 /products 或 /repositories 并携带 moduleId / moduleName / fromModuleDetail
 * - 目标已确定 → 跳转到 /products/:productId 或 /repositories/:repositoryId 并携带 moduleId / moduleName / fromModuleDetail
 *
 * 只读摘要展示：
 * - 继续展示已绑定的 Product 与 Repository 摘要
 * - 已绑定项可点击跳转到对应 Detail 页（目标已确定分支）
 * - 不再保留候选读取、选择器、提交按钮组成的第二主工作台
 */
export function ModuleBindingPanel({
  moduleId,
  moduleName,
  productBindings,
  repositoryMappings,
}: ModuleBindingPanelProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle>关联关系</CardTitle>
      </CardHeader>
      <CardContent className="space-y-4">
        {/* 产品绑定区 — 只读摘要 + 兼容跳转入口 */}
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm font-medium">产品绑定</span>
            {/* phase04-13 目标未确定：进入 Product Registry 列表主入口 */}
            <Button variant="outline" size="sm" asChild>
              <Link
                to="/products"
                search={{
                  fromModuleDetail: true,
                  moduleId,
                  moduleName,
                }}
              >
                <ArrowRight className="mr-2 h-4 w-4" />
                进入产品绑定
              </Link>
            </Button>
          </div>
          <div className="space-y-1">
            {productBindings.length === 0 ? (
              <p className="text-xs text-muted-foreground">未绑定产品</p>
            ) : (
              productBindings.map((pb) => (
                // phase04-13 目标已确定：跳转到 ProductDetailPage 并携带 moduleId / moduleName / fromModuleDetail
                <Link
                  key={pb.product_id}
                  to="/products/$productId"
                  params={{ productId: pb.product_id }}
                  search={{
                    fromModuleDetail: true,
                    moduleId,
                    moduleName,
                  }}
                  className="inline-block"
                >
                  <Badge variant="outline" className="cursor-pointer hover:bg-accent transition-colors">
                    {pb.product_name}
                  </Badge>
                </Link>
              ))
            )}
          </div>
        </div>

        {/* 仓库映射区 — 只读摘要 + 兼容跳转入口 */}
        <div className="space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-sm font-medium">仓库映射</span>
            {/* phase04-13 目标未确定：进入 Repository Binding 列表主入口 */}
            <Button variant="outline" size="sm" asChild>
              <Link
                to="/repositories"
                search={{
                  fromModuleDetail: true,
                  moduleId,
                  moduleName,
                }}
              >
                <ArrowRight className="mr-2 h-4 w-4" />
                进入仓库映射
              </Link>
            </Button>
          </div>
          <div className="space-y-1">
            {repositoryMappings.length === 0 ? (
              <p className="text-xs text-muted-foreground">未映射仓库</p>
            ) : (
              repositoryMappings.map((rm) => (
                // phase04-13 目标已确定：跳转到 RepositoryBindingDetailPage 并携带 moduleId / moduleName / fromModuleDetail
                <Link
                  key={rm.repository_id}
                  to="/repositories/$repositoryId"
                  params={{ repositoryId: rm.repository_id }}
                  search={{
                    fromModuleDetail: true,
                    moduleId,
                    moduleName,
                  }}
                  className="inline-block"
                >
                  <Badge variant="outline" className="cursor-pointer hover:bg-accent transition-colors">
                    {rm.repository_name}
                  </Badge>
                </Link>
              ))
            )}
          </div>
        </div>
      </CardContent>
    </Card>
  )
}
