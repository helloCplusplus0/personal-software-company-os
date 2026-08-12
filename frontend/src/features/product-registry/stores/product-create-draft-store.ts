/**
 * Product Create 草稿 Store
 *
 * phase09-06 / phase09-10 要求：
 * - 从 Module 补齐页返回 Product Create 时，必须恢复同一条正式 create form state 主线
 * - 草稿恢复不得依赖 search 参数混存字段，也不得退化为组件本地临时状态
 *
 * 持久化策略（sessionStorage）：
 * - 仅承接 Product Create 正式表单字段草稿
 * - 生命周期与当前标签页会话一致，关闭标签页后自动重置
 */
import { create } from 'zustand'
import { createJSONStorage, persist } from 'zustand/middleware'
import type { ProductStatus } from '../types'

export interface ProductCreateDraftSnapshot {
  name: string
  description: string
  status: ProductStatus
}

interface ProductCreateDraftStore {
  drafts: Record<string, ProductCreateDraftSnapshot>
  setDraft: (key: string, draft: ProductCreateDraftSnapshot) => void
  clearDraft: (key: string) => void
}

export const useProductCreateDraftStore = create<ProductCreateDraftStore>()(
  persist(
    (set) => ({
      drafts: {},
      setDraft: (key, draft) =>
        set((state) => ({
          drafts: {
            ...state.drafts,
            [key]: draft,
          },
        })),
      clearDraft: (key) =>
        set((state) => {
          const nextDrafts = { ...state.drafts }
          delete nextDrafts[key]
          return { drafts: nextDrafts }
        }),
    }),
    {
      name: 'psco-product-create-draft',
      storage: createJSONStorage(() => sessionStorage),
    },
  ),
)
