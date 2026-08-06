import { createRootRouteWithContext, Outlet, Link } from '@tanstack/react-router'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
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

/** 根布局：全局导航 + 页面出口 + Toast 通知 */
function RootComponent() {
  return (
    <QueryClientProvider client={queryClient}>
      <div className="min-h-screen bg-background">
        <header className="border-b bg-card">
          <div className="container mx-auto px-4 py-3">
            <nav className="flex items-center gap-6">
              <Link to="/modules" className="text-lg font-semibold text-foreground hover:text-primary transition-colors">
                PSCO Module Registry
              </Link>
              <Link to="/decisions" className="text-lg font-semibold text-foreground hover:text-primary transition-colors">
                Decision Center
              </Link>
            </nav>
          </div>
        </header>
        <main className="container mx-auto px-4 py-6">
          <Outlet />
        </main>
      </div>
      <Toaster />
    </QueryClientProvider>
  )
}
