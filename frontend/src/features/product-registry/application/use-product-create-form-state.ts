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
import { useProductCreateDraftStore } from '../stores/product-create-draft-store'

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
  clearDraft: () => void
}

// ============================================================================
// Hook
// ============================================================================

export interface UseProductCreateFormStateOptions {
  initialValues?: FormStateInitialValues
  draftKey?: string
}

export function useProductCreateFormState(
  options?: UseProductCreateFormStateOptions,
): UseProductCreateFormStateResult {
  const initialValues = options?.initialValues
  const draftKey = options?.draftKey
  const persistedDraft = useProductCreateDraftStore((state) =>
    draftKey ? state.drafts[draftKey] : undefined,
  )
  const setDraft = useProductCreateDraftStore((state) => state.setDraft)
  const clearStoredDraft = useProductCreateDraftStore((state) => state.clearDraft)
  const [name, setNameValue] = useState<string>(persistedDraft?.name ?? initialValues?.name ?? '')
  const [description, setDescriptionValue] = useState<string>(
    persistedDraft?.description ?? initialValues?.description ?? '',
  )
  const [status, setStatusValue] = useState<ProductStatus>(
    persistedDraft?.status ?? initialValues?.status ?? 'active',
  )
  const hasUserEditedRef = useRef(false)
  const lastAppliedInitialValuesRef = useRef<string>('')
  const lastHydratedDraftKeyRef = useRef<string>('')

  const initialValuesSignature = useMemo(
    () => JSON.stringify({
      name: initialValues?.name ?? '',
      description: initialValues?.description ?? '',
      status: initialValues?.status ?? 'active',
    }),
    [initialValues?.description, initialValues?.name, initialValues?.status],
  )

  useEffect(() => {
    if (draftKey && persistedDraft && lastHydratedDraftKeyRef.current !== draftKey) {
      setNameValue(persistedDraft.name)
      setDescriptionValue(persistedDraft.description)
      setStatusValue(persistedDraft.status)
      lastHydratedDraftKeyRef.current = draftKey
      lastAppliedInitialValuesRef.current = initialValuesSignature
      return
    }
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
  }, [
    draftKey,
    initialValues?.description,
    initialValues?.name,
    initialValues?.status,
    initialValuesSignature,
    persistedDraft,
  ])

  useEffect(() => {
    if (!draftKey) {
      return
    }

    setDraft(draftKey, { name, description, status })
  }, [description, draftKey, name, setDraft, status])

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

  const clearDraft = useCallback(() => {
    if (!draftKey) {
      return
    }

    clearStoredDraft(draftKey)
  }, [clearStoredDraft, draftKey])

  return {
    name,
    description,
    status,
    setName,
    setDescription,
    setStatus,
    isDirty,
    buildSubmitInput,
    clearDraft,
  }
}
