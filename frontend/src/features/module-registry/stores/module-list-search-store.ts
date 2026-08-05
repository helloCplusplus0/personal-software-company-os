/**
 * Module List 搜索上下文 Store
 *
 * §7.4 / §8.4 返回路径规则：
 * - 用户从 ModuleCreatePage 或 ModuleDetailPage 返回 ModuleListPage 时，
 *   必须恢复原有 queryText 与 statusFilter
 *
 * 列表搜索参数本身冻结在路由 URL 层（/modules 的 validateSearch），
 * 此 store 仅作为客户端侧的"最后一次列表搜索"缓存，
 * 供创建页 / 详情页在主动返回时读取并回传给 /modules 路由。
 *
 * 持久化策略（sessionStorage）：
 * - 使用 sessionStorage 而非内存，保证用户在详情页/创建页刷新后仍能恢复列表上下文
 * - sessionStorage 的生命周期与浏览器标签页会话一致，关闭标签页后重置为默认值
 * - 新会话（如从外部链接直达详情页）没有"原搜索上下文"，返回列表时使用默认值
 *
 * 定位：客户端状态（Zustand），不是服务端状态（TanStack Query）。
 * 不作为第二套数据主线，仅服务返回路径的上下文恢复。
 */
import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'
import type { ModuleStatus } from '../types'

/** 列表搜索上下文快照 */
export interface ModuleListSearchContext {
  queryText?: string
  statusFilter: ModuleStatus | 'all'
}

interface ModuleListSearchStore {
  /** 最后一次列表搜索上下文 */
  lastSearch: ModuleListSearchContext
  /** 由列表页在搜索参数变化时同步写入 */
  setLastSearch: (search: ModuleListSearchContext) => void
}

export const useModuleListSearchStore = create<ModuleListSearchStore>()(
  persist(
    (set) => ({
      lastSearch: { statusFilter: 'all' },
      setLastSearch: (search) => set({ lastSearch: search }),
    }),
    {
      name: 'psco-module-list-search',
      storage: createJSONStorage(() => sessionStorage),
    },
  ),
)
