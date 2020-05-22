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
	dpb "github.com/grafeas/grafeas/proto/v1beta1/deployment_go_proto"
)

func TestValidateDeployable(t *testing.T) {
	tests := []struct {
		desc     string
		d        *dpb.Deployable
		wantErrs bool
	}{
		{
			desc:     "missing resource URI, want error(s)",
			d:        &dpb.Deployable{},
			wantErrs: true,
		},
		{
			desc: "empty resource URI, want error(s)",
			d: &dpb.Deployable{
				ResourceUri: []string{},
			},
			wantErrs: true,
		},
		{
			desc: "invalid resource URI, want error(s)",
			d: &dpb.Deployable{
				ResourceUri: []string{""},
			},
			wantErrs: true,
		},
		{
			desc: "valid deployable, want success",
			d: &dpb.Deployable{
				ResourceUri: []string{"https://gcr.io/foo/bar"},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateDeployable(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateDeployable(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateDeployable(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
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
			desc:     "missing deployment, want error(s)",
			d:        &dpb.Details{},
			wantErrs: true,
		},
		{
			desc: "invalid deployment, want error(s)",
			d: &dpb.Details{
				Deployment: &dpb.Deployment{},
			},
			wantErrs: true,
		},
		{
			desc: "valid details, want success",
			d: &dpb.Details{
				Deployment: &dpb.Deployment{
					DeployTime: &tpb.Timestamp{},
				},
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

func TestValidateDeployment(t *testing.T) {
	tests := []struct {
		desc     string
		d        *dpb.Deployment
		wantErrs bool
	}{
		{
			desc:     "missing deploy time, want error(s)",
			d:        &dpb.Deployment{},
			wantErrs: true,
		},
		{
			desc: "valid deployment, want success",
			d: &dpb.Deployment{
				DeployTime: &tpb.Timestamp{},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateDeployment(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateDeployment(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateDeployment(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}
