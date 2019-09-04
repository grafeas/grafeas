// Package filter handles filtering results for list methods.
package filter

import (
	"github.com/google/logger"
	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
	"golang.org/x/net/context"
)

// ListNotesFilterFn is a function that efficiently handles a specific filter pattern. It parses the
// specified filter to determine if it can handle it, and if it can, it returns only notes that
// match the filter, and the next page token (if available). The function always returns a bool
// indicating whether it can answer the specified filter.
type ListNotesFilterFn func(ctx context.Context, projID, filter, pageToken string, pageSize int32) ([]*gpb.Note, string, bool, error)

// ListNotesFilterer holds functions on how to handle various filter patterns for listing notes.
type ListNotesFilterer struct {
	// FilterFns contain all functions that handle specific filter patterns.
	FilterFns []ListNotesFilterFn
	// DefaultFilterFn contains the fallback function to handle a filter if none of the filter
	// functions understand how to handle a filter.
	DefaultFilterFn ListNotesFilterFn
}

// Filter finds the appropriate filter function to handle the specified filter and executes it to
// return filtered notes.
func (f *ListNotesFilterer) Filter(ctx context.Context, projID, filter, pageToken string, pageSize int32) ([]*gpb.Note, string, error) {
	for _, filterFn := range f.FilterFns {
		notes, npt, ok, err := filterFn(ctx, projID, filter, pageToken, pageSize)
		if !ok {
			logger.Infof("Cannot handle filter %q", filter)
			continue
		}
		if err != nil {
			return nil, "", err
		}
		return notes, npt, nil
	}

	notes, npt, _, err := f.DefaultFilterFn(ctx, projID, filter, pageToken, pageSize)
	return notes, npt, err
}

// ListOccsFilterFn is a function that efficiently handles a specific filter pattern. It parses the
// specified filter to determine if it can handle it, and if it can, it returns only occurrences
// that match the filter, and the next page token (if available). The function always returns a bool
// indicating whether it can answer the specified filter.
type ListOccsFilterFn func(ctx context.Context, projID, filter, pageToken string, pageSize int32) ([]*gpb.Occurrence, string, bool, error)

// ListOccsFilterer holds functions on how to handle various filter patterns for listing
// occurrences.
type ListOccsFilterer struct {
	// FilterFns contain all functions that handle specific filter patterns.
	FilterFns []ListOccsFilterFn
	// DefaultFilterFn contains the fallback function to handle a filter if none of the filter
	// functions understand how to handle a filter.
	DefaultFilterFn ListOccsFilterFn
}

// Filter finds the appropriate filter function to handle the specified filter and executes it to
// return filtered occurrences.
func (f *ListOccsFilterer) Filter(ctx context.Context, projID, filter, pageToken string, pageSize int32) ([]*gpb.Occurrence, string, error) {
	for _, filterFn := range f.FilterFns {
		occs, npt, ok, err := filterFn(ctx, projID, filter, pageToken, pageSize)
		if !ok {
			logger.Infof("Cannot handle filter %q", filter)
			continue
		}
		if err != nil {
			return nil, "", err
		}
		return occs, npt, nil
	}

	occs, npt, _, err := f.DefaultFilterFn(ctx, projID, filter, pageToken, pageSize)
	return occs, npt, err
}
