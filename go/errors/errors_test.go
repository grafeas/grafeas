package errors

import (
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestNewf(t *testing.T) {
	err := Newf(codes.InvalidArgument, "invalid occurrence: missing field")

	s, ok := status.FromError(err)
	if !ok {
		t.Errorf("Unable to get a status from error %v", err)
	}
	if s.Code() != codes.InvalidArgument {
		t.Errorf("Got status code %v, want InvalidArgument", s.Code())
	}
}
