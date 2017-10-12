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

// package server is the implementation of a server that handles grafeas requests.
package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/grafeas/grafeas/samples/server/go-server/api"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/v1alpha1"
	"github.com/grafeas/grafeas/server-go/errors"
)

// Handler accepts httpRequests, converts them to Grafeas objects - calls into Grafeas to operation on them
// and converts responses to http responses.
type Handler struct {
	g v1alpha1.Grafeas
}

// CreateNote handles http requests to create notes in grafeas
func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	nIDs, ok := r.URL.Query()["noteId"]
	if !ok {
		log.Print("noteId is not specified")
		http.Error(w, "noteId must be specified in query", http.StatusBadRequest)
		return
	}
	if len(nIDs) != 1 {
		log.Print("noteId is not specified")
		http.Error(w, "Only one noteId should be specified in query", http.StatusBadRequest)
		return
	}
	nID := nIDs[0]
	k, pID, err := name.ParseResourceKindAndProjectFromPath(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		log.Printf("error parsing path %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	if k != name.Note {
		log.Printf("wrong object type %v", k)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", readErr)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	n := swagger.Note{}
	json.Unmarshal(body, &n)
	genName := name.FormatNote(pID, nID)
	if genName != n.Name {
		log.Printf("Mismatching names in n.Name field and request parameters.")
		http.Error(w, fmt.Sprintf("note.Name %v must specify match with request"+
			" url parameters with projectsId %v and noteID %v", n.Name, pID, nID),
			http.StatusBadRequest)
	}

	if err := h.g.CreateNote(&n); err != nil {
		log.Printf("Error creating note: %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}
	bytes, mErr := json.Marshal(&n)
	if err != nil {
		log.Printf("Error marshalling bytes: %v", mErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// CreateOccurrence handles http requests to create occurrences in grafeas
func (h *Handler) CreateOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	o := swagger.Occurrence{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	k, pID, parseErr := name.ParseResourceKindAndProjectFromPath(strings.Trim(r.URL.Path, "/"))
	if parseErr != nil {
		log.Printf("error parsing path %v", parseErr)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	if k != name.Occurrence {
		log.Printf("wrong object type %v", k)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	json.Unmarshal(body, &o)
	// Generate random occurrenceId to prevent collisions
	oID := uuid.New()
	oName := name.FormatOccurrence(pID, oID.String())

	// We replace the name in the specified occurrence with the name generated, as users shouldn't
	// specify an occurrence name.
	o.Name = oName
	if err := h.g.CreateOccurrence(&o); err != nil {
		log.Printf("Error creating occurrence: %v", err)
		http.Error(w, err.Err, err.StatusCode)
	}
	bytes, err := json.Marshal(&o)
	if err != nil {
		log.Printf("Error marshalling bytes: %v", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// CreateOperation handles http requests to create an operation in Grafeas
func (h *Handler) CreateOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	oIDs, ok := r.URL.Query()["operationId"]
	var oID string
	if !ok {
		oID = uuid.New().String()
	} else {
		if len(oIDs) != 1 {
			log.Print("too many operationIds specified")
			http.Error(w, "Only one operationId should be specified in query", http.StatusBadRequest)
			return
		} else {
			oID = oIDs[0]
		}
	}
	k, pID, err := name.ParseResourceKindAndProjectFromPath(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		log.Printf("error parsing path %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	if k != name.Operation {
		log.Printf("wrong object type %v", k)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	o := swagger.Operation{}

	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(body, &o)
	genName := name.FormatOperation(pID, oID)
	o.Name = genName
	if err := h.g.CreateOperation(&o); err != nil {
		log.Printf("Error creating occurrence: %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}

	bytes, mErr := json.Marshal(&o)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, nID, appErr := projectNoteIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	if err := h.g.DeleteNote(pID, nID); err != nil {
		log.Printf("Unable to delete note %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return

	}
	w.WriteHeader(http.StatusOK)
}

func projectNoteIDFromReq(r *http.Request) (string, string, *errors.AppError) {
	// We need to trim twice because the path may or may not contain the leading "/"
	nameString := strings.TrimPrefix(strings.TrimPrefix(r.URL.Path, "/"), "v1alpha1/")
	// Handle GetNoteOccurrences too
	nameString = strings.TrimSuffix(nameString, "/occurrences")
	pID, nID, err := name.ParseNote(nameString)
	if err != nil {
		log.Printf("error parsing path %v", err)
		return "", "", &errors.AppError{Err: err.Err, StatusCode: err.StatusCode}
	}
	return pID, nID, nil
}

func projectOccIDFromReq(r *http.Request) (string, string, *errors.AppError) {
	// We need to trim twice because the path may or may not contain the leading "/"
	nameString := strings.TrimPrefix(strings.TrimPrefix(r.URL.Path, "/"), "v1alpha1/")
	// Handle GetOccurrenceNotes too
	nameString = strings.TrimSuffix(nameString, "/notes")
	pID, oID, err := name.ParseOccurrence(nameString)
	if err != nil {
		log.Printf("error parsing path %v", err)
		return "", "", &errors.AppError{Err: err.Err, StatusCode: err.StatusCode}
	}
	return pID, oID, nil
}

func projectOperationIDFromReq(r *http.Request) (string, string, *errors.AppError) {
	// We need to trim twice because the path may or may not contain the leading "/"
	nameString := strings.TrimPrefix(strings.TrimPrefix(r.URL.Path, "/"), "v1alpha1/")

	pID, oID, err := name.ParseOperation(nameString)
	if err != nil {
		log.Printf("error parsing path %v", err)
		return "", "", &errors.AppError{Err: "Error processing request",
			StatusCode: http.StatusInternalServerError}
	}
	return pID, oID, nil
}

func (h *Handler) DeleteOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, nID, appErr := projectOperationIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	if err := h.g.DeleteOperation(pID, nID); err != nil {
		log.Printf("Unable to delete note %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return

	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, oID, appErr := projectOccIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	if err := h.g.DeleteOccurrence(pID, oID); err != nil {
		log.Printf("Unable to delete occurrence %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return

	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, nID, appErr := projectNoteIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	n, err := h.g.GetNote(pID, nID)
	if err != nil {
		log.Printf("Error getting note %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}

	bytes, mErr := json.Marshal(&n)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", mErr)
		http.Error(w, "Error getting Note", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, oID, appErr := projectOccIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	o, err := h.g.GetOccurrence(pID, oID)
	if err != nil {
		log.Printf("Error getting occurrence %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}

	bytes, mErr := json.Marshal(&o)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", mErr)
		http.Error(w, "Error getting Occurrence", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, oID, appErr := projectOperationIDFromReq(r)
	if appErr != nil {
		log.Printf("error parsing operation %v", appErr)
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	o, err := h.g.GetOperation(pID, oID)
	if err != nil {
		log.Printf("Error getting operation %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}

	bytes, mErr := json.Marshal(&o)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", mErr)
		http.Error(w, "Error getting operation", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetOccurrenceNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, oID, appErr := projectOccIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	n, err := h.g.GetOccurrenceNote(pID, oID)
	if err != nil {
		log.Printf("Error getting occurrence %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}

	bytes, mErr := json.Marshal(&n)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", mErr)
		http.Error(w, "Error getting Note", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListNoteOccurrences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Get project id
	pID, nID, err := projectNoteIDFromReq(r)
	if err != nil {
		log.Printf("error parsing path %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	// TODO: Support filters
	resp, err := h.g.ListNoteOccurrences(pID, nID, "")
	// Convert response to bytes
	bytes, mErr := json.Marshal(resp)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListOperations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Get project id
	k, pID, err := name.ParseResourceKindAndProjectFromPath(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		log.Printf("error parsing path %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	if k != name.Operation {
		log.Printf("wrong object type %v", k)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	// TODO: Support filters
	resp, err := h.g.ListOperations(pID, "")
	// Convert response to bytes
	bytes, mErr := json.Marshal(resp)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Get project id
	k, pID, err := name.ParseResourceKindAndProjectFromPath(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		log.Printf("error parsing path %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	if k != name.Note {
		log.Printf("wrong object type %v", k)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	// TODO: Support filters
	resp, err := h.g.ListNotes(pID, "")
	// Convert response to bytes
	bytes, mErr := json.Marshal(resp)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) ListOccurrences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// Get project id
	k, pID, err := name.ParseResourceKindAndProjectFromPath(strings.Trim(r.URL.Path, "/"))
	if err != nil {
		log.Printf("error parsing path %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	if k != name.Occurrence {
		log.Printf("wrong object type %v", k)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	// TODO: Support filters
	resp, err := h.g.ListOccurrences(pID, "")
	// Convert response to bytes
	bytes, mErr := json.Marshal(resp)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, nID, appErr := projectNoteIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", readErr)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	n := swagger.Note{}
	json.Unmarshal(body, &n)
	resp, err := h.g.UpdateNote(pID, nID, &n)
	if err != nil {
		log.Printf("Error updating note: %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}
	bytes, mErr := json.Marshal(&resp)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", mErr)
		http.Error(w, "Error getting Note", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateOccurrence(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, oID, appErr := projectOccIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", readErr)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	o := swagger.Occurrence{}
	json.Unmarshal(body, &o)
	resp, err := h.g.UpdateOccurrence(pID, oID, &o)
	if err != nil {
		log.Printf("Error updating occurrence: %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}
	bytes, mErr := json.Marshal(&resp)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", mErr)
		http.Error(w, "Error getting occurrence", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateOperation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	pID, oID, appErr := projectOperationIDFromReq(r)
	if appErr != nil {
		http.Error(w, appErr.Err, appErr.StatusCode)
		return
	}
	body, readErr := ioutil.ReadAll(r.Body)
	if readErr != nil {
		log.Printf("Error reading body: %v", readErr)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	o := swagger.Operation{}
	json.Unmarshal(body, &o)
	resp, err := h.g.UpdateOperation(pID, oID, &o)
	if err != nil {
		log.Printf("Error updating operation: %v", err)
		http.Error(w, err.Err, err.StatusCode)
		return
	}
	bytes, mErr := json.Marshal(&resp)
	if mErr != nil {
		log.Printf("Error marshalling bytes: %v", mErr)
		http.Error(w, "Error getting Operation", http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}
