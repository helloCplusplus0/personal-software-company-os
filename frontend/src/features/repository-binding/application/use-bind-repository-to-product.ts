/**
 * useBindRepositoryToProduct — Repository->Product 绑定的固定 mutation 承接位。
 *
 * phase07-10 §5.5：canonical 写动作单一正式 owner。
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { repositoryBindingClient } from '../data/connect-client'

export function useBindRepositoryToProduct(): UseMutationResult<void, Error, { repositoryId: string; productId: string }, unknown> {
  const queryClient = useQueryClient()
  return useMutation<void, Error, { repositoryId: string; productId: string }, unknown>({
    mutationFn: async ({ repositoryId, productId }) => {
      await repositoryBindingClient.bindRepositoryToProduct({ repositoryId, productId })
    },
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['repository-detail', variables.repositoryId] })
      queryClient.invalidateQueries({ queryKey: ['repository-product-candidates', variables.repositoryId] })
    },
  })
}