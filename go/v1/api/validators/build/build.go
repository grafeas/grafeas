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

// Package build implements functions to validate that the fields of build entities being passed
// into the API meet our requirements.
package build

import (
	"errors"
	"fmt"

	"github.com/grafeas/grafeas/go/v1/api/validators/provenance"
	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

// ValidateNote validates that a build has all its required fields filled in.
func ValidateNote(b *gpb.BuildNote) []error {
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

func validateSignature(s *gpb.BuildSignature) []error {
	errs := []error{}

	if s.GetSignature() == nil {
		errs = append(errs, errors.New("signature is required"))
	}

	return errs
}

// ValidateOccurrence validates that a details has all its required fields filled in.
func ValidateOccurrence(d *gpb.BuildOccurrence) []error {
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
