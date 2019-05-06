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

package grafeas

import (
	"testing"

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

func TestValidateNote(t *testing.T) {
	tests := []struct {
		desc    string
		n       *gpb.Note
		wantErr bool
	}{
		{
			desc:    "missing type, want error",
			n:       &gpb.Note{},
			wantErr: true,
		},
		{
			desc: "invalid vulnerability, want error",
			n: &gpb.Note{
				Type: &gpb.Note_Vulnerability{
					Vulnerability: &gpb.VulnerabilityNote{
						Details: []*gpb.VulnerabilityNote_Detail{nil},
					},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid build, want error",
			n: &gpb.Note{
				Type: &gpb.Note_Build{
					Build: &gpb.BuildNote{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid base image, want error",
			n: &gpb.Note{
				Type: &gpb.Note_BaseImage{
					BaseImage: &gpb.ImageNote{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid package, want error",
			n: &gpb.Note{
				Type: &gpb.Note_Package{
					Package: &gpb.PackageNote{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid deployable, want error",
			n: &gpb.Note{
				Type: &gpb.Note_Deployable{
					Deployable: &gpb.DeploymentNote{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid discovery, want error",
			n: &gpb.Note{
				Type: &gpb.Note_Discovery{
					Discovery: &gpb.DiscoveryNote{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid attestation authority, want error",
			n: &gpb.Note{
				Type: &gpb.Note_AttestationAuthority{
					AttestationAuthority: &gpb.AttestationNote{
						Hint: &gpb.AttestationNote_Hint{},
					},
				},
			},
			wantErr: true,
		},
		{
			desc: "valid note, want success",
			n: &gpb.Note{
				Type: &gpb.Note_Vulnerability{
					Vulnerability: &gpb.VulnerabilityNote{
						Severity: gpb.Severity_CRITICAL,
						Details: []*gpb.VulnerabilityNote_Detail{
							&gpb.VulnerabilityNote_Detail{
								CpeUri:       "cpe:/o:debian:debian_linux:7",
								Package:      "debian",
								SeverityName: "LOW",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		err := ValidateNote(tt.n)
		t.Logf("%q: error: %v", tt.desc, err)
		if err == nil && tt.wantErr {
			t.Errorf("%q: ValidateNote(%+v): got success, want error", tt.desc, tt.n)
		}
		if err != nil && !tt.wantErr {
			t.Errorf("%q: ValidateNote(%+v): got error %v, want success", tt.desc, tt.n, err)
		}
	}
}

func TestValidateOccurrence(t *testing.T) {
	tests := []struct {
		desc    string
		o       *gpb.Occurrence
		wantErr bool
	}{
		{
			desc:    "missing resource, want error",
			o:       &gpb.Occurrence{},
			wantErr: true,
		},
		{
			desc: "invalid resource, want error",
			o: &gpb.Occurrence{
				Resource: &gpb.Resource{},
			},
			wantErr: true,
		},
		{
			desc: "missing note name, want error",
			o: &gpb.Occurrence{
				Resource: &gpb.Resource{
					Uri: "goog://foo/bar",
				},
			},
			wantErr: true,
		},
		{
			desc: "missing details, want error",
			o: &gpb.Occurrence{
				Resource: &gpb.Resource{
					Uri: "goog://foo/bar",
				},
				NoteName: "projects/goog-vulnz/notes/CVE-UH-OH",
			},
			wantErr: true,
		},
		{
			desc: "invalid vulnerability, want error",
			o: &gpb.Occurrence{
				Details: &gpb.Occurrence_Vulnerability{
					Vulnerability: &gpb.VulnerabilityOccurrence{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid build, want error",
			o: &gpb.Occurrence{
				Details: &gpb.Occurrence_Build{
					Build: &gpb.BuildOccurrence{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid derived image, want error",
			o: &gpb.Occurrence{
				Details: &gpb.Occurrence_DerivedImage{
					DerivedImage: &gpb.ImageOccurrence{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid installation, want error",
			o: &gpb.Occurrence{
				Details: &gpb.Occurrence_Installation{
					Installation: &gpb.PackageOccurrence{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid deployment, want error",
			o: &gpb.Occurrence{
				Details: &gpb.Occurrence_Deployment{
					Deployment: &gpb.DeploymentOccurrence{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid discovered, want error",
			o: &gpb.Occurrence{
				Details: &gpb.Occurrence_Discovered{
					Discovered: &gpb.DiscoveryOccurrence{},
				},
			},
			wantErr: true,
		},
		{
			desc: "invalid attestation, want error",
			o: &gpb.Occurrence{
				Details: &gpb.Occurrence_Attestation{
					Attestation: &gpb.AttestationOccurrence{},
				},
			},
			wantErr: true,
		},
		{
			desc: "valid occurrence, want success",
			o: &gpb.Occurrence{
				Resource: &gpb.Resource{
					Uri: "goog://foo/bar",
				},
				NoteName: "projects/goog-vulnz/notes/CVE-UH-OH",
				Details:  &gpb.Occurrence_Vulnerability{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		err := ValidateOccurrence(tt.o)
		t.Logf("%q: error: %v", tt.desc, err)
		if err == nil && tt.wantErr {
			t.Errorf("%q: ValidateOccurrence(%+v): got success, want error", tt.desc, tt.o)
		}
		if err != nil && !tt.wantErr {
			t.Errorf("%q: ValidateOccurrence(%+v): got error %v, want success", tt.desc, tt.o, err)
		}
	}
}

func TestValidateResource(t *testing.T) {
	tests := []struct {
		desc     string
		r        *gpb.Resource
		wantErrs bool
	}{
		{
			desc:     "missing URI, want error(s)",
			r:        &gpb.Resource{},
			wantErrs: true,
		},
		{
			desc: "valid resource, want success",
			r: &gpb.Resource{
				Uri: "goog://foo/bar",
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateResource(tt.r)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateResource(%+v): got success, want error(s)", tt.desc, tt.r)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateResource(%+v): got error(s) %v, want success", tt.desc, tt.r, errs)
		}
	}
}
