// Package filter handles filtering the results of list methods.
package filter

import (
	"github.com/google/logger"
	"golang.org/x/net/context"
)

// Resource is the resource being filtered on.
type Resource interface{}

// Handler is a function that efficiently handles a specific filter pattern. It parses the specified
// filter to determine if it can handle it, and if it can, it returns only resources that match the
// filter, and the next page token (if available). The function always returns a bool indicating
// whether it can answer the specified filter.
//
// The base resource name
// (https://cloud.google.com/apis/design/resource_names#relative_resource_name) is an optional
// dimension the filter works on. For example it could be the project name if the list method this
// filter handler is for acts on a project level.
type Handler func(ctx context.Context, baseResourceName, filter, pageToken string, pageSize int32) (Resource, string, bool, error)

// Filterer holds functions on how to handle various filter patterns for listing resources.
type Filterer struct {
	// Handlers contain all functions that handle specific filter patterns.
	Handlers []Handler
	// DefaultHandler contains the fallback handler to handle a filter if none of the filter handlers
	// understand how to handle a filter.
	DefaultHandler Handler
}

// Filter finds the appropriate filter function to handle the specified filter and executes it to
// return filtered resources.
func (f *Filterer) Filter(ctx context.Context, baseResourceName, filter, pageToken string, pageSize int32) (Resource, string, error) {
	for _, handler := range f.Handlers {
		resources, npt, ok, err := handler(ctx, baseResourceName, filter, pageToken, pageSize)
		if !ok {
			logger.Infof("Cannot handle filter %q", filter)
			continue
		}
		if err != nil {
			return nil, "", err
		}
		return resources, npt, nil
	}

	resources, npt, _, err := f.DefaultHandler(ctx, baseResourceName, filter, pageToken, pageSize)
	return resources, npt, err
}
