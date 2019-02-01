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

package discovery

import (
	"testing"

	cpb "github.com/grafeas/grafeas/proto/v1/common_go_proto"
	dpb "github.com/grafeas/grafeas/proto/v1/discovery_go_proto"
)

func TestValidateDiscovery(t *testing.T) {
	tests := []struct {
		desc     string
		d        *dpb.Discovery
		wantErrs bool
	}{
		{
			desc:     "missing analysis kind, want error(s)",
			d:        &dpb.Discovery{},
			wantErrs: true,
		},
		{
			desc: "valid discovery, want success",
			d: &dpb.Discovery{
				AnalysisKind: cpb.NoteKind_VULNERABILITY,
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateDiscovery(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateDiscovery(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateDiscovery(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}

func TestValidateDetails(t *testing.T) {
	tests := []struct {
		desc     string
		d        *dpb.Details
		wantErrs bool
	}{
		{
			desc:     "missing discovered, want error(s)",
			d:        &dpb.Details{},
			wantErrs: true,
		},
		{
			desc: "valid details, want success",
			d: &dpb.Details{
				Discovered: &dpb.Discovered{},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateDetails(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateDetails(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateDetails(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}

func TestValidateDiscovered(t *testing.T) {
	tests := []struct {
		desc     string
		d        *dpb.Discovered
		wantErrs bool
	}{
		{
			desc:     "valid discovered, want success",
			d:        &dpb.Discovered{},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateDiscovered(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateDiscovered(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateDiscovered(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}
