/**
 * use-product-create-form-state — Product Create 表单的正式单一 form state owner。
 *
 * phase09-09 spec §"Product Create 的 form state 回收"：
 *   本 owner 承接 name / description / status 字段的单值状态，
 *   回收 ProductCreateForm 组件本地 useState 为正式 form state 主线。
 *
 * 职责：
 *   - 承接 name / description / status 字段状态
 *   - 接受 initialValues 以支持模板预填
 *   - 导出 isDirty 判断与 buildSubmitInput 组装
 *
 * 不承接：
 *   - useMutation / createProduct 调用（由 useCreateDraftProduct 承接）
 *   - 模板来源逻辑（由 use-product-create-template-handoff 承接）
 */
import { useCallback, useMemo, useState } from 'react'
import type { ProductStatus } from '@/features/product-registry/types'

// ============================================================================
// 类型
// ============================================================================

export interface FormStateInitialValues {
  name?: string
  description?: string
  status?: ProductStatus
}

export interface CreateProductSubmitInput {
  name: string
  description: string
  status: ProductStatus
}

export interface UseProductCreateFormStateResult {
  name: string
  description: string
  status: ProductStatus
  setName: (value: string) => void
  setDescription: (value: string) => void
  setStatus: (value: ProductStatus) => void
  isDirty: boolean
  buildSubmitInput: () => CreateProductSubmitInput
}

// ============================================================================
// Hook
// ============================================================================

export function useProductCreateFormState(
  initialValues?: FormStateInitialValues,
): UseProductCreateFormStateResult {
  const [name, setName] = useState<string>(initialValues?.name ?? '')
  const [description, setDescription] = useState<string>(initialValues?.description ?? '')
  const [status, setStatus] = useState<ProductStatus>(initialValues?.status ?? 'active')

  const isDirty = useMemo(
    () => name.trim() !== '' || description.trim() !== '',
    [name, description],
  )

  const buildSubmitInput = useCallback((): CreateProductSubmitInput => {
    return {
      name: name.trim(),
      description: description.trim(),
      status,
    }
  }, [name, description, status])

  return {
    name,
    description,
    status,
    setName,
    setDescription,
    setStatus,
    isDirty,
    buildSubmitInput,
  }
}