/**
 * use-product-create-template-handoff — Product Create 模板预填的单一 application owner。
 *
 * phase09-09 spec §"模板 handoff application owner 的最小落点"：
 *   本 owner 负责解析搜索参数、组合模板预填只读结果、拼装模板摘要 view model
 *   与生成返回/回流参数。
 *
 * 职责：
 *   - 解析 fromTemplateReuse / templateCandidateId / templateSource 搜索参数
 *   - 调用 use-template-prefill-read 获取模板预填数据
 *   - 导出 templateSummary view model（模板标题、描述、模块列表、来源标记）
 *   - 导出 prefillInitialValues（{ name: string; description: string }）
 *   - 导出 resolutionStatus（'resolved' | 'unavailable' | 'error'）
 *   - 导出 handleReturn() 基于 templateSource 的返回路径
 *   - 导出 buildSuccessSearch() 创建成功回流参数
 *   - 导出 templateSourceLabel 来源展示文案
 *
 * 不承接：
 *   - useMutation / createProduct 调用（由 useCreateDraftProduct 承接）
 *   - 直接 import TemplateReuseService 或 templateReuseClient
 */
import { useCallback, useMemo } from 'react'
import { useNavigate } from '@tanstack/react-router'
import {
  TemplateConsumerSurface,
} from '@/gen/proto/psco/template_reuse/v1/template_reuse_pb'
import { useTemplatePrefillRead } from '@/features/template-reuse/data/use-template-prefill-read'

// ============================================================================
// 搜索参数类型
// ============================================================================

export interface TemplateHandoffSearchParams {
  fromTemplateReuse?: boolean
  templateCandidateId?: string
  templateSource?: string
  /** 辅助返回链元数据（fromDashboard 等） */
  fromDashboard?: boolean
  dashboardSection?: string
  dashboardReturnTo?: string
}

// ============================================================================
// 返回类型
// ============================================================================

export type TemplateResolutionStatus = 'resolved' | 'unavailable' | 'error'

export interface TemplateSummary {
  templateTitle: string
  templateDescription: string
  moduleNames: string[]
  sourceLabel: string
  isFromTemplate: boolean
}

export interface TemplateHandoffResult {
  /** 模板摘要 view model */
  templateSummary: TemplateSummary | undefined
  /** 表单预填初始值 */
  prefillInitialValues: { name: string; description: string }
  /** 模板解析状态 */
  resolutionStatus: TemplateResolutionStatus
  /** 是否来自模板（fromTemplateReuse=true 且 templateCandidateId 非空） */
  isFromTemplate: boolean
  /** 模板来源展示文案 */
  templateSourceLabel: string
  /** 取消/返回处理函数 */
  handleReturn: () => void
  /** 创建成功回流参数构建函数 */
  buildSuccessSearch: () => Record<string, string | undefined>
  /** 模板候选 ID（用于 Product Detail 回流） */
  templateCandidateId: string
  /** 模板来源（用于 Product Detail 回流） */
  templateSource: string
}

// ============================================================================
// 辅助函数
// ============================================================================

const TEMPLATE_SOURCE_LABELS: Record<string, string> = {
  'weekly-review': 'Weekly Review',
  'dashboard': 'Dashboard',
  'product-detail': 'Product Detail',
}

const TEMPLATE_SOURCE_RETURN_PATHS: Record<string, string> = {
  'weekly-review': '/review/weekly',
  'dashboard': '/dashboard',
  'product-detail': '/products',
}

// ============================================================================
// Hook
// ============================================================================

export function useProductCreateTemplateHandoff(
  search: TemplateHandoffSearchParams,
): TemplateHandoffResult {
  const navigate = useNavigate()

  const fromTemplateReuse = search.fromTemplateReuse === true
  const templateCandidateId = search.templateCandidateId ?? ''
  const templateSource = search.templateSource ?? ''

  // 模板预填只读
  const prefillQuery = useTemplatePrefillRead(
    fromTemplateReuse && templateCandidateId !== '' ? templateCandidateId : '',
    TemplateConsumerSurface.PRODUCT_CREATE,
  )

  // 派生状态
  const isFromTemplate = useMemo(
    () => fromTemplateReuse && templateCandidateId !== '',
    [fromTemplateReuse, templateCandidateId],
  )

  const resolutionStatus = useMemo((): TemplateResolutionStatus => {
    if (!isFromTemplate) return 'unavailable'
    if (prefillQuery.isError) return 'error'
    if (prefillQuery.pageStatus === 'unavailable') return 'unavailable'
    if (prefillQuery.pageStatus === 'resolved') return 'resolved'
    return 'unavailable'
  }, [isFromTemplate, prefillQuery.isError, prefillQuery.pageStatus])

  const templateSummary = useMemo((): TemplateSummary | undefined => {
    if (!isFromTemplate) return undefined
    const prefill = prefillQuery.prefill
    if (!prefill) return undefined
    return {
      templateTitle: prefill.templateTitle,
      templateDescription: prefill.templateDescription,
      moduleNames: prefill.modules?.map((m) => m.moduleName) ?? [],
      sourceLabel: TEMPLATE_SOURCE_LABELS[templateSource] ?? templateSource,
      isFromTemplate: true,
    }
  }, [isFromTemplate, prefillQuery.prefill, templateSource])

  const prefillInitialValues = useMemo(() => {
    if (!isFromTemplate || resolutionStatus !== 'resolved') {
      return { name: '', description: '' }
    }
    const prefill = prefillQuery.prefill
    return {
      name: prefill?.suggestedProductName ?? '',
      description: prefill?.suggestedProductDescription ?? '',
    }
  }, [isFromTemplate, resolutionStatus, prefillQuery.prefill])

  const templateSourceLabel = useMemo(
    () => TEMPLATE_SOURCE_LABELS[templateSource] ?? templateSource,
    [templateSource],
  )

  const handleReturn = useCallback(() => {
    const returnPath = TEMPLATE_SOURCE_RETURN_PATHS[templateSource]
    if (returnPath) {
      navigate({ to: returnPath })
    } else {
      // 无模板来源回退：保持原有逻辑（回列表）
      navigate({ to: '/products' })
    }
  }, [navigate, templateSource])

  const buildSuccessSearch = useCallback((): Record<string, string | undefined> => {
    if (!isFromTemplate) return {}
    return {
      fromTemplateReuse: 'true',
      templateCandidateId,
      templateSource,
    }
  }, [isFromTemplate, templateCandidateId, templateSource])

  return {
    templateSummary,
    prefillInitialValues,
    resolutionStatus,
    isFromTemplate,
    templateSourceLabel,
    handleReturn,
    buildSuccessSearch,
    templateCandidateId,
    templateSource,
  }
}