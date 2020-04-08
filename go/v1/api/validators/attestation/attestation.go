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
	"strings"

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

// ValidateNote validates that an authority has all its required fields filled in.
func ValidateNote(a *gpb.AttestationNote) []error {
	errs := []error{}

	if h := a.GetHint(); h != nil {
		for _, err := range validateHint(h) {
			errs = append(errs, fmt.Errorf("hint.%s", err))
		}
	}

	return errs
}

func validateHint(h *gpb.AttestationNote_Hint) []error {
	errs := []error{}

	if h.GetHumanReadableName() == "" {
		errs = append(errs, errors.New("human_readable_name is required"))
	}

	return errs
}

// ValidateOccurrence validates that a details has all its required fields filled in.
func ValidateOccurrence(a *gpb.AttestationOccurrence) []error {
	errs := []error{}

	if sp := a.GetSerializedPayload(); sp == nil {
		errs = append(errs, errors.New("serialized payload is required"))
	}

	if s := a.GetSignatures(); s != nil {
		for _, err := range validateSignatures(s) {
			errs = append(errs, fmt.Errorf("signatures.%s", err))
		}
	}

	return errs
}

func validateSignatures(signatures []*gpb.Signature) []error {
	errs := []error{}

	for _, s := range signatures {
		if s.GetPublicKeyId() == "" {
			errs = append(errs, errors.New("public key ID is required"))
		} else if !strings.ContainsRune(s.GetPublicKeyId(), ':') {
			errs = append(errs, errors.New("public key ID must be an RFC3986 conformant URI"))
		}
		if s.GetSignature() == nil {
			errs = append(errs, errors.New("signature is required"))
		}
	}

	return errs
}
