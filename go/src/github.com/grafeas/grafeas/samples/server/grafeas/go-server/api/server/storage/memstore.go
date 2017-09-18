package storage

import (
	"context"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server"
	"fmt"
	"net/http"
)

type Store interface {
	CreateOccurrence(o *swagger.Occurrence) *server.AppError
	DeleteOccurrence(projectID, oID string) *server.AppError
	UpdateOccurrence(projectID, oID string, o *swagger.Occurrence) *server.AppError
	GetOccurrence(projectID, oID string) (*swagger.Occurrence, *server.AppError)
	ListOccurrences() *server.AppError
	CreateNote(nsf *swagger.Note) *server.AppError
	DeleteNote( providerID, nID string) *server.AppError
	UpdateNote(providerID, nID string, n *swagger.Note) *server.AppError
	GetNote(providerID, nID string, n *swagger.Note) *server.AppError
	GetNoteByOccurrence(projectID, oID string) (*swagger.Note, *server.AppError)
	ListNotes() *server.AppError
	ListNoteOccurrences() *server.AppError

	GetOperation(projectID, opID string) (*swagger.Operation, *server.AppError)
	CreateOperation( o *swagger.Operation) *server.AppError
	DeleteOperation(projectID, opID string) *server.AppError
	UpdateOperation( projectID, opID string, op *swagger.Operation) *server.AppError
	ListOperations() *server.AppError
}


type MemStore struct {
	occurrencesByID map[string]swagger.Occurrence
	notesByID map[string]swagger.Note
}

func (m *MemStore) CreateOccurrence(ctx context.Context, o *swagger.Occurrence) *server.AppError {
	if _, ok := m.occurrencesByID[o.Name]; ok {
		return &server.AppError{fmt.Sprintf("Occurrence with Name %v already exists", o.Name),
			http.StatusBadRequest}
	}
	m.occurrencesByID[o.Name] = *o
	return nil
}

func (m *MemStore) DeleteOccurrence(ctx context.Context, projectID, oID string) *server.AppError  {
	name := fmt.Sprintf("projects/%v/occurrences/%v", projectID, oID)
	if _, ok := m.occurrencesByID[name]; !ok {
		return &server.AppError{fmt.Sprintf("Occurrence with Name %v does not Exist", name),
			http.StatusBadRequest}
	}
	delete(m.occurrencesByID, name)
	return nil
}

func (m *MemStore) UpdateOccurrence(ctx context.Context, projectID, oID string, o *swagger.Occurrence) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) GetOccurrence(ctx context.Context, projectID, oID string) (*swagger.Occurrence, *server.AppError ) {
	panic("implement me")
}

func (m *MemStore) ListOccurrences(ctx context.Context) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) CreateNote(ctx context.Context, nsf *swagger.Note) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) DeleteNote(ctx context.Context, providerID, nID string) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) UpdateNote(ctx context.Context, providerID, nID string, n *swagger.Note) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) GetNote(ctx context.Context, providerID, nID string, n *swagger.Note) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) GetNoteByOccurrence(ctx context.Context, projectID, oID string) (*swagger.Note, *server.AppError ) {
	panic("implement me")
}

func (m *MemStore) ListNotes(ctx context.Context) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) ListNoteOccurrences(ctx context.Context) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) GetOperation(ctx context.Context, projectID, opID string) (*swagger.Operation, *server.AppError ) {
	panic("implement me")
}

func (m *MemStore) CreateOperation(ctx context.Context, o *swagger.Operation) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) DeleteOperation(ctx context.Context, projectID, opID string) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) UpdateOperation(ctx context.Context, projectID, opID string, op *swagger.Operation) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) ListOperations(ctx context.Context) *server.AppError  {
	panic("implement me")
}


