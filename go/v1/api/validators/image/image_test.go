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

package image

import (
	"testing"

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

func TestValidateNote(t *testing.T) {
	tests := []struct {
		desc     string
		b        *gpb.ImageNote
		wantErrs bool
	}{
		{
			desc:     "missing resource URL, want error(s)",
			b:        &gpb.ImageNote{},
			wantErrs: true,
		},
		{
			desc: "nil fingerprint, want error(s)",
			b: &gpb.ImageNote{
				ResourceUrl: "https://www.google.com",
			},
			wantErrs: true,
		},
		{
			desc: "invalid fingerprint, want error(s)",
			b: &gpb.ImageNote{
				ResourceUrl: "https://www.google.com",
				Fingerprint: &gpb.Fingerprint{},
			},
			wantErrs: true,
		},
		{
			desc: "valid fingerprint, want success",
			b: &gpb.ImageNote{
				ResourceUrl: "https://www.google.com",
				Fingerprint: &gpb.Fingerprint{
					V1Name: "foo",
					V2Blob: []string{"bar"},
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateNote(tt.b)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateNote(%+v): got success, want error(s)", tt.desc, tt.b)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateNote(%+v): got error(s) %v, want success", tt.desc, tt.b, errs)
		}
	}
}

func TestValidateFingerprint(t *testing.T) {
	tests := []struct {
		desc     string
		f        *gpb.Fingerprint
		wantErrs bool
	}{
		{
			desc:     "missing V1 name, want error(s)",
			f:        &gpb.Fingerprint{},
			wantErrs: true,
		},
		{
			desc: "missing V2 blob, want error(s)",
			f: &gpb.Fingerprint{
				V1Name: "foo",
			},
			wantErrs: true,
		},
		{
			desc: "empty V2 blob, want error(s)",
			f: &gpb.Fingerprint{
				V1Name: "foo",
				V2Blob: []string{},
			},
			wantErrs: true,
		},
		{
			desc: "invalid V2 blob, want error(s)",
			f: &gpb.Fingerprint{
				V1Name: "foo",
				V2Blob: []string{""},
			},
			wantErrs: true,
		},
		{
			desc: "valid fingerprint, want success",
			f: &gpb.Fingerprint{
				V1Name: "foo",
				V2Blob: []string{"bar"},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateFingerprint(tt.f)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateFingerprint(%+v): got success, want error(s)", tt.desc, tt.f)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateFingerprint(%+v): got error(s) %v, want success", tt.desc, tt.f, errs)
		}
	}
}

func TestValidateOccurrence(t *testing.T) {
	tests := []struct {
		desc     string
		d        *gpb.ImageOccurrence
		wantErrs bool
	}{
		{
			desc:     "missing derived image, want error(s)",
			d:        &gpb.ImageOccurrence{},
			wantErrs: true,
		},
		{
			desc: "invalid derived image, want error(s)",
			d: &gpb.ImageOccurrence{
				DerivedImage: &gpb.Derived{},
			},
			wantErrs: true,
		},
		{
			desc: "valid derived image, want success",
			d: &gpb.ImageOccurrence{
				DerivedImage: &gpb.Derived{
					Fingerprint: &gpb.Fingerprint{
						V1Name: "foo",
						V2Blob: []string{"bar"},
					},
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateOccurrence(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateOccurrence(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateOccurrence(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}

func TestValidateDerived(t *testing.T) {
	tests := []struct {
		desc     string
		d        *gpb.Derived
		wantErrs bool
	}{
		{
			desc:     "missing fingerprint, want error(s)",
			d:        &gpb.Derived{},
			wantErrs: true,
		},
		{
			desc: "invalid fingerprint, want error(s)",
			d: &gpb.Derived{
				Fingerprint: &gpb.Fingerprint{},
			},
			wantErrs: true,
		},
		{
			desc: "nil layer, want error(s)",
			d: &gpb.Derived{
				Fingerprint: &gpb.Fingerprint{
					V1Name: "foo",
					V2Blob: []string{"bar"},
				},
				LayerInfo: []*gpb.Layer{nil},
			},
			wantErrs: true,
		},
		{
			desc: "invalid layer, want error(s)",
			d: &gpb.Derived{
				Fingerprint: &gpb.Fingerprint{
					V1Name: "foo",
					V2Blob: []string{"bar"},
				},
				LayerInfo: []*gpb.Layer{
					{},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid derived, want success",
			d: &gpb.Derived{
				Fingerprint: &gpb.Fingerprint{
					V1Name: "foo",
					V2Blob: []string{"bar"},
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateDerived(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateDerived(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateDerived(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}

func TestValidateLayer(t *testing.T) {
	tests := []struct {
		desc     string
		l        *gpb.Layer
		wantErrs bool
	}{
		{
			desc:     "missing directive, want error(s)",
			l:        &gpb.Layer{},
			wantErrs: true,
		},
		{
			desc: "valid layer, want success",
			l: &gpb.Layer{
				Directive: gpb.Layer_ADD,
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateLayer(tt.l)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateLayer(%+v): got success, want error(s)", tt.desc, tt.l)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateLayer(%+v): got error(s) %v, want success", tt.desc, tt.l, errs)
		}
	}
}
