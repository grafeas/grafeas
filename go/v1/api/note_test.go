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
	"github.com/google/go-cmp/cmp"
	gpb "github.com/grafeas/grafeas/proto/v1/grafeas_go_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateNote(t *testing.T) {
	ctx := context.Background()
	g := &API{
		Storage:           newFakeStorage(),
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	req := &gpb.CreateNoteRequest{
		Parent: "projects/goog-vulnz",
		NoteId: "CVE-UH-OH",
		Note:   vulnzNote(t),
	}
	createdNote := &gpb.Note{}
	if err := g.CreateNote(ctx, req, createdNote); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	createdNote.Name = ""
	if diff := cmp.Diff(req.Note, createdNote, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("CreateNote(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestCreateNoteErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                                      string
		existingNotes                             map[string]*gpb.Note
		req                                       *gpb.CreateNoteRequest
		internalStorageErr, authErr, endUserIDErr bool
		wantErrStatus                             codes.Code
	}{
		{
			desc: "invalid project name",
			req: &gpb.CreateNoteRequest{
				Parent: "projectsIDONTBELIEVEINRULES",
				NoteId: "CVE-UH-OH",
				Note:   vulnzNote(t),
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "empty note ID",
			req: &gpb.CreateNoteRequest{
				Parent: "projects/goog-vulnz",
				NoteId: "",
				Note:   vulnzNote(t),
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "nil note",
			req: &gpb.CreateNoteRequest{
				Parent: "projects/goog-vulnz",
				NoteId: "CVE-UH-OH",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "auth error",
			req: &gpb.CreateNoteRequest{
				Parent: "projects/goog-vulnz",
				NoteId: "CVE-UH-OH",
				Note:   vulnzNote(t),
			},
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc: "end user ID error",
			req: &gpb.CreateNoteRequest{
				Parent: "projects/goog-vulnz",
				NoteId: "CVE-UH-OH",
				Note:   vulnzNote(t),
			},
			endUserIDErr:  true,
			wantErrStatus: codes.Internal,
		},
		{
			desc: "internal storage error",
			req: &gpb.CreateNoteRequest{
				Parent: "projects/goog-vulnz",
				NoteId: "CVE-UH-OH",
				Note:   vulnzNote(t),
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc: "note already exists, already exists error",
			existingNotes: map[string]*gpb.Note{
				"CVE-UH-OH": vulnzNote(t),
			},
			req: &gpb.CreateNoteRequest{
				Parent: "projects/goog-vulnz",
				NoteId: "CVE-UH-OH",
				Note:   vulnzNote(t),
			},
			wantErrStatus: codes.AlreadyExists,
		},
		{
			desc: "invalid vulnerability note",
			req: &gpb.CreateNoteRequest{
				Parent: "projects/goog-vulnz",
				NoteId: "CVE-UH-OH",
				Note:   invalidVulnzNote(t),
			},
			wantErrStatus: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.createNoteErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr, endUserIDErr: tt.endUserIDErr},
			EnforceValidation: true,
		}

		for nID, n := range tt.existingNotes {
			if _, err := s.CreateNote(ctx, "goog-vulnz", nID, "", n); err != nil {
				t.Fatalf("Failed to create note %+v", n)
			}
		}

		n := &gpb.Note{}
		err := g.CreateNote(ctx, tt.req, n)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestBatchCreateNotes(t *testing.T) {
	ctx := context.Background()
	g := &API{
		Storage:           newFakeStorage(),
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	req := &gpb.BatchCreateNotesRequest{
		Parent: "projects/goog-vulnz",
		Notes: map[string]*gpb.Note{
			"CVE-UH-OH": vulnzNote(t),
		},
	}
	resp := &gpb.BatchCreateNotesResponse{}
	if err := g.BatchCreateNotes(ctx, req, resp); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	if len(resp.Notes) != 1 {
		t.Fatalf("Got created notes of len %d, want 1", len(resp.Notes))
	}
	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	resp.Notes[0].Name = ""
	if diff := cmp.Diff(req.Notes["CVE-UH-OH"], resp.Notes[0], cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("BatchCreateNotes(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestBatchCreateNotesErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                                      string
		existingNotes                             map[string]*gpb.Note
		req                                       *gpb.BatchCreateNotesRequest
		internalStorageErr, authErr, endUserIDErr bool
		wantErrStatus                             codes.Code
	}{
		{
			desc: "invalid project name",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projectsdefinitelynotvalid",
				Notes: map[string]*gpb.Note{
					"CVE-UH-OH": vulnzNote(t),
				},
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "nil notes",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "empty notes",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes:  map[string]*gpb.Note{},
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "auth error",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes: map[string]*gpb.Note{
					"CVE-UH-OH": vulnzNote(t),
				},
			},
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc: "end user ID error",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes: map[string]*gpb.Note{
					"CVE-UH-OH": vulnzNote(t),
				},
			},
			endUserIDErr:  true,
			wantErrStatus: codes.Internal,
		},
		{
			desc: "internal storage error",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes: map[string]*gpb.Note{
					"CVE-UH-OH": vulnzNote(t),
				},
			},
			internalStorageErr: true,
			wantErrStatus:      codes.InvalidArgument,
		},
		{
			desc: "note already exists, invalid arg error",
			existingNotes: map[string]*gpb.Note{
				"CVE-UH-OH": vulnzNote(t),
			},
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes: map[string]*gpb.Note{
					"CVE-UH-OH": vulnzNote(t),
				},
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "invalid vulnerability note",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes: map[string]*gpb.Note{
					"CVE-UH-OH": invalidVulnzNote(t),
				},
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "two invalid vulnerability notes",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes: map[string]*gpb.Note{
					"CVE-UH-OH":  invalidVulnzNote(t),
					"CVE-UH-HUH": invalidVulnzNote(t),
				},
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "one valid, one invalid vulnerability notes",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes: map[string]*gpb.Note{
					"CVE-UH-OH":  vulnzNote(t),
					"CVE-UH-HUH": invalidVulnzNote(t),
				},
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "number of notes exceeds batch max",
			req: &gpb.BatchCreateNotesRequest{
				Parent: "projects/goog-vulnz",
				Notes:  vulnzNotes(t, 1001),
			},
			wantErrStatus: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.batchCreateNotesErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr, endUserIDErr: tt.endUserIDErr},
			EnforceValidation: true,
		}

		for nID, n := range tt.existingNotes {
			if _, err := s.CreateNote(ctx, "goog-vulnz", nID, "", n); err != nil {
				t.Fatalf("Failed to create note %+v", n)
			}
		}

		resp := &gpb.BatchCreateNotesResponse{}
		err := g.BatchCreateNotes(ctx, tt.req, resp)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestGetNote(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create the note to get.
	n := vulnzNote(t)
	if _, err := s.CreateNote(ctx, "goog-vulnz", "CVE-UH-OH", "", n); err != nil {
		t.Fatalf("Failed to create note %+v", n)
	}

	req := &gpb.GetNoteRequest{
		Name: "projects/goog-vulnz/notes/CVE-UH-OH",
	}
	gotN := &gpb.Note{}
	if err := g.GetNote(ctx, req, gotN); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	gotN.Name = ""
	if diff := cmp.Diff(n, gotN, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("GetNote(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestGetNoteErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                        string
		req                         *gpb.GetNoteRequest
		internalStorageErr, authErr bool
		wantErrStatus               codes.Code
	}{
		{
			desc: "invalid note name",
			req: &gpb.GetNoteRequest{
				Name: "projects/goog-vulnz",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "auth error",
			req: &gpb.GetNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-OH",
			},
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc: "internal storage error",
			req: &gpb.GetNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-OH",
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc: "note doesn't exist, not found error",
			req: &gpb.GetNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-HUH",
			},
			wantErrStatus: codes.NotFound,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.getNoteErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr},
			EnforceValidation: true,
		}

		n := &gpb.Note{}
		err := g.GetNote(ctx, tt.req, n)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestListNotes(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create the note to list.
	n := vulnzNote(t)
	if _, err := s.CreateNote(ctx, "goog-vulnz", "CVE-UH-OH", "", n); err != nil {
		t.Fatalf("Failed to create note %+v", n)
	}

	req := &gpb.ListNotesRequest{
		Parent: "projects/goog-vulnz",
	}
	resp := &gpb.ListNotesResponse{}
	if err := g.ListNotes(ctx, req, resp); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	resp.Notes[0].Name = ""
	if diff := cmp.Diff(n, resp.Notes[0], cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("ListNotes(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestListNotesErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                        string
		req                         *gpb.ListNotesRequest
		internalStorageErr, authErr bool
		wantErrStatus               codes.Code
	}{
		{
			desc: "invalid parent name",
			req: &gpb.ListNotesRequest{
				Parent: "projects",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "auth error",
			req: &gpb.ListNotesRequest{
				Parent: "projects/goog-vulnz",
			},
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc: "internal storage error",
			req: &gpb.ListNotesRequest{
				Parent: "projects/goog-vulnz",
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc: "invalid page size error",
			req: &gpb.ListNotesRequest{
				Parent:   "projects/goog-vulnz",
				PageSize: -1,
			},
			wantErrStatus: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.listNotesErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr},
			EnforceValidation: true,
		}

		resp := &gpb.ListNotesResponse{}
		err := g.ListNotes(ctx, tt.req, resp)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestUpdateNote(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create the note to update.
	n := vulnzNote(t)
	if _, err := s.CreateNote(ctx, "goog-vulnz", "CVE-UH-OH", "", n); err != nil {
		t.Fatalf("Failed to create note %+v", n)
	}
	n.ShortDescription = "a bad CVE"

	req := &gpb.UpdateNoteRequest{
		Name: "projects/goog-vulnz/notes/CVE-UH-OH",
		Note: n,
	}
	updatedN := &gpb.Note{}
	if err := g.UpdateNote(ctx, req, updatedN); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	updatedN.Name = ""
	if diff := cmp.Diff(n, updatedN, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("UpdateNote(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestUpdateNoteErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                        string
		req                         *gpb.UpdateNoteRequest
		internalStorageErr, authErr bool
		wantErrStatus               codes.Code
	}{
		{
			desc: "invalid note name",
			req: &gpb.UpdateNoteRequest{
				Name: "projects/totally/not/valid",
				Note: vulnzNote(t),
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "nil note",
			req: &gpb.UpdateNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-OH",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "auth error",
			req: &gpb.UpdateNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-OH",
				Note: vulnzNote(t),
			},
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc: "internal storage error",
			req: &gpb.UpdateNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-OH",
				Note: vulnzNote(t),
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc: "note doesn't exist, not found error",
			req: &gpb.UpdateNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-HUH",
				Note: vulnzNote(t),
			},
			wantErrStatus: codes.NotFound,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.updateNoteErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr},
			EnforceValidation: true,
		}

		n := &gpb.Note{}
		err := g.UpdateNote(ctx, tt.req, n)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestDeleteNote(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create the note to delete.
	n := vulnzNote(t)
	if _, err := s.CreateNote(ctx, "goog-vulnz", "CVE-UH-OH", "", n); err != nil {
		t.Fatalf("Failed to create note %+v", n)
	}

	req := &gpb.DeleteNoteRequest{
		Name: "projects/goog-vulnz/notes/CVE-UH-OH",
	}
	if err := g.DeleteNote(ctx, req, nil); err != nil {
		t.Errorf("Got err %v, want success", err)
	}
}

func TestDeleteNoteErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                                  string
		req                                   *gpb.DeleteNoteRequest
		internalStorageErr, authErr, purgeErr bool
		wantErr                               bool
		wantErrStatus                         codes.Code
	}{
		{
			desc: "invalid note name",
			req: &gpb.DeleteNoteRequest{
				Name: "projects/goog-vulnz",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "auth error",
			req: &gpb.DeleteNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-OH",
			},
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc: "internal storage error",
			req: &gpb.DeleteNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-OH",
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc: "purge policy err, fail open",
			req: &gpb.DeleteNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-OH",
			},
			purgeErr:      true,
			wantErrStatus: codes.OK,
		},
		{
			desc: "note doesn't exist, not found error",
			req: &gpb.DeleteNoteRequest{
				Name: "projects/goog-vulnz/notes/CVE-UH-HUH",
			},
			wantErrStatus: codes.NotFound,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.deleteNoteErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr, purgeErr: tt.purgeErr},
			EnforceValidation: true,
		}

		// Create the note to delete. It will only be deleted in the case of a purge policy error, which
		// fails open.
		n := vulnzNote(t)
		if _, err := s.CreateNote(ctx, "goog-vulnz", "CVE-UH-OH", "", n); err != nil {
			t.Fatalf("Failed to create note %+v", n)
		}

		err := g.DeleteNote(ctx, tt.req, nil)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestGetOccurrenceNote(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create the note we want to get and its occurrences.
	n := vulnzNote(t)
	if _, err := s.CreateNote(ctx, "goog-vulnz", "CVE-UH-OH", "", n); err != nil {
		t.Fatalf("Failed to create note %+v", n)
	}
	o := vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian")
	createdOcc, err := s.CreateOccurrence(ctx, "consumer1", "", o)
	if err != nil {
		t.Fatalf("Failed to create occurrence %+v", o)
	}

	req := &gpb.GetOccurrenceNoteRequest{
		Name: createdOcc.Name,
	}
	gotN := &gpb.Note{}
	if err := g.GetOccurrenceNote(ctx, req, gotN); err != nil {
		t.Fatalf("GetOccurrenceNote(%v): got err %v, want success", req, err)
	}

	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	gotN.Name = ""
	if diff := cmp.Diff(n, gotN, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("GetOccurrenceNote(%v): returned diff (want -> got):\n%s", req, diff)
	}
}

func TestGetOccurrenceNoteErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                        string
		req                         *gpb.GetOccurrenceNoteRequest
		internalStorageErr, authErr bool
		wantErrStatus               codes.Code
	}{
		{
			desc: "invalid occurrence name",
			req: &gpb.GetOccurrenceNoteRequest{
				Name: "projects/consumer1",
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc: "auth error",
			req: &gpb.GetOccurrenceNoteRequest{
				Name: "projects/consumer1/occurrences/1234-abcd-5678",
			},
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc: "internal storage error",
			req: &gpb.GetOccurrenceNoteRequest{
				Name: "projects/consumer1/occurrences/1234-abcd-5678",
			},
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc: "occurrence doesn't exist, not found error",
			req: &gpb.GetOccurrenceNoteRequest{
				Name: "projects/consumer1/occurrences/1234-abcd-5678",
			},
			wantErrStatus: codes.NotFound,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.getOccNoteErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr},
			EnforceValidation: true,
		}

		n := &gpb.Note{}
		err := g.GetOccurrenceNote(ctx, tt.req, n)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

// vulnzNote returns a fake v1 valid vulnerability note for testing.
func vulnzNote(t *testing.T) *gpb.Note {
	t.Helper()
	return &gpb.Note{
		Type: &gpb.Note_Vulnerability{
			Vulnerability: &gpb.VulnerabilityNote{
				Severity: gpb.Severity_CRITICAL,
				Details: []*gpb.VulnerabilityNote_Detail{
					{
						SeverityName:    "CRITICAL",
						AffectedCpeUri:  "cpe:/o:debian:debian_linux:7",
						AffectedPackage: "foobar",
						AffectedVersionStart: &gpb.Version{
							Kind: gpb.Version_MINIMUM,
						},
						AffectedVersionEnd: &gpb.Version{
							Kind: gpb.Version_MINIMUM,
						},
						FixedVersion: &gpb.Version{
							Kind: gpb.Version_MAXIMUM,
						},
					},
				},
			},
		},
	}
}

// invalidVulnzNote returns a fake v1 invalid vulnerability note for testing. Note has an empty
// detail.
func invalidVulnzNote(t *testing.T) *gpb.Note {
	t.Helper()
	return &gpb.Note{
		Type: &gpb.Note_Vulnerability{
			Vulnerability: &gpb.VulnerabilityNote{
				Severity: gpb.Severity_CRITICAL,
				Details: []*gpb.VulnerabilityNote_Detail{
					{},
				},
			},
		},
	}
}

// vulnzNotes creates the specified number of fake v1 valid vulnerability notes for testing.
func vulnzNotes(t *testing.T, num int) map[string]*gpb.Note {
	t.Helper()
	notes := map[string]*gpb.Note{}
	for i := 0; i < num; i++ {
		notes[fmt.Sprintf("CVE-UH-OH-%d", i)] = vulnzNote(t)
	}
	return notes
}
