/**
 * Product Detail 来源上下文透传辅助函数
 *
 * phase04-06 来源上下文透传规则：
 * - 从 Product Detail 进入 Repository Binding 时，必须继续携带 Product Detail 自己的来源上下文
 * - 使用 product 前缀区分，避免与 Repository Binding 自身的来源参数冲突
 * - 返回 Product Detail 时，基于这些参数恢复 Product Detail 的来源标记，不得退化为 direct-entry
 *
 * 参数映射：
 * - Product Detail 来源为 fromList → 透传 productFromList + productQueryText + productStatusFilter
 * - Product Detail 来源为 fromModuleDetail → 透传 productFromModuleDetail + productModuleId + productModuleName
 * - Product Detail 来源为 direct-entry → 不透传任何 product 前缀参数
 */

/**
 * Product Detail 的来源上下文参数（从 ProductDetailPage 的 search 派生）
 */
export interface ProductSourceContext {
  fromList?: boolean
  queryText?: string
  statusFilter?: string
  fromModuleDetail?: boolean
  moduleId?: string
  moduleName?: string
}

/**
 * Repository Binding 页面的完整 search 对象（包含自身来源 + Product Detail 来源透传）
 */
export interface RepositoryBindingSearch extends ProductSourceContext {
  // Repository Binding 自身的来源参数
  fromProductDetail?: boolean
  productId?: string
  productName?: string
  fromModuleDetail?: boolean
  moduleId?: string
  moduleName?: string
  // Product Detail 来源上下文透传
  productFromList?: boolean
  productQueryText?: string
  productStatusFilter?: string
  productFromModuleDetail?: boolean
  productModuleId?: string
  productModuleName?: string
}

/**
 * 从 Product Detail 的来源上下文构造透传参数
 *
 * 用于从 Product Detail 进入 Repository Binding（列表/详情/创建）时，
 * 把 Product Detail 自己的来源上下文以 product 前缀传递过去。
 *
 * @param productSource Product Detail 的来源上下文
 * @returns 透传参数对象（product 前缀），如果 Product Detail 是 direct-entry 则返回空对象
 */
export function buildProductSourceTransit(
  productSource: ProductSourceContext,
): Record<string, unknown> {
  const transit: Record<string, unknown> = {}

  if (productSource.fromList) {
    // Product Detail 来源是 Product List → 透传 productFromList + productQueryText / productStatusFilter
    transit.productFromList = true
    if (productSource.queryText) {
      transit.productQueryText = productSource.queryText
    }
    transit.productStatusFilter = productSource.statusFilter ?? 'all'
  } else if (productSource.fromModuleDetail) {
    // Product Detail 来源是 Module Detail → 透传 productFromModuleDetail + productModuleId / productModuleName
    transit.productFromModuleDetail = true
    if (productSource.moduleId) {
      transit.productModuleId = productSource.moduleId
    }
    if (productSource.moduleName) {
      transit.productModuleName = productSource.moduleName
    }
  }
  // direct-entry → 不透传任何 product 前缀参数

  return transit
}

/**
 * 从 Repository Binding 的 search 参数恢复 Product Detail 的来源搜索参数
 *
 * 用于从 Repository Binding（列表/详情/创建）返回到 Product Detail 时，
 * 基于透传参数恢复 Product Detail 的来源标记，使其能按真实来源继续返回。
 *
 * @param search Repository Binding 页面的 search 参数
 * @returns Product Detail 的来源搜索参数（fromList/fromModuleDetail + 相应参数），direct-entry 时返回空对象
 */
export function buildProductDetailSearchFromTransit(
  search: RepositoryBindingSearch,
): Record<string, unknown> {
  const productSearch: Record<string, unknown> = {}

  if (search.productFromList) {
    // Product Detail 来源是 Product List → 恢复 fromList + queryText / statusFilter
    productSearch.fromList = true
    if (search.productQueryText) {
      productSearch.queryText = search.productQueryText
    }
    productSearch.statusFilter = search.productStatusFilter ?? 'all'
  } else if (search.productFromModuleDetail) {
    // Product Detail 来源是 Module Detail → 恢复 fromModuleDetail + moduleId / moduleName
    productSearch.fromModuleDetail = true
    if (search.productModuleId) {
      productSearch.moduleId = search.productModuleId
    }
    if (search.productModuleName) {
      productSearch.moduleName = search.productModuleName
    }
  }
  // direct-entry → 不携带来源标记（Product Detail 返回列表时落默认筛选参数）

  return productSearch
}

/**
 * 从 Repository Binding 的 search 中提取需要继续透传的 product 前缀参数
 *
 * 用于在 Repository Binding 内部导航（列表→详情→创建）时，
 * 继续携带 Product Detail 的来源透传参数，不丢失上下文。
 *
 * @param search Repository Binding 页面的 search 参数
 * @returns 需要继续透传的 product 前缀参数对象
 */
export function extractProductSourceTransit(
  search: RepositoryBindingSearch,
): Record<string, unknown> {
  const transit: Record<string, unknown> = {}

  if (search.productFromList) {
    transit.productFromList = true
    if (search.productQueryText) {
      transit.productQueryText = search.productQueryText
    }
    transit.productStatusFilter = search.productStatusFilter ?? 'all'
  } else if (search.productFromModuleDetail) {
    transit.productFromModuleDetail = true
    if (search.productModuleId) {
      transit.productModuleId = search.productModuleId
    }
    if (search.productModuleName) {
      transit.productModuleName = search.productModuleName
    }
  }

  return transit
}
