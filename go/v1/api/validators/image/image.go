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

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

// ValidateNote validates that an image note has all its required fields filled in.
func ValidateNote(n *gpb.ImageNote) []error {
	errs := []error{}

	if n.GetResourceUrl() == "" {
		errs = append(errs, errors.New("resource_url is required"))
	}

	if f := n.GetFingerprint(); f == nil {
		errs = append(errs, errors.New("fingerprint is required"))
	} else {
		for _, err := range validateFingerprint(f) {
			errs = append(errs, fmt.Errorf("fingerprint.%s", err))
		}
	}

	return errs
}

func validateFingerprint(f *gpb.Fingerprint) []error {
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

// ValidateOccurrence validates that an image occurrence has all its required fields filled in.
func ValidateOccurrence(o *gpb.ImageOccurrence) []error {
	errs := []error{}

	if f := o.GetFingerprint(); f == nil {
		errs = append(errs, errors.New("fingerprint is required"))
	} else {
		for _, err := range validateFingerprint(f) {
			errs = append(errs, fmt.Errorf("fingerprint.%s", err))
		}
	}

	for i, l := range o.GetLayerInfo() {
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

func validateLayer(l *gpb.Layer) []error {
	errs := []error{}

	if l.GetDirective() == "" {
		errs = append(errs, errors.New("directive is required"))
	}

	return errs
}
