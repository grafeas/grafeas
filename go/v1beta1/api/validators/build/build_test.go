package build

import (
	"testing"

	bpb "github.com/grafeas/grafeas/proto/v1beta1/build_go_proto"
	ppb "github.com/grafeas/grafeas/proto/v1beta1/provenance_go_proto"
)

func TestValidateBuild(t *testing.T) {
	tests := []struct {
		desc     string
		b        *bpb.Build
		wantErrs bool
	}{
		{
			desc:     "missing builder version, want error(s)",
			b:        &bpb.Build{},
			wantErrs: true,
		},
		{
			desc: "invalid signature, want error(s)",
			b: &bpb.Build{
				BuilderVersion: "1.1.1",
				Signature:      &bpb.BuildSignature{},
			},
			wantErrs: true,
		},
		{
			desc: "valid signature, want success",
			b: &bpb.Build{
				BuilderVersion: "1.1.1",
				Signature: &bpb.BuildSignature{
					Signature: []byte("YmVhciByYXdyIHJhd3I="),
				},
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := ValidateBuild(tt.b)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: ValidateBuild(%+v): got success, want error(s)", tt.desc, tt.b)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: ValidateBuild(%+v): got error(s) %v, want success", tt.desc, tt.b, errs)
		}
	}
}

func TestValidateSignature(t *testing.T) {
	tests := []struct {
		desc     string
		s        *bpb.BuildSignature
		wantErrs bool
	}{
		{
			desc:     "missing signature, want error(s)",
			s:        &bpb.BuildSignature{},
			wantErrs: true,
		},
		{
			desc: "valid signature, want success",
			s: &bpb.BuildSignature{
				Signature: []byte("YmVhciByYXdyIHJhd3I="),
			},
			wantErrs: false,
		},
	}

	for _, tt := range tests {
		errs := validateSignature(tt.s)
		t.Logf("%q: error(s): %v", tt.desc, errs)
		if len(errs) == 0 && tt.wantErrs {
			t.Errorf("%q: validateSignature(%+v): got success, want error(s)", tt.desc, tt.s)
		}
		if len(errs) > 0 && !tt.wantErrs {
			t.Errorf("%q: validateSignature(%+v): got error(s) %v, want success", tt.desc, tt.s, errs)
		}
	}
}

func TestValidateDetails(t *testing.T) {
	tests := []struct {
		desc     string
		d        *bpb.Details
		wantErrs bool
	}{
		{
			desc:     "missing provenance, want error(s)",
			d:        &bpb.Details{},
			wantErrs: true,
		},
		{
			desc: "valid details, want success",
			d: &bpb.Details{
				Provenance: &ppb.BuildProvenance{
					Id: "8c0b1847-f78b-4bf7-8b2e-38e1bb48b125",
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
