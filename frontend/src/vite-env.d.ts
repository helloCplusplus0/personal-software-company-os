/// <reference types="vite/client" />

interface ImportMetaEnv {
  /**
   * 历史联调开关，当前 phase07-10 前端正式主线不再依赖它。
   *
   * 旧的 `api-adapter.ts / mock-adapter.ts` 已在 Connect 主线切换后退场，
   * 浏览器侧统一通过单一 `/api` 基址访问后端。
   */
  readonly VITE_USE_REAL_API?: string

  /**
   * 预留的后端 API 基础 URL。
   *
   * 当前 phase07-10 默认口径：
   * - 开发环境：前端通过 Vite `/api` proxy 同源转发到 `localhost:8081`
   * - 部署环境：由 Caddy 统一承接 `/api` 反代
   * - 隔离验收环境：允许显式指定后端 origin 或 /api 根路径，
   *   例如 `http://127.0.0.1:8082` 或 `http://127.0.0.1:8082/api`
   *
   * 正式业务调用不应再为不同切片注入第二套 API 基址。
   */
  readonly VITE_API_BASE_URL?: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
