package server

import (
	pb "github.com/grafeas/grafeas/v1alpha1/proto"
	opspb "google.golang.org/genproto/googleapis/longrunning"
)

// Storager is the interface that a Grafeas storage implementation would provide
type Storager interface {
	// CreateNote adds the specified note
	CreateNote(n *pb.Note) error

	// CreateOccurrence adds the specified occurrence
	CreateOccurrence(o *pb.Occurrence) error

	// CreateOperation adds the specified operation
	CreateOperation(o *opspb.Operation) error

	// DeleteNote deletes the note with the given pID and nID
	DeleteNote(pID, nID string) error

	// DeleteOccurrence deletes the occurrence with the given pID and oID
	DeleteOccurrence(pID, oID string) error

	// DeleteOperation deletes the operation with the given pID and oID
	DeleteOperation(pID, opID string) error

	// GetNote returns the note with project (pID) and note ID (nID)
	GetNote(pID, nID string) (*pb.Note, error)

	// GetNoteByOccurrence returns the note attached to occurrence with pID and oID
	GetNoteByOccurrence(pID, oID string) (*pb.Note, error)

	// GetOccurrence returns the occurrence with pID and oID
	GetOccurrence(pID, oID string) (*pb.Occurrence, error)

	// GetOperation returns the operation with pID and oID
	GetOperation(pID, opID string) (*opspb.Operation, error)

	// ListNoteOccurrences returns the occcurrences on the particular note (nID) for this project (pID)
	ListNoteOccurrences(pID, nID, filters string) ([]*pb.Occurrence, error)

	// ListNoteOccurrences returns the occcurrences on the particular note (nID) for this project (pID)
	ListNotes(pID, filters string) []*pb.Note

	// ListOccurrences returns the occurrences for this project ID (pID)
	ListOccurrences(pID, filters string) []*pb.Occurrence

	// ListOperations returns the operations for this project (pID)
	ListOperations(pID, filters string) []*opspb.Operation

	// UpdateNote updates the existing note with the given pID and nID
	UpdateNote(pID, nID string, n *pb.Note) error

	// UpdateOccurrence updates the existing occurrence with the given projectID and occurrenceID
	UpdateOccurrence(pID, oID string, o *pb.Occurrence) error

	// UpdateOperation updates the existing operation with the given pID and nID
	UpdateOperation(pID, opID string, op *opspb.Operation) error
}
