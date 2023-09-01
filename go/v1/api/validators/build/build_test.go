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

package build

import (
	"testing"

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

func TestValidateNote(t *testing.T) {
	tests := []struct {
		desc     string
		b        *gpb.BuildNote
		wantErrs bool
	}{
		{
			desc:     "missing builder version, want error(s)",
			b:        &gpb.BuildNote{},
			wantErrs: true,
		},
		{
			desc: "valid build note, want success",
			b: &gpb.BuildNote{
				BuilderVersion: "1.1.1",
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

func TestValidateOccurrence(t *testing.T) {
	tests := []struct {
		desc     string
		d        *gpb.BuildOccurrence
		wantErrs bool
	}{
		{
			desc:     "missing provenance, want error(s)",
			d:        &gpb.BuildOccurrence{},
			wantErrs: true,
		},
		{
			desc: "valid details with provenance, want success",
			d: &gpb.BuildOccurrence{
				Provenance: &gpb.BuildProvenance{
					Id: "8c0b1847-f78b-4bf7-8b2e-38e1bb48b125",
				},
			},
			wantErrs: false,
		},
		{
			desc: "valid details with intotoprovenance, want success",
			d: &gpb.BuildOccurrence{
				IntotoProvenance: &gpb.InTotoProvenance{},
			},
			wantErrs: false,
		},
		{
			desc: "valid details with intotostatement, want success",
			d: &gpb.BuildOccurrence{
				IntotoStatement: &gpb.InTotoStatement{
					Type: "my_type",
				},
			},
			wantErrs: false,
		},
		{
			desc: "valid details with intoto slsa provenance v1, want success",
			d: &gpb.BuildOccurrence{
				InTotoSlsaProvenanceV1: &gpb.InTotoSlsaProvenanceV1{
					Type: "my_type",
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
