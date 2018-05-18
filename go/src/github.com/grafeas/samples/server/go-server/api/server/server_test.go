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
	"github.com/grafeas/samples/server/go-server/api"
	"github.com/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/samples/server/go-server/api/server/storage"
	"github.com/grafeas/samples/server/go-server/api/server/testing"
	"github.com/grafeas/samples/server/go-server/api/server/v1alpha1"
)

func TestHandler_CreateNote(t *testing.T) {
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

func TestHandler_CreateOccurrence(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	n := testutil.Note()
	if err := createNote(n, h); err != nil {
		t.Fatalf("Error creating note: %v", err)
	}
	o := testutil.Occurrence(n.Name)
	if err := createOccurrence(o, h); err != nil {
		t.Errorf("%v", err)
	}
}

func TestHandler_CreateOperation(t *testing.T) {
	h := Handler{v1alpha1.Grafeas{S: storage.NewMemStore()}}
	o := testutil.Operation()
	if err := createOperation(o, h, ""); err != nil {
		t.Errorf("%v", err)
	}
	// Make sure we can specify operationId
	if err := createOperation(o, h, "testID"); err != nil {
		t.Errorf("%v", err)
	}
}

func createOccurrence(o swagger.Occurrence, g Handler) error {
	pID := "test-project"
	rawOcc, err := json.Marshal(&o)
	reader := bytes.NewReader(rawOcc)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		fmt.Sprintf("/v1alpha1/projects/%v/occurrences", pID), reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateOccurrence(w, r)
	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateOccurrence(%v) got %v want 200", o, w.Code))
	}
	got := swagger.Occurrence{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if got.Name == "" {
		return errors.New("got.Name got empty, want name")
	} else {
		if gotID, _, err := name.ParseOccurrence(got.Name); err != nil {
			return fmt.Errorf("Error parsing created occurrence name: %v", err)
		} else if gotID != pID {
			return fmt.Errorf("Created Occurrence projectID: got %v, want %v", gotID, pID)
		}
	}
	return nil
}

func createNote(n swagger.Note, g Handler) error {
	rawNote, err := json.Marshal(&n)
	reader := bytes.NewReader(rawNote)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	r, err := http.NewRequest("POST",
		"/v1alpha1/projects/vulnerability-scanner-a/notes?noteId=CVE-1999-0710", reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}

	w := httptest.NewRecorder()
	g.CreateNote(w, r)

	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateNote(%v) got %v want 200", n, w.Code))
	}
	return nil
}

func createOperation(o swagger.Operation, g Handler, oID string) error {
	rawOp, err := json.Marshal(&o)
	reader := bytes.NewReader(rawOp)
	pID := "vulnerability-scanner-a"
	if err != nil {
		return errors.New(fmt.Sprintf("error marshalling json: %v", err))
	}
	url := fmt.Sprintf("/v1alpha1/projects/%v/operations", pID)
	if oID != "" {
		url = fmt.Sprintf("%v?operationId=%v", url, oID)
	}
	r, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return errors.New(fmt.Sprintf("error creating http request %v", err))
	}
	w := httptest.NewRecorder()
	g.CreateOperation(w, r)
	if w.Code != 200 {
		return errors.New(fmt.Sprintf("CreateOperation(%v) got %v want 200 with error %v", o, w.Code, w.Body))
	}
	got := swagger.Operation{}
	json.Unmarshal(w.Body.Bytes(), &got)
	if got.Name == "" {
		return errors.New("got.Name got empty, want name")
	} else {
		if gotPID, gotOpID, err := name.ParseOperation(got.Name); err != nil {
			return fmt.Errorf("Error parsing created operation name: %v", err)
		} else if gotPID != pID || gotOpID == "" {
			return fmt.Errorf("Created Occurrence projectID: got projectID %v opID %v, want projectID %v, opID not empty",
				gotPID, gotOpID, pID)
		}
	}
	return nil
}
