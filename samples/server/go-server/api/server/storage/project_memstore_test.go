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
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/name"

	"fmt"
	"reflect"
	"sort"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateProject(t *testing.T) {
	ps := NewProjectMemStore()
	p := "myproject"
	if err := ps.CreateProject(p); err != nil {
		t.Errorf("CreateProject got %v want success", err)
	}
	// Try to insert the same project twice, expect failure.
	if err := ps.CreateProject(p); err == nil {
		t.Errorf("CreateProject got success, want Error")
	} else if s, _ := status.FromError(err); s.Code() != codes.AlreadyExists {
		t.Errorf("CreateProject got code %v want %v", s.Code(), codes.AlreadyExists)
	}
}

func TestDeleteProject(t *testing.T) {
	ps := NewProjectMemStore()
	pID := "myproject"
	// Delete before the note exists
	if err := ps.DeleteProject(pID); err == nil {
		t.Error("Deleting nonexistant note got success, want error")
	}
	if err := ps.CreateProject(pID); err != nil {
		t.Fatalf("CreateProject got %v want success", err)
	}

	if err := ps.DeleteProject(pID); err != nil {
		t.Errorf("DeleteProject got %v, want success ", err)
	}
}

func TestGetProject(t *testing.T) {
	ps := NewProjectMemStore()
	pID := "myproject"
	// Try to get project before it has been created, expect failure.
	if _, err := ps.GetProject(pID); err == nil {
		t.Errorf("GetProject got success, want Error")
	} else if s, _ := status.FromError(err); s.Code() != codes.NotFound {
		t.Errorf("GetProject got code %v want %v", s.Code(), codes.NotFound)
	}
	ps.CreateProject(pID)
	if p, err := ps.GetProject(pID); err != nil {
		t.Fatalf("GetProject got %v want success", err)
	} else if p.Name != name.FormatProject(pID) {
		t.Fatalf("Got %s want %s", p.Name, pID)
	}
}

func TestListProjects(t *testing.T) {
	ps := NewProjectMemStore()
	wantProjectNames := []string{}
	for i := 0; i < 20; i++ {
		pID := fmt.Sprint("Project", i)
		if err := ps.CreateProject(pID); err != nil {
			t.Fatalf("CreateProject got %v want success", err)
		}
		wantProjectNames = append(wantProjectNames, name.FormatProject(pID))
	}
	filter := "filters_are_yet_to_be_implemented"
	gotProjects := ps.ListProjects(filter)
	if len(gotProjects) != 20 {
		t.Errorf("ListProjects got %v operations, want 20", len(gotProjects))
	}
	gotProjectNames := make([]string, len(gotProjects))
	for i, project := range gotProjects {
		gotProjectNames[i] = project.Name
	}
	// Sort to handle that wantProjectNames are not guaranteed to be listed in insertion order
	sort.Strings(wantProjectNames)
	sort.Strings(gotProjectNames)
	if !reflect.DeepEqual(gotProjectNames, wantProjectNames) {
		t.Errorf("ListProjects got %v want %v", gotProjectNames, wantProjectNames)
	}
}
