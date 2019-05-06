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

// Package discovery implements functions to validate that the fields of discovery entities being
// passed into the API meet our requirements.
package discovery

import (
	"errors"

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

// ValidateDiscovery validates that a discovery has all its required fields filled in.
func ValidateDiscovery(d *gpb.DiscoveryNote) []error {
	errs := []error{}

	if d.GetAnalysisKind() == gpb.NoteKind_NOTE_KIND_UNSPECIFIED {
		errs = append(errs, errors.New("analysis_kind is required"))
	}

	return errs
}

// ValidateDetails validates that a details has all its required fields filled in.
func ValidateDetails(d *gpb.DiscoveryOccurrence) []error {
	errs := []error{}
	return errs
}
