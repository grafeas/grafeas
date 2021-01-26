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

package project

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/testing/protocmp"
)

// fakeStorage implements the projects storage interface using an in-memory map for tests. Filters
// and page tokens on list methods aren't supported.
type fakeStorage struct {
	// Map of project IDs to projects.
	projects map[string]*prpb.Project

	// The following errors are for simulating an internal database error.
	createProjErr, getProjErr, listProjErr, deleteProjErr bool
}

func newFakeStorage() *fakeStorage {
	return &fakeStorage{
		projects: map[string]*prpb.Project{},
	}
}

func (s *fakeStorage) CreateProject(ctx context.Context, pID string, p *prpb.Project) (*prpb.Project, error) {
	if s.createProjErr {
		return nil, status.Errorf(codes.Internal, "failed to create project %s", pID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.projects[pID]; !ok {
		s.projects[pID] = p
	}
	return s.projects[pID], nil
}

func (s *fakeStorage) GetProject(ctx context.Context, pID string) (*prpb.Project, error) {
	if s.getProjErr {
		return nil, status.Errorf(codes.Internal, "failed to get project %s", pID)
	}
	p, ok := s.projects[pID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "project %s not found", pID)
	}

	return p, nil
}

func (s *fakeStorage) ListProjects(ctx context.Context, filter string, pageSize int, pageToken string) ([]*prpb.Project, string, error) {
	if s.listProjErr {
		return nil, "", status.Error(codes.Internal, "failed to list projects")
	}

	// Ignore filter, pageSize, and pageToken params in these tests.
	projects := []*prpb.Project{}
	for _, p := range s.projects {
		projects = append(projects, p)
	}

	return projects, "", nil
}

func (s *fakeStorage) DeleteProject(ctx context.Context, pID string) error {
	if s.deleteProjErr {
		return status.Errorf(codes.Internal, "failed to delete project %s", pID)
	}
	delete(s.projects, pID)
	return nil
}

func TestCreateProject(t *testing.T) {
	ctx := context.Background()
	gp := &API{
		Storage: newFakeStorage(),
	}
	proj := &prpb.Project{
		Name: "projects/1234",
	}
	req := &prpb.CreateProjectRequest{
		Project: proj,
	}

	resp, err := gp.CreateProject(ctx, req)
	if err != nil {
		t.Errorf("Got err %v, want success", err)
	}

	if diff := cmp.Diff(proj, resp, protocmp.Transform()); diff != "" {
		t.Errorf("CreateProject(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestCreateProjectErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc               string
		req                *prpb.CreateProjectRequest
		internalStorageErr bool
		wantErrStatus      codes.Code
	}{
		{
			desc:          "empty project",
			req:           &prpb.CreateProjectRequest{},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "empty project name",
			req: &prpb.CreateProjectRequest{
				Project: &prpb.Project{},
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "invalid project name",
			req: &prpb.CreateProjectRequest{
				Project: &prpb.Project{
					Name: "invalid-project",
				},
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "internal storage error",
			req: &prpb.CreateProjectRequest{
				Project: &prpb.Project{
					Name: "projects/hello",
				},
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.createProjErr = tt.internalStorageErr
		gp := &API{
			Storage: s,
		}

		_, err := gp.CreateProject(ctx, tt.req)
		t.Logf("%q: error:%v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestGetProject(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	gp := &API{
		Storage: s,
	}

	// Create the project to get.
	proj := &prpb.Project{
		Name: "projects/1234",
	}

	if _, err := s.CreateProject(ctx, "1234", proj); err != nil {
		t.Errorf("Failed to create project %+v", proj)
	}

	req := &prpb.GetProjectRequest{
		Name: "projects/1234",
	}
	gotP, err := gp.GetProject(ctx, req)
	if err != nil {
		t.Errorf("Got err %v, want success", err)
	}

	if diff := cmp.Diff(proj, gotP, protocmp.Transform()); diff != "" {
		t.Errorf("GetProject(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestGetProjectErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc               string
		req                *prpb.GetProjectRequest
		internalStorageErr bool
		wantErrStatus      codes.Code
	}{
		{
			desc: "invalid project name",
			req: &prpb.GetProjectRequest{
				Name: "invalid-project",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "internal storage error",
			req: &prpb.GetProjectRequest{
				Name: "projects/hello",
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.getProjErr = tt.internalStorageErr
		gp := &API{
			Storage: s,
		}

		_, err := gp.GetProject(ctx, tt.req)
		t.Logf("%q: error:%v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestListProjects(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	gp := &API{
		Storage: s,
	}

	// Create the project to list.
	proj := &prpb.Project{
		Name: "projects/1234",
	}

	if _, err := s.CreateProject(ctx, "1234", proj); err != nil {
		t.Errorf("Failed to create project %+v", proj)
	}

	req := &prpb.ListProjectsRequest{}
	resp, err := gp.ListProjects(ctx, req)
	if err != nil {
		t.Errorf("Got err %v, want success", err)
	}

	if diff := cmp.Diff(proj, resp.Projects[0], protocmp.Transform()); diff != "" {
		t.Errorf("ListProjects(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestListProjectsErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc               string
		req                *prpb.ListProjectsRequest
		internalStorageErr bool
		wantErrStatus      codes.Code
	}{
		{
			desc:               "internal storage error",
			req:                &prpb.ListProjectsRequest{},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.listProjErr = tt.internalStorageErr
		gp := &API{
			Storage: s,
		}

		_, err := gp.ListProjects(ctx, tt.req)
		t.Logf("%q: error:%v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestDeleteProject(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	gp := &API{
		Storage: s,
	}

	// Create the project to delete.
	proj := &prpb.Project{
		Name: "projects/1234",
	}

	if _, err := s.CreateProject(ctx, "1234", proj); err != nil {
		t.Errorf("Failed to create project %+v", proj)
	}

	req := &prpb.DeleteProjectRequest{
		Name: "projects/1234",
	}
	if p, err := gp.DeleteProject(ctx, req); err != nil {
		t.Errorf("Got err %v, want success", err)
	} else if p == nil {
		t.Error("expected response to not be nil")
	}
}

func TestDeleteProjectErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc               string
		req                *prpb.DeleteProjectRequest
		internalStorageErr bool
		wantErrStatus      codes.Code
	}{
		{
			desc: "invalid project name",
			req: &prpb.DeleteProjectRequest{
				Name: "invalid-project",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "internal storage error",
			req: &prpb.DeleteProjectRequest{
				Name: "projects/hello",
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.deleteProjErr = tt.internalStorageErr
		gp := &API{
			Storage: s,
		}

		_, err := gp.DeleteProject(ctx, tt.req)
		t.Logf("%q: error:%v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}
