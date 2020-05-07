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

	"github.com/grafeas/grafeas/go/v1/api/validators/attestation"
	"github.com/grafeas/grafeas/go/v1/api/validators/build"
	"github.com/grafeas/grafeas/go/v1/api/validators/deployment"
	"github.com/grafeas/grafeas/go/v1/api/validators/discovery"
	"github.com/grafeas/grafeas/go/v1/api/validators/image"
	pkg "github.com/grafeas/grafeas/go/v1/api/validators/package"
	"github.com/grafeas/grafeas/go/v1/api/validators/vulnerability"
	vlib "github.com/grafeas/grafeas/go/validationlib"
	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ValidateNote validates that a note has all its required fields filled in.
func ValidateNote(n *gpb.Note) error {
	errs := []error{}

	if n.GetType() == nil {
		errs = append(errs, errors.New("type is required"))
	}

	if v := n.GetVulnerability(); v != nil {
		for _, err := range vulnerability.ValidateNote(v) {
			errs = append(errs, fmt.Errorf("vulnerability.%s", err))
		}
	}

	if b := n.GetBuild(); b != nil {
		for _, err := range build.ValidateNote(b) {
			errs = append(errs, fmt.Errorf("build.%s", err))
		}
	}

	if b := n.GetImage(); b != nil {
		for _, err := range image.ValidateNote(b) {
			errs = append(errs, fmt.Errorf("base_image.%s", err))
		}
	}

	if p := n.GetPackage(); p != nil {
		for _, err := range pkg.ValidateNote(p) {
			errs = append(errs, fmt.Errorf("package.%s", err))
		}
	}

	if d := n.GetDeployment(); d != nil {
		for _, err := range deployment.ValidateNote(d) {
			errs = append(errs, fmt.Errorf("deplyable.%s", err))
		}
	}

	if d := n.GetDiscovery(); d != nil {
		for _, err := range discovery.ValidateNote(d) {
			errs = append(errs, fmt.Errorf("discovery.%s", err))
		}
	}

	if a := n.GetAttestation(); a != nil {
		for _, err := range attestation.ValidateNote(a) {
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

	if l := len(o.GetResourceUri()); l > vlib.MaxResourceURILength {
		errs = append(errs, fmt.Errorf("resource_uri %s exceeds the limit %d", o.GetResourceUri(), vlib.MaxResourceURILength))
	}

	if o.GetResourceUri() == "" {
		errs = append(errs, errors.New("resource_uri is required"))
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
		for _, err := range vulnerability.ValidateOccurrence(v) {
			errs = append(errs, fmt.Errorf("vulnerability.%s", err))
		}
	}

	if v := o.GetBuild(); v != nil {
		for _, err := range build.ValidateOccurrence(v) {
			errs = append(errs, fmt.Errorf("build.%s", err))
		}
	}

	if i := o.GetImage(); i != nil {
		for _, err := range image.ValidateOccurrence(i) {
			errs = append(errs, fmt.Errorf("derived_image.%s", err))
		}
	}

	if i := o.GetPackage(); i != nil {
		for _, err := range pkg.ValidateOccurrence(i) {
			errs = append(errs, fmt.Errorf("installation.%s", err))
		}
	}

	if i := o.GetDeployment(); i != nil {
		for _, err := range deployment.ValidateOccurrence(i) {
			errs = append(errs, fmt.Errorf("deployment.%s", err))
		}
	}

	if i := o.GetDiscovery(); i != nil {
		for _, err := range discovery.ValidateOccurrence(i) {
			errs = append(errs, fmt.Errorf("discovered.%s", err))
		}
	}

	if i := o.GetAttestation(); i != nil {
		for _, err := range attestation.ValidateOccurrence(i) {
			errs = append(errs, fmt.Errorf("attestation.%s", err))
		}
	}

	if len(errs) > 0 {
		return status.Errorf(codes.InvalidArgument, "occurrence is invalid: %v", errs)
	}

	return nil
}
