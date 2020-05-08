// Copyright 2018 The Grafeas Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package grafeas implements functions to validate that the fields of Grafeas entities being passed
// into the API meet our requirements.
package grafeas

import (
	"errors"
	"fmt"

	"github.com/grafeas/grafeas/go/v1beta1/api/validators/attestation"
	"github.com/grafeas/grafeas/go/v1beta1/api/validators/build"
	"github.com/grafeas/grafeas/go/v1beta1/api/validators/deployment"
	"github.com/grafeas/grafeas/go/v1beta1/api/validators/discovery"
	"github.com/grafeas/grafeas/go/v1beta1/api/validators/image"
	pkg "github.com/grafeas/grafeas/go/v1beta1/api/validators/package"
	"github.com/grafeas/grafeas/go/v1beta1/api/validators/vulnerability"
	vlib "github.com/grafeas/grafeas/go/validationlib"
	gpb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateNote validates that a note has all its required fields filled in.
func ValidateNote(n *gpb.Note) error {
	errs := []error{}

	if n.GetType() == nil {
		errs = append(errs, errors.New("type is required"))
	}

	if len(n.GetShortDescription()) > vlib.MaxShortDescriptionLength {
		errs = append(errs, fmt.Errorf("short_description %s exceeds the limit %d", n.GetShortDescription(), vlib.MaxShortDescriptionLength))
	}

	if len(n.GetLongDescription()) > vlib.MaxDescriptionLength {
		errs = append(errs, fmt.Errorf("long_description %s exceeds the limit %d", n.GetLongDescription(), vlib.MaxDescriptionLength))
	}

	if len(n.GetRelatedNoteNames()) > vlib.MaxCollectionSize {
		errs = append(errs, fmt.Errorf("size of related_note_names exceeds limit %d", vlib.MaxCollectionSize))
	}

	for _, noteName := range n.GetRelatedNoteNames() {
		if len(noteName) > vlib.MaxResourceURILength {
			errs = append(errs, fmt.Errorf("length of %s exceeds limit %d", noteName, vlib.MaxResourceURILength))
		}
	}

	if len(n.GetRelatedUrl()) > vlib.MaxCollectionSize {
		errs = append(errs, fmt.Errorf("size of related_url exceeds limit %d", vlib.MaxCollectionSize))
	}

	for i, relatedURL := range n.GetRelatedUrl() {
		if url := relatedURL.GetUrl(); len(url) > vlib.MaxResourceURILength {
			errs = append(errs, fmt.Errorf("length of related_url[%d].url=%s exceeds limit %d", i, url, vlib.MaxResourceURILength))
		}
		if label := relatedURL.GetLabel(); len(label) > vlib.MaxDescriptionLength {
			errs = append(errs, fmt.Errorf("length of related_url[%d].label=%s exceeds limit %d", i, label, vlib.MaxDescriptionLength))
		}
	}

	if v := n.GetVulnerability(); v != nil {
		for _, err := range vulnerability.ValidateVulnerability(v) {
			errs = append(errs, fmt.Errorf("vulnerability.%s", err))
		}
	}

	if b := n.GetBuild(); b != nil {
		for _, err := range build.ValidateBuild(b) {
			errs = append(errs, fmt.Errorf("build.%s", err))
		}
	}

	if b := n.GetBaseImage(); b != nil {
		for _, err := range image.ValidateBasis(b) {
			errs = append(errs, fmt.Errorf("base_image.%s", err))
		}
	}

	if p := n.GetPackage(); p != nil {
		for _, err := range pkg.ValidatePackage(p) {
			errs = append(errs, fmt.Errorf("package.%s", err))
		}
	}

	if d := n.GetDeployable(); d != nil {
		for _, err := range deployment.ValidateDeployable(d) {
			errs = append(errs, fmt.Errorf("deplyable.%s", err))
		}
	}

	if d := n.GetDiscovery(); d != nil {
		for _, err := range discovery.ValidateDiscovery(d) {
			errs = append(errs, fmt.Errorf("discovery.%s", err))
		}
	}

	if a := n.GetAttestationAuthority(); a != nil {
		for _, err := range attestation.ValidateAuthority(a) {
			errs = append(errs, fmt.Errorf("attestation_authority.%s", err))
		}
	}

	if len(errs) > 0 {
		return status.Errorf(codes.InvalidArgument, "note is invalid: %v", errs)
	}

	return nil
}

// ValidateOccurrence validates that an occurrence has all its required fields filled in.
func ValidateOccurrence(o *gpb.Occurrence) error {
	errs := []error{}

	if r := o.GetResource(); r == nil {
		errs = append(errs, errors.New("resource is required"))
	} else {
		for _, err := range validateResource(r) {
			errs = append(errs, fmt.Errorf("resource.%s", err))
		}
	}

	if o.GetNoteName() == "" {
		errs = append(errs, errors.New("note_name is required"))
	}

	if l := len(o.GetRemediation()); l > vlib.MaxDescriptionLength {
		errs = append(errs, fmt.Errorf("remediation %s exceeds the limit %d", o.GetRemediation(), vlib.MaxDescriptionLength))
	}

	if o.GetDetails() == nil {
		errs = append(errs, errors.New("details is required"))
	}

	if v := o.GetVulnerability(); v != nil {
		for _, err := range vulnerability.ValidateDetails(v) {
			errs = append(errs, fmt.Errorf("vulnerability.%s", err))
		}
	}

	if v := o.GetBuild(); v != nil {
		for _, err := range build.ValidateDetails(v) {
			errs = append(errs, fmt.Errorf("build.%s", err))
		}
	}

	if i := o.GetDerivedImage(); i != nil {
		for _, err := range image.ValidateDetails(i) {
			errs = append(errs, fmt.Errorf("derived_image.%s", err))
		}
	}

	if i := o.GetInstallation(); i != nil {
		for _, err := range pkg.ValidateDetails(i) {
			errs = append(errs, fmt.Errorf("installation.%s", err))
		}
	}

	if i := o.GetDeployment(); i != nil {
		for _, err := range deployment.ValidateDetails(i) {
			errs = append(errs, fmt.Errorf("deployment.%s", err))
		}
	}

	if i := o.GetDiscovered(); i != nil {
		for _, err := range discovery.ValidateDetails(i) {
			errs = append(errs, fmt.Errorf("discovered.%s", err))
		}
	}

	if i := o.GetAttestation(); i != nil {
		for _, err := range attestation.ValidateDetails(i) {
			errs = append(errs, fmt.Errorf("attestation.%s", err))
		}
	}

	if len(errs) > 0 {
		return status.Errorf(codes.InvalidArgument, "occurrence is invalid: %v", errs)
	}

	return nil
}

func validateResource(r *gpb.Resource) []error {
	errs := []error{}

	if r.GetUri() == "" {
		errs = append(errs, errors.New("uri is required"))
	}

	return errs
}
