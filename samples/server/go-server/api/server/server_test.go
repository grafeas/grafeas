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
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"

	"errors"
	"fmt"
	"github.com/grafeas/grafeas/samples/server/go-server/api"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/storage"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/testing"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/v1alpha1"
)

func TestCreateNote(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	// Test that note.Name must match values in path.
	badN := testutil.Note()
	badN.Name = "/projects/foo/notes/wrong"
	if err := createNote(n, h); err == nil {
		t.Error("CreateNote with mismatched url/name got success, want error")
	}

}

func TestCreateOccurrence(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Fatalf("Error creating note: %v", err)
	}
	o := testutil.Occurrence(n.Name)
	if _, err := createOccurrence(o, h); err != nil {
		t.Errorf("%v", err)
	}
}

func TestDeleteOccurrence(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	pID := "project"
	oID := "occurrence"
	r, err := http.NewRequest("DELETE", fmt.Sprintf("/v1alpha1/projects/%v/occurrences/%v", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.DeleteOccurrence(w, r)
	if w.Code != 400 {
		t.Errorf("DeleteOccurrence with no occurrence got %v, want 400", w.Code)
	}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	o := testutil.Occurrence(n.Name)
	got, err := createOccurrence(o, h)
	if err != nil {
		t.Fatalf("%v", err)
	}
	pID, oID, aErr := name.ParseOccurrence(got.Name)
	if aErr != nil {
		t.Fatalf("Error parsing note name: %v", aErr)
	}
	r, err = http.NewRequest("DELETE", fmt.Sprintf("/v1alpha1/projects/%v/occurrences/%v", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.DeleteOccurrence(w, r)
	if w.Code != 200 {
		t.Errorf("DeleteOccurrence got %v; %v, want 200", w.Code, w.Body)
	}

}

func TestCreateOperation(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	o := testutil.Operation()
	if _, err := createOperation(o, h, ""); err != nil {
		t.Errorf("%v", err)
	}
	// Make sure we can specify operationId
	if _, err := createOperation(o, h, "testID"); err != nil {
		t.Errorf("%v", err)
	}
}

func TestDeleteOperation(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	pID := "vulnerability-scanner-a"
	oID := "operation"
	r, err := http.NewRequest("DELETE", fmt.Sprintf("/v1alpha1/projects/%v/operations/%v", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.DeleteOperation(w, r)
	if w.Code != 400 {
		t.Errorf("DeleteOperation with no note got %v, want 400", w.Code)
	}
	o := testutil.Operation()
	if _, err := createOperation(o, h, oID); err != nil {
		t.Errorf("%v", err)
	}

	r, err = http.NewRequest("DELETE", fmt.Sprintf("/v1alpha1/projects/%v/operations/%v", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.DeleteOperation(w, r)
	if w.Code != 200 {
		t.Errorf("DeleteOperation got %v; %v, want 200", w.Code, w.Body)
	}
}

func TestDeleteNote(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	pID := "project"
	nID := "note"
	r, err := http.NewRequest("DELETE", fmt.Sprintf("/v1alpha1/projects/%v/notes/%v", pID, nID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.DeleteNote(w, r)
	if w.Code != 400 {
		t.Errorf("DeleteNote with no note got %v, want 400", w.Code)
	}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	pID, nID, aErr := name.ParseNote(n.Name)
	if aErr != nil {
		t.Fatalf("Error parsing note name: %v", aErr)
	}
	r, err = http.NewRequest("DELETE", fmt.Sprintf("/v1alpha1/projects/%v/notes/%v", pID, nID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.DeleteNote(w, r)
	if w.Code != 200 {
		t.Errorf("DeleteNote got %v; %v, want 200", w.Code, w.Body)
	}
}

func TestGetNote(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	pID := "project"
	nID := "note"
	r, err := http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/notes/%v", pID, nID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.GetNote(w, r)
	if w.Code != 400 {
		t.Errorf("GetNote with no note got %v, want 400", w.Code)
	}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	pID, nID, aErr := name.ParseNote(n.Name)
	if aErr != nil {
		t.Fatalf("Error parsing note name: %v", aErr)
	}
	r, err = http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/notes/%v", pID, nID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.GetNote(w, r)
	if w.Code != 200 {
		t.Errorf("GetNote  got %v, want 200", w.Code)
	}
}

func TestGetOperation(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	pID := "project"
	oID := "operation"
	r, err := http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/operations/%v", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.GetOperation(w, r)
	if w.Code != 400 {
		t.Errorf("GetOperation with no operation got %v, want 400", w.Code)
	}
	o := testutil.Operation()
	got, err := createOperation(o, h, "")
	if err != nil {
		t.Errorf("%v", err)
	}
	pID, nID, aErr := name.ParseOperation(got.Name)
	if aErr != nil {
		t.Fatalf("Error parsing note name: %v", aErr)
	}
	r, err = http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/operations/%v", pID, nID), nil)
	w = httptest.NewRecorder()
	h.GetOperation(w, r)
	if w.Code != 200 {
		t.Errorf("GetOperation got %v, want 200", w.Code)
	}
}

func TestGetOccurrence(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	pID := "project"
	oID := "occurrence"
	r, err := http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/occurrences/%v", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.GetOccurrence(w, r)
	if w.Code != 400 {
		t.Errorf("GetOccurrence with no occurrence got %v, want 400", w.Code)
	}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	o := testutil.Occurrence(n.Name)
	got, err := createOccurrence(o, h)
	if err != nil {
		t.Errorf("%v", err)
	}

	pID, oID, aErr := name.ParseOccurrence(got.Name)
	if aErr != nil {
		t.Errorf("Error parsing created occurrence name: %v", aErr)
	}
	r, err = http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/occurrences/%v", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.GetOccurrence(w, r)
	if w.Code != 200 {
		t.Errorf("GetOccurrence got %v, want 200", w.Code)
	}

}

func createOccurrence(o swagger.Occurrence, g Handler) (*swagger.Occurrence, error) {
	pID := "test-project"
	rawOcc, err := json.Marshal(&o)
	reader := bytes.NewReader(rawOcc)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		fmt.Sprintf("/v1alpha1/projects/%v/occurrences", pID), reader)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateOccurrence(w, r)
	if w.Code != 200 {
		return nil, errors.New(fmt.Sprintf("CreateOccurrence(%v) got %v want 200", o, w.Code))
	}
	got := swagger.Occurrence{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if got.Name == "" {
		return nil, errors.New("got.Name got empty, want name")
	} else {
		if gotID, _, err := name.ParseOccurrence(got.Name); err != nil {
			return nil, fmt.Errorf("Error parsing created occurrence name: %v", err)
		} else if gotID != pID {
			return nil, fmt.Errorf("Created Occurrence projectID: got %v, want %v", gotID, pID)
		}
	}
	return &got, nil
}

func createNote(n swagger.Note, g Handler) error {
	rawNote, err := json.Marshal(&n)
	reader := bytes.NewReader(rawNote)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	pID, nID, aErr := name.ParseNote(n.Name)
	if aErr != nil {
		return errors.New(fmt.Sprintf("error parsing name %v", err))
	}
	r, err := http.NewRequest("POST",
		fmt.Sprintf("/v1alpha1/projects/%v/notes?noteId=%v", pID, nID), reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v with path %v", err, r.URL))
	}

	w := httptest.NewRecorder()
	g.CreateNote(w, r)

	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateNote(%v) got %v want 200", r, w.Code))
	}
	return nil
}

func createOperation(o swagger.Operation, g Handler, oID string) (*swagger.Operation, error) {
	rawOp, err := json.Marshal(&o)
	reader := bytes.NewReader(rawOp)
	pID := "vulnerability-scanner-a"
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	url := fmt.Sprintf("/v1alpha1/projects/%v/operations", pID)
	if oID != "" {
		url = fmt.Sprintf("%v?operationId=%v", url, oID)
	}
	r, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateOperation(w, r)
	if w.Code != 200 {
		return nil, errors.New(fmt.Sprintf("CreateOperation(%v) got %v want 200 with error %v", o, w.Code, w.Body))
	}
	got := swagger.Operation{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if got.Name == "" {
		return nil, errors.New("got.Name got empty, want name")
	} else {
		if gotPID, gotOpID, err := name.ParseOperation(got.Name); err != nil {
			return nil, fmt.Errorf("Error parsing created operation name: %v", err)
		} else if gotPID != pID || gotOpID == "" {
			return nil, fmt.Errorf("Created Occurrence projectID: got projectID %v opID %v, want projectID %v, opID not empty",
				gotPID, gotOpID, pID)
		}
	}
	return &got, nil
}

func TestGetOccurrenceNote(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	pID := "project"
	oID := "note"
	r, err := http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/occurrences/%v/notes", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.GetOccurrenceNote(w, r)
	if w.Code != 400 {
		t.Errorf("GetOccurrenceNote with no occurrence got %v, want 400", w.Code)
	}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Fatalf("Error creating note: %v", err)
	}
	o := testutil.Occurrence(n.Name)
	occ, err := createOccurrence(o, h)
	if err != nil {
		t.Errorf("%v", err)
	}
	pID, oID, e := name.ParseOccurrence(occ.Name)
	if e != nil {
		t.Fatalf("Error parsing occurrences %v", e)
	}
	w = httptest.NewRecorder()
	r, err = http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/occurrences/%v/notes", pID, oID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	h.GetOccurrenceNote(w, r)
	if w.Code != 200 {
		t.Errorf("GetOccurrenceNote with no occurrence got %v, want 200", w.Code)
	}
}

func TestListNotes(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	n := testutil.Note()
	pID, _, aErr := name.ParseNote(n.Name)
	r, err := http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/notes/", pID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.ListNotes(w, r)
	if w.Code != 200 {
		t.Errorf("ListNotes with no notes got %v, want 200", w.Code)
	}
	got := swagger.ListNotesResponse{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if len(got.Notes) != 0 {
		t.Errorf("ListNotes got %d, want 0", len(got.Notes))
	}

	if aErr != nil {
		t.Fatalf("Error parsing note name: %v", aErr)
	}
	for i := 0; i < 20; i++ {
		n := testutil.Note()
		n.Name = fmt.Sprintf("%v-%d", n.Name, i)
		if err := createNote(n, h); err != nil {
			t.Errorf("%v", err)
		}
	}

	r, err = http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/notes/", pID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.ListNotes(w, r)
	if w.Code != 200 {
		t.Errorf("ListNotes got %v, want 200", w.Code)
	}
	got = swagger.ListNotesResponse{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if len(got.Notes) != 20 {
		t.Errorf("ListNotes got %d, want 20", len(got.Notes))
	}
}

func TestListOccurrences(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	n := testutil.Note()
	o := testutil.Occurrence(n.Name)
	pID, _, aErr := name.ParseOccurrence(o.Name)
	if aErr != nil {
		t.Fatalf("Error parsing occurrence name: %v", aErr)
	}
	r, err := http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/occurrences/", pID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.ListOccurrences(w, r)
	if w.Code != 200 {
		t.Errorf("ListOccurrences no occurrences got %v, want 200", w.Code)
	}
	got := swagger.ListOccurrencesResponse{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if len(got.Occurrences) != 0 {
		t.Errorf("ListOccurrences got %d, want 0", len(got.Occurrences))
	}
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	for i := 0; i < 20; i++ {
		o := testutil.Occurrence(n.Name)
		if _, err := createOccurrence(o, h); err != nil {
			t.Errorf("%v", err)
		}
	}

	r, err = http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/occurrences/", pID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.ListOccurrences(w, r)
	if w.Code != 200 {
		t.Errorf("ListOccurrences got %v, want 200", w.Code)
	}
	got = swagger.ListOccurrencesResponse{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if len(got.Occurrences) != 20 {
		t.Errorf("ListOccurrences got %d, want 20", len(got.Occurrences))
	}
}

func TestListOperations(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	o := testutil.Operation()
	pID, _, aErr := name.ParseOperation(o.Name)
	r, err := http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/operations/", pID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.ListOperations(w, r)
	if w.Code != 200 {
		t.Errorf("ListOperations got %v, want 200", w.Code)
	}
	got := swagger.ListOperationsResponse{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if len(got.Operations) != 0 {
		t.Errorf("ListOccurrences got %d, want 0", len(got.Operations))
	}
	if aErr != nil {
		t.Fatalf("Error parsing operation name: %v", aErr)
	}
	for i := 0; i < 20; i++ {
		o := testutil.Operation()
		if _, err := createOperation(o, h, ""); err != nil {
			t.Errorf("%v", err)
		}
	}

	r, err = http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/operations/", pID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.ListOperations(w, r)
	if w.Code != 200 {
		t.Errorf("ListOperations got %v, want 200", w.Code)
	}
	got = swagger.ListOperationsResponse{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if len(got.Operations) != 20 {
		t.Errorf("ListOperations got %d, want 20", len(got.Operations))
	}
}

func TestListNoteOccurrences(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	n := testutil.Note()
	pID, nID, aErr := name.ParseNote(n.Name)
	if aErr != nil {
		t.Fatalf("Error parsing note name: %v", aErr)
	}
	r, err := http.NewRequest("GET", fmt.Sprintf("/v1alpha1/projects/%v/notes/%v/occurrences", pID, nID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.ListNoteOccurrences(w, r)
	if w.Code != 200 {
		t.Errorf("ListNoteOccurrences no occurrences got %v, want 200", w.Code)
	}
	got := swagger.ListNoteOccurrencesResponse{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if len(got.Occurrences) != 0 {
		t.Errorf("ListNoteOccurrences got %d, want 0", len(got.Occurrences))
	}
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	for i := 0; i < 20; i++ {
		o := testutil.Occurrence(n.Name)
		if _, err := createOccurrence(o, h); err != nil {
			t.Errorf("%v", err)
		}
	}

	r, err = http.NewRequest("GET",
		fmt.Sprintf("/v1alpha1/projects/%v/notes/%v/occurrences", pID, nID), nil)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.ListNoteOccurrences(w, r)
	if w.Code != 200 {
		t.Fatalf("ListNoteOccurrences got %v, want 200", w.Code)
	}
	got = swagger.ListNoteOccurrencesResponse{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if len(got.Occurrences) != 20 {
		t.Errorf("ListNoteOccurrences got %d, want 20", len(got.Occurrences))
	}
}

func TestUpdateNote(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	// update note Name, verify 400 error
	update := testutil.Note()
	update.Name = "projects/p/notes/thisisbad"
	rawNote, err := json.Marshal(&update)
	reader := bytes.NewReader(rawNote)
	if err != nil {
		t.Errorf(fmt.Sprintf("error marshalling json: %v", err))
	}
	pID, nID, pErr := name.ParseNote(n.Name)
	if pErr != nil {
		t.Fatalf("Error parsing Note; %v", pErr)
	}
	r, err := http.NewRequest("PUT",
		fmt.Sprintf("/v1alpha1/projects/%v/notes/%v", pID, nID), reader)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()

	h.UpdateNote(w, r)
	if w.Code != 400 {
		t.Errorf("UpdateNote with no new name got %v, want 400", w.Code)
	}
	// update note description, verify 200
	update = testutil.Note()
	update.LongDescription = "This note needs a new long description"
	rawNote, err = json.Marshal(&update)
	reader = bytes.NewReader(rawNote)
	if err != nil {
		t.Errorf(fmt.Sprintf("error marshalling json: %v", err))
	}

	r, err = http.NewRequest("PUT",
		fmt.Sprintf("/v1alpha1/projects/%v/notes/%v", pID, nID), reader)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.UpdateNote(w, r)
	if w.Code != 200 {
		t.Errorf("UpdateNote got %v, want 200", w.Code)
	}
}

func TestUpdateOperation(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	o := testutil.Operation()
	created, err := createOperation(o, h, "")
	if err != nil {
		t.Errorf("%v", err)
	}
	// update operation Name, verify 400 error
	update := testutil.Operation()
	update.Name = "projects/p/operations/thisisbad"
	rawOperation, err := json.Marshal(&update)
	reader := bytes.NewReader(rawOperation)
	if err != nil {
		t.Errorf(fmt.Sprintf("error marshalling json: %v", err))
	}
	pID, oID, pErr := name.ParseOperation(created.Name)
	if pErr != nil {
		t.Fatalf("Error parsing operation; %v", pErr)
	}
	r, err := http.NewRequest("PUT",
		fmt.Sprintf("/v1alpha1/projects/%v/operations/%v", pID, oID), reader)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.UpdateOperation(w, r)
	if w.Code != 400 {
		t.Errorf("UpdateOperation with no new name got %v, want 400", w.Code)
	}
	// update operation done, verify 200
	update = testutil.Operation()
	update.Done = true
	update.Name = created.Name
	rawOperation, err = json.Marshal(&update)
	reader = bytes.NewReader(rawOperation)
	if err != nil {
		t.Errorf(fmt.Sprintf("error marshalling json: %v", err))
	}

	r, err = http.NewRequest("PUT",
		fmt.Sprintf("/v1alpha1/projects/%v/operations/%v", pID, oID), reader)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()
	h.UpdateOperation(w, r)
	if w.Code != 200 {
		t.Errorf("UpdateOperation got %v, want 200", w.Code)
	}
}

func TestUpdateOccurrence(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Errorf("%v", err)
	}
	o := testutil.Occurrence(n.Name)
	created, err := createOccurrence(o, h)
	if err != nil {
		t.Errorf("%v", err)
	}
	// update name, verify 400 error
	update := testutil.Occurrence(n.Name)
	update.Name = "projects/p/occurrences/thisisbad"
	rawOccurrence, err := json.Marshal(&update)
	reader := bytes.NewReader(rawOccurrence)
	if err != nil {
		t.Errorf(fmt.Sprintf("error marshalling json: %v", err))
	}
	pID, oID, pErr := name.ParseOccurrence(created.Name)
	if pErr != nil {
		t.Fatalf("Error parsing Occurrence; %v", pErr)
	}
	r, err := http.NewRequest("PUT",
		fmt.Sprintf("/v1alpha1/projects/%v/occurrences/%v", pID, oID), reader)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w := httptest.NewRecorder()
	h.UpdateOccurrence(w, r)
	if w.Code != 400 {
		t.Errorf("UpdateOccurrence with no new name got %v, want 400", w.Code)
	}
	// update a different field, verify 200
	update = testutil.Occurrence(n.Name)
	update.Remediation = "updgrade to latest"
	update.Name = created.Name
	rawOccurrence, err = json.Marshal(&update)
	reader = bytes.NewReader(rawOccurrence)
	if err != nil {
		t.Errorf(fmt.Sprintf("error marshalling json: %v", err))
	}

	r, err = http.NewRequest("PUT",
		fmt.Sprintf("/v1alpha1/projects/%v/occurrences/%v", pID, oID), reader)
	if err != nil {
		t.Fatalf("Could not create httprequest %v", err)
	}
	w = httptest.NewRecorder()

	h.UpdateOccurrence(w, r)
	if w.Code != 200 {
		t.Errorf("UpdateOccurrence got %v, want 200", w.Code)
	}
}
