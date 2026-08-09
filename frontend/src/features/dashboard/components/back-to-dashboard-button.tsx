/**
 * BackToDashboardButton — 共享的"返回 Dashboard"按钮
 *
 * phase05-13 §"既有页面必须支持返回 Dashboard 导航"
 * phase05-03 / phase05-10 §8.3 返回路径规则：
 *   - 从四类 Detail（Module / Product / Repository / Decision）返回 → /dashboard
 *   - 从四类 List（Module / Product / Repository / Decision）返回 → /dashboard
 *
 * 此组件封装 useDashboardBackButton hook，仅在 fromDashboard=true 时渲染。
 * 既有 List / Detail / Create 页面组件可直接插入此组件，
 * 不需要各自重复实现 fromDashboard 检测与 navigate 逻辑。
 *
 * 多跳返回约束（phase05-10 §8.3 / phase05-13 spec）：
 *   - Detail 页同时携带 fromList=true 与 fromDashboard=true 时，
 *     必须同时保留"返回列表"与"返回 Dashboard"两个导航入口
 *   - "返回列表"使用 fromList 上下文
 *   - "返回 Dashboard"使用 fromDashboard 上下文（本组件承接）
 *
 * 一次性路由 state 承接（phase05-13 §"Dashboard 主动返回的一次性状态承接必须用路由 state"）：
 *   - 点击本按钮时，dashboardSection 通过 TanStack Router 路由 state 一次性承接
 *   - 不持久化，刷新 /dashboard 后不保留
 */
import { Button } from '@/components/ui/button'
import { LayoutDashboard } from 'lucide-react'
import { useDashboardBackButton } from '../lib/dashboard-source'

interface BackToDashboardButtonProps {
  // 按钮尺寸，默认 'sm'，与既有"返回列表"按钮对齐
  size?: 'default' | 'sm' | 'lg' | 'icon'
  // 按钮变体，默认 'ghost'，与既有"返回列表"按钮对齐
  variant?: 'default' | 'destructive' | 'outline' | 'secondary' | 'ghost' | 'link'
  // 是否在 fromDashboard=false 时强制隐藏（默认 true）
  // 若设置为 false，则即使非 Dashboard 来源也展示（当前阶段不建议）
  autoHide?: boolean
}

/**
 * BackToDashboardButton — 返回 Dashboard 按钮。
 *
 * 行为：
 * - autoHide=true（默认）：仅当 fromDashboard=true 时渲染
 * - autoHide=false：始终渲染（当前阶段未使用）
 *
 * 点击行为：
 * - 调用 useDashboardBackButton 提供的 handleBack
 * - handleBack 内部通过 navigate({ to: '/dashboard', state: { dashboardSection } }) 一次性承接
 */
export function BackToDashboardButton({
  size = 'sm',
  variant = 'ghost',
  autoHide = true,
}: BackToDashboardButtonProps) {
  const { showBackButton, handleBack } = useDashboardBackButton()

  // autoHide=true 且非 Dashboard 来源时不渲染
  if (autoHide && !showBackButton) {
    return null
  }

  return (
    <Button variant={variant} size={size} onClick={handleBack}>
      <LayoutDashboard className="mr-2 h-4 w-4" />
      返回 Dashboard
    </Button>
  )
}
