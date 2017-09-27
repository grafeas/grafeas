package storage

import (
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/name"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/testing"
	"net/http"
	"testing"
	"reflect"
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
	// Try to insert the same note twice, expect failure.
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
	} else if reflect.DeepEqual(got, o){
		t.Errorf("GetOccurrence got %v, want %v",got, o)
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
		t.Error("Deleting not-existant occurrence got success, want error")
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
		t.Fatalf("UpdateOccurrence got success want error", err)
	}
	if err := s.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence got %v want success", err)
	}
	if got, err := s.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence got %v, want success", err)
	} else if reflect.DeepEqual(got, o){
		t.Errorf("GetOccurrence got %v, want %v",got, o)
	}

	o2 := o
	o2.VulnerabilityDetails.CvssScore = 1.0
	if err := s.UpdateOccurrence(pID, oID, &o2); err != nil {
		t.Fatalf("UpdateOccurrence got %v want success", err)
	}

	if got, err := s.GetOccurrence(pID, oID); err != nil {
		t.Fatalf("GetOccurrence got %v, want success", err)
	} else if reflect.DeepEqual(got, o2){
		t.Errorf("GetOccurrence got %v, want %v",got, o2)
	}
}
