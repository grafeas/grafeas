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
	gpb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	// Copied from ../greafeas.go since that const was local to the module.
	// TODO(liupen): unify the two.
	defaultPageSize = 20
	maxPageSize     = 1000
	maxBatchSize    = 1000
)

// ValidateCreateNoteRequest validates that a CreateNoteRequest message has all valid fields.
// Note: the tests are in ../note_test.go, which create various messages and run CreateNote.
func ValidateCreateNoteRequest(req *gpb.CreateNoteRequest) error {
	errs := []error{}
	if req.NoteId == "" {
		errs = append(errs, status.Errorf(codes.InvalidArgument, "a noteId must be specified"))
	}
	if req.Note == nil {
		errs = append(errs, status.Errorf(codes.InvalidArgument, "a note must be specified"))
	}
	if err := ValidateNote(req.Note); err != nil {
		errs = append(errs, err)
	}
	if len(errs) > 0 {
		return status.Errorf(codes.InvalidArgument, "CreateNoteRequest message is invalid: %v", errs)
	}
	return nil
}

// ValidateBatchCreateNotesRequest validates that a BatchCreateNotesRequest message has all valid fields.
// Note: the tests are in ../note_test.go.
func ValidateBatchCreateNotesRequest(req *gpb.BatchCreateNotesRequest) error {
	errs := []error{}
	if len(req.Notes) == 0 {
		errs = append(errs, status.Errorf(codes.InvalidArgument, "at least one note must be specified"))
	}
	if len(req.Notes) > maxBatchSize {
		errs = append(errs, status.Errorf(codes.InvalidArgument, "%d is too many notes to batch create, a maximum of %d notes is allowed per batch create", len(req.Notes), maxBatchSize))
	}

	for i, n := range req.Notes {
		if err := ValidateNote(n); err != nil {
			errs = append(errs, fmt.Errorf("notes[%q]: %v", i, err))
		}
	}

	if len(errs) > 0 {
		return status.Errorf(codes.InvalidArgument, "BatchCreateNotesRequest message is invalid: %v", errs)
	}
	return nil
}

// ValidateUpdateNoteRequest validates that UpdateNoteRequest has all valid fields.
func ValidateUpdateNoteRequest(req *gpb.UpdateNoteRequest) error {
	if req.Note == nil {
		return status.Errorf(codes.InvalidArgument, "an note must be specified")
	}
	return nil
}

// ValidateGetNoteRequest validates that GetNoteRequest has all valid fields.
func ValidateGetNoteRequest(req *gpb.GetNoteRequest) error {
	return nil
}

// ValidateDeleteNoteRequest validates that DeleteNoteRequest has all valid fields.
func ValidateDeleteNoteRequest(req *gpb.DeleteNoteRequest) error {
	return nil
}

// ValidateListNotesRequest validates that ListNotesRequest has all valid fields.
func ValidateListNotesRequest(req *gpb.ListNotesRequest) error {
	errs := []error{}
	_, err := ValidatePageSize(req.PageSize)
	if err != nil {
		errs = append(errs, err)
	}
	// TODO(liupen): add validation of req.Filter here.
	if len(errs) > 0 {
		return status.Errorf(codes.InvalidArgument, "ListNotesRequest message is invalid: %v", errs)
	}
	return nil
}

// ValidateGetOccurrenceNoteRequest validates that GetOccurrenceNoteRequest has all valid fields.
func ValidateGetOccurrenceNoteRequest(req *gpb.GetOccurrenceNoteRequest) error {
	return nil
}

// ValidateNote validates that a note has all its required fields filled in.
func ValidateNote(n *gpb.Note) error {
	errs := []error{}

	if n.GetType() == nil {
		errs = append(errs, errors.New("type is required"))
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

// ValidateCreateOccurrenceRequest validates that CreateOccurrenceRequest has all valid fields.
func ValidateCreateOccurrenceRequest(req *gpb.CreateOccurrenceRequest) error {
	if err := ValidateOccurrence(req.Occurrence); err != nil {
		return err
	}

	return nil
}

// ValidateBatchCreateOccurrencesRequest validates that BatchCreateOccurrencesRequest has all valid fields.
func ValidateBatchCreateOccurrencesRequest(req *gpb.BatchCreateOccurrencesRequest) error {
	errs := []error{}
	if len(req.Occurrences) == 0 {
		errs = append(errs, status.Errorf(codes.InvalidArgument, "at least one occurrence must be specified"))
	}
	if len(req.Occurrences) > maxBatchSize {
		errs = append(errs, status.Errorf(codes.InvalidArgument, "%d is too many occurrence to batch create, a maximum of %d occurrence is allowed per batch create", len(req.Occurrences), maxBatchSize))
	}

	for i, o := range req.Occurrences {
		if err := ValidateOccurrence(o); err != nil {
			errs = append(errs, fmt.Errorf("occurrences[%d]: %v", i, err))
		}
	}
	if len(errs) > 0 {
		return status.Errorf(codes.InvalidArgument, "BatchCreateOccurrencesRequest message is invalid: %v", errs)
	}
	return nil
}

// ValidateUpdateOccurrenceRequest validates that UpdateOccurrenceRequest has all valid fields.
func ValidateUpdateOccurrenceRequest(req *gpb.UpdateOccurrenceRequest) error {
	if req.Occurrence == nil {
		return status.Errorf(codes.InvalidArgument, "an occurrence must be specified")
	}

	return nil
}

// ValidateDeleteOccurrenceRequest validates that DeleteOccurrenceRequest has all valid fields.
func ValidateDeleteOccurrenceRequest(req *gpb.DeleteOccurrenceRequest) error {
	return nil
}

// ValidateGetOccurrenceRequest validates that GetOccurrenceRequest has all valid fields.
func ValidateGetOccurrenceRequest(req *gpb.GetOccurrenceRequest) error {
	return nil
}

// ValidateListOccurrencesRequest validates ListOccurrencesRequest has all valid fields.
func ValidateListOccurrencesRequest(req *gpb.ListOccurrencesRequest) error {
	_, err := ValidatePageSize(req.PageSize)
	if err != nil {
		return err
	}
	return nil
}

// ValidateListNoteOccurrencesRequest validates ListNoteOccurrencesRequest has all valid fields.
func ValidateListNoteOccurrencesRequest(req *gpb.ListNoteOccurrencesRequest) error {
	return nil
}

// ValidateGetVulnerabilityOccurrencesSummaryRequest validates GetVulnerabilityOccurrencesSummaryRequest has all valid fields.
func ValidateGetVulnerabilityOccurrencesSummaryRequest(req *gpb.GetVulnerabilityOccurrencesSummaryRequest) error {
	return nil
}

// ValidateOccurrence validates that an Occurrence has all its required fields filled in.
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

// ValidatePageSize returns the default page size if the specified page size is 0, otherwise it
// validates the specified page size.
func ValidatePageSize(ps int32) (int32, error) {
	switch {
	case ps == 0:
		return defaultPageSize, nil
	case ps > maxPageSize:
		return 0, status.Errorf(codes.InvalidArgument, "page size %d cannot be large than max page size %d", ps, maxPageSize)
	case ps < 0:
		return 0, status.Errorf(codes.InvalidArgument, "page size %d cannot be negative", ps)
	}

	return ps, nil
}
