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

	emptypb "github.com/golang/protobuf/ptypes/empty"
	"github.com/grafeas/grafeas/go/name"
	validator "github.com/grafeas/grafeas/go/v1beta1/api/validators/grafeas"
	gpb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetOccurrence gets the specified occurrence.
func (g *API) GetOccurrence(ctx context.Context, req *gpb.GetOccurrenceRequest) (*gpb.Occurrence, error) {
	pID, oID, err := name.ParseOccurrence(req.Name)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateGetOccurrenceRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating GetOccurrenceRequest request %+v in project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, oID, OccurrencesGet); err != nil {
		return nil, err
	}

	o, err := g.Storage.GetOccurrence(ctx, pID, oID)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// ListOccurrences lists occurrences for the specified project.
func (g *API) ListOccurrences(ctx context.Context, req *gpb.ListOccurrencesRequest) (*gpb.ListOccurrencesResponse, error) {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateListOccurrencesRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating ListOccurrencesRequest %+v in project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", OccurrencesList); err != nil {
		return nil, err
	}

	// The following call is not for validation purpose. Instead, its main goal is to get an adjusted PageSize value.
	ps, err := validator.ValidatePageSize(req.PageSize)
	if err != nil {
		return nil, err
	}
	if err := g.Filter.Validate(req.Filter); err != nil {
		return nil, err
	}

	occs, npt, err := g.Storage.ListOccurrences(ctx, pID, req.Filter, req.PageToken, ps)
	if err != nil {
		return nil, err
	}

	resp := &gpb.ListOccurrencesResponse{
		Occurrences:   occs,
		NextPageToken: npt,
	}
	return resp, nil
}

// CreateOccurrence creates the specified occurrence.
func (g *API) CreateOccurrence(ctx context.Context, req *gpb.CreateOccurrenceRequest) (*gpb.Occurrence, error) {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if req.Occurrence == nil {
		return nil, status.Errorf(codes.InvalidArgument, "an occurrence must be specified")
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", OccurrencesCreate); err != nil {
		return nil, err
	}

	// Creating occurrences requires an additional notes attacher permissions check before we can
	// continue validation.
	notePID, nID, err := name.ParseNote(req.Occurrence.NoteName)
	if err != nil {
		return nil, err
	}
	if err := g.Auth.CheckAccessAndProject(ctx, notePID, nID, NotesAttachOccurrence); err != nil {
		return nil, err
	}

	if err := validator.ValidateCreateOccurrenceRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating CreateOccurrenceRequest %+v for project %q: %v", req, pID, err)
	}

	uID, err := g.Auth.EndUserID(ctx)
	if err != nil {
		return nil, err
	}

	o, err := g.Storage.CreateOccurrence(ctx, pID, uID, req.Occurrence)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// BatchCreateOccurrences batch creates the specified occurrences.
func (g *API) BatchCreateOccurrences(ctx context.Context, req *gpb.BatchCreateOccurrencesRequest) (*gpb.BatchCreateOccurrencesResponse, error) {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", OccurrencesCreate); err != nil {
		return nil, err
	}

	if err := validator.ValidateBatchCreateOccurrencesRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, status.Errorf(codes.InvalidArgument, "one or more occurrences are invalid, no occurrences were created: %v", err)
		}
		g.Logger.Warningf(ctx, "Error in validating BatchCreateOccurrencesRequest %+v for project %q: %v", req, pID, err)
	}

	// Creating occurrences requires an additional notes attacher permissions check before we can
	// continue validation.
	authErrs := []error{}
	for i, o := range req.Occurrences {
		notePID, nID, err := name.ParseNote(o.NoteName)
		if err != nil {
			return nil, err
		}
		if err := g.Auth.CheckAccessAndProject(ctx, notePID, nID, NotesAttachOccurrence); err != nil {
			authErrs = append(authErrs, fmt.Errorf("occurrences[%d]: %s", i, err))
		}
	}
	if len(authErrs) > 0 {
		return nil, status.Errorf(codes.PermissionDenied, "one or more occurrences had auth errors, no occurrences were created: %v", authErrs)
	}

	uID, err := g.Auth.EndUserID(ctx)
	if err != nil {
		return nil, err
	}

	created, errs := g.Storage.BatchCreateOccurrences(ctx, pID, uID, req.Occurrences)
	if len(errs) != 0 {
		// Report any storage layer errors as invalid argument for now, find a better way to do this.
		return nil, status.Errorf(codes.InvalidArgument, "errors encountered when batch creating occurrences: %d of %d occurrences failed: %v", len(errs), len(req.Occurrences), errs)
	}

	resp := &gpb.BatchCreateOccurrencesResponse{
		Occurrences: created,
	}
	return resp, nil
}

// UpdateOccurrence updates the specified occurrence.
func (g *API) UpdateOccurrence(ctx context.Context, req *gpb.UpdateOccurrenceRequest) (*gpb.Occurrence, error) {
	pID, oID, err := name.ParseOccurrence(req.Name)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateUpdateOccurrenceRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating UpdateOccurrenceRequest %+v for project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, oID, OccurrencesUpdate); err != nil {
		return nil, err
	}
	notePID, nID, err := name.ParseNote(req.Occurrence.NoteName)
	if err != nil {
		return nil, err
	}
	if err := g.Auth.CheckAccessAndProject(ctx, notePID, nID, NotesAttachOccurrence); err != nil {
		return nil, err
	}

	o, err := g.Storage.UpdateOccurrence(ctx, pID, oID, req.Occurrence, req.UpdateMask)
	if err != nil {
		return nil, err
	}

	return o, nil
}

// DeleteOccurrence deletes the specified occurrence.
func (g *API) DeleteOccurrence(ctx context.Context, req *gpb.DeleteOccurrenceRequest) (*emptypb.Empty, error) {
	pID, oID, err := name.ParseOccurrence(req.Name)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateDeleteOccurrenceRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating DeleteOccurrenceRequest %+v in project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, oID, OccurrencesDelete); err != nil {
		return nil, err
	}

	o, err := g.Storage.GetOccurrence(ctx, pID, oID)
	if err != nil {
		return nil, err
	}
	if o.NoteName != "" {
		notePID, nID, err := name.ParseNote(o.NoteName)
		if err != nil {
			return nil, err
		}
		if err := g.Auth.CheckAccessAndProject(ctx, notePID, nID, NotesAttachOccurrence); err != nil {
			return nil, err
		}
	}

	if err := g.Storage.DeleteOccurrence(ctx, pID, oID); err != nil {
		return nil, err
	}

	// Purge any IAM policies set on this entity.
	if err := g.Auth.PurgePolicy(ctx, pID, oID, Occurrences); err != nil {
		// This fails open, should not block on policy deletion failure.
		g.Logger.Warningf(ctx, "Error deleting policies for occurrence %q in project %q: %v", oID, pID, err)
	}

	return nil, nil
}

// ListNoteOccurrences lists occurrences for the specified note.
func (g *API) ListNoteOccurrences(ctx context.Context, req *gpb.ListNoteOccurrencesRequest) (*gpb.ListNoteOccurrencesResponse, error) {
	pID, nID, err := name.ParseNote(req.Name)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateListNoteOccurrencesRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating ListNoteOccurrencesRequest %+v in project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, nID, NotesListOccurrences); err != nil {
		return nil, err
	}

	if err := g.Filter.Validate(req.Filter); err != nil {
		return nil, err
	}

	occs, npt, err := g.Storage.ListNoteOccurrences(ctx, pID, nID, req.Filter, req.PageToken, req.PageSize)
	if err != nil {
		return nil, err
	}

	resp := &gpb.ListNoteOccurrencesResponse{
		Occurrences:   occs,
		NextPageToken: npt,
	}
	return resp, nil
}

// GetVulnerabilityOccurrencesSummary produces a summary of vulnerability
// occurrences grouped by severity that match the specified filter.
func (g *API) GetVulnerabilityOccurrencesSummary(ctx context.Context, req *gpb.GetVulnerabilityOccurrencesSummaryRequest) (*gpb.VulnerabilityOccurrencesSummary, error) {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return nil, err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := validator.ValidateGetVulnerabilityOccurrencesSummaryRequest(req); err != nil {
		if g.EnforceValidation {
			return nil, err
		}
		g.Logger.Warningf(ctx, "Error in validating GetVulnerabilityOccurrencesSummaryRequest %+v in project %q: %v", req, pID, err)
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", OccurrencesList); err != nil {
		return nil, err
	}

	if err := g.Filter.Validate(req.Filter); err != nil {
		return nil, err
	}

	summary, err := g.Storage.GetVulnerabilityOccurrencesSummary(ctx, pID, req.Filter)
	if err != nil {
		return nil, err
	}
	return summary, nil
}
