package storage

import (
	"context"
)

type Store interface {
	CreateOccurrence(ctx context.Context, o Occurrence) error
	DeleteOccurrence(ctx context.Context, projectID, oID string) error
	UpdateOccurrence(ctx context.Context, projectID, oID string, o Occurrence) error
	GetOccurrence(ctx context.Context, projectID, oID string, o Occurrence) error
	ListOccurrences(ctx context.Context, filters OccurrenceFilterFields, req *ListOccurrencesRequest, resp *alpha_pb.ListOccurrencesResponse) error
	CreateNote(ctx context.Context, nsf NoteStorageFields) error
	DeleteNote(ctx context.Context, providerID, nID string) error
	UpdateNote(ctx context.Context, providerID, nID string, r *alpha_pb.Note) error
	GetNote(ctx context.Context, providerID, nID string, resp *alpha_pb.Note) error
	GetNoteByOccurrence(ctx context.Context, projectID, oID string, resp *alpha_pb.Note) error
	ListNotes(ctx context.Context, filters NoteFilterFields, req *alpha_pb.ListNotesRequest, resp *alpha_pb.ListNotesResponse) error
	ListNoteOccurrences(ctx context.Context, filters NoteOccurrenceFilterFields, req *alpha_pb.ListNoteOccurrencesRequest, resp *alpha_pb.ListNoteOccurrencesResponse) error

	GetOperation(ctx context.Context, projectID, opID string) (Operation, error)
	CreateOperation(ctx context.Context, o *Operation) error
	DeleteOperation(ctx context.Context, projectID, opID string) error
	UpdateOperation(ctx context.Context, projectID, opID string, op *Operation) error
	ListOperations(ctx context.Context, filters OperationFilterFields, pageToken string, pageSize int32, resp *opspb.ListOperationsResponse) error
}