// Copyright 2019 The Grafeas Authors. All rights reserved.
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

package main

import (
	"context"
	"log"

	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	buildpb "github.com/grafeas/grafeas/proto/v1beta1/build_go_proto"
	deppb "github.com/grafeas/grafeas/proto/v1beta1/deployment_go_proto"
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	packpb "github.com/grafeas/grafeas/proto/v1beta1/package_go_proto"
	projpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	provpb "github.com/grafeas/grafeas/proto/v1beta1/provenance_go_proto"
	sourcepb "github.com/grafeas/grafeas/proto/v1beta1/source_go_proto"
	vulpb "github.com/grafeas/grafeas/proto/v1beta1/vulnerability_go_proto"

	"google.golang.org/grpc"
)

func main() {
	// demonstrates creation of two projects,
	// creation of notes in one project and occurences in the other
	// shows creation of build, vulnerability and deploy types
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	defer conn.Close()

	projectClient := projpb.NewProjectsClient(conn)

	log.Println("Create a new project to store notes")
	_, err = projectClient.CreateProject(context.Background(), &projpb.CreateProjectRequest{
		Project: &projpb.Project{Name: "projects/provider_example"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Get a single project
	log.Println("Get the details of the new project")
	project, err := projectClient.GetProject(context.Background(), &projpb.GetProjectRequest{
		Name: "projects/provider_example",
	})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(project)
	}

	// Create a second project
	log.Println("now create a second project for the occurrences")
	_, err = projectClient.CreateProject(context.Background(), &projpb.CreateProjectRequest{
		Project: &projpb.Project{Name: "projects/occurrence_example"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// List projects
	log.Println("list out all the projects")
	projectResp, err := projectClient.ListProjects(context.Background(),
		&projpb.ListProjectsRequest{})

	if err != nil {
		log.Fatal(err)
	}

	if len(projectResp.Projects) != 0 {
		log.Println(projectResp.Projects)
	} else {
		log.Printf("no projects")
	}

	client := pb.NewGrafeasV1Beta1Client(conn)

	log.Println("create a vulnerability note")
	vulDetails := vulpb.Vulnerability_Detail{
		Package: "libexempi3",
		CpeUri:  "cpe:/o:debian:debian_linux:7",
		MinAffectedVersion: &packpb.Version{
			Name:     "2.5.7",
			Revision: "1",
			Kind:     1,
		},
	}

	noteReq := pb.CreateNoteRequest{
		Parent: "projects/provider_example",
		NoteId: "testVulnerabilityNote",
		Note: &pb.Note{
			Name:             "projects/provider_example/notes/testVulnerabilityNote",
			ShortDescription: "A vulnerability note",
			LongDescription:  "A longer description vulnerability note",
			Kind:             1,
			Type: &pb.Note_Vulnerability{
				Vulnerability: &vulpb.Vulnerability{
					CvssScore: 1.23,
					Severity:  3,
					Details: []*vulpb.Vulnerability_Detail{
						&vulDetails,
					},
				},
			},
		},
	}

	_, err = client.CreateNote(context.Background(), &noteReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("create a deployment note")
	noteReq = pb.CreateNoteRequest{
		Parent: "projects/provider_example",
		NoteId: "testDeploymentNote",
		Note: &pb.Note{
			Name:             "projects/provider_example/notes/testDeploymentNote",
			ShortDescription: "A deployment note",
			LongDescription:  "A longer description deployment note",
			Kind:             5,
			Type: &pb.Note_Deployable{
				Deployable: &deppb.Deployable{
					ResourceUri: []string{
						"http://somewhere",
					},
				},
			},
		},
	}

	_, err = client.CreateNote(context.Background(), &noteReq)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("create a build note")
	noteReq = pb.CreateNoteRequest{
		Parent: "projects/provider_example",
		NoteId: "testBuildNote",
		Note: &pb.Note{
			Name:             "projects/provider_example/notes/testBuildNote",
			ShortDescription: "A build note",
			LongDescription:  "A longer description build note",
			Kind:             2,
			Type: &pb.Note_Build{
				Build: &buildpb.Build{
					BuilderVersion: "some version",
					Signature: &buildpb.BuildSignature{
						PublicKey: "some public key",
						Signature: []byte("Z3JhZmVhcw=="),
						KeyId:     "some key identifier",
						KeyType:   0,
					},
				},
			},
		},
	}
	_, err = client.CreateNote(context.Background(), &noteReq)
	if err != nil {
		log.Fatal(err)
	}

	// List notes
	log.Println("list all the notes for the provider project")
	resp, err := client.ListNotes(context.Background(),
		&pb.ListNotesRequest{
			Parent: "projects/provider_example",
		})
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Notes) != 0 {
		log.Println(resp.Notes)
	} else {
		log.Printf("Project 'provider_example' does not contain any notes")
	}

	// --- now occurrernces ---

	// Create occurence for the deployment note in the second project
	log.Println("create a deployment occurrence")

	deploymentDetails := deppb.Deployment{
		UserEmail: "some@email.com",
		DeployTime: &timestamp.Timestamp{
			Seconds: time.Now().UTC().UnixNano(),
		},
		Config:   "some deployment config",
		Address:  "some hosting platform",
		Platform: deppb.Deployment_CUSTOM,
	}
	occurenceDeployment := pb.Occurrence_Deployment{
		Deployment: &deppb.Details{Deployment: &deploymentDetails},
	}
	occurenceDetails := pb.Occurrence{
		Name: "projects/occurrence_example/occurences/testDeploymentOccurrence",
		Resource: &pb.Resource{
			Name: "some os",
			Uri:  "http://someuri",
		},
		NoteName: "projects/occurrence_example/notes/testDeploymentOccurrence",
		Kind:     5,
		Details:  &occurenceDeployment,
	}

	occurenceRequest := pb.CreateOccurrenceRequest{
		Parent:     "projects/occurrence_example",
		Occurrence: &occurenceDetails,
	}

	_, err = client.CreateOccurrence(context.Background(), &occurenceRequest)
	if err != nil {
		log.Fatal(err)
	}

	//create a vulnerability occurrence
	log.Println("create a vulnerability occurrence")
	packageIssue := vulpb.PackageIssue{
		AffectedLocation: &vulpb.VulnerabilityLocation{
			CpeUri:  "7",
			Package: "a",
			Version: &packpb.Version{
				Name:     "v1.1.1",
				Kind:     3,
				Revision: "r",
			},
		},
		FixedLocation: &vulpb.VulnerabilityLocation{
			CpeUri:  "cpe:/o:debian:debian_linux:7",
			Package: "a",
			Version: &packpb.Version{
				Name:     "namestring",
				Kind:     3,
				Revision: "1",
			},
		},
	}
	vulnerability := vulpb.Details{
		Type: "the type of package",
		PackageIssue: []*vulpb.PackageIssue{
			&packageIssue,
		},
	}
	occurrenceVulnerabilityDetails := pb.Occurrence_Vulnerability{
		Vulnerability: &vulnerability,
	}

	occurenceDetails = pb.Occurrence{
		Name: "projects/occurrence_example/occurences/testVulnerabilityOccurrence",
		Resource: &pb.Resource{
			Name: "some os",
			Uri:  "http://someuri",
		},
		NoteName: "projects/occurrence_example/notes/testVulnerabilityOccurrence",
		Kind:     5,
		Details:  &occurrenceVulnerabilityDetails,
	}

	occurenceRequest = pb.CreateOccurrenceRequest{
		Parent:     "projects/occurrence_example",
		Occurrence: &occurenceDetails,
	}

	_, err = client.CreateOccurrence(context.Background(), &occurenceRequest)
	if err != nil {
		log.Fatal(err)
	}

	// now create a build occurrence
	log.Println("create a build occurrence")
	build := buildpb.Details{
		Provenance: &provpb.BuildProvenance{
			Id:        "build identifier",
			ProjectId: "some project identifier",
			Commands:  []*provpb.Command{},
			BuiltArtifacts: []*provpb.Artifact{
				&provpb.Artifact{
					Checksum: "123",
					Id:       "some identifier for the artifact",
					Names: []string{
						"name of the related artifact",
					},
				},
			},
			CreateTime: &timestamp.Timestamp{
				Seconds: time.Now().UTC().UnixNano(),
			},
			StartTime: &timestamp.Timestamp{
				Seconds: time.Now().UTC().UnixNano(),
			},
			EndTime: &timestamp.Timestamp{
				Seconds: time.Now().UTC().UnixNano(),
			},
			Creator: "User initiating the build",
			LogsUri: "location of the build logs",
			SourceProvenance: &provpb.Source{
				ArtifactStorageSourceUri: "input binary artifacts from this build",
				Context: &sourcepb.SourceContext{
					Context: &sourcepb.SourceContext_Git{
						Git: &sourcepb.GitSourceContext{
							Url:        "the git repo url",
							RevisionId: "git commit hash",
						},
					},
				},
			},
			TriggerId:      "triggered by code commit 123",
			BuilderVersion: "some version of the builder",
		},
		ProvenanceBytes: "Z3JhZmVhcw==",
	}
	occurrenceBuildDetails := pb.Occurrence_Build{
		Build: &build,
	}

	occurenceDetails = pb.Occurrence{
		Name: "projects/occurrence_example/occurences/testBuildOccurrence",
		Resource: &pb.Resource{
			Name: "some os",
			Uri:  "http://someuri",
		},
		NoteName: "projects/occurrence_example/notes/testBuildOccurrence",
		Kind:     2,
		Details:  &occurrenceBuildDetails,
	}

	_, err = client.CreateOccurrence(context.Background(), &occurenceRequest)
	if err != nil {
		log.Fatal(err)
	}

	// List notes
	log.Println("list all the occurrences for the provider project")
	respOccurrences, err := client.ListOccurrences(context.Background(),
		&pb.ListOccurrencesRequest{
			Parent: "projects/occurrence_example",
		})
	if err != nil {
		log.Fatal(err)
	}
	if len(respOccurrences.Occurrences) != 0 {
		log.Println(respOccurrences.Occurrences)
	} else {
		log.Printf("Project 'occurrence_example' does not contain any occurrences")
	}

}
