package connecterrors

import (
	"errors"
	"fmt"
	"testing"

	"connectrpc.com/connect"

	"github.com/psco/backend/internal/projectcontext"
)

func TestMapToConnectErrorProjectContextRepositoryErrors(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		err  error
		want connect.Code
	}{
		{
			name: "repository not found maps to not found",
			err:  projectcontext.ErrRepositoryNotFound,
			want: connect.CodeNotFound,
		},
		{
			name: "binding incomplete maps to failed precondition",
			err:  projectcontext.ErrRepositoryBindingIncomplete,
			want: connect.CodeFailedPrecondition,
		},
		{
			name: "wrapped binding incomplete still maps to failed precondition",
			err:  fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, projectcontext.ErrRepositoryBindingIncomplete),
			want: connect.CodeFailedPrecondition,
		},
		{
			name: "wrapped repository not found still maps to not found",
			err:  fmt.Errorf("%w: %w", projectcontext.ErrProjectContextReadFailed, projectcontext.ErrRepositoryNotFound),
			want: connect.CodeNotFound,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := MapToConnectError(tt.err)
			if got.Code() != tt.want {
				t.Fatalf("MapToConnectError(%v) code = %v, want %v", tt.err, got.Code(), tt.want)
			}

			if !errors.Is(got, tt.err) {
				t.Fatalf("MapToConnectError(%v) should preserve wrapped error", tt.err)
			}
		})
	}
}
