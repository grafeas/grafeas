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

// Package pkg implements functions to validate that the fields of package entities being passed
// into the API meet our requirements.
package pkg

import (
	"errors"
	"fmt"

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

// ValidateNote validates that a package note all its required fields filled in.
func ValidateNote(n *gpb.PackageNote) []error {
	errs := []error{}

	if n.GetName() == "" {
		errs = append(errs, errors.New("name is required"))
	}

	for i, d := range n.GetDistribution() {
		if d == nil {
			errs = append(errs, fmt.Errorf("distribution[%d] distribution cannot be null", i))
		} else {
			for _, err := range validateDistribution(d) {
				errs = append(errs, fmt.Errorf("distribution[%d].%s", i, err))
			}
		}
	}

	return errs
}

func validateDistribution(d *gpb.Distribution) []error {
	errs := []error{}

	if d.GetCpeUri() == "" {
		errs = append(errs, errors.New("cpe_uri is required"))
	}

	if ver := d.GetLatestVersion(); ver != nil {
		for _, err := range ValidateVersion(ver) {
			errs = append(errs, fmt.Errorf("version.%s", err))
		}
	}

	return errs
}

// ValidateVersion validates that a version has all its required fields filled in.
func ValidateVersion(v *gpb.Version) []error {
	errs := []error{}

	// MAXIMUM and MINIMUM version kinds are valid without a Name
	if v.GetKind() == gpb.Version_NORMAL && v.GetName() == "" {
		errs = append(errs, errors.New("name is required"))
	}
	if v.GetKind() == gpb.Version_VERSION_KIND_UNSPECIFIED {
		errs = append(errs, errors.New("kind is required"))
	}

	return errs
}

// ValidateOccurrence validates that a package occurrence has all its required fields filled in.
func ValidateOccurrence(o *gpb.PackageOccurrence) []error {
	errs := []error{}

	if loc := o.GetLocation(); loc == nil {
		errs = append(errs, errors.New("location is required"))
	} else if len(loc) == 0 {
		errs = append(errs, errors.New("location requires at least 1 element"))
	} else {
		for i, l := range loc {
			if l == nil {
				errs = append(errs, fmt.Errorf("location[%d] location cannot be null", i))
			} else {
				for _, err := range validateLocation(l) {
					errs = append(errs, fmt.Errorf("location[%d].%s", i, err))
				}
			}
		}
	}

	return errs
}

func validateLocation(l *gpb.Location) []error {
	errs := []error{}

	if l.GetCpeUri() == "" {
		errs = append(errs, errors.New("cpe_uri is required"))
	}

	if v := l.GetVersion(); v != nil {
		for _, err := range ValidateVersion(v) {
			errs = append(errs, fmt.Errorf("version.%s", err))
		}
	}

	return errs
}
