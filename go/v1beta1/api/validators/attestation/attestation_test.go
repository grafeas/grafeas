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

package attestation

import (
	"testing"

	apb "github.com/grafeas/grafeas/proto/v1beta1/attestation_go_proto"
)

func TestValidateAuthority(t *testing.T) {
	tests := []struct {
		desc     string
		a        *apb.Authority
		wantErrs bool
	}{
		{
			desc: "invalid hint, want error(s)",
			a: &apb.Authority{
				Hint: &apb.Authority_Hint{},
			},
			wantErrs: true,
		},
		{
			desc: "valid authority, want success",
			a: &apb.Authority{
				Hint: &apb.Authority_Hint{
					HumanReadableName: "QA tests run",
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateAuthority(tt.a)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateAuthority(%+v): got success, want error(s)", tt.desc, tt.a)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateAuthority(%+v): got error(s) %v, want success", tt.desc, tt.a, errs)
		}
	}
}

func TestValidateHint(t *testing.T) {
	tests := []struct {
		desc     string
		h        *apb.Authority_Hint
		wantErrs bool
	}{
		{
			desc:     "invalid hint, want error(s)",
			h:        &apb.Authority_Hint{},
			wantErrs: true,
		},
		{
			desc: "valid hint, want success",
			h: &apb.Authority_Hint{
				HumanReadableName: "QA tests run",
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateHint(tt.h)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateHint(%+v): got success, want error(s)", tt.desc, tt.h)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateHint(%+v): got error(s) %v, want success", tt.desc, tt.h, errs)
		}
	}
}

func TestValidateDetails(t *testing.T) {
	tests := []struct {
		desc     string
		d        *apb.Details
		wantErrs bool
	}{
		{
			desc:     "missing attestation, want error(s)",
			d:        &apb.Details{},
			wantErrs: true,
		},
		{
			desc: "invalid attestation, want error(s)",
			d: &apb.Details{
				Attestation: &apb.Attestation{},
			},
			wantErrs: true,
		},
		{
			desc: "valid details, want success",
			d: &apb.Details{
				Attestation: &apb.Attestation{
					Signature: &apb.Attestation_PgpSignedAttestation{
						PgpSignedAttestation: &apb.PgpSignedAttestation{
							Signature: "-----BEGIN PGP SIGNATURE-----\n\nxA0DARh\n-----END PGP SIGNATURE-----",
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

func TestValidateAttestation(t *testing.T) {
	tests := []struct {
		desc     string
		a        *apb.Attestation
		wantErrs bool
	}{
		{
			desc:     "missing signature, want error(s)",
			a:        &apb.Attestation{},
			wantErrs: true,
		},
		{
			desc: "invalid PGP signed attestation, want error(s)",
			a: &apb.Attestation{
				Signature: &apb.Attestation_PgpSignedAttestation{
					PgpSignedAttestation: &apb.PgpSignedAttestation{},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid attestation, want success",
			a: &apb.Attestation{
				Signature: &apb.Attestation_PgpSignedAttestation{
					PgpSignedAttestation: &apb.PgpSignedAttestation{
						Signature: "-----BEGIN PGP SIGNATURE-----\n\nxA0DARh\n-----END PGP SIGNATURE-----",
					},
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateAttestation(tt.a)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateAttestation(%+v): got success, want error(s)", tt.desc, tt.a)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateAttestation(%+v): got error(s) %v, want success", tt.desc, tt.a, errs)
		}
	}
}

func TestValidatePgpSignedAttestation(t *testing.T) {
	tests := []struct {
		desc     string
		p        *apb.PgpSignedAttestation
		wantErrs bool
	}{
		{
			desc:     "missing signature, want error(s)",
			p:        &apb.PgpSignedAttestation{},
			wantErrs: true,
		},
		{
			desc: "valid PGP signed attestation, want success",
			p: &apb.PgpSignedAttestation{
				Signature: "-----BEGIN PGP SIGNATURE-----\n\nxA0DARh\n-----END PGP SIGNATURE-----",
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validatePgpSignedAttestation(tt.p)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validatePgpSignedAttestation(%+v): got success, want error(s)", tt.desc, tt.p)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validatePgpSignedAttestation(%+v): got error(s) %v, want success", tt.desc, tt.p, errs)
		}
	}
}
