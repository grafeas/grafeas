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

package server

import (
	pb "github.com/grafeas/grafeas/v1alpha1/proto"
)

// ProjectStorager is the interface that a Grafeas storage implementation would provide
// for Project enteties
type ProjectStorager interface {
	// CreateProject adds the specified project
	CreateProject(pID string) error

	// DeleteNote deletes the project with the given pID
	DeleteProject(pID string) error

	// GetProject returns the project with the given pID
	GetProject(pID string) (*pb.Project, error)

	// ListProjects returns the project id for all projects
	ListProjects(filters string) []*pb.Project
}
