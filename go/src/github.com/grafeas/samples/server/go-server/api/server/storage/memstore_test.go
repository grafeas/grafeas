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
	"github.com/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/samples/server/go-server/api/server/testing"
	"net/http"
	"reflect"
	"testing"
)

func TestMemStore_CreateNote(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()
	if err := s.CreateNote(&n); err != nil {
		t.Errorf("CreateNote got %v want success", err)
	}
	// Try to insert the same note twice, expect failure.
	if err := s.CreateNote(&n); err == nil {
		t.Errorf("CreateNote got success, want Error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateNote got code %v want %v", err.StatusCode, http.StatusBadRequest)
	}
}

func TestMemStore_CreateOccurrence(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()
	if err := s.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote got %v want success", err)
	}
	o := testutil.Occurrence(n.Name)
	if err := s.CreateOccurrence(&o); err != nil {
		t.Errorf("CreateOccurrence got %v want success", err)
	}
	// Try to insert the same occurrence twice, expect failure.
	if err := s.CreateOccurrence(&o); err == nil {
		t.Errorf("CreateOccurrence got success, want Error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOccurrence got code %v want %v", err.StatusCode, http.StatusBadRequest)
	}
	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing projectID and occurrenceID %v", err)
	}
	if got, err := s.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence got %v, want success", err)
	} else if reflect.DeepEqual(got, o) {
		t.Errorf("GetOccurrence got %v, want %v", got, o)
	}
}

func TestMemStore_CreateOperation(t *testing.T) {
	s := NewMemStore()
	op := testutil.Operation()
	if err := s.CreateOperation(&op); err != nil {
		t.Errorf("CreateNote got %v want success", err)
	}
	// Try to insert the same note twice, expect failure.
	if err := s.CreateOperation(&op); err == nil {
		t.Errorf("CreateOperation got success, want Error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOperation got code %v want %v", err.StatusCode, http.StatusBadRequest)
	}
}

func TestMemStore_DeleteOccurrence(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()
	if err := s.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote got %v want success", err)
	}
	o := testutil.Occurrence(n.Name)
	// Delete before the occurrence exists
	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence %v", err)
	}
	if err := s.DeleteOccurrence(pID, oID); err == nil {
		t.Error("Deleting nonexistant occurrence got success, want error")
	}
	if err := s.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence got %v want success", err)
	}
	if err := s.DeleteOccurrence(pID, oID); err != nil {
		t.Errorf("DeleteOccurrence got %v, want success ", err)
	}
}

func TestMemStore_UpdateOccurrence(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()
	if err := s.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote got %v want success", err)
	}
	o := testutil.Occurrence(n.Name)
	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing projectID and occurrenceID %v", err)
	}
	if err := s.UpdateOccurrence(pID, oID, &o); err == nil {
		t.Fatal("UpdateOccurrence got success want error")
	}
	if err := s.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence got %v want success", err)
	}
	if got, err := s.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence got %v, want success", err)
	} else if reflect.DeepEqual(got, o) {
		t.Errorf("GetOccurrence got %v, want %v", got, o)
	}

	o2 := o
	o2.VulnerabilityDetails.CvssScore = 1.0
	if err := s.UpdateOccurrence(pID, oID, &o2); err != nil {
		t.Fatalf("UpdateOccurrence got %v want success", err)
	}

	if got, err := s.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence got %v, want success", err)
	} else if reflect.DeepEqual(got, o2) {
		t.Errorf("GetOccurrence got %v, want %v", got, o2)
	}
}

func TestMemStore_DeleteNote(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()
	// Delete before the note exists
	pID, oID, err := name.ParseNote(n.Name)
	if err != nil {
		t.Fatalf("Error parsing note %v", err)
	}
	if err := s.DeleteNote(pID, oID); err == nil {
		t.Error("Deleting nonexistant note got success, want error")
	}
	if err := s.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote got %v want success", err)
	}

	if err := s.DeleteNote(pID, oID); err != nil {
		t.Errorf("DeleteNote got %v, want success ", err)
	}
}

func TestMemStore_UpdateNote(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()

	pID, nID, err := name.ParseNote(n.Name)
	if err != nil {
		t.Fatalf("Error parsing projectID and noteID %v", err)
	}
	if err := s.UpdateNote(pID, nID, &n); err == nil {
		t.Fatal("UpdateNote got success want error")
	}
	if err := s.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote got %v want success", err)
	}
	if got, err := s.GetNote(pID, nID); err != nil {
		t.Fatalf("GetNote got %v, want success", err)
	} else if reflect.DeepEqual(got, n) {
		t.Errorf("GetNote got %v, want %v", got, n)
	}

	n2 := n
	n2.VulnerabilityType.CvssScore = 1.0
	if err := s.UpdateNote(pID, nID, &n2); err != nil {
		t.Fatalf("UpdateNote got %v want success", err)
	}

	if got, err := s.GetNote(pID, nID); err != nil {
		t.Fatalf("GetNote got %v, want success", err)
	} else if reflect.DeepEqual(got, n2) {
		t.Errorf("GetNote got %v, want %v", got, n2)
	}
}

func TestMemStore_GetOccurrence(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()
	if err := s.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote got %v want success", err)
	}
	o := testutil.Occurrence(n.Name)
	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence %v", err)
	}
	if _, err := s.GetOccurrence(pID, oID); err == nil {
		t.Fatal("GetOccurrence got success, want error")
	}
	if err := s.CreateOccurrence(&o); err != nil {
		t.Errorf("CreateOccurrence got %v, want Success", err)
	}
	if got, err := s.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence got %v, want success", err)
	} else if reflect.DeepEqual(got, o) {
		t.Errorf("GetOccurrence got %v, want %v", got, o)
	}
}

func TestMemStore_GetNote(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()

	pID, nID, err := name.ParseNote(n.Name)
	if err != nil {
		t.Fatalf("Error parsing note %v", err)
	}
	if _, err := s.GetNote(pID, nID); err == nil {
		t.Fatal("GetNote got success, want error")
	}
	if err := s.CreateNote(&n); err != nil {
		t.Errorf("CreateNote got %v, want Success", err)
	}
	if got, err := s.GetNote(pID, nID); err != nil {
		t.Fatalf("GetNote got %v, want success", err)
	} else if reflect.DeepEqual(got, n) {
		t.Errorf("GetNote got %v, want %v", got, n)
	}
}

func TestMemStore_GetNoteByOccurrence(t *testing.T) {
	s := NewMemStore()
	n := testutil.Note()
	if err := s.CreateNote(&n); err != nil {
		t.Fatalf("CreateNote got %v want success", err)
	}
	o := testutil.Occurrence(n.Name)
	pID, oID, err := name.ParseOccurrence(o.Name)
	if err != nil {
		t.Fatalf("Error parsing occurrence %v", err)
	}
	if _, err := s.GetNoteByOccurrence(pID, oID); err == nil {
		t.Fatal("GetNoteByOccurrence got success, want error")
	}
	if err := s.CreateOccurrence(&o); err != nil {
		t.Errorf("CreateOccurrence got %v, want Success", err)
	}
	if got, err := s.GetNoteByOccurrence(pID, oID); err != nil {
		t.Fatalf("GetNoteByOccurrence got %v, want success", err)
	} else if reflect.DeepEqual(got, n) {
		t.Errorf("GetNoteByOccurrence got %v, want %v", got, n)
	}
}

func TestMemStore_GetOperation(t *testing.T) {
	s := NewMemStore()
	o := testutil.Operation()

	pID, oID, err := name.ParseOperation(o.Name)
	if err != nil {
		t.Fatalf("Error parsing operation %v", err)
	}
	if _, err := s.GetOperation(pID, oID); err == nil {
		t.Fatal("GetOperation got success, want error")
	}
	if err := s.CreateOperation(&o); err != nil {
		t.Errorf("CreateOperation got %v, want Success", err)
	}
	if got, err := s.GetOperation(pID, oID); err != nil {
		t.Fatalf("GetOperation got %v, want success", err)
	} else if reflect.DeepEqual(got, o) {
		t.Errorf("GetOperation got %v, want %v", got, o)
	}
}

func TestMemStore_DeleteOperation(t *testing.T) {
	s := NewMemStore()
	o := testutil.Operation()
	// Delete before the operation exists
	pID, oID, err := name.ParseOperation(o.Name)
	if err != nil {
		t.Fatalf("Error parsing note %v", err)
	}
	if err := s.DeleteOperation(pID, oID); err == nil {
		t.Error("Deleting nonexistant operation got success, want error")
	}
	if err := s.CreateOperation(&o); err != nil {
		t.Fatalf("CreateOperation got %v want success", err)
	}

	if err := s.DeleteOperation(pID, oID); err != nil {
		t.Errorf("DeleteOperation got %v, want success ", err)
	}
}

func TestMemStore_UpdateOperation(t *testing.T) {
	s := NewMemStore()
	o := testutil.Operation()

	pID, oID, err := name.ParseOperation(o.Name)
	if err != nil {
		t.Fatalf("Error parsing projectID and operationID %v", err)
	}
	if err := s.UpdateOperation(pID, oID, &o); err == nil {
		t.Fatal("UpdateOperation got success want error")
	}
	if err := s.CreateOperation(&o); err != nil {
		t.Fatalf("CreateOperation got %v want success", err)
	}
	if got, err := s.GetOperation(pID, oID); err != nil {
		t.Fatalf("GetOperation got %v, want success", err)
	} else if reflect.DeepEqual(got, o) {
		t.Errorf("GetOperation got %v, want %v", got, o)
	}

	o2 := o
	o2.Done = true
	if err := s.UpdateOperation(pID, oID, &o2); err != nil {
		t.Fatalf("UpdateOperation got %v want success", err)
	}

	if got, err := s.GetOperation(pID, oID); err != nil {
		t.Fatalf("GetOperation got %v, want success", err)
	} else if reflect.DeepEqual(got, o2) {
		t.Errorf("GetOperation got %v, want %v", got, o2)
	}
}
