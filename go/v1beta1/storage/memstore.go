// Copyright 2019 The Grafeas Authors. All rights reserved.
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
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/grafeas/grafeas/go/errors"
	"github.com/grafeas/grafeas/go/name"
	gpb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	"golang.org/x/net/context"
	fieldmaskpb "google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
)

// MemStore is an in-memory storage solution for Grafeas
type MemStore struct {
	sync.RWMutex
	occurrencesByID map[string]*gpb.Occurrence
	notesByName     map[string]*gpb.Note
	projects        map[string]*prpb.Project
}

// NewMemStore creates a MemStore with all maps initialized.
func NewMemStore() *MemStore {
	return &MemStore{
		occurrencesByID: map[string]*gpb.Occurrence{},
		notesByName:     map[string]*gpb.Note{},
		projects:        map[string]*prpb.Project{},
	}
}

// CreateProject creates the specified project in memstore.
func (m *MemStore) CreateProject(ctx context.Context, pID string, p *prpb.Project) (*prpb.Project, error) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.projects[pID]; ok {
		return nil, errors.Newf(codes.AlreadyExists, "Project with name %q already exists", pID)
	}
	m.projects[pID] = p
	return m.projects[pID], nil
}

// GetProject gets the specified project from memstore.
func (m *MemStore) GetProject(ctx context.Context, pID string) (*prpb.Project, error) {
	m.RLock()
	defer m.RUnlock()
	p, ok := m.projects[pID]
	if !ok {
		return nil, errors.Newf(codes.NotFound, "Project with name %q does not Exist", pID)
	}
	return p, nil
}

// ListProjects returns up to pageSize number of projects beginning at pageToken, or from
// start if pageToken is the empty string.
func (m *MemStore) ListProjects(ctx context.Context, filter string, pageSize int, pageToken string) ([]*prpb.Project, string, error) {
	m.RLock()
	defer m.RUnlock()
	projects := make([]*prpb.Project, len(m.projects))
	i := 0
	for k := range m.projects {
		projects[i] = m.projects[k]
		i++
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Name < projects[j].Name
	})
	startPos := parsePageToken(pageToken, 0)
	endPos := min(startPos+pageSize, len(projects))
	return projects[startPos:endPos], nextPageToken(endPos, len(projects)), nil
}

// DeleteProject deletes the specified project from memstore.
func (m *MemStore) DeleteProject(ctx context.Context, pID string) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.projects[pID]; !ok {
		return errors.Newf(codes.NotFound, "Project with name %q does not Exist", pID)
	}
	delete(m.projects, pID)
	return nil
}

// GetOccurrence gets the specified occurrence from memstore.
func (m *MemStore) GetOccurrence(ctx context.Context, pID, oID string) (*gpb.Occurrence, error) {
	m.RLock()
	defer m.RUnlock()
	o, ok := m.occurrencesByID[oID]
	if !ok {
		return nil, errors.Newf(codes.NotFound, "Occurrence with ID %s does not Exist", oID)
	}

	// Set the output-only field before returning
	o.Name = name.FormatOccurrence(pID, oID)
	return o, nil
}

// ListOccurrences returns up to pageSize number of occurrences for this project beginning
// at pageToken, or from start if pageToken is the empty string.
func (m *MemStore) ListOccurrences(ctx context.Context, pID, filter, pageToken string, pageSize int32) ([]*gpb.Occurrence, string, error) {
	os := []*gpb.Occurrence{}
	m.RLock()
	defer m.RUnlock()
	for _, o := range m.occurrencesByID {
		if strings.HasPrefix(o.Name, fmt.Sprintf("projects/%v", pID)) {
			os = append(os, o)
		}
	}
	sort.Slice(os, func(i, j int) bool {
		return os[i].Name < os[j].Name
	})
	startPos := parsePageToken(pageToken, 0)
	endPos := min(startPos+int(pageSize), len(os))
	return os[startPos:endPos], nextPageToken(endPos, len(os)), nil
}

// CreateOccurrence creates the specified occurrence in memstore.
func (m *MemStore) CreateOccurrence(ctx context.Context, pID, uID string, o *gpb.Occurrence) (*gpb.Occurrence, error) {
	var id string
	o = proto.Clone(o).(*gpb.Occurrence)
	if nr, err := uuid.NewRandom(); err != nil {
		return nil, errors.Newf(codes.Internal, "Failed to generate UUID")
	} else {
		id = nr.String()
	}

	m.Lock()
	defer m.Unlock()
	if _, ok := m.occurrencesByID[id]; ok {
		return nil, errors.Newf(codes.AlreadyExists, "Occurrence with ID %q already exists", id)
	}
	o.CreateTime = ptypes.TimestampNow()
	o.UpdateTime = o.CreateTime
	o.Name = name.FormatOccurrence(pID, id)
	m.occurrencesByID[id] = o
	return o, nil
}

// BatchCreateOccurrence batch creates the specified occurrences in memstore.
func (m *MemStore) BatchCreateOccurrences(ctx context.Context, pID string, uID string, occs []*gpb.Occurrence) ([]*gpb.Occurrence, []error) {
	clonedOccs := []*gpb.Occurrence{}
	for _, o := range occs {
		clonedOccs = append(clonedOccs, proto.Clone(o).(*gpb.Occurrence))
	}
	occs = clonedOccs

	errs := []error{}
	created := []*gpb.Occurrence{}
	for _, o := range occs {
		occ, err := m.CreateOccurrence(ctx, pID, uID, o)
		if err != nil {
			// Occurrence already exists, skipping.
			continue
		} else {
			created = append(created, occ)
		}
	}

	return created, errs
}

// UpdateOccurrence updates the specified occurrence in memstore.
func (m *MemStore) UpdateOccurrence(ctx context.Context, pID, oID string, o *gpb.Occurrence, mask *fieldmaskpb.FieldMask) (*gpb.Occurrence, error) {
	o = proto.Clone(o).(*gpb.Occurrence)

	m.Lock()
	defer m.Unlock()
	if _, ok := m.occurrencesByID[oID]; !ok {
		return nil, errors.Newf(codes.NotFound, "Occurrence with ID %s does not exist", oID)
	}

	// TODO(#312): implement the update operation
	o.UpdateTime = ptypes.TimestampNow()
	m.occurrencesByID[oID] = o
	return o, nil
}

// DeleteOccurrence deletes the specified occurrence in memstore.
func (m *MemStore) DeleteOccurrence(ctx context.Context, pID, oID string) error {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.occurrencesByID[oID]; !ok {
		return errors.Newf(codes.NotFound, "Occurrence with ID %s does not Exist", oID)
	}
	delete(m.occurrencesByID, oID)
	return nil
}

// GetNote gets the specified note from memstore.
func (m *MemStore) GetNote(ctx context.Context, pID, nID string) (*gpb.Note, error) {
	nName := name.FormatNote(pID, nID)
	m.RLock()
	defer m.RUnlock()
	n, ok := m.notesByName[nName]
	if !ok {
		return nil, errors.Newf(codes.NotFound, "Note with name %q does not Exist", nName)
	}

	// Set the output-only field before returning
	n.Name = name.FormatNote(pID, nID)
	return n, nil
}

// ListNotes returns up to pageSize number of notes for the project pID beginning
// at pageToken, or from start if pageToken is the empty string.
func (m *MemStore) ListNotes(ctx context.Context, pID, filter, pageToken string, pageSize int32) ([]*gpb.Note, string, error) {
	ns := []*gpb.Note{}
	m.RLock()
	defer m.RUnlock()
	for _, n := range m.notesByName {
		if strings.HasPrefix(n.Name, fmt.Sprintf("projects/%v", pID)) {
			ns = append(ns, n)
		}
	}
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Name < ns[j].Name
	})
	startPos := parsePageToken(pageToken, 0)
	endPos := min(startPos+int(pageSize), len(ns))
	return ns[startPos:endPos], nextPageToken(endPos, len(ns)), nil
}

// CreateNote creates the specified note in memstore.
func (m *MemStore) CreateNote(ctx context.Context, pID, nID, uID string, n *gpb.Note) (*gpb.Note, error) {
	n = proto.Clone(n).(*gpb.Note)
	nName := name.FormatNote(pID, nID)

	m.Lock()
	defer m.Unlock()
	if _, ok := m.notesByName[nName]; ok {
		return nil, errors.Newf(codes.AlreadyExists, "Note with name %q already exists", n.Name)
	}

	n.Name = nName
	n.CreateTime = ptypes.TimestampNow()
	n.UpdateTime = n.CreateTime
	m.notesByName[nName] = n
	return n, nil
}

// BatchCreateNotes batch creates the specified notes in memstore.
func (m *MemStore) BatchCreateNotes(ctx context.Context, pID, uID string, notes map[string]*gpb.Note) ([]*gpb.Note, []error) {
	clonedNotes := map[string]*gpb.Note{}
	for nID, n := range notes {
		clonedNotes[nID] = proto.Clone(n).(*gpb.Note)
	}
	notes = clonedNotes

	errs := []error{}
	created := []*gpb.Note{}
	for nID, n := range notes {
		note, err := m.CreateNote(ctx, pID, nID, uID, n)
		if err != nil {
			// Note already exists, skipping.
			continue
		} else {
			created = append(created, note)
		}

	}

	return created, errs
}

// UpdateNote updates the specified note in memstore.
func (m *MemStore) UpdateNote(ctx context.Context, pID, nID string, n *gpb.Note, mask *fieldmaskpb.FieldMask) (*gpb.Note, error) {
	n = proto.Clone(n).(*gpb.Note)
	nName := name.FormatNote(pID, nID)

	m.Lock()
	defer m.Unlock()
	if _, ok := m.notesByName[nName]; !ok {
		return nil, errors.Newf(codes.NotFound, "Note with name %q does not Exist", nName)
	}

	// TODO(#312): implement the update operation
	n.UpdateTime = ptypes.TimestampNow()
	n.Name = nName
	m.notesByName[nName] = n
	return n, nil
}

// DeleteNote deletes the specified note in memstore.
func (m *MemStore) DeleteNote(ctx context.Context, pID, nID string) error {
	nName := name.FormatNote(pID, nID)
	m.Lock()
	defer m.Unlock()
	if _, ok := m.notesByName[nName]; !ok {
		return errors.Newf(codes.NotFound, "Note with name %q does not Exist", nName)
	}
	delete(m.notesByName, nName)
	return nil
}

// GetOccurrenceNote gets the note for the specified occurrence from memstore.
func (m *MemStore) GetOccurrenceNote(ctx context.Context, pID, oID string) (*gpb.Note, error) {
	m.RLock()
	defer m.RUnlock()
	o, ok := m.occurrencesByID[oID]
	if !ok {
		return nil, errors.Newf(codes.NotFound, "Occurrence with ID %s does not Exist", oID)
	}
	n, ok := m.notesByName[o.NoteName]
	if !ok {
		return nil, errors.Newf(codes.NotFound, "Note with name %q does not Exist", o.NoteName)
	}

	return n, nil
}

// ListNoteOccurrences returns up to pageSize number of occurrences on the note
// for the project beginning at pageToken, or from start if pageToken is empty.
func (m *MemStore) ListNoteOccurrences(ctx context.Context, pID, nID, filter, pageToken string, pageSize int32) ([]*gpb.Occurrence, string, error) {
	// TODO: use filters
	m.RLock()
	defer m.RUnlock()
	// Verify that note exists
	if _, err := m.GetNote(ctx, pID, nID); err != nil {
		return nil, "", err
	}
	nName := name.FormatNote(pID, nID)
	os := []*gpb.Occurrence{}
	for _, o := range m.occurrencesByID {
		if o.NoteName == nName {
			os = append(os, o)
		}
	}
	sort.Slice(os, func(i, j int) bool {
		return os[i].Name < os[j].Name
	})
	startPos := parsePageToken(pageToken, 0)
	endPos := min(startPos+int(pageSize), len(os))
	return os[startPos:endPos], nextPageToken(endPos, len(os)), nil
}

// GetVulnerabilityOccurrencesSummary gets a summary of vulnerability occurrences from storage.
func (m *MemStore) GetVulnerabilityOccurrencesSummary(ctx context.Context, projectID, filter string) (*gpb.VulnerabilityOccurrencesSummary, error) {
	return &gpb.VulnerabilityOccurrencesSummary{}, nil
}

// Parses the page token to an int. Returns defaultValue if parsing fails
func parsePageToken(pageToken string, defaultValue int) int {
	if pageToken == "" {
		return defaultValue
	}
	parsed, err := strconv.Atoi(pageToken)
	if err != nil {
		return defaultValue
	}
	return parsed
}

// Returns the smallest of a and b
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// nextPageToken returns the next page token (the next item index or empty if not more items are left)
func nextPageToken(lastPage, total int) string {
	if lastPage == total {
		return ""
	}
	return strconv.Itoa(lastPage)
}
