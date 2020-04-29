// Copyright 2018 The Grafeas Authors. All rights reserved.
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

package grafeas

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/grafeas/grafeas/go/iam"
	"github.com/grafeas/grafeas/go/name"
	validator "github.com/grafeas/grafeas/go/v1beta1/api/validators/grafeas"
	gpb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	provpb "github.com/grafeas/grafeas/proto/v1beta1/provenance_go_proto"
	vulnpb "github.com/grafeas/grafeas/proto/v1beta1/vulnerability_go_proto"
	"golang.org/x/net/context"
	fieldmaskpb "google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// fakeStorage implements the Grafeas storage interface using an in-memory map for tests. Filters
// and page tokens on list methods aren't supported. Update masks aren't supported. It only fills in
// resource name output only fields.
type fakeStorage struct {
	// Map of project IDs to a map of note IDs to their note.
	notes map[string]map[string]*gpb.Note
	// Map of project IDs to a map of occurrence IDs to their occurrence.
	occurrences map[string]map[string]*gpb.Occurrence

	// The following errors are for simulating an internal database error.
	getOccErr, listOccsErr, createOccErr, batchCreateOccsErr, updateOccErr, deleteOccErr       bool
	getNoteErr, listNotesErr, createNoteErr, batchCreateNotesErr, updateNoteErr, deleteNoteErr bool
	getOccNoteErr, listNoteOccsErr, getVulnSummaryErr                                          bool
}

func newFakeStorage() *fakeStorage {
	return &fakeStorage{
		notes:       map[string]map[string]*gpb.Note{},
		occurrences: map[string]map[string]*gpb.Occurrence{},
	}
}

func (s *fakeStorage) GetOccurrence(ctx context.Context, pID, oID string) (*gpb.Occurrence, error) {
	if s.getOccErr {
		return nil, status.Errorf(codes.Internal, "failed to get occurrence %q", oID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.occurrences[pID]; !ok {
		s.occurrences[pID] = map[string]*gpb.Occurrence{}
	}

	o, ok := s.occurrences[pID][oID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "occurrence %q not found", oID)
	}

	// Set the output-only field before returning
	o.Name = name.FormatOccurrence(pID, oID)
	return o, nil
}

func (s *fakeStorage) ListOccurrences(ctx context.Context, pID, filter, pageToken string, pageSize int32) ([]*gpb.Occurrence, string, error) {
	if s.listOccsErr {
		return nil, "", status.Errorf(codes.Internal, "failed to list occurrences for project %q", pID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.occurrences[pID]; !ok {
		s.occurrences[pID] = map[string]*gpb.Occurrence{}
	}

	occurrences := []*gpb.Occurrence{}
	for oID, o := range s.occurrences[pID] {
		// Set the output-only field before adding
		o.Name = name.FormatOccurrence(pID, oID)
		occurrences = append(occurrences, o)
	}

	return occurrences, "", nil
}

func (s *fakeStorage) CreateOccurrence(ctx context.Context, pID string, userID string, o *gpb.Occurrence) (*gpb.Occurrence, error) {
	o = proto.Clone(o).(*gpb.Occurrence)

	if s.createOccErr {
		return nil, status.Errorf(codes.Internal, "failed to create occurrence %+v", o)
	}

	// Create project if it doesn't exist.
	if _, ok := s.occurrences[pID]; !ok {
		s.occurrences[pID] = map[string]*gpb.Occurrence{}
	}

	oID := uuid.New().String()
	o.Name = name.FormatOccurrence(pID, oID)
	s.occurrences[pID][oID] = o

	return o, nil
}

func (s *fakeStorage) BatchCreateOccurrences(ctx context.Context, pID string, userID string, occs []*gpb.Occurrence) ([]*gpb.Occurrence, []error) {
	clonedOccs := []*gpb.Occurrence{}
	for _, o := range occs {
		clonedOccs = append(clonedOccs, proto.Clone(o).(*gpb.Occurrence))
	}
	occs = clonedOccs

	errs := []error{}
	if s.batchCreateOccsErr {
		for _, o := range occs {
			errs = append(errs, status.Errorf(codes.Internal, "failed to create occurrence %+v", o))
		}
		return nil, errs
	}

	// Create project if it doesn't exist.
	if _, ok := s.occurrences[pID]; !ok {
		s.occurrences[pID] = map[string]*gpb.Occurrence{}
	}

	created := []*gpb.Occurrence{}
	for _, o := range occs {
		oID := uuid.New().String()
		s.occurrences[pID][oID] = o
		o.Name = name.FormatOccurrence(pID, oID)
		created = append(created, o)
	}

	return created, errs
}

func (s *fakeStorage) UpdateOccurrence(ctx context.Context, pID, oID string, o *gpb.Occurrence, mask *fieldmaskpb.FieldMask) (*gpb.Occurrence, error) {
	o = proto.Clone(o).(*gpb.Occurrence)

	if s.updateOccErr {
		return nil, status.Errorf(codes.Internal, "failed to update occurrence %+v", o)
	}

	// Create project if it doesn't exist.
	if _, ok := s.occurrences[pID]; !ok {
		s.occurrences[pID] = map[string]*gpb.Occurrence{}
	}

	if _, ok := s.occurrences[pID][oID]; !ok {
		return nil, status.Errorf(codes.NotFound, "occurrence %q not found", oID)
	}

	o.Name = name.FormatOccurrence(pID, oID)
	s.occurrences[pID][oID] = o

	return o, nil
}

func (s *fakeStorage) DeleteOccurrence(ctx context.Context, pID, oID string) error {
	if s.deleteOccErr {
		return status.Errorf(codes.Internal, "failed to delete occurrence %q", oID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.occurrences[pID]; !ok {
		s.occurrences[pID] = map[string]*gpb.Occurrence{}
	}

	if _, ok := s.occurrences[pID][oID]; !ok {
		return status.Errorf(codes.NotFound, "occurrence %q not found", oID)
	}

	delete(s.occurrences[pID], oID)

	return nil
}

func (s *fakeStorage) GetNote(ctx context.Context, pID, nID string) (*gpb.Note, error) {
	if s.getNoteErr {
		return nil, status.Errorf(codes.Internal, "failed to get note %q", nID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.notes[pID]; !ok {
		s.notes[pID] = map[string]*gpb.Note{}
	}

	n, ok := s.notes[pID][nID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "note %q not found", nID)
	}

	// Set the output-only field before returning
	n.Name = name.FormatNote(pID, nID)
	return n, nil
}

func (s *fakeStorage) ListNotes(ctx context.Context, pID, filter, pageToken string, pageSize int32) ([]*gpb.Note, string, error) {
	if s.listNotesErr {
		return nil, "", status.Errorf(codes.Internal, "failed to list notes for project %q", pID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.notes[pID]; !ok {
		s.notes[pID] = map[string]*gpb.Note{}
	}

	notes := []*gpb.Note{}
	for _, n := range s.notes[pID] {
		notes = append(notes, n)
	}

	return notes, "", nil
}

func (s *fakeStorage) CreateNote(ctx context.Context, pID, nID string, userID string, n *gpb.Note) (*gpb.Note, error) {
	n = proto.Clone(n).(*gpb.Note)

	if s.createNoteErr {
		return nil, status.Errorf(codes.Internal, "failed to create note %+v", n)
	}

	// Create project if it doesn't exist.
	if _, ok := s.notes[pID]; !ok {
		s.notes[pID] = map[string]*gpb.Note{}
	}

	if _, ok := s.notes[pID][nID]; ok {
		return nil, status.Errorf(codes.AlreadyExists, "note %q already exists", nID)
	}

	n.Name = name.FormatNote(pID, nID)
	s.notes[pID][nID] = n

	return n, nil
}

func (s *fakeStorage) BatchCreateNotes(ctx context.Context, pID string, uID string, notes map[string]*gpb.Note) ([]*gpb.Note, []error) {
	clonedNotes := map[string]*gpb.Note{}
	for nID, n := range notes {
		clonedNotes[nID] = proto.Clone(n).(*gpb.Note)
	}
	notes = clonedNotes

	errs := []error{}
	if s.batchCreateNotesErr {
		for _, n := range notes {
			errs = append(errs, status.Errorf(codes.Internal, "failed to create note %+v", n))
		}
		return nil, errs
	}

	// Create project if it doesn't exist.
	if _, ok := s.notes[pID]; !ok {
		s.notes[pID] = map[string]*gpb.Note{}
	}

	created := []*gpb.Note{}
	for nID, n := range notes {
		if _, ok := s.notes[pID][nID]; ok {
			errs = append(errs, status.Errorf(codes.AlreadyExists, "note %q already exists", nID))
			continue
		}

		s.notes[pID][nID] = n
		n.Name = name.FormatNote(pID, nID)
		created = append(created, n)
	}

	return created, errs
}

func (s *fakeStorage) UpdateNote(ctx context.Context, pID, nID string, n *gpb.Note, mask *fieldmaskpb.FieldMask) (*gpb.Note, error) {
	n = proto.Clone(n).(*gpb.Note)

	if s.updateNoteErr {
		return nil, status.Errorf(codes.Internal, "failed to update note %+v", n)
	}

	// Create project if it doesn't exist.
	if _, ok := s.notes[pID]; !ok {
		s.notes[pID] = map[string]*gpb.Note{}
	}

	if _, ok := s.notes[pID][nID]; !ok {
		return nil, status.Errorf(codes.NotFound, "note %q not found", nID)
	}

	s.notes[pID][nID] = n
	n.Name = name.FormatNote(pID, nID)

	return n, nil
}

func (s *fakeStorage) DeleteNote(ctx context.Context, pID, nID string) error {
	if s.deleteNoteErr {
		return status.Errorf(codes.Internal, "failed to delete note %q", nID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.notes[pID]; !ok {
		s.notes[pID] = map[string]*gpb.Note{}
	}

	if _, ok := s.notes[pID][nID]; !ok {
		return status.Errorf(codes.NotFound, "note %q not found", nID)
	}

	delete(s.notes[pID], nID)

	return nil
}

func (s *fakeStorage) GetOccurrenceNote(ctx context.Context, pID, oID string) (*gpb.Note, error) {
	if s.getOccNoteErr {
		return nil, status.Errorf(codes.Internal, "failed to get note for occurrence %q", oID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.occurrences[pID]; !ok {
		s.occurrences[pID] = map[string]*gpb.Occurrence{}
	}

	// Get the occurrence and parse its note name.
	o, ok := s.occurrences[pID][oID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "occurrence %q not found", oID)
	}
	provID, nID, err := name.ParseNote(o.NoteName)
	if err != nil {
		return nil, err
	}

	// Create project if it doesn't exist.
	if _, ok := s.notes[pID]; !ok {
		s.notes[pID] = map[string]*gpb.Note{}
	}

	// Look up the note for the specified occurrence.
	n, ok := s.notes[provID][nID]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "note %q not found", nID)
	}

	// Set the output-only field before returning
	n.Name = name.FormatNote(pID, nID)
	return n, nil
}

func (s *fakeStorage) ListNoteOccurrences(ctx context.Context, pID, nID, filter, pageToken string, pageSize int32) ([]*gpb.Occurrence, string, error) {
	if s.listNoteOccsErr {
		return nil, "", status.Errorf(codes.Internal, "failed to get occurrences for note %q", nID)
	}

	// Create project if it doesn't exist.
	if _, ok := s.occurrences[pID]; !ok {
		s.occurrences[pID] = map[string]*gpb.Occurrence{}
	}

	foundOccs := []*gpb.Occurrence{}
	for _, occs := range s.occurrences {
		for _, o := range occs {
			if o.NoteName == name.FormatNote(pID, nID) {
				foundOccs = append(foundOccs, o)
			}
		}
	}

	return foundOccs, "", nil
}

func (s *fakeStorage) GetVulnerabilityOccurrencesSummary(ctx context.Context, projectID, filter string) (*gpb.VulnerabilityOccurrencesSummary, error) {
	if s.getVulnSummaryErr {
		return nil, fmt.Errorf("failed to get vulnerability occurrences summary for project %q", projectID)
	}

	return &gpb.VulnerabilityOccurrencesSummary{
		Counts: []*gpb.VulnerabilityOccurrencesSummary_FixableTotalByDigest{
			{
				Resource: &gpb.Resource{
					Name: "debian9",
					Uri:  "https://eu.gcr.io/consumer1/debian9@sha256:dbc96ed51bc598faeec0901bad307ebb5d1d7259b33e2d7d7296c28f439dc777",
					ContentHash: &provpb.Hash{
						Type:  provpb.Hash_SHA256,
						Value: []byte("dbc96ed51bc598faeec0901bad307ebb5d1d7259b33e2d7d7296c28f439dc777"),
					},
				},
				Severity:     vulnpb.Severity_CRITICAL,
				FixableCount: 1,
				TotalCount:   3,
			},
			{
				Resource: &gpb.Resource{
					Name: "debian9",
					Uri:  "https://eu.gcr.io/consumer1/debian9@sha256:dbc96ed51bc598faeec0901bad307ebb5d1d7259b33e2d7d7296c28f439dc777",
					ContentHash: &provpb.Hash{
						Type:  provpb.Hash_SHA256,
						Value: []byte("dbc96ed51bc598faeec0901bad307ebb5d1d7259b33e2d7d7296c28f439dc777"),
					},
				},
				Severity:     vulnpb.Severity_LOW,
				FixableCount: 4,
				TotalCount:   10,
			},
		},
	}, nil
}

type fakeAuth struct {
	// Whether auth calls return an error to exercise err code paths.
	authErr, endUserIDErr, purgeErr bool
}

func (a *fakeAuth) CheckAccessAndProject(ctx context.Context, projectID string, entityID string, p iam.Permission) error {
	if a.authErr {
		return status.Errorf(codes.PermissionDenied, "permission %q denied for %q or %q", p, projectID, entityID)
	}
	return nil
}

func (a *fakeAuth) EndUserID(ctx context.Context) (string, error) {
	if a.endUserIDErr {
		return "", status.Errorf(codes.Internal, "failed to get user ID")
	}
	return "23", nil
}

func (a *fakeAuth) PurgePolicy(ctx context.Context, projectID string, entityID string, r iam.Resource) error {
	if a.purgeErr {
		return status.Errorf(codes.Internal, "failed to purge policy for entity ID %q of resource type %q", entityID, r)
	}
	return nil
}

type fakeFilter struct {
	// Whether filter calls return an error to exercise err code paths.
	err bool
}

func (f *fakeFilter) Validate(filter string) error {
	if f.err {
		return status.Errorf(codes.InvalidArgument, "failed to parse filter %q", filter)
	}
	return nil
}

type fakeLogger struct{}

func (fakeLogger) PrepareCtx(ctx context.Context, projectID string) context.Context {
	return ctx
}
func (fakeLogger) Info(ctx context.Context, args ...interface{})                    {}
func (fakeLogger) Infof(ctx context.Context, format string, args ...interface{})    {}
func (fakeLogger) Warning(ctx context.Context, args ...interface{})                 {}
func (fakeLogger) Warningf(ctx context.Context, format string, args ...interface{}) {}
func (fakeLogger) Error(ctx context.Context, args ...interface{})                   {}
func (fakeLogger) Errorf(ctx context.Context, format string, args ...interface{})   {}

func TestValidatePageSize(t *testing.T) {
	tests := []struct {
		desc       string
		ps, wantPS int32
	}{
		{
			desc:   "page size of 0, want default page size (20)",
			ps:     0,
			wantPS: 20,
		},
		{
			desc:   "page size of 1, valid",
			ps:     1,
			wantPS: 1,
		},
		{
			desc:   "page size of 7, valid",
			ps:     7,
			wantPS: 7,
		},
		{
			desc:   "page size of 500, valid",
			ps:     500,
			wantPS: 500,
		},
		{
			desc:   "page size of 1000 (the max allowed), valid",
			ps:     1000,
			wantPS: 1000,
		},
	}

	for _, tt := range tests {
		ps, err := validator.ValidatePageSize(tt.ps)
		if err != nil {
			t.Errorf("%q: validatePageSize(%d): got error %v, want success", tt.desc, tt.ps, err)
		}
		if ps != tt.wantPS {
			t.Errorf("%q: validatePageSize(%d): got page size %d, want %d", tt.desc, tt.ps, ps, tt.wantPS)
		}
	}
}

func TestValidatePageSizeErrors(t *testing.T) {
	tests := []struct {
		desc        string
		ps, wantPS  int32
		wantErrCode codes.Code
	}{
		{
			desc:        "page size of -1, want error",
			ps:          -1,
			wantErrCode: codes.InvalidArgument,
		},
		{
			desc:        "page size of -500, want error",
			ps:          -500,
			wantErrCode: codes.InvalidArgument,
		},
		{
			desc:        "page size of 1001 (larger than max allowed), want error",
			ps:          1001,
			wantErrCode: codes.InvalidArgument,
		},
		{
			desc:        "page size of 5000 (larger than max allowed), want error",
			ps:          5000,
			wantErrCode: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		_, err := validator.ValidatePageSize(tt.ps)
		t.Logf("%q: error: %v", tt.desc, err)
		if err == nil {
			t.Errorf("%q: validatePageSize(%d): got success, want error code %q", tt.desc, tt.ps, tt.wantErrCode)
		}
		if status.Code(err) != tt.wantErrCode {
			t.Errorf("%q: validatePageSize(%d): got error code %q, want %q", tt.desc, tt.ps, status.Code(err), tt.wantErrCode)
		}
	}
}
