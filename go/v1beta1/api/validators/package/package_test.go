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

package pkg

import (
	"testing"

	ppb "github.com/grafeas/grafeas/proto/v1beta1/package_go_proto"
)

func TestValidatePackage(t *testing.T) {
	tests := []struct {
		desc     string
		p        *ppb.Package
		wantErrs bool
	}{
		{
			desc:     "missing name, want error(s)",
			p:        &ppb.Package{},
			wantErrs: true,
		},
		{
			desc: "nil distribution, want error(s)",
			p: &ppb.Package{
				Name: "debian",
				Distribution: []*ppb.Distribution{
					nil,
				},
			},
			wantErrs: true,
		},
		{
			desc: "invalid distribution, want error(s)",
			p: &ppb.Package{
				Name: "debian",
				Distribution: []*ppb.Distribution{
					{},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid package, want success",
			p: &ppb.Package{
				Name: "debian",
				Distribution: []*ppb.Distribution{
					{
						CpeUri: "cpe:/o:debian:debian_linux:7",
					},
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidatePackage(tt.p)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidatePackage(%+v): got success, want error(s)", tt.desc, tt.p)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidatePackage(%+v): got error(s) %v, want success", tt.desc, tt.p, errs)
		}
	}
}

func TestValidateDistribution(t *testing.T) {
	tests := []struct {
		desc     string
		d        *ppb.Distribution
		wantErrs bool
	}{
		{
			desc:     "missing CPE URI, want error(s)",
			d:        &ppb.Distribution{},
			wantErrs: true,
		},
		{
			desc: "invalid latest version, want error(s)",
			d: &ppb.Distribution{
				CpeUri:        "cpe:/o:debian:debian_linux:7",
				LatestVersion: &ppb.Version{},
			},
			wantErrs: true,
		},
		{
			desc: "valid distribution, want success",
			d: &ppb.Distribution{
				CpeUri: "cpe:/o:debian:debian_linux:7",
				LatestVersion: &ppb.Version{
					Name: "1.1.2",
					Kind: ppb.Version_NORMAL,
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateDistribution(tt.d)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateDistribution(%+v): got success, want error(s)", tt.desc, tt.d)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateDistribution(%+v): got error(s) %v, want success", tt.desc, tt.d, errs)
		}
	}
}

func TestValidateVersion(t *testing.T) {
	tests := []struct {
		desc     string
		v        *ppb.Version
		wantErrs bool
	}{
		{
			desc: "missing name, want error(s)",
			v: &ppb.Version{
				Kind: ppb.Version_NORMAL,
			},
			wantErrs: true,
		},
		{
			desc: "missing kind, want error(s)",
			v: &ppb.Version{
				Name: "debian",
			},
			wantErrs: true,
		},
		{
			desc: "valid version, want success",
			v: &ppb.Version{
				Name: "1.1.2",
				Kind: ppb.Version_NORMAL,
			},
			wantErrs: false,
		},
		{
			desc: "valid maximum version, want success",
			v: &ppb.Version{
				Kind: ppb.Version_MAXIMUM,
			},
			wantErrs: false,
		},
		{
			desc: "valid minimum version, want success",
			v: &ppb.Version{
				Kind: ppb.Version_MINIMUM,
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateVersion(tt.v)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateVersion(%+v): got success, want error(s)", tt.desc, tt.v)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateVersion(%+v): got error(s) %v, want success", tt.desc, tt.v, errs)
		}
	}
}

func TestValidateDetails(t *testing.T) {
	tests := []struct {
		desc     string
		d        *ppb.Details
		wantErrs bool
	}{
		{
			desc:     "missing installation, want error(s)",
			d:        &ppb.Details{},
			wantErrs: true,
		},
		{
			desc: "invalid installation, want error(s)",
			d: &ppb.Details{
				Installation: &ppb.Installation{},
			},
			wantErrs: true,
		},
		{
			desc: "valid details, want success",
			d: &ppb.Details{
				Installation: &ppb.Installation{
					Location: []*ppb.Location{
						{
							CpeUri: "cpe:/o:debian:debian_linux:7",
						},
					},
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

func TestValidateInstallation(t *testing.T) {
	tests := []struct {
		desc     string
		i        *ppb.Installation
		wantErrs bool
	}{
		{
			desc:     "missing location, want error(s)",
			i:        &ppb.Installation{},
			wantErrs: true,
		},
		{
			desc: "empty location, want error(s)",
			i: &ppb.Installation{
				Location: []*ppb.Location{},
			},
			wantErrs: true,
		},
		{
			desc: "nil location, want error(s)",
			i: &ppb.Installation{
				Location: []*ppb.Location{nil},
			},
			wantErrs: true,
		},
		{
			desc: "invalid location, want error(s)",
			i: &ppb.Installation{
				Location: []*ppb.Location{
					{},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid installation, want success",
			i: &ppb.Installation{
				Location: []*ppb.Location{
					{
						CpeUri: "cpe:/o:debian:debian_linux:7",
					},
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateInstallation(tt.i)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateInstallation(%+v): got success, want error(s)", tt.desc, tt.i)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateInstallation(%+v): got error(s) %v, want success", tt.desc, tt.i, errs)
		}
	}
}

func TestValidateLocation(t *testing.T) {
	tests := []struct {
		desc     string
		l        *ppb.Location
		wantErrs bool
	}{
		{
			desc:     "missing CPE URI, want error(s)",
			l:        &ppb.Location{},
			wantErrs: true,
		},
		{
			desc: "invalid version, want error(s)",
			l: &ppb.Location{
				CpeUri:  "cpe:/o:debian:debian_linux:7",
				Version: &ppb.Version{},
			},
			wantErrs: true,
		},
		{
			desc: "valid location, want success",
			l: &ppb.Location{
				CpeUri: "cpe:/o:debian:debian_linux:7",
				Version: &ppb.Version{
					Name: "1.1.2",
					Kind: ppb.Version_NORMAL,
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateLocation(tt.l)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateLocation(%+v): got success, want error(s)", tt.desc, tt.l)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateInstallation(%+v): got error(s) %v, want success", tt.desc, tt.l, errs)
		}
	}
}
