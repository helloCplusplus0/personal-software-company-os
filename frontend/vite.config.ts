import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { tanstackRouter } from '@tanstack/router-plugin/vite'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    tailwindcss(),
    // TanStack Router 文件路由插件：自动从 src/routes/ 生成 routeTree.gen.ts
    // 必须放在 react() 之前
    tanstackRouter({ target: 'react', autoCodeSplitting: true }),
    react(),
  ],
  resolve: {
    alias: {
      '@': new URL('./src', import.meta.url).pathname,
    },
  },
  server: {
    // phase03-14 修复：开发期间通过 Vite proxy 转发 /api 到后端，
    // 使前端请求走同源 (localhost:5173)，彻底消除 CORS preflight 与 net::ERR_ABORTED。
    // 生产环境由 Caddy 反代统一接入，不依赖此 proxy。
    proxy: {
      '/api': {
        target: 'http://localhost:8081',
        changeOrigin: true,
        // phase03-14 修复：Vite proxy 默认在响应中添加 Connection: close，
        // 当 mutation 的 onSuccess 立即触发 GET refetch 时，Chrome 检测到 POST
        // 连接被关闭会记录 net::ERR_ABORTED（尽管 POST 实际已成功）。
        // 通过移除 Connection: close 头，允许浏览器复用连接，消除 ERR_ABORTED。
        configure: (proxy) => {
          proxy.on('proxyRes', (proxyRes) => {
            // 移除 Connection: close，允许 keep-alive 连接复用
            delete proxyRes.headers['connection']
          })
        },
      },
    },
  },
})
