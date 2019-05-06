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

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

func TestValidatePackage(t *testing.T) {
	tests := []struct {
		desc     string
		p        *gpb.PackageNote
		wantErrs bool
	}{
		{
			desc:     "missing name, want error(s)",
			p:        &gpb.PackageNote{},
			wantErrs: true,
		},
		{
			desc: "nil distribution, want error(s)",
			p: &gpb.PackageNote{
				Name: "debian",
				Distribution: []*gpb.Distribution{
					nil,
				},
			},
			wantErrs: true,
		},
		{
			desc: "invalid distribution, want error(s)",
			p: &gpb.PackageNote{
				Name: "debian",
				Distribution: []*gpb.Distribution{
					{},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid package, want success",
			p: &gpb.PackageNote{
				Name: "debian",
				Distribution: []*gpb.Distribution{
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
		d        *gpb.Distribution
		wantErrs bool
	}{
		{
			desc:     "missing CPE URI, want error(s)",
			d:        &gpb.Distribution{},
			wantErrs: true,
		},
		{
			desc: "invalid latest version, want error(s)",
			d: &gpb.Distribution{
				CpeUri:        "cpe:/o:debian:debian_linux:7",
				LatestVersion: &gpb.Version{},
			},
			wantErrs: true,
		},
		{
			desc: "valid distribution, want success",
			d: &gpb.Distribution{
				CpeUri: "cpe:/o:debian:debian_linux:7",
				LatestVersion: &gpb.Version{
					Name: "1.1.2",
					Kind: gpb.Version_NORMAL,
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
		v        *gpb.Version
		wantErrs bool
	}{
		{
			desc: "missing name, want error(s)",
			v: &gpb.Version{
				Kind: gpb.Version_NORMAL,
			},
			wantErrs: true,
		},
		{
			desc: "missing kind, want error(s)",
			v: &gpb.Version{
				Name: "debian",
			},
			wantErrs: true,
		},
		{
			desc: "valid version, want success",
			v: &gpb.Version{
				Name: "1.1.2",
				Kind: gpb.Version_NORMAL,
			},
			wantErrs: false,
		},
		{
			desc: "valid maximum version, want success",
			v: &gpb.Version{
				Kind: gpb.Version_MAXIMUM,
			},
			wantErrs: false,
		},
		{
			desc: "valid minimum version, want success",
			v: &gpb.Version{
				Kind: gpb.Version_MINIMUM,
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
		d        *gpb.PackageOccurrence
		wantErrs bool
	}{
		{
			desc:     "missing installation, want error(s)",
			d:        &gpb.PackageOccurrence{},
			wantErrs: true,
		},
		{
			desc: "invalid installation, want error(s)",
			d: &gpb.PackageOccurrence{
				Installation: &gpb.Installation{},
			},
			wantErrs: true,
		},
		{
			desc: "valid details, want success",
			d: &gpb.PackageOccurrence{
				Installation: &gpb.Installation{
					Location: []*gpb.Location{
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
		i        *gpb.Installation
		wantErrs bool
	}{
		{
			desc:     "missing location, want error(s)",
			i:        &gpb.Installation{},
			wantErrs: true,
		},
		{
			desc: "empty location, want error(s)",
			i: &gpb.Installation{
				Location: []*gpb.Location{},
			},
			wantErrs: true,
		},
		{
			desc: "nil location, want error(s)",
			i: &gpb.Installation{
				Location: []*gpb.Location{nil},
			},
			wantErrs: true,
		},
		{
			desc: "invalid location, want error(s)",
			i: &gpb.Installation{
				Location: []*gpb.Location{
					{},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid installation, want success",
			i: &gpb.Installation{
				Location: []*gpb.Location{
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
		l        *gpb.Location
		wantErrs bool
	}{
		{
			desc:     "missing CPE URI, want error(s)",
			l:        &gpb.Location{},
			wantErrs: true,
		},
		{
			desc: "invalid version, want error(s)",
			l: &gpb.Location{
				CpeUri:  "cpe:/o:debian:debian_linux:7",
				Version: &gpb.Version{},
			},
			wantErrs: true,
		},
		{
			desc: "valid location, want success",
			l: &gpb.Location{
				CpeUri: "cpe:/o:debian:debian_linux:7",
				Version: &gpb.Version{
					Name: "1.1.2",
					Kind: gpb.Version_NORMAL,
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
