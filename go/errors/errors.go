// Package errors contains utils for returning gRPC errors from the Grafeas API.
package errors

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Newf creates a new gRPC error with the specified code and message.
func Newf(c codes.Code, format string, a ...interface{}) error {
	return status.Errorf(c, format, a...)
}
