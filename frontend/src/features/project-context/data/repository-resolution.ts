/**
 * 受控 repository 候选解析工具。
 *
 * phase12-05 / 06：间接 repository-scoped 页面只能在唯一候选成立时
 * 回到共享只读主线，禁止在多候选场景下静默挑选一个 repository。
 */
export interface RepositoryCandidateLike {
  repository_id: string
  repository_name?: string
}

export function resolveUniqueRepositoryCandidate<T extends RepositoryCandidateLike>(
  candidates: readonly T[],
): T | null {
  const uniqueCandidates = Array.from(
    new Map(
      candidates
        .filter((candidate) => candidate.repository_id !== '')
        .map((candidate) => [candidate.repository_id, candidate]),
    ).values(),
  )

  return uniqueCandidates.length === 1 ? uniqueCandidates[0] : null
}
