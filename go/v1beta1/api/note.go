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
	emptypb "github.com/golang/protobuf/ptypes/empty"
	"github.com/grafeas/grafeas/go/name"
	validator "github.com/grafeas/grafeas/go/v1beta1/api/validators/grafeas"
	gpb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateNote creates the specified note.
func (g *API) CreateNote(ctx context.Context, req *gpb.CreateNoteRequest) (*gpb.Note, error) {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateCreateNoteRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating CreateNoteRequest %+v for project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", NotesCreate); err != nil {
		return nil, err
	}

	uID, err := g.Auth.EndUserID(ctx)
	if err != nil {
		return nil, err
	}

	n, err := g.Storage.CreateNote(ctx, pID, req.NoteId, uID, req.Note)
	if err != nil {
		return nil, err
	}

	return n, nil
}

// BatchCreateNotes batch creates the specified notes.
func (g *API) BatchCreateNotes(ctx context.Context, req *gpb.BatchCreateNotesRequest) (*gpb.BatchCreateNotesResponse, error) {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateBatchCreateNotesRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, status.Errorf(codes.InvalidArgument, "one or more notes are invalid, no notes were created: %v", err)
		}
		g.Logger.Warningf(ctx, "Error in validating BatchCreateNotesRequest %+v for project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", NotesCreate); err != nil {
		return nil, err
	}

	uID, err := g.Auth.EndUserID(ctx)
	if err != nil {
		return nil, err
	}

	created, errs := g.Storage.BatchCreateNotes(ctx, pID, uID, req.Notes)
	if len(errs) > 0 {
		// Report any storage layer errors as invalid argument for now, find a better way to do this.
		return nil, status.Errorf(codes.InvalidArgument, "errors encountered when batch creating notes: %d of %d notes failed: %v", len(errs), len(req.Notes), errs)
	}

	resp := &gpb.BatchCreateNotesResponse{
		Notes: created,
	}
	return resp, nil
}

// GetNote gets the specified note.
func (g *API) GetNote(ctx context.Context, req *gpb.GetNoteRequest) (*gpb.Note, error) {
	pID, nID, err := name.ParseNote(req.Name)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateGetNoteRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating GetNoteRequest %+v in project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, nID, NotesGet); err != nil {
		return nil, err
	}

	n, err := g.Storage.GetNote(ctx, pID, nID)
	if err != nil {
		return nil, err
	}

	return n, nil
}

// UpdateNote updates the specified note.
func (g *API) UpdateNote(ctx context.Context, req *gpb.UpdateNoteRequest) (*gpb.Note, error) {
	pID, nID, err := name.ParseNote(req.Name)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateUpdateNoteRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating UpdateNoteRequest %+v for project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, nID, NotesUpdate); err != nil {
		return nil, err
	}

	n, err := g.Storage.UpdateNote(ctx, pID, nID, req.Note, req.UpdateMask)
	if err != nil {
		return nil, err
	}

	return n, nil
}

// DeleteNote deletes the specified note.
func (g *API) DeleteNote(ctx context.Context, req *gpb.DeleteNoteRequest) (*emptypb.Empty, error) {
	pID, nID, err := name.ParseNote(req.Name)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateDeleteNoteRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating DeleteNoteRequest %+v in project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, nID, NotesDelete); err != nil {
		return nil, err
	}

	if err := g.Storage.DeleteNote(ctx, pID, nID); err != nil {
		return nil, err
	}

	// Purge any IAM policies set on this entity.
	if err := g.Auth.PurgePolicy(ctx, pID, nID, Notes); err != nil {
		// This fails open, should not block on policy deletion failure.
		g.Logger.Warningf(ctx, "Error deleting policies for note %q in project %q: %v", nID, pID, err)
	}

	return nil, nil
}

// ListNotes lists notes for the specified project.
func (g *API) ListNotes(ctx context.Context, req *gpb.ListNotesRequest) (*gpb.ListNotesResponse, error) {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateListNotesRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating ListNotesRequest %+v in project %q: %v", req, pID, err)
	}

	if err := g.Filter.Validate(req.Filter); err != nil {
		return nil, err
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", NotesList); err != nil {
		return nil, err
	}

	// The following call is not for validation purpose. Instead, its main goal is to get an adjusted PageSize value.
	ps, err := validator.ValidatePageSize(req.PageSize)
	if err != nil {
		return nil, err
	}
	notes, npt, err := g.Storage.ListNotes(ctx, pID, req.Filter, req.PageToken, ps)
	if err != nil {
		return nil, err
	}

	resp := &gpb.ListNotesResponse{
		Notes:         notes,
		NextPageToken: npt,
	}
	return resp, nil
}

// GetOccurrenceNote gets the note for the specified occurrence.
func (g *API) GetOccurrenceNote(ctx context.Context, req *gpb.GetOccurrenceNoteRequest) (*gpb.Note, error) {
	pID, oID, err := name.ParseOccurrence(req.Name)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateGetOccurrenceNoteRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating GetOccurrenceNoteRequest %+v in project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, oID, OccurrencesGet); err != nil {
		return nil, err
	}

	n, err := g.Storage.GetOccurrenceNote(ctx, pID, oID)
	if err != nil {
		return nil, err
	}

	return n, nil
}
