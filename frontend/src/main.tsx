import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { createRouter, RouterProvider } from '@tanstack/react-router'
import { routeTree } from './routeTree.gen'
import './index.css'

/** TanStack Router 实例 — 文件路由自动生成 routeTree.gen.ts */
const router = createRouter({ routeTree })

// 类型注册：让 useNavigate / useSearch / Link 等获得完整类型推导
declare module '@tanstack/react-router' {
  interface Register {
    router: typeof router
  }
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <RouterProvider router={router} />
  </StrictMode>,
)
