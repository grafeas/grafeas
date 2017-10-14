package server

import (
	"github.com/grafeas/grafeas/server-go/api/v1alpha1"
	"github.com/grafeas/grafeas/server-go/errors"
)

// Storager is the interface that a Grafeas storage implementation would provide
type Storager interface {
	// CreateNote adds the specified note
	CreateNote(n *api.Note) *errors.AppError

	// CreateOccurrence adds the specified occurrence
	CreateOccurrence(o *api.Occurrence) *errors.AppError

	// CreateOperation adds the specified operation
	CreateOperation(o *api.Operation) *errors.AppError

	// DeleteNote deletes the note with the given pID and nID
	DeleteNote(pID, nID string) *errors.AppError

	// DeleteOccurrence deletes the occurrence with the given pID and oID
	DeleteOccurrence(pID, oID string) *errors.AppError

	// DeleteOperation deletes the operation with the given pID and oID
	DeleteOperation(pID, opID string) *errors.AppError

	// GetNote returns the note with project (pID) and note ID (nID)
	GetNote(pID, nID string) (*api.Note, *errors.AppError)

	// GetNoteByOccurrence returns the note attached to occurrence with pID and oID
	GetNoteByOccurrence(pID, oID string) (*api.Note, *errors.AppError)

	// GetOccurrence returns the occurrence with pID and oID
	GetOccurrence(pID, oID string) (*api.Occurrence, *errors.AppError)

	// GetOperation returns the operation with pID and oID
	GetOperation(pID, opID string) (*api.Operation, *errors.AppError)

	// ListNoteOccurrences returns the occcurrences on the particular note (nID) for this project (pID)
	ListNoteOccurrences(pID, nID, filters string) ([]api.Occurrence, *errors.AppError)

	// ListNoteOccurrences returns the occcurrences on the particular note (nID) for this project (pID)
	ListNotes(pID, filters string) []api.Note

	// ListOccurrences returns the occurrences for this project ID (pID)
	ListOccurrences(pID, filters string) []api.Occurrence

	// ListOperations returns the operations for this project (pID)
	ListOperations(pID, filters string) []api.Operation

	// UpdateNote updates the existing note with the given pID and nID
	UpdateNote(pID, nID string, n *api.Note) *errors.AppError

	// UpdateOccurrence updates the existing occurrence with the given projectID and occurrenceID
	UpdateOccurrence(pID, oID string, o *api.Occurrence) *errors.AppError

	// UpdateOperation updates the existing operation with the given pID and nID
	UpdateOperation(pID, opID string, op *api.Operation) *errors.AppError
}
