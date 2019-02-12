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

// Package image implements functions to validate that the fields of image entities being passed
// into the API meet our requirements.
package image

import (
	"errors"
	"fmt"

	ipb "github.com/grafeas/grafeas/proto/v1/image_go_proto"
)

// ValidateBasis validates that an image basis has all its required fields filled in.
func ValidateBasis(b *ipb.Basis) []error {
	errs := []error{}

	if b.GetResourceUrl() == "" {
		errs = append(errs, errors.New("resource_url is required"))
	}

	if f := b.GetFingerprint(); f == nil {
		errs = append(errs, errors.New("fingerprint is required"))
	} else {
		for _, err := range validateFingerprint(f) {
			errs = append(errs, fmt.Errorf("fingerprint.%s", err))
		}
	}

	return errs
}

func validateFingerprint(f *ipb.Fingerprint) []error {
	errs := []error{}

	if f.GetV1Name() == "" {
		errs = append(errs, errors.New("v1_name is required"))
	}

	if blob := f.GetV2Blob(); blob == nil {
		errs = append(errs, errors.New("v2_blob is required"))
	} else if len(blob) == 0 {
		errs = append(errs, errors.New("v2_blob requires at least 1 element"))
	} else {
		for i, b := range blob {
			if b == "" {
				errs = append(errs, fmt.Errorf("v2_blob[%d] cannot be empty", i))
			}
		}
	}

	return errs
}

// ValidateDetails validates that a details has all its required fields filled in.
func ValidateDetails(d *ipb.Details) []error {
	errs := []error{}

	if d := d.GetDerivedImage(); d == nil {
		errs = append(errs, errors.New("derived_image is required"))
	} else {
		for _, err := range validateDerived(d) {
			errs = append(errs, fmt.Errorf("derived_image.%s", err))
		}
	}

	return errs
}

func validateDerived(d *ipb.Derived) []error {
	errs := []error{}

	if f := d.GetFingerprint(); f == nil {
		errs = append(errs, errors.New("fingerprint is required"))
	} else {
		for _, err := range validateFingerprint(f) {
			errs = append(errs, fmt.Errorf("fingerprint.%s", err))
		}
	}

	for i, l := range d.GetLayerInfo() {
		if l == nil {
			errs = append(errs, fmt.Errorf("layer_info[%d] layer cannot be null", i))
		} else {
			for _, err := range validateLayer(l) {
				errs = append(errs, fmt.Errorf("layer_info[%d].%s", i, err))
			}
		}
	}

	return errs
}

func validateLayer(l *ipb.Layer) []error {
	errs := []error{}

	if l.GetDirective() == ipb.Layer_DIRECTIVE_UNSPECIFIED {
		errs = append(errs, errors.New("directive is required"))
	}

	return errs
}
