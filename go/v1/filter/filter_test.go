package filter_test

import (
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-cmp/cmp"
	"github.com/grafeas/grafeas/go/v1/filter"
	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFilter(t *testing.T) {
	byID := func(ctx context.Context, baseResourceName, filter, pageToken string, pageSize int32) (filter.Resource, string, bool, error) {
		if strings.HasPrefix(filter, "noteId = ") {
			return []*gpb.Note{{Name: "CVE-UH-OH-01"}}, "", true, nil
		}
		return nil, "", false, nil
	}
	defaultHandler := func(ctx context.Context, baseResourceName, filter, pageToken string, pageSize int32) (filter.Resource, string, bool, error) {
		return []*gpb.Note{{Name: "CVE-UH-OH-99"}}, "", true, nil
	}

	f := filter.Filterer{
		Handlers:       []filter.Handler{byID},
		DefaultHandler: defaultHandler,
	}

	tests := []struct {
		desc      string
		filter    string
		wantNotes []*gpb.Note
	}{
		{
			desc:   "no filter functions match, default filter function is used",
			filter: `updateTime = "2019-01-01"`,
			wantNotes: []*gpb.Note{
				{
					Name: "CVE-UH-OH-99",
				},
			},
		},
		{
			desc:   "a filter function match",
			filter: `noteId = "CVE-UH-OH-01"`,
			wantNotes: []*gpb.Note{
				{
					Name: "CVE-UH-OH-01",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx := context.Background()

			resources, _, err := f.Filter(ctx, "projects/my-proj", tt.filter, "", 0)
			if err != nil {
				t.Fatalf(`Filter("projects/my-proj", %q, "", 0) failed with: %v`, tt.filter, err)
			}
			notes := resources.([]*gpb.Note)

			if diff := cmp.Diff(tt.wantNotes, notes, cmp.Comparer(proto.Equal)); diff != "" {
				t.Errorf("Filter(\"projects/my-proj\", %q, \"\", 0) returned diff -want +got\n%s", tt.filter, diff)
			}
		})
	}
}

func TestFilterErrors(t *testing.T) {
	byID := func(ctx context.Context, baseResourceName, filter, pageToken string, pageSize int32) (filter.Resource, string, bool, error) {
		if strings.HasPrefix(filter, "noteId = ") {
			return nil, "", true, status.Errorf(codes.Internal, "error executing filter")
		}
		return nil, "", false, nil
	}
	defaultHandler := func(ctx context.Context, baseResourceName, filter, pageToken string, pageSize int32) (filter.Resource, string, bool, error) {
		return nil, "", true, status.Errorf(codes.InvalidArgument, "argument not valid")
	}

	f := filter.Filterer{
		Handlers:       []filter.Handler{byID},
		DefaultHandler: defaultHandler,
	}

	tests := []struct {
		desc        string
		filter      string
		wantErrCode codes.Code
	}{
		{
			desc:        "no filter functions match, default filter function is used, but there is an error",
			filter:      `updateTime = "2019-01-01"`,
			wantErrCode: codes.InvalidArgument,
		},
		{
			desc:        "a filter function match, but there is an error",
			filter:      `noteId = "CVE-UH-OH-01"`,
			wantErrCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			ctx := context.Background()

			_, _, err := f.Filter(ctx, "projects/my-proj", tt.filter, "", 0)
			if c := status.Code(err); c != tt.wantErrCode {
				t.Errorf(`Filter("projects/my-proj", %q, "", 0) got error code %v, want %v`, tt.filter, c, tt.wantErrCode)
			}
		})
	}
}
