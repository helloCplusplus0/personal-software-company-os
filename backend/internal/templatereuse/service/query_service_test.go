package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/psco/backend/internal/reusesummary"
	"github.com/psco/backend/internal/templatereuse"
	"github.com/psco/backend/internal/templatereuse/candidate"
)

type fakeCandidateReaders struct {
	rows []candidate.ProductModuleRow
	err  error
}

func (f *fakeCandidateReaders) ReadAllProductModuleBindings(ctx context.Context) ([]candidate.ProductModuleRow, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.rows, nil
}

type fakeReuseSummaryReader struct {
	result *reusesummary.ReuseSummaryReadResult
	err    error
}

func (f *fakeReuseSummaryReader) ReadReuseSummary(ctx context.Context, scope reusesummary.ReuseSummaryScope, moduleID, productID string) (*reusesummary.ReuseSummaryReadResult, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.result, nil
}

func TestGetDerivedInsightHintsRequiresReviewScopeForWeeklyReview(t *testing.T) {
	svc := NewQueryService(&fakeCandidateReaders{}, &fakeReuseSummaryReader{})

	_, err := svc.GetDerivedInsightHints(context.Background(), "candidate-1", consumerSurfaceWeeklyReview, "")
	if !errors.Is(err, templatereuse.ErrInvalidInput) {
		t.Fatalf("expected invalid input error, got %v", err)
	}
}

func TestGetDerivedInsightHintsReturnsOnlyUncoveredCapabilityGaps(t *testing.T) {
	now := time.Now().UTC()
	rows := []candidate.ProductModuleRow{
		{
			ProductID:          "product-1",
			ProductName:        "Source Product",
			ProductDescription: "desc",
			ProductCreatedAt:   now,
			ModuleID:           "module-1",
			ModuleName:         "Frontend Module",
			CapabilityKey:      stringPtr("web_frontend"),
		},
	}
	templateCandidateID := candidate.ComputeTemplateCandidateID([]string{"module-1"})
	svc := NewQueryService(
		&fakeCandidateReaders{rows: rows},
		&fakeReuseSummaryReader{
			result: &reusesummary.ReuseSummaryReadResult{
				CapabilitySummary: []reusesummary.CapabilitySummary{
					{CapabilityKey: "web_frontend", CapabilityLabel: "Web Frontend", SupportingModuleCount: 3},
					{CapabilityKey: "backend_api", CapabilityLabel: "Backend API", SupportingModuleCount: 2},
				},
			},
		},
	)

	result, err := svc.GetDerivedInsightHints(context.Background(), templateCandidateID, consumerSurfaceWeeklyReview, "weekly-review")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result.ResolutionStatus != "RESOLVED" {
		t.Fatalf("expected RESOLVED status, got %s", result.ResolutionStatus)
	}
	if len(result.Hints) != 1 {
		t.Fatalf("expected 1 uncovered capability gap hint, got %d", len(result.Hints))
	}

	hint := result.Hints[0]
	if hint.HintType != "CAPABILITY_GAP" {
		t.Fatalf("expected CAPABILITY_GAP, got %s", hint.HintType)
	}
	if hint.CapabilityKey != "backend_api" {
		t.Fatalf("expected uncovered capability backend_api, got %s", hint.CapabilityKey)
	}
	if hint.CTAKind != "GO_TO_MODULE_DETAIL" {
		t.Fatalf("expected GO_TO_MODULE_DETAIL CTA, got %s", hint.CTAKind)
	}
}

func stringPtr(v string) *string {
	return &v
}
