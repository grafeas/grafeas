// Package build implements functions to validate that the fields of build entities being passed
// into the API meet our requirements.
package build

import (
	"errors"
	"fmt"

	"github.com/grafeas/grafeas/go/v1beta1/api/validators/provenance"
	bpb "github.com/grafeas/grafeas/proto/v1beta1/build_go_proto"
)

// ValidateBuild validates that a build has all its required fields filled in.
func ValidateBuild(b *bpb.Build) []error {
	errs := []error{}

	if b.GetBuilderVersion() == "" {
		errs = append(errs, errors.New("builder_version is required"))
	}

	if s := b.GetSignature(); s != nil {
		for _, err := range validateSignature(s) {
			errs = append(errs, fmt.Errorf("signature.%s", err))
		}
	}

	return errs
}

func validateSignature(s *bpb.BuildSignature) []error {
	errs := []error{}

	if s.GetSignature() == nil {
		errs = append(errs, errors.New("signature is required"))
	}

	return errs
}

// ValidateDetails validates that a details has all its required fields filled in.
func ValidateDetails(d *bpb.Details) []error {
	errs := []error{}

	if p := d.GetProvenance(); p == nil {
		errs = append(errs, errors.New("provenance is required"))
	} else {
		for _, err := range provenance.ValidateBuildProvenance(p) {
			errs = append(errs, fmt.Errorf("provenance.%s", err))
		}
	}

	return errs
}
