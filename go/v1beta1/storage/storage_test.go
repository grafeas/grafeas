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

package storage_test

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/grafeas/grafeas/go/name"
	grafeas "github.com/grafeas/grafeas/go/v1beta1/api"
	"github.com/grafeas/grafeas/go/v1beta1/project"
	cpb "github.com/grafeas/grafeas/proto/v1beta1/common_go_proto"
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	pkgpb "github.com/grafeas/grafeas/proto/v1beta1/package_go_proto"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	vpb "github.com/grafeas/grafeas/proto/v1beta1/vulnerability_go_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	testNoteID = "CVE-1999-0710"
)

var (
	opt = cmp.FilterPath(
		func(p cmp.Path) bool {
			ignoreCreate := p.String() == "CreateTime"
			ignoreUpdate := p.String() == "UpdateTime"
			// Remove ignoring of the fields below once go-cmp is able to ignore generated fields.
			// See https://github.com/google/go-cmp/issues/153
			ignoreXXXCache :=
				p.String() == "XXX_sizecache" ||
					p.String() == "Resource.XXX_sizecache" ||
					p.String() == "Details.Vulnerability.XXX_sizecache" ||
					p.String() == "Details.Vulnerability.PackageIssue.XXX_sizecache" ||
					p.String() == "Details.Vulnerability.PackageIssue.AffectedLocation.XXX_sizecache" ||
					p.String() == "Details.Vulnerability.PackageIssue.AffectedLocation.Version.XXX_sizecache" ||
					p.String() == "Details.Vulnerability.PackageIssue.FixedLocation.XXX_sizecache" ||
					p.String() == "Details.Vulnerability.PackageIssue.FixedLocation.Version.XXX_sizecache"
			return ignoreCreate || ignoreUpdate || ignoreXXXCache
		}, cmp.Ignore())
)

// Tests implementations of grafeas.Storage and project.Storage
// createStore is a function that creates new grafeas.Storage and project.Storage instances and
// a corresponding cleanUp function that will be run at the end of each
// test case.
// TODO: add testing for CreateTime and UpdateTime
func doTestStorage(t *testing.T, createStore func(t *testing.T) (grafeas.Storage, project.Storage, func())) {
	t.Run("CreateProject", func(t *testing.T) {
		_, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		pID := "project"
		p := &prpb.Project{}
		p.Name = name.FormatProject(pID)
		if _, err := gp.CreateProject(ctx, pID, p); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}
	})

	t.Run("CreateSameProjectTwice", func(t *testing.T) {
		_, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		pID := "project"
		p := &prpb.Project{}
		p.Name = name.FormatProject(pID)
		if _, err := gp.CreateProject(ctx, pID, p); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}
		// Try to insert the same project twice, expect failure.
		if _, err := gp.CreateProject(ctx, pID, p); err == nil {
			t.Errorf("CreateProject got success, want Error")
		} else if s, _ := status.FromError(err); s.Code() != codes.AlreadyExists {
			t.Errorf("CreateProject got code %v want %v", s.Code(), codes.AlreadyExists)
		}
	})

	t.Run("CreateNote", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Errorf("CreateNote got %v want success", err)
		}
	})

	t.Run("CreateSameNoteTwice", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Errorf("CreateNote got %v want success", err)
		}
		// Try to insert the same note twice, expect failure.
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err == nil {
			t.Errorf("CreateNote got success, want Error")
		} else if s, _ := status.FromError(err); s.Code() != codes.AlreadyExists {
			t.Errorf("CreateNote got code %v want %v", s.Code(), codes.AlreadyExists)
		}
	})

	t.Run("CreateOccurrence", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Errorf("CreateNote got %v want success", err)
		}

		oPID := "occurrence-project"
		o := createTestOccurrence(oPID, n.Name)
		oo, err := g.CreateOccurrence(ctx, oPID, "userID", o)
		if err != nil {
			t.Errorf("CreateOccurrence got %v want success", err)
		}

		pID, oID, err := name.ParseOccurrence(oo.Name)
		if err != nil {
			t.Fatalf("Error parsing projectID and occurrenceID %v", err)
		}

		got, err := g.GetOccurrence(ctx, pID, oID)
		if err != nil {
			t.Fatalf("GetOccurrence got %v, want success", err)
		}
		if diff := cmp.Diff(got, oo, opt); diff != "" {
			t.Errorf("GetOccurrence returned diff (want -> got):\n%s", diff)
		}
	})

	t.Run("CreateSameOccurrenceTwice", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Errorf("CreateNote got %v want success", err)
		}

		oPID := "occurrence-project"
		o := createTestOccurrence(oPID, n.Name)
		if _, err := g.CreateOccurrence(ctx, nPID, "userID", o); err != nil {
			t.Errorf("CreateOccurrence got %v want success", err)
		}
		// Try to insert the same occurrence twice, expect success, because different IDs are generated.
		if _, err := g.CreateOccurrence(ctx, nPID, "userID", o); err != nil {
			t.Errorf("CreateOccurrence got %v, want success", err)
		}
	})

	t.Run("DeleteProject", func(t *testing.T) {
		_, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		pID := "myproject"
		// Delete before the project exists
		if err := gp.DeleteProject(ctx, pID); err == nil {
			t.Error("Deleting nonexistant note got success, want error")
		}
		if _, err := gp.CreateProject(ctx, pID, &prpb.Project{}); err != nil {
			t.Fatalf("CreateProject got %v want success", err)
		}

		if err := gp.DeleteProject(ctx, pID); err != nil {
			t.Errorf("DeleteProject got %v, want success ", err)
		}
	})

	t.Run("DeleteOccurrence", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		oPID := "occurrence-project"
		o := createTestOccurrence(oPID, n.Name)
		oo, err := g.CreateOccurrence(ctx, oPID, "userID", o)
		if err != nil {
			t.Fatalf("CreateOccurrence got %v want success", err)
		}
		pID, oID, err := name.ParseOccurrence(oo.Name)
		if err != nil {
			t.Fatalf("Error parsing occurrence %v", err)
		}
		if err := g.DeleteOccurrence(ctx, pID, oID); err != nil {
			t.Errorf("DeleteOccurrence got %v, want success ", err)
		}
	})

	t.Run("UpdateOccurrence", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		oPID := "occurrence-project"
		o := createTestOccurrence(oPID, n.Name)
		oo, err := g.CreateOccurrence(ctx, oPID, "userID", o)
		if err != nil {
			t.Fatalf("CreateOccurrence got %v want success", err)
		}
		pID, oID, err := name.ParseOccurrence(oo.Name)
		if err != nil {
			t.Fatalf("Error parsing projectID and occurrenceID %v", err)
		}
		got, err := g.GetOccurrence(ctx, pID, oID)
		if err != nil {
			t.Fatalf("GetOccurrence got %v, want success", err)
		}

		if diff := cmp.Diff(got, oo, opt); diff != "" {
			t.Errorf("GetOccurrence returned diff (want -> got):\n%s", diff)
		}

		o2 := oo
		o2.GetVulnerability().CvssScore = 1.0
		// TODO(#312): check the result of the update
		// TODO(#312): use fieldmask in the param
		if _, err := g.UpdateOccurrence(ctx, pID, oID, o2, nil); err != nil {
			t.Fatalf("UpdateOccurrence got %v want success", err)
		}

		got, err = g.GetOccurrence(ctx, pID, oID)
		if err != nil {
			t.Fatalf("GetOccurrence got %v, want success", err)
		}
		if diff := cmp.Diff(got, o2, opt); diff != "" {
			t.Errorf("GetOccurrence returned diff (want -> got):\n%s", diff)
		}
	})

	t.Run("DeleteNote", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		// Delete before the note exists
		pID, nID, err := name.ParseNote(n.Name)
		if err != nil {
			t.Fatalf("Error parsing note %v", err)
		}
		if err := g.DeleteNote(ctx, pID, nID); err == nil {
			t.Error("Deleting nonexistant note got success, want error")
		}
		if _, err := g.CreateNote(ctx, pID, nID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		if err := g.DeleteNote(ctx, pID, nID); err != nil {
			t.Errorf("DeleteNote got %v, want success ", err)
		}
	})

	t.Run("UpdateNote", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		pID, nID, err := name.ParseNote(n.Name)
		if err != nil {
			t.Fatalf("Error parsing projectID and noteID %v", err)
		}
		if _, err := g.UpdateNote(ctx, pID, nID, n, nil); err == nil {
			t.Fatal("UpdateNote got success want error")
		}
		if _, err := g.CreateNote(ctx, pID, nID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}
		got, err := g.GetNote(ctx, pID, nID)
		if err != nil {
			t.Fatalf("GetNote got %v, want success", err)
		}
		if diff := cmp.Diff(got, n, opt); diff != "" {
			t.Errorf("GetNote returned diff (want -> got):\n%s", diff)
		}

		n2 := n
		n2.GetVulnerability().CvssScore = 1.0
		// TODO(#312): check the result of the update
		// TODO(#312): use fieldmask in the param
		if _, err := g.UpdateNote(ctx, pID, nID, n2, nil); err != nil {
			t.Fatalf("UpdateNote got %v want success", err)
		}

		got, err = g.GetNote(ctx, pID, nID)
		if err != nil {
			t.Fatalf("GetNote got %v, want success", err)
		}
		if diff := cmp.Diff(got, n2, opt); diff != "" {
			t.Errorf("GetNote returned diff (want -> got):\n%s", diff)
		}
	})

	t.Run("GetProject", func(t *testing.T) {
		_, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		pID := "myproject"
		// Try to get project before it has been created, expect failure.
		if _, err := gp.GetProject(ctx, pID); err == nil {
			t.Errorf("GetProject got success, want Error")
		} else if s, _ := status.FromError(err); s.Code() != codes.NotFound {
			t.Errorf("GetProject got code %v want %v", s.Code(), codes.NotFound)
		}

		p := &prpb.Project{}
		p.Name = name.FormatProject(pID)
		_, err := gp.CreateProject(ctx, pID, p)
		if err != nil {
			t.Fatalf("CreateProject got %v want success", err)
		}

		if proj, err := gp.GetProject(ctx, pID); err != nil {
			t.Fatalf("GetProject got %v want success", err)
		} else if p.Name != proj.Name {
			t.Fatalf("Got %s want %s", p.Name, pID)
		}
	})

	t.Run("GetOccurrence", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		oPID := "occurrence-project"
		o := createTestOccurrence(oPID, n.Name)
		oo, err := g.CreateOccurrence(ctx, oPID, "userID", o)
		if err != nil {
			t.Errorf("CreateOccurrence got %v, want Success", err)
		}
		pID, oID, err := name.ParseOccurrence(oo.Name)
		if err != nil {
			t.Fatalf("Error parsing occurrence %v", err)
		}

		got, err := g.GetOccurrence(ctx, pID, oID)
		if err != nil {
			t.Fatalf("GetOccurrence got %v, want success", err)
		}
		if diff := cmp.Diff(got, oo, opt); diff != "" {
			t.Errorf("GetOccurrence returned diff (want -> got):\n%s", diff)
		}
	})

	t.Run("GetNote", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		pID, nID, err := name.ParseNote(n.Name)
		if err != nil {
			t.Fatalf("Error parsing note %v", err)
		}
		if _, err := g.GetNote(ctx, pID, nID); err == nil {
			t.Fatal("GetNote got success, want error")
		}
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Errorf("CreateNote got %v, want Success", err)
		}

		got, err := g.GetNote(ctx, pID, nID)
		if err != nil {
			t.Fatalf("GetNote got %v, want success", err)
		}
		if diff := cmp.Diff(got, n, opt); diff != "" {
			t.Errorf("GetNote returned diff (want -> got):\n%s", diff)
		}
	})

	t.Run("GetOccurrenceNote", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		nPID := "vulnerability-scanner-a"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		oPID := "occurrence-project"
		o := createTestOccurrence(oPID, n.Name)
		oo, err := g.CreateOccurrence(ctx, oPID, "userID", o)
		if err != nil {
			t.Errorf("CreateOccurrence got %v, want Success", err)
		}
		pID, oID, err := name.ParseOccurrence(oo.Name)
		if err != nil {
			t.Fatalf("Error parsing occurrence %v", err)
		}

		got, err := g.GetOccurrenceNote(ctx, pID, oID)
		if err != nil {
			t.Fatalf("GetOccurrenceNote got %v, want success", err)
		}
		if diff := cmp.Diff(got, n, opt); diff != "" {
			t.Errorf("GetOccurrenceNote returned diff (want -> got):\n%s", diff)
		}
	})

	t.Run("ListProjects", func(t *testing.T) {
		_, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		wantProjectNames := []string{}
		for i := 0; i < 20; i++ {
			pID := fmt.Sprint("Project", i)
			p := &prpb.Project{}
			p.Name = name.FormatProject(pID)
			p, err := gp.CreateProject(ctx, pID, p)
			if err != nil {
				t.Fatalf("CreateProject got %v want success", err)
			}
			wantProjectNames = append(wantProjectNames, p.Name)
		}

		filter := "filters_are_yet_to_be_implemented"
		gotProjects, pageToken, err := gp.ListProjects(ctx, filter, 100, "")
		if err != nil {
			t.Fatalf("ListProjects got %v want success", err)
		}
		if pageToken != "" {
			t.Errorf("Got %s want empty page token", pageToken)
		}
		if len(gotProjects) != 20 {
			t.Errorf("ListProjects got %v projects, want 20", len(gotProjects))
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
	})

	t.Run("ListNotes", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		findProject := "findThese"
		if _, err := gp.CreateProject(ctx, findProject, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}
		dontFind := "dontFind"
		if _, err := gp.CreateProject(ctx, dontFind, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		ns := []*pb.Note{}
		for i := 0; i < 20; i++ {
			n := createTestNote("")
			nPID := ""
			if i < 5 {
				n.Name = name.FormatNote(findProject, strconv.Itoa(i))
				nPID = findProject
			} else {
				n.Name = name.FormatNote(dontFind, strconv.Itoa(i))
				nPID = dontFind
			}
			if _, err := g.CreateNote(ctx, nPID, n.Name, "userID", n); err != nil {
				t.Fatalf("CreateNote got %v want success", err)
			}
			ns = append(ns, n)
		}

		filter := "filters_are_yet_to_be_implemented"
		gotNs, _, err := g.ListNotes(ctx, findProject, filter, "", 100)
		if err != nil {
			t.Fatalf("ListNotes got %v want success", err)
		}
		if len(gotNs) != 5 {
			t.Errorf("ListNotes got %v notes, want 5", len(gotNs))
		}
		for _, n := range gotNs {
			want := name.FormatProject(findProject)
			if !strings.HasPrefix(n.Name, want) {
				t.Errorf("ListNotes got %v want %v", n.Name, want)
			}
		}
	})

	t.Run("ListOccurrences", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		findProject := "findThese"
		if _, err := gp.CreateProject(ctx, findProject, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}
		dontFind := "dontFind"
		if _, err := gp.CreateProject(ctx, dontFind, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		nFind := createTestNote(findProject)
		if _, err := g.CreateNote(ctx, findProject, testNoteID, "userID", nFind); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}
		nDontFind := createTestNote(dontFind)
		if _, err := g.CreateNote(ctx, dontFind, testNoteID, "userID", nDontFind); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		os := []*pb.Occurrence{}
		for i := 0; i < 20; i++ {
			oPID := ""
			o := createTestOccurrence("", "")
			if i < 5 {
				oPID = findProject
				o.NoteName = nFind.Name
			} else {
				oPID = dontFind
				o.NoteName = nDontFind.Name
			}
			oo, err := g.CreateOccurrence(ctx, oPID, "userID", o)
			if err != nil {
				t.Fatalf("CreateOccurrence got %v want success", err)
			}
			os = append(os, oo)
		}

		filter := "filters_are_yet_to_be_implemented"
		gotOs, _, err := g.ListOccurrences(ctx, findProject, filter, "", 100)
		if err != nil {
			t.Fatalf("ListOccurrences got %v want success", err)
		}
		if len(gotOs) != 5 {
			t.Errorf("ListOccurrences got %v Occurrences, want 5", len(gotOs))
		}
		for _, o := range gotOs {
			want := name.FormatProject(findProject)
			if !strings.HasPrefix(o.Name, want) {
				t.Errorf("ListOccurrences got %v want  %v", o.Name, want)
			}
		}
	})

	t.Run("ListNoteOccurrences", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		findProject := "findThese"
		if _, err := gp.CreateProject(ctx, findProject, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}
		dontFind := "dontFind"
		if _, err := gp.CreateProject(ctx, dontFind, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(findProject)
		if _, err := g.CreateNote(ctx, findProject, testNoteID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		os := []*pb.Occurrence{}
		for i := 0; i < 20; i++ {
			oPID := ""
			o := createTestOccurrence("", "")

			if i < 5 {
				oPID = findProject
			} else {
				oPID = dontFind
			}
			o.NoteName = n.Name
			oo, err := g.CreateOccurrence(ctx, oPID, "userID", o)
			if err != nil {
				t.Fatalf("CreateOccurrence got %v want success", err)
			}
			os = append(os, oo)
		}

		pID, nID, err := name.ParseNote(n.Name)
		if err != nil {
			t.Fatalf("Error parsing note name %v", err)
		}
		filter := "filters_are_yet_to_be_implemented"
		gotOs, _, err := g.ListNoteOccurrences(ctx, pID, nID, filter, "", 100)
		if err != nil {
			t.Fatalf("ListNoteOccurrences got %v want success", err)
		}
		if len(gotOs) != 20 {
			t.Errorf("ListNoteOccurrences got %v Occurrences, want 20", len(gotOs))
		}
		for _, o := range gotOs {
			if o.NoteName != n.Name {
				t.Errorf("ListNoteOccurrences got %v want  %v", o.Name, o.NoteName)
			}
		}
	})

	t.Run("ProjectPagination", func(t *testing.T) {
		_, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		p1 := &prpb.Project{
			Name: "projects/project1",
		}
		p1ID := "project1"
		if _, err := gp.CreateProject(ctx, p1ID, p1); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		p2 := &prpb.Project{
			Name: "projects/project2",
		}
		p2ID := "project2"
		if _, err := gp.CreateProject(ctx, p2ID, p2); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		p3 := &prpb.Project{
			Name: "projects/project3",
		}
		p3ID := "project3"
		if _, err := gp.CreateProject(ctx, p3ID, p3); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}
		filter := "filters_are_yet_to_be_implemented"
		// Get projects
		gotProjects, lastPage, err := gp.ListProjects(ctx, filter, 2, "")
		if err != nil {
			t.Fatalf("ListProjects got %v want success", err)
		}
		if len(gotProjects) != 2 {
			t.Errorf("ListProjects got %v projects, want 2", len(gotProjects))
		}
		if p := gotProjects[0]; p.Name != p1.Name {
			t.Errorf("Got %s want %s", p.Name, p1.Name)
		}
		if p := gotProjects[1]; p.Name != p2.Name {
			t.Errorf("Got %s want %s", p.Name, p2.Name)
		}
		// Get projects again
		gotProjects, pageToken, err := gp.ListProjects(ctx, filter, 100, lastPage)
		if err != nil {
			t.Fatalf("ListProjects got %v want success", err)
		}
		if pageToken != "" {
			t.Fatalf("Got %s want empty page token", pageToken)
		}
		if len(gotProjects) != 1 {
			t.Errorf("ListProjects got %v projects, want 1", len(gotProjects))
		}
		if p := gotProjects[0]; p.Name != p3.Name {
			t.Fatalf("Got %s want %s", p.Name, p3.Name)
		}
	})

	t.Run("NotesPagination", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		pID := "project"
		if _, err := gp.CreateProject(ctx, pID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		nID1 := "note1"
		op1 := createTestNote(pID)
		op1.Name = name.FormatNote(pID, nID1)
		if _, err := g.CreateNote(ctx, pID, nID1, "userID", op1); err != nil {
			t.Errorf("CreateNote got %v want success", err)
		}

		nID2 := "note2"
		op2 := createTestNote(pID)
		op2.Name = name.FormatNote(pID, nID2)
		if _, err := g.CreateNote(ctx, pID, nID2, "userID", op2); err != nil {
			t.Errorf("CreateNote got %v want success", err)
		}

		nID3 := "note3"
		op3 := createTestNote(pID)
		op3.Name = name.FormatNote(pID, nID3)
		if _, err := g.CreateNote(ctx, pID, nID3, "userID", op3); err != nil {
			t.Errorf("CreateNote got %v want success", err)
		}
		filter := "filters_are_yet_to_be_implemented"
		// Get occurrences
		gotNotes, lastPage, err := g.ListNotes(ctx, pID, filter, "", 2)
		if err != nil {
			t.Fatalf("ListNotes got %v want success", err)
		}
		if len(gotNotes) != 2 {
			t.Errorf("ListNotes got %v notes, want 2", len(gotNotes))
		}
		if p := gotNotes[0]; p.Name != name.FormatNote(pID, nID1) {
			t.Fatalf("Got %s want %s", p.Name, name.FormatNote(pID, nID1))
		}
		if p := gotNotes[1]; p.Name != name.FormatNote(pID, nID2) {
			t.Fatalf("Got %s want %s", p.Name, name.FormatNote(pID, nID2))
		}
		// Get occurrences again
		gotNotes, pageToken, err := g.ListNotes(ctx, pID, filter, lastPage, 100)
		if err != nil {
			t.Fatalf("ListNotes got %v want success", err)
		}
		if pageToken != "" {
			t.Errorf("Got %s want empty page token", pageToken)
		}
		if len(gotNotes) != 1 {
			t.Errorf("ListNotes got %v notes, want 1", len(gotNotes))
		}
		if p := gotNotes[0]; p.Name != name.FormatNote(pID, nID3) {
			t.Fatalf("Got %s want %s", p.Name, name.FormatNote(pID, nID3))
		}
	})

	t.Run("OccurrencePagination", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		pID := "project"
		if _, err := gp.CreateProject(ctx, pID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(pID)
		if _, err := g.CreateNote(ctx, pID, testNoteID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		op1 := createTestOccurrence(pID, n.Name)
		if _, err := g.CreateOccurrence(ctx, pID, "userID", op1); err != nil {
			t.Errorf("CreateOccurrence got %v want success", err)
		}

		op2 := createTestOccurrence(pID, n.Name)
		if _, err := g.CreateOccurrence(ctx, pID, "userID", op2); err != nil {
			t.Errorf("CreateOccurrence got %v want success", err)
		}

		op3 := createTestOccurrence(pID, n.Name)
		if _, err := g.CreateOccurrence(ctx, pID, "userID", op3); err != nil {
			t.Errorf("CreateOccurrence got %v want success", err)
		}
		filter := "filters_are_yet_to_be_implemented"
		// Get occurrences
		gotOccurrences, lastPage, err := g.ListOccurrences(ctx, pID, filter, "", 2)
		if err != nil {
			t.Fatalf("ListOccurrences got %v want success", err)
		}
		if len(gotOccurrences) != 2 {
			t.Errorf("ListOccurrences got %v occurrences, want 2", len(gotOccurrences))
		}
		// Get occurrences again
		gotOccurrences, pageToken, err := g.ListOccurrences(ctx, pID, filter, lastPage, 100)
		if err != nil {
			t.Fatalf("ListOccurrences got %v want success", err)
		}
		if pageToken != "" {
			t.Errorf("Got %s want empty page token", pageToken)
		}
		if len(gotOccurrences) != 1 {
			t.Errorf("ListOccurrences got %v operations, want 1", len(gotOccurrences))
		}
	})

	t.Run("NoteOccurrencePagination", func(t *testing.T) {
		g, gp, cleanUp := createStore(t)
		defer cleanUp()

		ctx := context.Background()
		pID := "project"
		if _, err := gp.CreateProject(ctx, pID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}
		nPID := "noteproject"
		if _, err := gp.CreateProject(ctx, nPID, &prpb.Project{}); err != nil {
			t.Errorf("CreateProject got %v want success", err)
		}

		n := createTestNote(nPID)
		if _, err := g.CreateNote(ctx, nPID, testNoteID, "userID", n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}

		op1 := createTestOccurrence(pID, n.Name)
		if _, err := g.CreateOccurrence(ctx, pID, "userID", op1); err != nil {
			t.Errorf("CreateOccurrence got %v want success", err)
		}

		op2 := createTestOccurrence(pID, n.Name)
		if _, err := g.CreateOccurrence(ctx, pID, "userID", op2); err != nil {
			t.Errorf("CreateOccurrence got %v want success", err)
		}

		op3 := createTestOccurrence(pID, n.Name)
		if _, err := g.CreateOccurrence(ctx, pID, "userID", op3); err != nil {
			t.Errorf("CreateOccurrence got %v want success", err)
		}
		filter := "filters_are_yet_to_be_implemented"
		_, nID, err := name.ParseNote(n.Name)
		// Get occurrences
		gotOccurrences, lastPage, err := g.ListNoteOccurrences(ctx, nPID, nID, filter, "", 2)
		if err != nil {
			t.Fatalf("ListNoteOccurrences got %v want success", err)
		}
		if len(gotOccurrences) != 2 {
			t.Errorf("ListNoteOccurrences got %v occurrences, want 2", len(gotOccurrences))
		}
		// Get occurrences again
		gotOccurrences, pageToken, err := g.ListNoteOccurrences(ctx, nPID, nID, filter, lastPage, 100)
		if err != nil {
			t.Fatalf("ListNoteOccurrences got %v want success", err)
		}
		if pageToken != "" {
			t.Fatalf("Got %s want empty page token", pageToken)
		}
		if len(gotOccurrences) != 1 {
			t.Errorf("ListNoteOccurrences got %v operations, want 1", len(gotOccurrences))
		}
	})
}

func createTestOccurrence(pID, noteName string) *pb.Occurrence {
	return &pb.Occurrence{
		Name:     fmt.Sprintf("projects/%s/occurrences/134", pID),
		Resource: &pb.Resource{Uri: "gcr.io/foo/bar"},
		NoteName: noteName,
		Kind:     cpb.NoteKind_VULNERABILITY,
		Details: &pb.Occurrence_Vulnerability{
			Vulnerability: &vpb.Details{
				Severity:  vpb.Severity_HIGH,
				CvssScore: 7.5,
				PackageIssue: []*vpb.PackageIssue{
					{
						SeverityName: "HIGH",
						AffectedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:8",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "52.1",
								Revision: "8+deb8u3",
							},
						},
						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:8",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "52.1",
								Revision: "8+deb8u4",
							},
						},
					},
				},
			},
		},
	}
}

func createTestNote(pID string) *pb.Note {
	return &pb.Note{
		Name:             fmt.Sprintf("projects/%s/notes/%s", pID, testNoteID),
		ShortDescription: "CVE-2014-9911",
		LongDescription:  "NIST vectors: AV:N/AC:L/Au:N/C:P/I:P",
		Kind:             cpb.NoteKind_VULNERABILITY,
		Type: &pb.Note_Vulnerability{
			Vulnerability: &vpb.Vulnerability{
				CvssScore: 7.5,
				Severity:  vpb.Severity_HIGH,
				Details: []*vpb.Vulnerability_Detail{
					{
						CpeUri:  "cpe:/o:debian:debian_linux:7",
						Package: "icu",
						Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
							"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
							"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
						MinAffectedVersion: &pkgpb.Version{
							Kind: pkgpb.Version_MINIMUM,
						},
						SeverityName: "HIGH",

						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:7",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "4.8.1.1",
								Revision: "12+deb7u6",
							},
						},
					},
					{
						CpeUri:  "cpe:/o:debian:debian_linux:8",
						Package: "icu",
						Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
							"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
							"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
						MinAffectedVersion: &pkgpb.Version{
							Kind: pkgpb.Version_MINIMUM,
						},
						SeverityName: "HIGH",

						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:8",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "52.1",
								Revision: "8+deb8u4",
							},
						},
					},
					{
						CpeUri:  "cpe:/o:debian:debian_linux:9",
						Package: "icu",
						Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
							"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
							"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
						MinAffectedVersion: &pkgpb.Version{
							Kind: pkgpb.Version_MINIMUM,
						},
						SeverityName: "HIGH",

						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:debian:debian_linux:9",
							Package: "icu",
							Version: &pkgpb.Version{
								Name:     "55.1",
								Revision: "3",
							},
						},
					},
					{
						CpeUri:  "cpe:/o:canonical:ubuntu_linux:14.04",
						Package: "android",
						Description: "Stack-based buffer overflow in the ures_getByKeyWithFallback function in " +
							"common/uresbund.cpp in International Components for Unicode (ICU) before 54.1 for C/C++ allows " +
							"remote attackers to cause a denial of service or possibly have unspecified other impact via a crafted uloc_getDisplayName call.",
						MinAffectedVersion: &pkgpb.Version{
							Kind: pkgpb.Version_MINIMUM,
						},
						SeverityName: "MEDIUM",

						FixedLocation: &vpb.VulnerabilityLocation{
							CpeUri:  "cpe:/o:canonical:ubuntu_linux:14.04",
							Package: "android",
							Version: &pkgpb.Version{
								Kind: pkgpb.Version_MINIMUM,
							},
						},
					},
				},
			},
		},
		RelatedUrl: []*cpb.RelatedUrl{
			{
				Url:   "https://security-tracker.debian.org/tracker/CVE-2014-9911",
				Label: "More Info",
			},
			{
				Url:   "http://people.ubuntu.com/~ubuntu-security/cve/CVE-2014-9911",
				Label: "More Info",
			},
		},
	}
}
