import { Link } from '@tanstack/react-router'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Badge } from '@/components/ui/badge'
import { Plus } from 'lucide-react'
import type { Release } from '../types'

interface ModuleReleaseListSectionProps {
  releases: Release[]
  moduleId: string
}

/**
 * ModuleReleaseListSection — 版本列表区
 * §5.7 详情读取承接版本列表
 * §8.3 默认归属于 ModuleDetailPage
 */
export function ModuleReleaseListSection({ releases, moduleId }: ModuleReleaseListSectionProps) {
  return (
    <Card>
      <CardHeader>
        <CardTitle className="flex items-center justify-between">
          <span>版本列表</span>
          <Button variant="outline" size="sm" asChild>
            <Link to="/modules/$moduleId/releases/new" params={{ moduleId }}>
              <Plus className="mr-1 h-3 w-3" />
              登记版本
            </Link>
          </Button>
        </CardTitle>
      </CardHeader>
      <CardContent>
        {releases.length === 0 ? (
          <p className="text-xs text-muted-foreground py-3 text-center">
            尚无版本登记
          </p>
        ) : (
          <div className="space-y-2">
            {releases.map((release) => (
              <div
                key={release.id}
                className="flex items-center justify-between rounded-md border px-3 py-2"
              >
                <div className="flex items-center gap-3">
                  <span className="font-mono text-sm font-medium">{release.version}</span>
                  <Badge variant="outline">{release.status}</Badge>
                </div>
                <span className="text-xs text-muted-foreground">
                  {new Date(release.released_at).toLocaleDateString()}
                </span>
              </div>
            ))}
          </div>
        )}
      </CardContent>
    </Card>
  )
}
