/**
 * useBindModuleToProduct — Module->Product 绑定的固定 mutation 承接位。
 *
 * phase07-10 §5.5：canonical 写动作单一正式 owner。
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { productRegistryClient } from '../data/connect-client'

export function useBindModuleToProduct(): UseMutationResult<void, Error, { productId: string; moduleId: string }, unknown> {
  const queryClient = useQueryClient()
  return useMutation<void, Error, { productId: string; moduleId: string }, unknown>({
    mutationFn: async ({ productId, moduleId }) => {
      await productRegistryClient.bindModuleToProduct({ productId, moduleId })
    },
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['product-detail', variables.productId] })
      queryClient.invalidateQueries({ queryKey: ['product-module-candidates', variables.productId] })
    },
  })
}