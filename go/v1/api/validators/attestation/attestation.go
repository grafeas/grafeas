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

// Package attestation implements functions to validate that the fields of attestation entities
// being passed into the API meet our requirements.
package attestation

import (
	"errors"
	"fmt"

	apb "github.com/grafeas/grafeas/proto/v1/attestation_go_proto"
)

// ValidateAuthority validates that an authority has all its required fields filled in.
func ValidateAuthority(a *apb.Authority) []error {
	errs := []error{}

	if h := a.GetHint(); h != nil {
		for _, err := range validateHint(h) {
			errs = append(errs, fmt.Errorf("hint.%s", err))
		}
	}

	return errs
}

func validateHint(h *apb.Authority_Hint) []error {
	errs := []error{}

	if h.GetHumanReadableName() == "" {
		errs = append(errs, errors.New("human_readable_name is required"))
	}

	return errs
}

// ValidateDetails validates that a details has all its required fields filled in.
func ValidateDetails(d *apb.Details) []error {
	errs := []error{}

	if a := d.GetAttestation(); a == nil {
		errs = append(errs, errors.New("attestation is required"))
	} else {
		for _, err := range validateAttestation(a) {
			errs = append(errs, fmt.Errorf("attestation.%s", err))
		}
	}

	return errs
}

func validateAttestation(a *apb.Attestation) []error {
	errs := []error{}

	if s := a.GetSignature(); s == nil {
		errs = append(errs, errors.New("signature is required"))
	}

	if p := a.GetPgpSignedAttestation(); p != nil {
		for _, err := range validatePgpSignedAttestation(p) {
			errs = append(errs, fmt.Errorf("pgp_signed_attestation.%s", err))
		}
	}

	return errs
}

func validatePgpSignedAttestation(p *apb.PgpSignedAttestation) []error {
	errs := []error{}

	if p.GetSignature() == "" {
		errs = append(errs, errors.New("signature is required"))
	}

	return errs
}
