package intoto

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	pb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
	intoto "github.com/in-toto/in-toto-golang/in_toto"
	slsa "github.com/in-toto/in-toto-golang/in_toto/slsa_provenance/v0.2"

	"google.golang.org/protobuf/testing/protocmp"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestToProto_FullExample(t *testing.T) {
	slsaZeroTwoProvenance, err := provenanceFromFile("testdata/slsa02Provenance.json")
	if err != nil {
		t.Fatal("Unable to read the provenance from a file: ", err)
	}

	got, err := ToProto(slsaZeroTwoProvenance)
	if err != nil {
		t.Fatal("Unable to convert the slsa provenance to proto: ", err)
	}

	expectedSteps, err := structpb.NewValue([]interface{}{
		map[string]interface{}{
			"entryPoint": "set -e\necho \"FROM alpine@sha256:69e70a79f2d41ab5d637de98c1e0b055206ba40a8145e7bddb55ccc04e13cf8f\" | tee $(params.DOCKERFILE)\n",
			"arguments":  nil,
			"environment": map[string]interface{}{
				"container": "add-dockerfile",
				"image":     "docker.io/library/bash@sha256:b3abe4255706618c550e8db5ec0875328333a14dbf663e6f1e2b6875f45521e5",
			},
			"annotations": nil,
		},
		map[string]interface{}{
			"entryPoint": "",
			"arguments": []interface{}{
				"$(params.EXTRA_ARGS)",
				"--dockerfile=$(params.DOCKERFILE)",
				"--context=$(workspaces.source.path)/$(params.CONTEXT)",
				"--destination=$(params.IMAGE)",
				"--digest-file=$(results.IMAGE_DIGEST.path)",
			},
			"environment": map[string]interface{}{
				"container": "build-and-push",
				"image":     "gcr.io/kaniko-project/executor@sha256:c6166717f7fe0b7da44908c986137ecfeab21f31ec3992f6e128fff8a94be8a5",
			},
			"annotations": nil,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	expected := &pb.InTotoStatement{
		Type: "https://in-toto.io/Statement/v0.1",
		Subject: []*pb.Subject{{
			Name:   "us-central1-docker.pkg.dev/test/kaniko-example/kaniko-chains",
			Digest: map[string]string{"sha256": "a2e500bebfe16cf12fc56316ba72c645e1d29054541dc1ab6c286197434170a9"},
		}},
		PredicateType: "https://slsa.dev/provenance/v0.2",
		Predicate: &pb.InTotoStatement_SlsaProvenanceZeroTwo{
			SlsaProvenanceZeroTwo: &pb.SlsaProvenanceZeroTwo{
				Builder: &pb.SlsaProvenanceZeroTwo_SlsaBuilder{
					Id: "https://tekton.dev/chains/v2",
				},
				BuildType: "https://tekton.dev/attestations/chains@v2",
				Invocation: &pb.SlsaProvenanceZeroTwo_SlsaInvocation{
					ConfigSource: &pb.SlsaProvenanceZeroTwo_SlsaConfigSource{},
					Parameters: &structpb.Struct{
						Fields: map[string]*structpb.Value{
							"BUILDER_IMAGE": structpb.NewStringValue("gcr.io/kaniko-project/executor:v1.5.1@sha256:c6166717f7fe0b7da44908c986137ecfeab21f31ec3992f6e128fff8a94be8a5"),
							"CONTEXT":       structpb.NewStringValue("./"),
							"DOCKERFILE":    structpb.NewStringValue("./Dockerfile"),
						},
					},
				},
				BuildConfig: &structpb.Struct{
					Fields: map[string]*structpb.Value{
						"steps": expectedSteps,
					},
				},
				Metadata: &pb.SlsaProvenanceZeroTwo_SlsaMetadata{
					BuildStartedOn:  timestamppb.New(time.Date(2022, time.Month(5), 25, 18, 30, 35, 0, time.UTC)),
					BuildFinishedOn: timestamppb.New(time.Date(2022, time.Month(5), 25, 18, 30, 46, 0, time.UTC)),
					Completeness:    &pb.SlsaProvenanceZeroTwo_SlsaCompleteness{},
					Reproducible:    false,
				},
				Materials: []*pb.SlsaProvenanceZeroTwo_SlsaMaterial{},
			},
		},
	}

	if diff := cmp.Diff(got, expected, protocmp.Transform()); diff != "" {
		t.Errorf("Wrong converted SLSAv0.2 provenance received, diff=%s", diff)
	}
}

func TestToProto_WithSomeEmptyFields(t *testing.T) {
	tests := []struct {
		name     string
		original *intoto.ProvenanceStatement
		expected *pb.InTotoStatement
	}{
		{
			name:     "nil input",
			original: nil,
			expected: nil,
		},
		{
			name:     "all empty",
			original: &intoto.ProvenanceStatement{},
			expected: &pb.InTotoStatement{
				Predicate: &pb.InTotoStatement_SlsaProvenanceZeroTwo{
					SlsaProvenanceZeroTwo: &pb.SlsaProvenanceZeroTwo{
						Builder: &pb.SlsaProvenanceZeroTwo_SlsaBuilder{},
						Invocation: &pb.SlsaProvenanceZeroTwo_SlsaInvocation{
							ConfigSource: &pb.SlsaProvenanceZeroTwo_SlsaConfigSource{},
						},
					},
				}},
		},
		{
			name: "empty BuildStartedOn and BuildFinishedOn in the metadata fields",
			original: &intoto.ProvenanceStatement{
				Predicate: slsa.ProvenancePredicate{
					Metadata: &slsa.ProvenanceMetadata{},
				},
			},
			expected: &pb.InTotoStatement{
				Predicate: &pb.InTotoStatement_SlsaProvenanceZeroTwo{
					SlsaProvenanceZeroTwo: &pb.SlsaProvenanceZeroTwo{
						Builder: &pb.SlsaProvenanceZeroTwo_SlsaBuilder{},
						Invocation: &pb.SlsaProvenanceZeroTwo_SlsaInvocation{
							ConfigSource: &pb.SlsaProvenanceZeroTwo_SlsaConfigSource{},
						},
						Metadata: &pb.SlsaProvenanceZeroTwo_SlsaMetadata{
							Completeness: &pb.SlsaProvenanceZeroTwo_SlsaCompleteness{},
						},
					},
				}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToProto(tt.original)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(got, tt.expected, protocmp.Transform()); diff != "" {
				t.Errorf("Wrong converted SLSAv0.2 provenance received, diff=%s", diff)
			}
		})
	}
}

func provenanceFromFile(f string) (*intoto.ProvenanceStatement, error) {
	rawPayload, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}

	result := intoto.ProvenanceStatement{}
	if err := json.Unmarshal(rawPayload, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
