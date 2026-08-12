package connect

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	connectrpc "connectrpc.com/connect"

	pbc "github.com/psco/backend/internal/gen/connect/psco/template_reuse/v1/template_reusev1connect"
	pb "github.com/psco/backend/internal/gen/proto/psco/template_reuse/v1"
	"github.com/psco/backend/internal/templatereuse/candidate"
	"github.com/psco/backend/internal/templatereuse/service"
)

type fakeQueryService struct {
	derivedHintsArgs struct {
		templateCandidateID string
		consumerSurface     string
		reviewScopeKey      string
	}
	listCandidatesResult *service.ListTemplateCandidatesResult
	prefillResult        *service.TemplateCandidatePrefillResult
	derivedHintsResult   *service.DerivedInsightHintsResult
	sourceSummaryResult  *service.TemplateSourceSummaryResult
}

func (f *fakeQueryService) ListTemplateCandidates(ctx context.Context) (*service.ListTemplateCandidatesResult, error) {
	if f.listCandidatesResult != nil {
		return f.listCandidatesResult, nil
	}
	return &service.ListTemplateCandidatesResult{}, nil
}

func (f *fakeQueryService) GetTemplateCandidatePrefill(ctx context.Context, templateCandidateID string) (*service.TemplateCandidatePrefillResult, error) {
	if f.prefillResult != nil {
		return f.prefillResult, nil
	}
	return &service.TemplateCandidatePrefillResult{}, nil
}

func (f *fakeQueryService) GetDerivedInsightHints(ctx context.Context, templateCandidateID, consumerSurface, reviewScopeKey string) (*service.DerivedInsightHintsResult, error) {
	f.derivedHintsArgs.templateCandidateID = templateCandidateID
	f.derivedHintsArgs.consumerSurface = consumerSurface
	f.derivedHintsArgs.reviewScopeKey = reviewScopeKey
	if f.derivedHintsResult != nil {
		return f.derivedHintsResult, nil
	}
	return &service.DerivedInsightHintsResult{
		ResolutionStatus: "RESOLVED",
		Hints:            []service.DerivedHintData{},
	}, nil
}

func (f *fakeQueryService) GetTemplateSourceSummary(ctx context.Context, templateCandidateID, templateSource string) (*service.TemplateSourceSummaryResult, error) {
	if f.sourceSummaryResult != nil {
		return f.sourceSummaryResult, nil
	}
	return &service.TemplateSourceSummaryResult{}, nil
}

func TestListTemplateCandidatesRejectsInvalidSurface(t *testing.T) {
	client := newTemplateReuseTestClient(t, &fakeQueryService{})

	_, err := client.ListTemplateCandidates(context.Background(), &pb.ListTemplateCandidatesRequest{
		ConsumerSurface: pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_PRODUCT_CREATE,
	})
	if err == nil {
		t.Fatal("expected invalid argument error")
	}
	if got := connectrpc.CodeOf(err); got != connectrpc.CodeInvalidArgument {
		t.Fatalf("expected CodeInvalidArgument, got %v", got)
	}
}

func TestGetDerivedInsightHintsForwardsConsumerSurfaceAndScope(t *testing.T) {
	fakeSvc := &fakeQueryService{}
	client := newTemplateReuseTestClient(t, fakeSvc)

	_, err := client.GetDerivedInsightHints(context.Background(), &pb.GetDerivedInsightHintsRequest{
		TemplateCandidateId: "candidate-1",
		ConsumerSurface:     pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_WEEKLY_REVIEW,
		ReviewScopeKey:      "weekly-review",
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if fakeSvc.derivedHintsArgs.templateCandidateID != "candidate-1" {
		t.Fatalf("expected candidate-1, got %s", fakeSvc.derivedHintsArgs.templateCandidateID)
	}
	if fakeSvc.derivedHintsArgs.consumerSurface != "weekly-review" {
		t.Fatalf("expected weekly-review surface, got %s", fakeSvc.derivedHintsArgs.consumerSurface)
	}
	if fakeSvc.derivedHintsArgs.reviewScopeKey != "weekly-review" {
		t.Fatalf("expected weekly-review scope, got %s", fakeSvc.derivedHintsArgs.reviewScopeKey)
	}
}

func TestTemplateReuseConnectSmokeCoversAllReadRPCs(t *testing.T) {
	fakeSvc := &fakeQueryService{
		listCandidatesResult: &service.ListTemplateCandidatesResult{
			Candidates: []candidate.TemplateCandidateData{},
		},
		prefillResult: &service.TemplateCandidatePrefillResult{
			ResolutionStatus:    "RESOLVED",
			TemplateCandidateID: "candidate-1",
			TemplateTitle:       "Template One",
		},
		derivedHintsResult: &service.DerivedInsightHintsResult{
			ResolutionStatus:      "UNAVAILABLE",
			UnavailableReasonText: "template candidate drifted",
			Hints:                 []service.DerivedHintData{},
		},
		sourceSummaryResult: &service.TemplateSourceSummaryResult{
			ResolutionStatus:    "RESOLVED",
			TemplateCandidateID: "candidate-1",
			TemplateTitle:       "Template One",
			TemplateSource:      "weekly-review",
		},
	}
	client := newTemplateReuseTestClient(t, fakeSvc)

	listRes, err := client.ListTemplateCandidates(context.Background(), &pb.ListTemplateCandidatesRequest{
		ConsumerSurface: pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_WEEKLY_REVIEW,
	})
	if err != nil {
		t.Fatalf("list candidates smoke failed: %v", err)
	}
	if listRes == nil {
		t.Fatal("expected list candidates response")
	}

	prefillRes, err := client.GetTemplateCandidatePrefill(context.Background(), &pb.GetTemplateCandidatePrefillRequest{
		TemplateCandidateId: "candidate-1",
		ConsumerSurface:     pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_PRODUCT_CREATE,
	})
	if err != nil {
		t.Fatalf("prefill smoke failed: %v", err)
	}
	if prefillRes.GetPrefill().GetResolutionStatus() != pb.TemplateResolutionStatus_TEMPLATE_RESOLUTION_STATUS_RESOLVED {
		t.Fatalf("expected resolved prefill, got %v", prefillRes.GetPrefill().GetResolutionStatus())
	}

	hintsRes, err := client.GetDerivedInsightHints(context.Background(), &pb.GetDerivedInsightHintsRequest{
		TemplateCandidateId: "candidate-1",
		ConsumerSurface:     pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_WEEKLY_REVIEW,
		ReviewScopeKey:      "weekly-review",
	})
	if err != nil {
		t.Fatalf("derived hints smoke failed: %v", err)
	}
	if hintsRes.GetResolutionStatus() != pb.TemplateResolutionStatus_TEMPLATE_RESOLUTION_STATUS_UNAVAILABLE {
		t.Fatalf("expected unavailable hints status, got %v", hintsRes.GetResolutionStatus())
	}

	sourceRes, err := client.GetTemplateSourceSummary(context.Background(), &pb.GetTemplateSourceSummaryRequest{
		TemplateCandidateId: "candidate-1",
		TemplateSource:      pb.TemplateSource_TEMPLATE_SOURCE_WEEKLY_REVIEW,
		ConsumerSurface:     pb.TemplateConsumerSurface_TEMPLATE_CONSUMER_SURFACE_PRODUCT_DETAIL,
	})
	if err != nil {
		t.Fatalf("source summary smoke failed: %v", err)
	}
	if sourceRes.GetSourceSummary().GetResolutionStatus() != pb.TemplateResolutionStatus_TEMPLATE_RESOLUTION_STATUS_RESOLVED {
		t.Fatalf("expected resolved source summary, got %v", sourceRes.GetSourceSummary().GetResolutionStatus())
	}
}

func newTemplateReuseTestClient(t *testing.T, querySvc queryServicer) pbc.TemplateReuseServiceClient {
	t.Helper()

	mux := http.NewServeMux()
	path, handler := pbc.NewTemplateReuseServiceHandler(NewServer(querySvc))
	mux.Handle(path, handler)

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)

	return pbc.NewTemplateReuseServiceClient(server.Client(), server.URL)
}
