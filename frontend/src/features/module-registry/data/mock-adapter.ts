/**
 * Module Registry mock 适配层
 *
 * phase02-10 阶段的 mock 数据实现，通过内存数据支撑前端运行与演示。
 * phase02-11 后端主线完成后，可通过 module-registry-adapter.ts 的切换开关
 * 切换到 api-adapter.ts 的真实后端调用。
 *
 * 适配层输出的数据结构严格对齐 phase02-09 正式规格正文，
 * 不得发明第二套对象字段、状态含义或返回路径语义。
 */
import type {
  Module,
  ModuleListItem,
  ModuleDetail,
  ModuleListSearch,
  CreateModuleInput,
  CreateReleaseInput,
  Release,
  BindModuleToProductInput,
  MapModuleToRepositoryInput,
  ProductCandidate,
  RepositoryCandidate,
  ProductBinding,
  RepositoryMapping,
} from '../types'
import {
  mockModules,
  mockReleases,
  mockProductBindings,
  mockRepositoryMappings,
  mockDecisionLinks,
  mockProductCandidates,
  mockRepositoryCandidates,
} from './mock-data'

/** 模拟网络延迟 */
const delay = (ms: number = 300) => new Promise((resolve) => setTimeout(resolve, ms))

/**
 * ModuleListRead — 列表读取
 * §5.7 列表读取至少承接 name / description / status / latest_release / product_bind_count / repository_bind_count
 */
export async function fetchModuleList(search: ModuleListSearch): Promise<ModuleListItem[]> {
  await delay()
  let items = [...mockModules]

  // 筛选：statusFilter
  if (search.statusFilter && search.statusFilter !== 'all') {
    items = items.filter((m) => m.status === search.statusFilter)
  }

  // 筛选：queryText（模糊匹配 name 与 description）
  if (search.queryText) {
    const q = search.queryText.toLowerCase()
    items = items.filter(
      (m) => m.name.toLowerCase().includes(q) || m.description.toLowerCase().includes(q),
    )
  }

  // 组装列表读取模型
  return items.map((m) => {
    const moduleReleases = mockReleases.filter((r) => r.module_id === m.id)
    const latestRelease = moduleReleases.length > 0
      ? moduleReleases.sort((a, b) => b.released_at.localeCompare(a.released_at))[0].version
      : null

    return {
      id: m.id,
      name: m.name,
      description: m.description,
      status: m.status,
      latest_release: latestRelease,
      // 按真实模块-绑定关系过滤计数
      product_bind_count: mockProductBindings.filter((b) => b.module_id === m.id).length,
      repository_bind_count: mockRepositoryMappings.filter((b) => b.module_id === m.id).length,
    }
  })
}

/**
 * ModuleDetailRead — 详情读取（统一读模型宿主）
 * §5.7 详情读取至少承接核心字段、版本列表、产品绑定、仓库映射与 Decision 入口
 * §6.3 Decision 读取内嵌于此，不设独立读接口
 */
export async function fetchModuleDetail(moduleId: string): Promise<ModuleDetail> {
  await delay()
  const module = mockModules.find((m) => m.id === moduleId)
  if (!module) {
    throw new Error(`Module ${moduleId} not found`)
  }

  const releases = mockReleases
    .filter((r) => r.module_id === moduleId)
    .sort((a, b) => b.released_at.localeCompare(a.released_at))

  // 按模块过滤产品绑定与仓库映射（剥离 mock 内部的 module_id，输出对齐详情读取模型）
  const productBindings: ProductBinding[] = mockProductBindings
    .filter((b) => b.module_id === moduleId)
    .map((b) => ({ product_id: b.product_id, product_name: b.product_name }))
  const repositoryMappings: RepositoryMapping[] = mockRepositoryMappings
    .filter((b) => b.module_id === moduleId)
    .map((b) => ({ repository_id: b.repository_id, repository_name: b.repository_name }))

  // Decision 附属读取 — 内嵌于 ModuleDetailRead，不设独立读接口组
  // 按模块过滤（剥离 mock 内部的 module_id，输出对齐详情读取模型）
  const decisionLinks = mockDecisionLinks
    .filter((d) => d.module_id === moduleId)
    .map((d) => ({ decision_id: d.decision_id, decision_title: d.decision_title }))

  return {
    module,
    releases,
    product_bindings: productBindings,
    repository_mappings: repositoryMappings,
    decision_links: decisionLinks,
  }
}

/**
 * ModuleCreateWrite — 创建模块
 * §5.7 创建写入承接 CreateModule（最小字段 name / description / status）
 * 返回新建模块标识以支持前端回流到 ModuleDetailPage
 */
export async function createModule(input: CreateModuleInput): Promise<Module> {
  await delay(500)

  // 名称唯一校验 — §5.6 模块准入规则
  if (mockModules.some((m) => m.name === input.name)) {
    throw new Error(`Module name "${input.name}" already exists`)
  }

  const newModule: Module = {
    id: `mod-${Date.now()}`,
    name: input.name,
    description: input.description,
    status: input.status,
    created_at: new Date().toISOString(),
  }

  mockModules.push(newModule)
  return newModule
}

/**
 * ModuleReleaseWrite — 版本登记
 * §5.7 版本写入承接 CreateRelease（最小字段 version / status / released_at，module_id 由上下文隐式承接）
 */
export async function createRelease(input: CreateReleaseInput): Promise<Release> {
  await delay(500)

  // 校验当前模块是否存在 — Release Create 必须依附有效当前模块上下文，不得创建孤儿 release
  const module = mockModules.find((m) => m.id === input.moduleId)
  if (!module) {
    throw new Error(`Module ${input.moduleId} not found`)
  }

  const newRelease: Release = {
    id: `rel-${Date.now()}`,
    module_id: input.moduleId,
    version: input.version,
    status: input.status,
    released_at: input.releasedAt,
  }

  mockReleases.push(newRelease)
  return newRelease
}

/**
 * ModuleBindingWrite — 绑定产品
 * §4.1 BindModuleToProduct 归属 Module Registry 后端模块
 */
export async function bindModuleToProduct(input: BindModuleToProductInput): Promise<void> {
  await delay(500)

  const candidate = mockProductCandidates.find((p) => p.id === input.productId)
  if (!candidate) {
    throw new Error(`Product ${input.productId} not found`)
  }

  // 避免重复绑定
  if (!mockProductBindings.some((p) => p.module_id === input.moduleId && p.product_id === input.productId)) {
    mockProductBindings.push({
      module_id: input.moduleId,
      product_id: candidate.id,
      product_name: candidate.name,
    })
  }
}

/**
 * ModuleBindingWrite — 映射仓库
 * §4.1 MapModuleToRepository 归属 Module Registry 后端模块
 */
export async function mapModuleToRepository(input: MapModuleToRepositoryInput): Promise<void> {
  await delay(500)

  const candidate = mockRepositoryCandidates.find((r) => r.id === input.repositoryId)
  if (!candidate) {
    throw new Error(`Repository ${input.repositoryId} not found`)
  }

  if (!mockRepositoryMappings.some((r) => r.module_id === input.moduleId && r.repository_id === input.repositoryId)) {
    mockRepositoryMappings.push({
      module_id: input.moduleId,
      repository_id: candidate.id,
      repository_name: candidate.name,
    })
  }
}

/**
 * ProductBindingCandidateRead — Product 候选读取
 * §6.2 候选读取（phase02 由 Module Registry 临时承接）
 */
export async function fetchProductCandidates(): Promise<ProductCandidate[]> {
  await delay()
  return [...mockProductCandidates]
}

/**
 * RepositoryBindingCandidateRead — Repository 候选读取
 * §6.2 候选读取（phase02 由 Module Registry 临时承接）
 */
export async function fetchRepositoryCandidates(): Promise<RepositoryCandidate[]> {
  await delay()
  return [...mockRepositoryCandidates]
}
