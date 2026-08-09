import { createRootRouteWithContext, Outlet, Link } from '@tanstack/react-router'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { useEffect, useState } from 'react'
import { Menu, X } from 'lucide-react'
import { Toaster } from '@/components/ui/sonner'

/** TanStack Query 客户端 — phase02-10 阶段使用默认配置 */
const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 0,
      refetchOnWindowFocus: false,
    },
  },
})

export const Route = createRootRouteWithContext()({
  component: RootComponent,
})

/**
 * 全局导航项。
 * Dashboard 作为一级导航首项（phase05-13），其余为四大功能模块入口。
 */
const NAV_ITEMS = [
  { to: '/dashboard', label: 'Dashboard' },
  { to: '/modules', label: 'PSCO Module Registry' },
  { to: '/decisions', label: 'Decision Center' },
  { to: '/products', label: 'Product Registry' },
  { to: '/repositories', label: 'Repository Binding' },
] as const

/** 根布局：全局导航 + 页面出口 + Toast 通知 */
function RootComponent() {
  // 移动端汉堡菜单展开状态（PC 端导航常驻，无需此状态）
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false)

  // Escape 键关闭移动端菜单（无障碍最佳实践：键盘用户需有关闭出口）
  useEffect(() => {
    if (!mobileMenuOpen) return
    const handleKey = (e: KeyboardEvent) => {
      if (e.key === 'Escape') setMobileMenuOpen(false)
    }
    document.addEventListener('keydown', handleKey)
    return () => document.removeEventListener('keydown', handleKey)
  }, [mobileMenuOpen])

  return (
    <QueryClientProvider client={queryClient}>
      <div className="min-h-screen bg-background">
        {/*
          导航容器：固定行高 h-14，PC 端水平导航常驻，移动端汉堡按钮。
          phase05-13 体验修复：原 flex items-center gap-6 的水平导航在窄屏
          （5 项 × text-lg + gap-6 ≈ 476px）溢出 header，撑出横向滚动条，
          与 main（响应式不溢出）错位。改为响应式：md+ 水平、md- 汉堡菜单。

          phase05-13 移动端菜单 overlay 化（解决内容跳动）：
          - 菜单用 absolute 浮在 header 下方，脱离文档流，main 不下压/上弹
          - 配合 fixed backdrop 点击关闭 + Escape 键关闭
          - header 设 relative 作为菜单 absolute 定位上下文
        */}
        <header className="relative border-b bg-card">
          <div className="container mx-auto px-4">
            <div className="flex h-14 items-center justify-between gap-4">
              {/* PC 端主导航：md+ 水平排列 */}
              <nav aria-label="主导航" className="hidden items-center gap-6 md:flex">
                {NAV_ITEMS.map((item) => (
                  <Link
                    key={item.to}
                    to={item.to}
                    className="text-lg font-semibold text-foreground hover:text-primary transition-colors"
                  >
                    {item.label}
                  </Link>
                ))}
              </nav>

              {/*
                移动端汉堡按钮：md 以下显示。
                用原生 button + lucide 图标，避免引入第二套 UI 框架。
              */}
              <button
                type="button"
                className="inline-flex items-center justify-center rounded-md p-2 text-foreground hover:bg-muted transition-colors md:hidden"
                aria-label={mobileMenuOpen ? '关闭导航菜单' : '打开导航菜单'}
                aria-expanded={mobileMenuOpen}
                aria-controls="mobile-nav"
                onClick={() => setMobileMenuOpen((v) => !v)}
              >
                {mobileMenuOpen ? (
                  <X className="h-5 w-5" aria-hidden="true" />
                ) : (
                  <Menu className="h-5 w-5" aria-hidden="true" />
                )}
              </button>
            </div>
          </div>

          {/*
            移动端菜单 overlay：absolute 浮在 header 下方，脱离文档流不推动 main。
            - backdrop：fixed 覆盖整个视口，点击关闭（aria-hidden，纯鼠标交互；
              键盘用户由汉堡按钮的 X 图标 + Escape 键承担关闭）
            - 菜单：absolute top-full 贴在 header 底部，z-50 在 backdrop 之上，
              border-b + shadow-lg 强化浮层视觉
            仅在 mobileMenuOpen 时渲染，PC 端 md:hidden 不显示。
          */}
          {mobileMenuOpen && (
            <>
              <div
                aria-hidden="true"
                onClick={() => setMobileMenuOpen(false)}
                className="fixed inset-0 z-40 bg-black/40 md:hidden"
              />
              <nav
                id="mobile-nav"
                aria-label="移动端导航"
                className="absolute left-0 right-0 top-full z-50 border-b bg-card shadow-lg md:hidden"
              >
                <div className="container mx-auto px-4 py-2">
                  {NAV_ITEMS.map((item) => (
                    <Link
                      key={item.to}
                      to={item.to}
                      className="block rounded-md px-3 py-2 text-base font-medium text-foreground hover:bg-muted transition-colors"
                      onClick={() => setMobileMenuOpen(false)}
                    >
                      {item.label}
                    </Link>
                  ))}
                </div>
              </nav>
            </>
          )}
        </header>
        <main className="container mx-auto px-4 py-6">
          <Outlet />
        </main>
      </div>
      <Toaster />
    </QueryClientProvider>
  )
}
