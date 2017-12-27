// Copyright 2017 The Grafeas Authors. All rights reserved.
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

package storage

import (
	"fmt"

	"github.com/grafeas/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/grafeas/server-go"
	pb "github.com/grafeas/grafeas/v1alpha1/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// projectMemStore is an in-memory storage solution for Grafeas
type projectMemStore struct {
	projects map[string]bool
}

// NewMemStore creates a memStore with all maps initialized.
func NewProjectMemStore() server.ProjectStorager {
	return &projectMemStore{make(map[string]bool)}
}

// CreateProject adds the specified project to the mem store
func (m *projectMemStore) CreateProject(pID string) error {
	if _, ok := m.projects[pID]; ok {
		return status.Error(codes.AlreadyExists, fmt.Sprintf("Project with name %q already exists", pID))
	}
	m.projects[pID] = true
	return nil
}

// DeleteProject deletes the project with the given pID from the mem store
func (m *projectMemStore) DeleteProject(pID string) error {
	if _, ok := m.projects[pID]; !ok {
		return status.Error(codes.NotFound, fmt.Sprintf("Project with name %q does not Exist", pID))
	}
	delete(m.projects, pID)
	return nil
}

// GetProject returns the project with the given pID from the mem store
func (m *projectMemStore) GetProject(pID string) (*pb.Project, error) {
	if _, ok := m.projects[pID]; !ok {
		return nil, status.Error(codes.NotFound, fmt.Sprintf("Project with name %q does not Exist", pID))
	}
	return &pb.Project{Name: name.FormatProject(pID)}, nil
}

// ListProjects returns the project id for all projects from the mem store
func (m *projectMemStore) ListProjects(filters string) []*pb.Project {
	projects := make([]*pb.Project, len(m.projects))
	i := 0
	for k := range m.projects {
		projects[i] = &pb.Project{Name: name.FormatProject(k)}
		i++
	}
	return projects
}
