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

// package v1alpha1 is an implementation of the v1alpha1 version of Grafeas.
package v1alpha1

import (
	"fmt"
	"github.com/grafeas/samples/server/go-server/api"
	"github.com/grafeas/samples/server/go-server/api/server/errors"
	"github.com/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/samples/server/go-server/api/server/storage"
	"log"
	"net/http"
	"github.com/grafeas/grafeas/samples/server/go-server/api"
)

// Grafeas is an implementation of the Grafeas API, which should be called by handler methods for verification of logic
// and storage.
type Grafeas struct {
	S *storage.MemStore
}

// CreateNote validates that a note is valid and then creates a note in the backing datastore.
func (g *Grafeas) CreateNote(n *swagger.Note) *errors.AppError {
	if n.Name == "" {
		log.Printf("Invalid note name: %v", n.Name)
		return &errors.AppError{Err: "Invalid note name", StatusCode: http.StatusBadRequest}
	}
	// TODO: Validate that operation exists if it is specified when get methods are implmented
	return g.S.CreateNote(n)
}

// CreateOccurrence validates that a note is valid and then creates an occurrence in the backing datastore.
func (g *Grafeas) CreateOccurrence(o *swagger.Occurrence) *errors.AppError {
	if o.Name == "" {
		log.Printf("Invalid occurrence name: %v", o.Name)
		return &errors.AppError{Err: "Invalid occurrence name", StatusCode: http.StatusBadRequest}
	}
	if o.NoteName == "" {
		log.Print("No note is associated with this occurrence")
	}
	pID, nID, err := name.ParseNote(o.NoteName)
	if err != nil {
		log.Printf("Invalid note name: %v", o.Name)
		return &errors.AppError{Err: "Invalid note name", StatusCode: http.StatusBadRequest}
	}
	if n, err := g.S.GetNote(pID, nID); n == nil || err != nil {
		log.Printf("Unable to getnote %v, err: %v", n, err)
		return &errors.AppError{Err: fmt.Sprintf("Note %v not found", o.NoteName), StatusCode: http.StatusBadRequest}
	}
	// TODO: Validate that operation exists if it is specified
	return g.S.CreateOccurrence(o)
}

// CreateOperation validates that a note is valid and then creates an operation note in the backing datastore.
func (g *Grafeas) CreateOperation(o *swagger.Operation) *errors.AppError {
	if o.Name == "" {
		log.Printf("Invalid occurrence name: %v", o.Name)
		return &errors.AppError{Err: "Invalid occurrence name", StatusCode: http.StatusBadRequest}
	}
	return g.S.CreateOperation(o)
}

// DeleteOccurrence deletes an occurrence from the datastore.
func (g *Grafeas) DeleteOccurrence(pID, oID string) *errors.AppError {
	return g.S.DeleteOccurrence(pID, oID)
}

// DeleteNote deletes a note from the datastore.
func (g *Grafeas) DeleteNote(pID, nID string) *errors.AppError {
	// TODO: Check for occurrences tied to this note, and return an error if there are any before deletion.
	return g.S.DeleteNote(pID, nID)
}

// DeleteOperation deletes an operation from the datastore.
func (g *Grafeas) DeleteOperation(pID, nID string) *errors.AppError {
	// TODO: Check for occurrences and notes tied to this operation, and return an error if there are any before deletion.
	return g.S.DeleteOperation(pID, nID)
}

// GetNote gets a note from the datastore.
func (g *Grafeas) GetNote(pID, nID string) (*swagger.Note, *errors.AppError) {
	return g.S.GetNote(pID, nID)
}

// GetOccurrence gets a occurrence from the datastore.
func (g *Grafeas) GetOccurrence(pID, oID string) (*swagger.Occurrence, *errors.AppError) {
	return g.S.GetOccurrence(pID, oID)
}

// GetOccurrence gets a occurrence from the datastore.
func (g *Grafeas) GetOperation(pID, oID string) (*swagger.Operation, *errors.AppError) {
	return g.S.GetOperation(pID, oID)
}

// GetOccurrenceNote gets a the note for the provided occurrence from the datastore.
func (g *Grafeas) GetOccurrenceNote(pID, oID string) (*swagger.Note, *errors.AppError) {
	o, err := g.S.GetOccurrence(pID, oID)
	if err != nil {
		return nil, err
	}
	npID, nID, err := name.ParseNote(o.NoteName)
	if err != nil {
		log.Printf("Invalid note name: %v", o.Name)
		return nil, &errors.AppError{Err: fmt.Sprintf("Invalid note name: %v", o.NoteName),
			StatusCode: http.StatusBadRequest}
	}
	return g.S.GetNote(npID, nID)
}

func (g *Grafeas) UpdateNote(pID, nID string, n *swagger.Note) (*swagger.Note, *errors.AppError) {
	// get existing note
	existing, err := g.GetNote(pID, nID)
	if err != nil {
		return nil, err
	}
	// verify that name didnt change
	if n.Name != existing.Name {
		log.Printf("Cannot change note name: %v", n.Name)
		return nil, &errors.AppError{Err: fmt.Sprintf("Cannot change note name: %v", n.Name),
			StatusCode: http.StatusBadRequest}
	}

	// update note
	if err = g.S.UpdateNote(pID, nID, n); err != nil {
		log.Printf("Cannot update note : %v", n.Name)
		return nil, &errors.AppError{Err: fmt.Sprintf("Cannot change note name: %v", n.Name),
			StatusCode: http.StatusInternalServerError}
	}
	return n, nil
}

func (g *Grafeas) UpdateOccurrence(pID, oID string, o *swagger.Occurrence) (*swagger.Occurrence, *errors.AppError) {
	// get existing Occurrence
	existing, err := g.GetOccurrence(pID, oID)
	if err != nil {
		return nil, err
	}

	// verify that name didnt change
	if o.Name != existing.Name {
		log.Printf("Cannot change occurrence name: %v", o.Name)
		return nil, &errors.AppError{Err: fmt.Sprintf("Cannot change occurrence name: %v", o.Name),
			StatusCode: http.StatusBadRequest}
	}
	// verify that if note name changed, it still exists
	if o.NoteName != existing.NoteName {
		npID, nID, err := name.ParseNote(o.NoteName)
		if err != nil {
			return nil, err
		}
		if newN, err := g.GetNote(npID, nID); newN == nil || err != nil {
			return nil, err
		}
	}

	// update Occurrence
	if err = g.S.UpdateOccurrence(pID, oID, o); err != nil {
		log.Printf("Cannot update occurrence : %v", o.Name)
		return nil, &errors.AppError{Err: fmt.Sprintf("Cannot update Occurrences: %v", err),
			StatusCode: http.StatusInternalServerError}
	}
	return o, nil
}

func (g *Grafeas) UpdateOperation(pID, oID string, o *swagger.Operation) (*swagger.Operation, *errors.AppError) {
	// get existing operation
	existing, err := g.GetOperation(pID, oID)
	if err != nil {
		return nil, err
	}

	// verify that operation isn't marked done
	if o.Done != existing.Done && existing.Done {
		log.Printf("Trying to update a done operation")
		return nil, &errors.AppError{Err: fmt.Sprintf("Cannot update operation in status done: %v", o.Name),
			StatusCode: http.StatusBadRequest}
	}

	// verify that name didnt change
	if o.Name != existing.Name {
		log.Printf("Cannot change operation name: %v", o.Name)
		return nil, &errors.AppError{Err: fmt.Sprintf("Cannot change operation name: %v", o.Name),
			StatusCode: http.StatusBadRequest}
	}

	// update operation
	if err = g.S.UpdateOperation(pID, oID, o); err != nil {
		log.Printf("Cannot update operation : %v", o.Name)
		return nil, &errors.AppError{Err: fmt.Sprintf("Cannot update Opreation: %v", err),
			StatusCode: http.StatusInternalServerError}
	}
	return o, nil
}
func (g *Grafeas) ListOperations(pID, fs string) (*swagger.ListOperationsResponse, *errors.AppError) {
	// TODO: support filters
	ops := g.S.ListOperations(pID, fs)
	return &swagger.ListOperationsResponse{Operations: ops}, nil
}

func (g *Grafeas) ListNotes(pID, fs string) (*swagger.ListNotesResponse, *errors.AppError) {
	// TODO: support filters
	ns := g.S.ListNotes(pID, fs)
	return &swagger.ListNotesResponse{Notes: ns}, nil

}

func (g *Grafeas) ListOccurrences(pID, fs string) (*swagger.ListOccurrencesResponse, *errors.AppError) {
	// TODO: support filters - prioritizing resource url
	os := g.S.ListOccurrences(pID, fs)
	return &swagger.ListOccurrencesResponse{Occurrences: os}, nil
}
