/**
 * useMapModuleToRepository — Module->Repository 映射的固定 mutation 承接位。
 *
 * phase07-10 §5.5：canonical 写动作单一正式 owner。
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { repositoryBindingClient } from '../data/connect-client'

export function useMapModuleToRepository(): UseMutationResult<void, Error, { repositoryId: string; moduleId: string }, unknown> {
  const queryClient = useQueryClient()
  return useMutation<void, Error, { repositoryId: string; moduleId: string }, unknown>({
    mutationFn: async ({ repositoryId, moduleId }) => {
      await repositoryBindingClient.mapModuleToRepository({ repositoryId, moduleId })
    },
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['repository-detail', variables.repositoryId] })
      queryClient.invalidateQueries({ queryKey: ['repository-module-candidates', variables.repositoryId] })
    },
  })
}