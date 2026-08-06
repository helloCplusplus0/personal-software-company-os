/// <reference types="vite/client" />

interface ImportMetaEnv {
  /**
   * 是否切换到真实后端 API。
   * - 'true'   → 使用 api-adapter.ts（当前 phase03 联调默认入口）
   * - 'false'  → 使用 mock-adapter.ts（仅保留给局部演示/历史 phase02 场景）
   *
   * 当前阶段（phase03-14）要求真实前后端联调与可重复复核，
   * 因此默认样例配置应显式开启真实 API。
   */
  readonly VITE_USE_REAL_API?: string

  /**
   * 后端 API 基础 URL。
   * 开发期间默认留空，前端通过 Vite /api proxy 同源转发到 localhost:8081；
   * 生产环境按部署入口注入实际地址。
   * 仅在 VITE_USE_REAL_API=true 时生效。
   */
  readonly VITE_API_BASE_URL?: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
