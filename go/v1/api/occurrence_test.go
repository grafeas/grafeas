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

func TestGetOccurrence(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create the occurrence to get.
	o := vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian")
	createdOcc, err := s.CreateOccurrence(ctx, "consumer1", "", o)
	if err != nil {
		t.Fatalf("Failed to create occurrence %+v", o)
	}

	req := &gpb.GetOccurrenceRequest{
		Name: createdOcc.Name,
	}
	gotOcc := &gpb.Occurrence{}
	if err := g.GetOccurrence(ctx, req, gotOcc); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	gotOcc.Name = ""
	if diff := cmp.Diff(o, gotOcc, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("GetOccurrence(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestGetOccurrenceErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                        string
		occName                     string
		internalStorageErr, authErr bool
		wantErrStatus               codes.Code
	}{
		{
			desc:          "invalid occurrence name",
			occName:       "projects/consumer1",
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "auth error",
			occName:       "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc:               "internal storage error",
			occName:            "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc:          "occurrence doesn't exist, not found error",
			occName:       "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			wantErrStatus: codes.NotFound,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.getOccErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr},
			EnforceValidation: true,
		}

		req := &gpb.GetOccurrenceRequest{
			Name: tt.occName,
		}
		gotOcc := &gpb.Occurrence{}
		err := g.GetOccurrence(ctx, req, gotOcc)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestListOccurrences(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create the occurrence to list.
	o := vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian")
	if _, err := s.CreateOccurrence(ctx, "consumer1", "", o); err != nil {
		t.Fatalf("Failed to create occurrence %+v", o)
	}

	req := &gpb.ListOccurrencesRequest{
		Parent: "projects/consumer1",
	}
	resp := &gpb.ListOccurrencesResponse{}
	if err := g.ListOccurrences(ctx, req, resp); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	if len(resp.Occurrences) != 1 {
		t.Fatalf("Got occurrences of len %d, want 1", len(resp.Occurrences))
	}
	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	resp.Occurrences[0].Name = ""
	if diff := cmp.Diff(o, resp.Occurrences[0], cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("ListOccurrences(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestListOccurrencesErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                        string
		parent                      string
		pageSize                    int32
		internalStorageErr, authErr bool
		wantErrStatus               codes.Code
	}{
		{
			desc:          "invalid parent name",
			parent:        "projects",
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "auth error",
			parent:        "projects/consumer1",
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc:               "internal storage error",
			parent:             "projects/consumer1",
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc:          "invalid page size error",
			parent:        "projects/consumer1",
			pageSize:      -1,
			wantErrStatus: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		s := newFakeStorage()
		s.listOccsErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr},
			EnforceValidation: true,
		}

		req := &gpb.ListOccurrencesRequest{
			Parent:   tt.parent,
			PageSize: tt.pageSize,
		}
		resp := &gpb.ListOccurrencesResponse{}
		err := g.ListOccurrences(ctx, req, resp)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestCreateOccurrence(t *testing.T) {
	ctx := context.Background()
	g := &API{
		Storage:           newFakeStorage(),
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	req := &gpb.CreateOccurrenceRequest{
		Parent:     "projects/consumer1",
		Occurrence: vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
	}
	createdOcc := &gpb.Occurrence{}
	if err := g.CreateOccurrence(ctx, req, createdOcc); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	createdOcc.Name = ""
	if diff := cmp.Diff(req.Occurrence, createdOcc, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("CreateOccurrence(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestCreateOccurrenceErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                                      string
		parent                                    string
		occ                                       *gpb.Occurrence
		internalStorageErr, authErr, endUserIDErr bool
		wantErrStatus                             codes.Code
	}{
		{
			desc:          "invalid project name",
			parent:        "projects",
			occ:           vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "nil occurrence",
			parent:        "projects/consumer1",
			occ:           nil,
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "invalid note name",
			parent:        "projects/consumer1",
			occ:           vulnzOcc(t, "consumer1", "foobar", "debian"),
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "auth error",
			parent:        "projects/consumer1",
			occ:           vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc:          "end user ID error",
			parent:        "projects/consumer1",
			occ:           vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			endUserIDErr:  true,
			wantErrStatus: codes.Internal,
		},
		{
			desc:               "internal storage error",
			parent:             "projects/consumer1",
			occ:                vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc:          "invalid vulnerability occurrence",
			parent:        "projects/goog-vulnz",
			occ:           invalidVulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH"),
			wantErrStatus: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.createOccErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr, endUserIDErr: tt.endUserIDErr},
			EnforceValidation: true,
		}

		req := &gpb.CreateOccurrenceRequest{
			Parent:     tt.parent,
			Occurrence: tt.occ,
		}
		o := &gpb.Occurrence{}
		err := g.CreateOccurrence(ctx, req, o)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestBatchCreateOccurrences(t *testing.T) {
	ctx := context.Background()
	g := &API{
		Storage:           newFakeStorage(),
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	req := &gpb.BatchCreateOccurrencesRequest{
		Parent: "projects/consumer1",
		Occurrences: []*gpb.Occurrence{
			vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
		},
	}
	resp := &gpb.BatchCreateOccurrencesResponse{}
	if err := g.BatchCreateOccurrences(ctx, req, resp); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	if len(resp.Occurrences) != 1 {
		t.Fatalf("Got created occurrences of len %d, want 1", len(resp.Occurrences))
	}
	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	resp.Occurrences[0].Name = ""
	if diff := cmp.Diff(req.Occurrences[0], resp.Occurrences[0], cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("BatchCreateOccurrences(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestBatchCreateOccurrencesErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                                      string
		parent                                    string
		occs                                      []*gpb.Occurrence
		internalStorageErr, authErr, endUserIDErr bool
		wantErrStatus                             codes.Code
	}{
		{
			desc:   "invalid project name",
			parent: "projects",
			occs: []*gpb.Occurrence{
				vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "nil occurrences",
			parent:        "projects/consumer1",
			occs:          nil,
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "empty occurrences",
			parent:        "projects/consumer1",
			occs:          []*gpb.Occurrence{},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:   "invalid note name",
			parent: "projects/consumer1",
			occs: []*gpb.Occurrence{
				vulnzOcc(t, "consumer1", "foobar", "debian"),
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:   "auth error",
			parent: "projects/consumer1",
			occs: []*gpb.Occurrence{
				vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			},
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc:   "end user ID error",
			parent: "projects/consumer1",
			occs: []*gpb.Occurrence{
				vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			},
			endUserIDErr:  true,
			wantErrStatus: codes.Internal,
		},
		{
			desc:   "internal storage error",
			parent: "projects/consumer1",
			occs: []*gpb.Occurrence{
				vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			},
			internalStorageErr: true,
			wantErrStatus:      codes.InvalidArgument,
		},
		{
			desc:   "invalid vulnerability occurrence",
			parent: "projects/consumer1",
			occs: []*gpb.Occurrence{
				invalidVulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH"),
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:   "two invalid vulnerability occurrences",
			parent: "projects/consumer1",
			occs: []*gpb.Occurrence{
				invalidVulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH"),
				invalidVulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH"),
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:   "one valid, one invalid vulnerability occurrences",
			parent: "projects/consumer1",
			occs: []*gpb.Occurrence{
				vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
				invalidVulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH"),
			},
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "number of occurrences exceeds batch max",
			parent:        "projects/consumer1",
			occs:          vulnzOccs(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian-", 1001),
			wantErrStatus: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.batchCreateOccsErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr, endUserIDErr: tt.endUserIDErr},
			EnforceValidation: true,
		}

		req := &gpb.BatchCreateOccurrencesRequest{
			Parent:      tt.parent,
			Occurrences: tt.occs,
		}
		resp := &gpb.BatchCreateOccurrencesResponse{}
		err := g.BatchCreateOccurrences(ctx, req, resp)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestUpdateOccurrence(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create the occurrence to update.
	o := vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian")
	createdOcc, err := s.CreateOccurrence(ctx, "consumer1", "", o)
	if err != nil {
		t.Fatalf("Failed to create occurrence %+v", o)
	}
	o.Remediation = "update to latest version"

	req := &gpb.UpdateOccurrenceRequest{
		Name:       createdOcc.Name,
		Occurrence: o,
	}
	updatedOcc := &gpb.Occurrence{}
	if err := g.UpdateOccurrence(ctx, req, updatedOcc); err != nil {
		t.Fatalf("Got err %v, want success", err)
	}

	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	updatedOcc.Name = ""
	if diff := cmp.Diff(o, updatedOcc, cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("UpdateOccurrence(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestUpdateOccurrenceErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                        string
		occName                     string
		occ                         *gpb.Occurrence
		internalStorageErr, authErr bool
		wantErrStatus               codes.Code
	}{
		{
			desc:          "invalid occurrence name",
			occName:       "projects",
			occ:           vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "nil occurrence",
			occName:       "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			occ:           nil,
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "invalid note name",
			occName:       "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			occ:           vulnzOcc(t, "consumer1", "projects/foobar", "debian"),
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "auth error",
			occName:       "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			occ:           vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc:               "internal storage error",
			occName:            "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			occ:                vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc:          "occurrence doesn't exist, not found error",
			occName:       "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			occ:           vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			wantErrStatus: codes.NotFound,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.updateOccErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr},
			EnforceValidation: true,
		}

		req := &gpb.UpdateOccurrenceRequest{
			Name:       tt.occName,
			Occurrence: tt.occ,
		}
		o := &gpb.Occurrence{}
		err := g.UpdateOccurrence(ctx, req, o)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestDeleteOccurrence(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		noteName string
	}{
		{noteName: "projects/goog-vulnz/notes/CVE-UH-OH"},
		{noteName: ""},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{},
			EnforceValidation: true,
		}

		// Create the occurrence to delete.
		o := vulnzOcc(t, "consumer1", tt.noteName, "debian")
		createdOcc, err := s.CreateOccurrence(ctx, "consumer1", "", o)
		if err != nil {
			t.Fatalf("Failed to create occurrence %+v", o)
		}

		req := &gpb.DeleteOccurrenceRequest{
			Name: createdOcc.Name,
		}
		if err := g.DeleteOccurrence(ctx, req, nil); err != nil {
			t.Errorf("Got err %v, want success", err)
		}
	}
}

func TestDeleteOccurrenceErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                                  string
		existingOcc                           *gpb.Occurrence
		occToDeleteOverride                   string
		internalStorageErr, authErr, purgeErr bool
		wantErrStatus                         codes.Code
	}{
		{
			desc:                "invalid occurrence name in delete request",
			existingOcc:         vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			occToDeleteOverride: "projects/consumer1",
			wantErrStatus:       codes.InvalidArgument,
		},
		{
			desc:          "occurrence to delete has invalid note name",
			existingOcc:   vulnzOcc(t, "consumer1", "projects/goog-vulnz", "debian"),
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "auth error",
			existingOcc:   vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc:               "internal storage error",
			existingOcc:        vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc:          "purge policy err, fail open",
			existingOcc:   vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian"),
			purgeErr:      true,
			wantErrStatus: codes.OK,
		},
		{
			desc:                "occurrence doesn't exist, not found error",
			occToDeleteOverride: "projects/consumer1/occurrences/1234-abcd-3456-wxyz",
			wantErrStatus:       codes.NotFound,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.deleteOccErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr, purgeErr: tt.purgeErr},
			EnforceValidation: true,
		}

		var createdOcc *gpb.Occurrence
		if tt.existingOcc != nil {
			var err error
			createdOcc, err = s.CreateOccurrence(ctx, "consumer1", "", tt.existingOcc)
			if err != nil {
				t.Fatalf("Failed to create occurrence %+v: %v", tt.existingOcc, err)
			}
		}

		occToDelete := createdOcc.GetName()
		if tt.occToDeleteOverride != "" {
			occToDelete = tt.occToDeleteOverride
		}

		req := &gpb.DeleteOccurrenceRequest{
			Name: occToDelete,
		}
		err := g.DeleteOccurrence(ctx, req, nil)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

func TestListNoteOccurrences(t *testing.T) {
	ctx := context.Background()
	s := newFakeStorage()
	g := &API{
		Storage:           s,
		Auth:              &fakeAuth{},
		EnforceValidation: true,
	}

	// Create a note and its occurrences we want to list.
	n := vulnzNote(t)
	if _, err := s.CreateNote(ctx, "goog-vulnz", "CVE-UH-OH", "", n); err != nil {
		t.Fatalf("Failed to create note %+v", n)
	}
	o := vulnzOcc(t, "consumer1", "projects/goog-vulnz/notes/CVE-UH-OH", "debian")
	createdOcc, err := s.CreateOccurrence(ctx, "consumer1", "", o)
	if err != nil {
		t.Fatalf("Failed to create occurrence %+v", o)
	}

	req := &gpb.ListNoteOccurrencesRequest{
		Name: "projects/goog-vulnz/notes/CVE-UH-OH",
	}
	resp := &gpb.ListNoteOccurrencesResponse{}
	if err := g.ListNoteOccurrences(ctx, req, resp); err != nil {
		t.Fatalf("ListNoteOccurrences(%v) got err %v, want success", req, err)
	}

	if len(resp.Occurrences) != 1 {
		t.Fatalf("Got occurrences of len %d, want 1", len(resp.Occurrences))
	}
	// TODO: migrate to protocolbuffers/protobuf-go when it is stable so we can use
	// protocmp.IgnoreFields instead.
	resp.Occurrences[0].Name = ""
	if diff := cmp.Diff(createdOcc, resp.Occurrences[0], cmp.Comparer(proto.Equal)); diff != "" {
		t.Errorf("ListNoteOccurrences(%v) returned diff (want -> got):\n%s", req, diff)
	}
}

func TestListNoteOccurrencesErrors(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		desc                        string
		noteName                    string
		pageSize                    int32
		internalStorageErr, authErr bool
		wantErrStatus               codes.Code
	}{
		{
			desc:          "invalid note name",
			noteName:      "projects/google-vulnz/notes/",
			wantErrStatus: codes.InvalidArgument,
		},
		{
			desc:          "auth error",
			noteName:      "projects/goog-vulnz/notes/CVE-UH-OH",
			authErr:       true,
			wantErrStatus: codes.PermissionDenied,
		},
		{
			desc:               "internal storage error",
			noteName:           "projects/goog-vulnz/notes/CVE-UH-OH",
			internalStorageErr: true,
			wantErrStatus:      codes.Internal,
		},
		{
			desc:          "invalid page size error",
			noteName:      "projects/goog-vulnz/notes/CVE-UH-OH",
			pageSize:      -1,
			wantErrStatus: codes.InvalidArgument,
		},
	}

	for _, tt := range tests {
		s := newFakeStorage()
		s.listNoteOccsErr = tt.internalStorageErr
		g := &API{
			Storage:           s,
			Auth:              &fakeAuth{authErr: tt.authErr},
			EnforceValidation: true,
		}

		req := &gpb.ListNoteOccurrencesRequest{
			Name:     tt.noteName,
			PageSize: tt.pageSize,
		}
		resp := &gpb.ListNoteOccurrencesResponse{}
		err := g.ListNoteOccurrences(ctx, req, resp)
		t.Logf("%q: error: %v", tt.desc, err)
		if status.Code(err) != tt.wantErrStatus {
			t.Errorf("%q: got error status %v, want %v", tt.desc, status.Code(err), tt.wantErrStatus)
		}
	}
}

// vulnzOcc returns a fake v1 valid vulnerability occurrence for testing.
func vulnzOcc(t *testing.T, pID, noteName, imageName string) *gpb.Occurrence {
	t.Helper()
	return &gpb.Occurrence{
		ResourceUri: fmt.Sprintf("https://us.gcr.io/%s/%s@sha256:0baa7a935c0cba530xxx03af85770cb52b26bfe570a9ff09e17c1a02c6b0bd9a", pID, imageName),
		NoteName:    noteName,
		Details: &gpb.Occurrence_Vulnerability{
			Vulnerability: &gpb.VulnerabilityOccurrence{
				PackageIssue: []*gpb.VulnerabilityOccurrence_PackageIssue{
					{
						AffectedCpeUri:  "cpe:/o:debian:debian_linux:8",
						AffectedPackage: "abc",
						AffectedVersion: &gpb.Version{
							Name: "0.2.0",
							Kind: gpb.Version_NORMAL,
						},
						FixedCpeUri:  "cpe:/o:debian:debian_linux:8",
						FixedPackage: "abc",
						FixedVersion: &gpb.Version{
							Name: "1.0.0",
							Kind: gpb.Version_NORMAL,
						},
					},
				},
			},
		},
	}
}

// invalidVulnzOcc returns a fake v1 invalid vulnerability occurrence for testing. Occurrence
// is missing resource.
func invalidVulnzOcc(t *testing.T, pID, noteName string) *gpb.Occurrence {
	t.Helper()
	return &gpb.Occurrence{
		NoteName: noteName,
		Details: &gpb.Occurrence_Vulnerability{
			Vulnerability: &gpb.VulnerabilityOccurrence{
				PackageIssue: []*gpb.VulnerabilityOccurrence_PackageIssue{
					{
						AffectedCpeUri:  "cpe:/o:debian:debian_linux:8",
						AffectedPackage: "abc",
						AffectedVersion: &gpb.Version{
							Name: "0.2.0",
							Kind: gpb.Version_NORMAL,
						},
						FixedCpeUri:  "cpe:/o:debian:debian_linux:8",
						FixedPackage: "abc",
						FixedVersion: &gpb.Version{
							Name: "1.0.0",
							Kind: gpb.Version_NORMAL,
						},
					},
				},
			},
		},
	}
}

// vulnzOccs creates the specified number of fake v1 valid vulnerability occurrences for
// testing.
func vulnzOccs(t *testing.T, pID, noteName, imageNamePrefix string, num int) []*gpb.Occurrence {
	t.Helper()
	occs := []*gpb.Occurrence{}
	for i := 0; i < num; i++ {
		occs = append(occs, vulnzOcc(t, pID, noteName, fmt.Sprintf("%s%d", imageNamePrefix, i)))
	}
	return occs
}
