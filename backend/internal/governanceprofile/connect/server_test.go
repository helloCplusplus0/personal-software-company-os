package connect

import (
	"testing"
	"time"

	"github.com/psco/backend/internal/governanceprofile"
)

func TestDomainResultToProtoIncludesMarkdownResolvable(t *testing.T) {
	timestamp := time.Date(2026, 8, 17, 12, 0, 0, 0, time.UTC)
	summary := "规则摘要"

	result := &governanceprofile.GovernanceProfileReadResult{
		Record: governanceprofile.GovernanceProfileRecord{
			RepositoryID:          "repo-1",
			ProjectProfileVersion: governanceprofile.CurrentProfileVersion,
			TrackType:             governanceprofile.RootFrozenTrackType,
			DocsWorkflowLayout:    governanceprofile.RootFrozenDocsWorkflowLayout,
			CurrentPhaseName:      governanceprofile.RootFrozenCurrentPhaseName,
			CurrentPhaseRef:       governanceprofile.RootFrozenCurrentPhaseRef,
			CurrentPhaseStatus:    governanceprofile.RootFrozenCurrentPhaseStatus,
			CreatedAt:             timestamp,
			UpdatedAt:             timestamp,
		},
		GlobalAssetBindings: []governanceprofile.GlobalAssetBinding{
			{
				Name:               "project_rules.md",
				Kind:               "rules",
				EntryRef:           "project_rules.md",
				Role:               "rules",
				StructuredSummary:  &summary,
				MarkdownResolvable: true,
			},
		},
	}

	profile := DomainResultToProto(result)
	if profile == nil {
		t.Fatal("expected governance profile proto")
	}
	if len(profile.GetGlobalAssetBindings()) != 1 {
		t.Fatalf("expected 1 global asset binding, got %d", len(profile.GetGlobalAssetBindings()))
	}
	if !profile.GetGlobalAssetBindings()[0].GetMarkdownResolvable() {
		t.Fatal("expected markdown_resolvable to be true")
	}
}
