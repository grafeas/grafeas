// Package discovery implements functions to validate that the fields of discovery entities being
// passed into the API meet our requirements.
package discovery

import (
	"errors"
	"fmt"

	cpb "github.com/grafeas/grafeas/proto/v1beta1/common_go_proto"
	dpb "github.com/grafeas/grafeas/proto/v1beta1/discovery_go_proto"
)

// ValidateDiscovery validates that a discovery has all its required fields filled in.
func ValidateDiscovery(d *dpb.Discovery) []error {
	errs := []error{}

	if d.GetAnalysisKind() == cpb.NoteKind_NOTE_KIND_UNSPECIFIED {
		errs = append(errs, errors.New("analysis_kind is required"))
	}

	return errs
}

// ValidateDetails validates that a details has all its required fields filled in.
func ValidateDetails(d *dpb.Details) []error {
	errs := []error{}

	if di := d.GetDiscovered(); di == nil {
		errs = append(errs, errors.New("discovered is required"))
	} else {
		for _, err := range validateDiscovered(di) {
			errs = append(errs, fmt.Errorf("discovered.%s", err))
		}
	}

	return errs
}

func validateDiscovered(d *dpb.Discovered) []error {
	errs := []error{}
	return errs
}
