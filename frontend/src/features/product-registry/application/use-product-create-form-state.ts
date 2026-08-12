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
import { useCallback, useEffect, useMemo, useRef, useState } from 'react'
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
  const [name, setNameValue] = useState<string>(initialValues?.name ?? '')
  const [description, setDescriptionValue] = useState<string>(initialValues?.description ?? '')
  const [status, setStatusValue] = useState<ProductStatus>(initialValues?.status ?? 'active')
  const hasUserEditedRef = useRef(false)
  const lastAppliedInitialValuesRef = useRef<string>('')

  const initialValuesSignature = useMemo(
    () => JSON.stringify({
      name: initialValues?.name ?? '',
      description: initialValues?.description ?? '',
      status: initialValues?.status ?? 'active',
    }),
    [initialValues?.description, initialValues?.name, initialValues?.status],
  )

  useEffect(() => {
    if (hasUserEditedRef.current) {
      return
    }
    if (lastAppliedInitialValuesRef.current === initialValuesSignature) {
      return
    }

    setNameValue(initialValues?.name ?? '')
    setDescriptionValue(initialValues?.description ?? '')
    setStatusValue(initialValues?.status ?? 'active')
    lastAppliedInitialValuesRef.current = initialValuesSignature
  }, [initialValues?.description, initialValues?.name, initialValues?.status, initialValuesSignature])

  const setName = useCallback((value: string) => {
    hasUserEditedRef.current = true
    setNameValue(value)
  }, [])

  const setDescription = useCallback((value: string) => {
    hasUserEditedRef.current = true
    setDescriptionValue(value)
  }, [])

  const setStatus = useCallback((value: ProductStatus) => {
    hasUserEditedRef.current = true
    setStatusValue(value)
  }, [])

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
