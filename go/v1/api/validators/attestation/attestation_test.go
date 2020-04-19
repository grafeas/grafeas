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

	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
)

func TestValidateNote(t *testing.T) {
	tests := []struct {
		desc     string
		a        *gpb.AttestationNote
		wantErrs bool
	}{
		{
			desc: "invalid hint, want error(s)",
			a: &gpb.AttestationNote{
				Hint: &gpb.AttestationNote_Hint{},
			},
			wantErrs: true,
		},
		{
			desc: "valid authority, want success",
			a: &gpb.AttestationNote{
				Hint: &gpb.AttestationNote_Hint{
					HumanReadableName: "QA tests run",
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateNote(tt.a)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateNote(%+v): got success, want error(s)", tt.desc, tt.a)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateNote(%+v): got error(s) %v, want success", tt.desc, tt.a, errs)
		}
	}
}

func TestValidateHint(t *testing.T) {
	tests := []struct {
		desc     string
		h        *gpb.AttestationNote_Hint
		wantErrs bool
	}{
		{
			desc:     "invalid hint, want error(s)",
			h:        &gpb.AttestationNote_Hint{},
			wantErrs: true,
		},
		{
			desc: "valid hint, want success",
			h: &gpb.AttestationNote_Hint{
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

func TestValidateOccurrence(t *testing.T) {
	tests := []struct {
		desc     string
		a        *gpb.AttestationOccurrence
		wantErrs bool
	}{
		{
			desc:     "missing serialized payload, want error(s)",
			a:        &gpb.AttestationOccurrence{},
			wantErrs: true,
		},
		{
			desc: "missing public key ID in signature, want error(s)",
			a: &gpb.AttestationOccurrence{
				SerializedPayload: []byte("bar"),
				Signatures: []*gpb.Signature{
					{
						Signature: []byte("foo"),
					},
				},
			},
			wantErrs: true,
		},
		{
			desc: "missing signature, want error(s)",
			a: &gpb.AttestationOccurrence{
				SerializedPayload: []byte("bar"),
				Signatures: []*gpb.Signature{
					{
						PublicKeyId: "openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA",
					},
				},
			},
			wantErrs: true,
		},
		{
			desc: "one invalid signature in attestation, want error(s)",
			a: &gpb.AttestationOccurrence{
				SerializedPayload: []byte("bar"),
				Signatures: []*gpb.Signature{
					{
						Signature:   []byte("foo"),
						PublicKeyId: "openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA",
					},
					{
						Signature:   []byte("foo"),
						PublicKeyId: "ni:///sha-256;cD9o9Cq6LG3jD0iKXqEi_vdjJGecm_iXkbqVoScViaU",
					},
					{
						Signature: []byte("foo"),
					},
				},
			},
			wantErrs: true,
		},
		{
			desc: "invalid public key format, want errors",
			a: &gpb.AttestationOccurrence{
				SerializedPayload: []byte("bar"),
				Signatures: []*gpb.Signature{
					{
						Signature:   []byte("foo"),
						PublicKeyId: "74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA",
					},
				},
			},
			wantErrs: true,
		},
		{
			desc: "valid attestation, want success",
			a: &gpb.AttestationOccurrence{
				SerializedPayload: []byte("bar"),
				Signatures: []*gpb.Signature{
					{
						Signature:   []byte("foo"),
						PublicKeyId: "openpgp4fpr:74FAF3B861BDA0870C7B6DEF607E48D2A663AEEA",
					},
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateOccurrence(tt.a)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateOccurrence(%+v): got success, want error(s)", tt.desc, tt.a)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateOccurrence(%+v): got error(s) %v, want success", tt.desc, tt.a, errs)
		}
	}
}
