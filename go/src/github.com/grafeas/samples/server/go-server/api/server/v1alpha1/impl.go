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
