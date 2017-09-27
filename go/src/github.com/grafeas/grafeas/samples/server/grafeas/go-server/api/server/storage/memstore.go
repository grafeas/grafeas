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

package storage

import (
	"fmt"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/errors"

	"net/http"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/name"
)

// Memstore is an in-memory storage solution for Grafeas
type MemStore struct {
	occurrencesByID map[string]swagger.Occurrence
	notesByID       map[string]swagger.Note
	opsByID         map[string]swagger.Operation
}

// NewMemStore creates a memstore with all maps initialized.
func NewMemStore() *MemStore {
	return &MemStore{make(map[string]swagger.Occurrence), make(map[string]swagger.Note),
		make(map[string]swagger.Operation)}
}

// CreateOccurrence adds the specified occurrence to the mem store
func (m *MemStore) CreateOccurrence(o *swagger.Occurrence) *errors.AppError {
	if _, ok := m.occurrencesByID[o.Name]; ok {
		return &errors.AppError{fmt.Sprintf("Occurrence with Name %v already exists", o.Name),
			http.StatusBadRequest}
	}
	m.occurrencesByID[o.Name] = *o
	return nil
}

// Delete Occurrence deletes the occurrence with the given projectID and occurrenceID from the memstore
func (m *MemStore) DeleteOccurrence(pID, oID string) *errors.AppError {
	name := name.OccurrenceName(pID, oID)
	if _, ok := m.occurrencesByID[name]; !ok {
		return &errors.AppError{fmt.Sprintf("Occurrence with Name %v does not Exist", name),
			http.StatusBadRequest}
	}
	delete(m.occurrencesByID, name)
	return nil
}

// UpdateOccurrence makes changes in o to the existing occurrence with the given projectID and occurrenceID
func (m *MemStore) UpdateOccurrence(pID, oID string, o *swagger.Occurrence) *errors.AppError {
	name := name.OccurrenceName(pID, oID)
	if _, ok := m.occurrencesByID[name]; !ok {
		return &errors.AppError{fmt.Sprintf("Occurrence with Name %v does not Exist", name),
			http.StatusBadRequest}
	}
	m.occurrencesByID[name] = *o
	return nil
}

func (m *MemStore) GetOccurrence(pID, oID string) (*swagger.Occurrence, *errors.AppError) {
	name := name.OccurrenceName(pID, oID)
	o, ok := m.occurrencesByID[name]
	if !ok {
		return nil, &errors.AppError{fmt.Sprintf("Occurrence with Name %v does not Exist", name),
			http.StatusBadRequest}
	}
	return &o, nil
}

func (m *MemStore) ListOccurrences() *errors.AppError {
	panic("implement me")
}

func (m *MemStore) CreateNote(n *swagger.Note) *errors.AppError {
	if _, ok := m.notesByID[n.Name]; ok {
		return &errors.AppError{fmt.Sprintf("Occurrence with Name %v already exists", n.Name),
			http.StatusBadRequest}
	}
	m.notesByID[n.Name] = *n
	return nil
}

func (m *MemStore) DeleteNote(providerID, nID string) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) UpdateNote(providerID, nID string, n *swagger.Note) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) GetNote(providerID, nID string, n *swagger.Note) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) GetNoteByOccurrence(projectID, oID string) (*swagger.Note, *errors.AppError) {
	panic("implement me")
}

func (m *MemStore) ListNotes() *errors.AppError {
	panic("implement me")
}

func (m *MemStore) ListNoteOccurrences() *errors.AppError {
	panic("implement me")
}

func (m *MemStore) GetOperation(projectID, opID string) (*swagger.Operation, *errors.AppError) {
	panic("implement me")
}

func (m *MemStore) CreateOperation(o *swagger.Operation) *errors.AppError {
	if _, ok := m.opsByID[o.Name]; ok {
		return &errors.AppError{fmt.Sprintf("Operation with Name %v already exists", o.Name),
			http.StatusBadRequest}
	}
	m.opsByID[o.Name] = *o
	return nil
}

func (m *MemStore) DeleteOperation(projectID, opID string) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) UpdateOperation(projectID, opID string, op *swagger.Operation) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) ListOperations() *errors.AppError {
	panic("implement me")
}
