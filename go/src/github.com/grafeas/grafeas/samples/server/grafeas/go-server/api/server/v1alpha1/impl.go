// package v1alpha1 is an implementation of the v1alpha1 version of Grafeas.
package v1alpha1

import (
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/errors"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/storage"
	"log"
	"net/http"
)

// Grafeas is an implementation of the Grafeas API, which should be called by handler methods for verification of logic
// and storage.
type Grafeas struct {
	s *storage.MemStore
}

// CreateNote validates that a note is valid and then creates a note in the backing datastore.
func (g *Grafeas) CreateNote(n *swagger.Note) *errors.AppError {
	if n.Name == "" {
		log.Printf("Invalid note name: %v", n.Name)
		return &errors.AppError{"Invalid note name", http.StatusBadRequest}
	}
	// TODO: Validate that operation exists if it is specified when get methods are implmented
	return g.s.CreateNote(n)
}

// CreateOccurrence validates that a note is valid and then creates an occurrence in the backing datastore.
func (g *Grafeas) CreateOccurrence(o *swagger.Occurrence) *errors.AppError {
	if o.Name == "" {
		log.Printf("Invalid occurrence name: %v", o.Name)
		return &errors.AppError{"Invalid occurrence name", http.StatusBadRequest}
	}
	// TODO: Validate that note exists
	// TODO: Validate that operation exists if it is specified
	return g.s.CreateOccurrence(o)
}

// CreateOperation validates that a note is valid and then creates an operation note in the backing datastore.
func (g *Grafeas) CreateOperation(o *swagger.Operation) *errors.AppError {
	if o.Name == "" {
		log.Printf("Invalid occurrence name: %v", o.Name)
		return &errors.AppError{"Invalid occurrence name", http.StatusBadRequest}
	}
	return g.s.CreateOperation(o)
}

