package storage

import (
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/testing"
	"net/http"
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
	// Try to insert the same note twice, expect failure.
	if err := s.CreateOccurrence(&o); err == nil {
		t.Errorf("CreateOccurrence got success, want Error")
	} else if err.StatusCode != http.StatusBadRequest {
		t.Errorf("CreateOccurrence got code %v want %v", err.StatusCode, http.StatusBadRequest)
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

	if s.DeleteOccurrence(
	if err := s.CreateOccurrence(&o); err != nil {
		t.Fatalf("CreateOccurrence got %v want success", err)
	}
}
