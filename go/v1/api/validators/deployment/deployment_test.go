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

package deployment

import (
	"testing"

	tpb "github.com/golang/protobuf/ptypes/timestamp"
	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

func TestValidateNote(t *testing.T) {
	tests := []struct {
		desc     string
		d        *gpb.DeploymentNote
		wantErrs bool
	}{
		{
			desc:     "missing resource URI, want error(s)",
			d:        &gpb.DeploymentNote{},
			wantErrs: true,
		},
		{
			desc: "empty resource URI, want error(s)",
			d: &gpb.DeploymentNote{
				ResourceUri: []string{},
			},
			wantErrs: true,
		},
		{
			desc: "invalid resource URI, want error(s)",
			d: &gpb.DeploymentNote{
				ResourceUri: []string{""},
			},
			wantErrs: true,
		},
		{
			desc: "valid deployable, want success",
			d: &gpb.DeploymentNote{
				ResourceUri: []string{"https://gcr.io/foo/bar"},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateNote(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateNote(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateNote(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}

func TestValidateOccurrence(t *testing.T) {
	tests := []struct {
		desc     string
		d        *gpb.DeploymentOccurrence
		wantErrs bool
	}{
		{
			desc:     "missing deploy time, want error(s)",
			d:        &gpb.DeploymentOccurrence{},
			wantErrs: true,
		},
		{
			desc: "valid deployment, want success",
			d: &gpb.DeploymentOccurrence{
				DeployTime: &tpb.Timestamp{},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateOccurrence(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateDetails(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateDetails(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}
