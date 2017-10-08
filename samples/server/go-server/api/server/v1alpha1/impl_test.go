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

package v1alpha1

import (
	"github.com/grafeas/samples/server/go-server/api"
	"github.com/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/samples/server/go-server/api/server/storage"
	"github.com/grafeas/samples/server/go-server/api/server/testing"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateOperation(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	op := swagger.Operation{}
	if err := g.CreateOperation(&op); err == nil {
		t.Error("CreateOperation(empty operation): got success, want error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOperation(empty operation): got %v, want BadRequest(400)", err.StatusCode)
	}
	op = testutil.Operation()
	if err := g.CreateOperation(&op); err != nil {
		t.Errorf("CreateOperation(%v) got %v, want success", op, err)
	}
}

func TestCreateOccurrence(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	o := swagger.Occurrence{}
	if err := g.CreateOccurrence(&o); err == nil {
		t.Error("CreateOccurrence(empty occ): got success, want error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOccurrence(empty occ): got %v, want BadRequest(400)", err.StatusCode)
	}
	o = testutil.Occurrence(n.Name)
	if err := g.CreateOccurrence(&o); err != nil {
		t.Errorf("CreateOccurrence(%v) got %v, want success", n, err)
	}

	// Try to insert an occurrence for a note that does not exist.
	o.Name = "projects/testproject/occurrences/nonote"
	o.NoteName = "projects/scan-provider/notes/notthere"
	if err := g.CreateOccurrence(&o); err == nil {
		t.Errorf("CreateOccurrence got success, want Error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOccurrence got code %v want %v", err.StatusCode, http.StatusBadRequest)
	}
}

func TestCreateNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := swagger.Note{}
	if err := g.CreateNote(&n); err == nil {
		t.Error("CreateNote(empty note): got success, want error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateNote(empty note): got %v, want %v", err.StatusCode, http.StatusBadRequest)
	}
	n = testutil.Note()
	if err := g.CreateNote(&n); err != nil {
		t.Errorf("CreateNote(%v) got %v, want success", n, err)
	}
}

func TestDeleteNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	pID, nID, err := name.ParseNote(n.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	if err := g.DeleteNote(pID, nID); err == nil {
		t.Error("DeleteNote that doesn't exist got success, want err")
	}
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if err := g.DeleteNote(pID, nID); err != nil {
		t.Errorf("DeleteNote  got %v, want success", err)
	}
}

func TestDeleteOccurrence(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	// CreateNote so we can create an occurrence
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	o := testutil.Occurrence(n.Name)
	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence name %v", err)
	}
	if err := g.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence(%v) got %v, want success", n, err)
	}
	if err := g.DeleteOccurrence(pID, oID); err != nil {
		t.Errorf("DeleteOccurrence  got %v, want success", err)
	}
}

func TestDeleteOperation(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	o := testutil.Operation()
	pID, oID, err := name.ParseOperation(o.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	if err := g.DeleteOperation(pID, oID); err == nil {
		t.Error("DeleteOperation that doesn't exist got success, want err")
	}
	if err := g.CreateOperation(&o); err != nil {
		t.Fatalf("CreateOperation(%v) got %v, want success", o, err)
	}
	if err := g.DeleteOperation(pID, oID); err != nil {
		t.Errorf("DeleteOperation  got %v, want success", err)
	}
}

func TestGetNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	pID, nID, err := name.ParseNote(n.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	if _, err := g.GetNote(pID, nID); err == nil {
		t.Error("GetNote that doesn't exist got success, want err")
	}
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if got, err := g.GetNote(pID, nID); err != nil {
		t.Fatalf("GetNote(%v) got %v, want success", n, err)
	} else if n.Name != got.Name || !reflect.DeepEqual(n.VulnerabilityType, got.VulnerabilityType) {
		t.Errorf("GetNote got %v, want %v", *got, n)
	}
}

func TestGetOccurrence(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()

	o := testutil.Occurrence(n.Name)

	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence name %v", err)
	}
	if _, err := g.GetOccurrence(pID, oID); err == nil {
		t.Error("GetOccurrence that doesn't exist got success, want err")
	}
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if err := g.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence(%v) got %v, want success", n, err)
	}
	if got, err := g.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence(%v) got %v, want success", o, err)
	} else if o.Name != got.Name || !reflect.DeepEqual(o.VulnerabilityDetails, got.VulnerabilityDetails) {
		t.Errorf("GetOccurrence got %v, want %v", *got, o)
	}
}

func TestGetOperation(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	o := testutil.Operation()
	pID, oID, err := name.ParseOperation(o.Name)
	if err != nil {
		t.Fatalf("Error parsing note name %v", err)
	}
	if _, err := g.GetOperation(pID, oID); err == nil {
		t.Error("GetOperation that doesn't exist got success, want err")
	}
	if err := g.CreateOperation(&o); err != nil {
		t.Fatalf("CreateOperation(%v) got %v, want success", o, err)
	}
	if got, err := g.GetOperation(pID, oID); err != nil {
		t.Fatalf("GetOperation(%v) got %v, want success", o, err)
	} else if o.Name != got.Name || !reflect.DeepEqual(*got, o) {
		t.Errorf("GetOperation got %v, want %v", *got, o)
	}
}

func TestGetOccurrenceNote(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()

	o := testutil.Occurrence(n.Name)

	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence name %v", err)
	}
	if _, err := g.GetOccurrenceNote(pID, oID); err == nil {
		t.Error("GetOccurrenceNote that doesn't exist got success, want err")
	}
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote(%v) got %v, want success", n, err)
	}
	if err := g.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence(%v) got %v, want success", n, err)
	}
	if got, err := g.GetOccurrenceNote(pID, oID); err != nil {
		t.Fatalf("GetOccurrenceNote(%v) got %v, want success", n, err)
	} else if n.Name != got.Name || !reflect.DeepEqual(n.VulnerabilityType, got.VulnerabilityType) {
		t.Errorf("GetOccurrenceNote got %v, want %v", *got, n)
	}
}

func TestListOccurrences(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	n := testutil.Note()
	if err := g.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote got %v want success", err)
	}
	os := []swagger.Occurrence{}
	findProject := "findThese"
	dontFind := "dontFind"
	for i := 0; i < 20; i++ {
		o := testutil.Occurrence(n.Name)
		if i < 5 {
			o.Name = name.FormatOccurrence(findProject, string(i))
		} else {
			o.Name = name.FormatOccurrence(dontFind, string(i))
		}
		if err := g.CreateOccurrence(&o); err != nil {
			t.Fatalf("CreateOccurrence got %v want success", err)
		}
		os = append(os, o)
	}

	resp, err := g.ListOccurrences(findProject, "")
	if err != nil {
		t.Fatalf("ListOccurrences got %v want success", err)
	}
	if len(resp.Occurrences) != 5 {
		t.Errorf("resp.Occurrences got %d, want 5", len(resp.Occurrences))
	}
}

func TestListOperations(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	os := []swagger.Operation{}
	findProject := "findThese"
	dontFind := "dontFind"
	for i := 0; i < 20; i++ {
		o := testutil.Operation()
		if i < 5 {
			o.Name = name.FormatOperation(findProject, string(i))
		} else {
			o.Name = name.FormatOperation(dontFind, string(i))
		}
		if err := g.CreateOperation(&o); err != nil {
			t.Fatalf("CreateOperation got %v want success", err)
		}
		os = append(os, o)
	}

	resp, err := g.ListOperations(findProject, "")
	if err != nil {
		t.Fatalf("ListOperations got %v want success", err)
	}
	if len(resp.Operations) != 5 {
		t.Errorf("resp.Operations got %d, want 5", len(resp.Operations))
	}
}

func TestListNotes(t *testing.T) {
	g := Grafeas{storage.NewMemStore()}
	ns := []swagger.Note{}
	findProject := "findThese"
	dontFind := "dontFind"
	for i := 0; i < 20; i++ {
		n := testutil.Note()
		if i < 5 {
			n.Name = name.FormatNote(findProject, string(i))
		} else {
			n.Name = name.FormatNote(dontFind, string(i))
		}
		if err := g.CreateNote(&n); err != nil {
			t.Fatalf("CreateNote got %v want success", err)
		}
		ns = append(ns, n)
	}

	resp, err := g.ListNotes(findProject, "")
	if err != nil {
		t.Fatalf("ListNotes got %v want success", err)
	}
	if len(resp.Notes) != 5 {
		t.Errorf("resp.Notes got %d, want 5", len(resp.Notes))
	}
}
