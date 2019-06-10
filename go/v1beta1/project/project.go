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
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grafeas/grafeas/go/errors"
	"github.com/grafeas/grafeas/go/name"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

// Storage provides storage functions for this API.
type Storage interface {
	// CreateProject creates a project with the specified project ID
	CreateProject(ctx context.Context, pID string, p *prpb.Project) (*prpb.Project, error)
	// GetProject gets a project from the datastore
	GetProject(ctx context.Context, pID string) (*prpb.Project, error)
	// ListProjects returns the project IDs for all projects in the datastore
	ListProjects(ctx context.Context, filter string, pageSize int, pageToken string) ([]*prpb.Project, string, error)
	// DeleteProject deletes the specified project
	DeleteProject(ctx context.Context, pID string) error
}

type API struct {
	Storage Storage
}

func (gp *API) CreateProject(ctx context.Context, req *prpb.CreateProjectRequest, resp *prpb.Project) error {
	proj := req.Project
	if proj == nil {
		log.Print("Project must not be empty.")
		return errors.Newf(codes.InvalidArgument, "Project must not be empty")
	}
	if proj.Name == "" {
		log.Printf("Project name must not be empty: %v", proj.Name)
		return errors.Newf(codes.InvalidArgument, "Project name must not be empty")
	}
	pID, err := name.ParseProject(proj.Name)
	if err != nil {
		log.Printf("Invalid project name: %v", proj.Name)
		return errors.Newf(codes.InvalidArgument, "Invalid project name")
	}

	p, err := gp.Storage.CreateProject(ctx, pID, proj)
	if err != nil {
		return err
	}
	*resp = *p
	return nil
}

// GetProject gets a project from the datastore.
func (gp *API) GetProject(ctx context.Context, req *prpb.GetProjectRequest, resp *prpb.Project) error {
	pID, err := name.ParseProject(req.Name)
	if err != nil {
		log.Printf("Error parsing project name: %v", req.Name)
		return errors.Newf(codes.InvalidArgument, "Invalid Project name")
	}
	p, err := gp.Storage.GetProject(ctx, pID)
	if err != nil {
		return err
	}
	*resp = *p
	return nil
}

// ListProjects returns the project id for all projects in the backing datastore.
func (gp *API) ListProjects(ctx context.Context, req *prpb.ListProjectsRequest, resp *prpb.ListProjectsResponse) error {
	// TODO: support filters
	if req.PageSize == 0 {
		req.PageSize = 100
	}
	ps, nextToken, err := gp.Storage.ListProjects(ctx, req.Filter, int(req.PageSize), req.PageToken)
	if err != nil {
		return err
	}
	resp.Projects = ps
	resp.NextPageToken = nextToken
	return nil
}

// DeleteProject deletes a project from the datastore.
func (gp *API) DeleteProject(ctx context.Context, req *prpb.DeleteProjectRequest, _ *empty.Empty) error {
	pID, err := name.ParseProject(req.Name)
	if err != nil {
		log.Printf("Error parsing project name: %v", req.Name)
		return errors.Newf(codes.InvalidArgument, "Invalid Project name")
	}
	if err := gp.Storage.DeleteProject(ctx, pID); err != nil {
		return err
	}
	return nil
}
