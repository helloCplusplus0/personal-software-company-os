/// <reference types="vite/client" />

interface ImportMetaEnv {
  /**
   * 是否切换到真实后端 API。
   * - 'true'  → 使用 api-adapter.ts（调用 Go 后端）
   * - 未设置  → 使用 mock-adapter.ts（内存 mock 数据）
   *
   * 上游规格：phase02-11 spec §"前端临时适配层必须能切换到真实后端"
   */
  readonly VITE_USE_REAL_API?: string

  /**
   * 后端 API 基础 URL，默认 http://localhost:8080。
   * 仅在 VITE_USE_REAL_API=true 时生效。
   */
  readonly VITE_API_BASE_URL?: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
