/**
 * useCreateRelease — Release Create 的固定 mutation 承接位。
 *
 * phase07-10 §5.5：canonical 写动作单一正式 owner。
 */
import { useMutation, useQueryClient, type UseMutationResult } from '@tanstack/react-query'
import { timestampDate, timestampFromDate } from '@bufbuild/protobuf/wkt'
import { moduleRegistryClient } from '../data/connect-client'
import { ReleaseStatus } from '@/gen/proto/psco/module_registry/v1/module_registry_pb'
import type { CreateReleaseInput, Release } from '../types'

export type UseCreateRelease = UseMutationResult<Release, Error, CreateReleaseInput & { moduleId: string }, unknown>

export function useCreateRelease(): UseCreateRelease {
  const queryClient = useQueryClient()
  return useMutation<Release, Error, CreateReleaseInput & { moduleId: string }, unknown>({
    mutationFn: async (input) => {
      const res = await moduleRegistryClient.createRelease({
        moduleId: input.moduleId,
        version: input.version,
        status: input.status === 'active' ? ReleaseStatus.ACTIVE : input.status === 'archived' ? ReleaseStatus.ARCHIVED : ReleaseStatus.UNSPECIFIED,
          releasedAt: timestampFromDate(new Date(input.releasedAt)),
        })
      const r = res.release
      if (!r) throw new Error('版本创建失败')
      return {
        id: r.id ?? '',
        module_id: r.moduleId ?? '',
        version: r.version ?? '',
        status: (r.status === ReleaseStatus.ACTIVE ? 'active' : r.status === ReleaseStatus.ARCHIVED ? 'archived' : '') as Release['status'],
        released_at: r.releasedAt ? timestampDate(r.releasedAt).toISOString() : '',
      }
    },
    onSuccess: (_data, variables) => {
      queryClient.invalidateQueries({ queryKey: ['module-detail', variables.moduleId] })
      queryClient.invalidateQueries({ queryKey: ['module-list'] })
    },
  })
}
