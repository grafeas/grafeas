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

func occurrenceName(pID, oID string) string {
	return fmt.Sprintf("projects/%v/occurrences/%v", pID, oID)
}

func (m *MemStore) CreateOccurrence(o *swagger.Occurrence) *server.AppError {
	if _, ok := m.occurrencesByID[o.Name]; ok {
		return &server.AppError{fmt.Sprintf("Occurrence with Name %v already exists", o.Name),
			http.StatusBadRequest}
	}
	m.occurrencesByID[o.Name] = *o
	return nil
}

func (m *MemStore) DeleteOccurrence(pID, oID string) *server.AppError  {
	name := occurrenceName(pID, oID)
	if _, ok := m.occurrencesByID[name]; !ok {
		return &server.AppError{fmt.Sprintf("Occurrence with Name %v does not Exist", name),
			http.StatusBadRequest}
	}
	delete(m.occurrencesByID, name)
	return nil
}

func (m *MemStore) UpdateOccurrence(projectID, oID string, o *swagger.Occurrence) *server.AppError  {
	name := fmt.Sprintf("projects/%v/occurrences/%v", projectID, oID)
}

func (m *MemStore) GetOccurrence(projectID, oID string) (*swagger.Occurrence, *server.AppError ) {
	panic("implement me")
}

func (m *MemStore) ListOccurrences() *server.AppError  {
	panic("implement me")
}

func (m *MemStore) CreateNote(nsf *swagger.Note) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) DeleteNote(providerID, nID string) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) UpdateNote(providerID, nID string, n *swagger.Note) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) GetNote(providerID, nID string, n *swagger.Note) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) GetNoteByOccurrence(projectID, oID string) (*swagger.Note, *server.AppError ) {
	panic("implement me")
}

func (m *MemStore) ListNotes() *server.AppError  {
	panic("implement me")
}

func (m *MemStore) ListNoteOccurrences() *server.AppError  {
	panic("implement me")
}

func (m *MemStore) GetOperation(projectID, opID string) (*swagger.Operation, *server.AppError ) {
	panic("implement me")
}

func (m *MemStore) CreateOperation(o *swagger.Operation) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) DeleteOperation(projectID, opID string) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) UpdateOperation(projectID, opID string, op *swagger.Operation) *server.AppError  {
	panic("implement me")
}

func (m *MemStore) ListOperations() *server.AppError  {
	panic("implement me")
}


