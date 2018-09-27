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
	"github.com/grafeas/grafeas/go/errors"
	"github.com/grafeas/grafeas/go/name"
	"github.com/grafeas/grafeas/go/v1beta1/api/validators/grafeas"
	gpb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
)

// GetOccurrence gets the specified occurrence.
func (g *API) GetOccurrence(ctx context.Context, req *gpb.GetOccurrenceRequest, resp *gpb.Occurrence) error {
	pID, oID, err := name.ParseOccurrence(req.Name)
	if err != nil {
		return err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := g.Auth.CheckAccessAndProject(ctx, pID, oID, OccurrencesGet); err != nil {
		return err
	}

	o, err := g.Storage.GetOccurrence(ctx, pID, oID)
	if err != nil {
		return err
	}
	*resp = *o

	return nil
}

// ListOccurrences lists occurrences for the specified project.
func (g *API) ListOccurrences(ctx context.Context, req *gpb.ListOccurrencesRequest, resp *gpb.ListOccurrencesResponse) error {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", OccurrencesList); err != nil {
		return err
	}

	ps, err := validatePageSize(req.PageSize)
	if err != nil {
		return err
	}
	if err := g.Filter.Validate(req.Filter); err != nil {
		return err
	}

	occs, npt, err := g.Storage.ListOccurrences(ctx, pID, req.Filter, req.PageToken, ps)
	if err != nil {
		return err
	}
	resp.Occurrences = occs
	resp.NextPageToken = npt

	return nil
}

// CreateOccurrence creates the specified occurrence.
func (g *API) CreateOccurrence(ctx context.Context, req *gpb.CreateOccurrenceRequest, resp *gpb.Occurrence) error {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if req.Occurrence == nil {
		return errors.Newf(codes.InvalidArgument, "an occurrence must be specified")
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", OccurrencesCreate); err != nil {
		return err
	}

	// Creating occurrences requires an additional notes attacher permissions check before we can
	// continue validation.
	notePID, nID, err := name.ParseNote(req.Occurrence.NoteName)
	if err != nil {
		return err
	}
	if err := g.Auth.CheckAccessAndProject(ctx, notePID, nID, NotesAttachOccurrence); err != nil {
		return err
	}

	if err := grafeas.ValidateOccurrence(req.Occurrence); err != nil {
		if g.EnforceValidation {
			return err
		}
		g.Logger.Warningf(ctx, "CreateOccurrence %+v for project %q: invalid occurrence, fail open, would have failed with: %v", req.Occurrence, pID, err)
	}

	uID, err := g.Auth.EndUserID(ctx)
	if err != nil {
		return err
	}

	o, err := g.Storage.CreateOccurrence(ctx, pID, uID, req.Occurrence)
	if err != nil {
		return err
	}
	*resp = *o

	return nil
}

// BatchCreateOccurrences batch creates the specified occurrences.
func (g *API) BatchCreateOccurrences(ctx context.Context, req *gpb.BatchCreateOccurrencesRequest, resp *gpb.BatchCreateOccurrencesResponse) error {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", OccurrencesCreate); err != nil {
		return err
	}

	if len(req.Occurrences) == 0 {
		return errors.Newf(codes.InvalidArgument, "at least one occurrence must be specified")
	}
	if len(req.Occurrences) > maxBatchSize {
		return errors.Newf(codes.InvalidArgument, "%d is too many occurrence to batch create, a maximum of %d occurrence is allowed per batch create", len(req.Occurrences), maxBatchSize)
	}

	// Creating occurrences requires an additional notes attacher permissions check before we can
	// continue validation.
	authErrs := []error{}
	for i, o := range req.Occurrences {
		notePID, nID, err := name.ParseNote(o.NoteName)
		if err != nil {
			return err
		}
		if err := g.Auth.CheckAccessAndProject(ctx, notePID, nID, NotesAttachOccurrence); err != nil {
			authErrs = append(authErrs, fmt.Errorf("occurrences[%d]: %s", i, err))
		}
	}
	if len(authErrs) > 0 {
		return errors.Newf(codes.PermissionDenied, "one or more occurrences had auth errors, no occurrences were created: %v", authErrs)
	}

	validationErrs := []error{}
	for i, o := range req.Occurrences {
		if err := grafeas.ValidateOccurrence(o); err != nil {
			validationErrs = append(validationErrs, fmt.Errorf("occurrences[%d]: %v", i, err))
		}
	}
	if len(validationErrs) > 0 {
		if g.EnforceValidation {
			return errors.Newf(codes.InvalidArgument, "one or more occurrences are invalid, no occurrences were created: %v", validationErrs)
		}
		g.Logger.Warningf(ctx, "BatchCreateOccurrences %+v for project %q: invalid occurrences(s), fail open, would have failed with: %v", req.Occurrences, pID, validationErrs)
	}

	uID, err := g.Auth.EndUserID(ctx)
	if err != nil {
		return err
	}

	created, errs := g.Storage.BatchCreateOccurrences(ctx, pID, uID, req.Occurrences)
	resp.Occurrences = created
	if len(errs) != 0 {
		// Report any storage layer errors as invalid argument for now, find a better way to do this.
		return errors.Newf(codes.InvalidArgument, "errors encountered when batch creating occurrences: %d of %d occurrences failed: %v", len(errs), len(req.Occurrences), errs)
	}

	return nil
}

// UpdateOccurrence updates the specified occurrence.
func (g *API) UpdateOccurrence(ctx context.Context, req *gpb.UpdateOccurrenceRequest, resp *gpb.Occurrence) error {
	pID, oID, err := name.ParseOccurrence(req.Name)
	if err != nil {
		return err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if req.Occurrence == nil {
		return errors.Newf(codes.InvalidArgument, "an occurrence must be specified")
	}

	if err := g.Auth.CheckAccessAndProject(ctx, pID, oID, OccurrencesUpdate); err != nil {
		return err
	}
	notePID, nID, err := name.ParseNote(req.Occurrence.NoteName)
	if err != nil {
		return err
	}
	if err := g.Auth.CheckAccessAndProject(ctx, notePID, nID, NotesAttachOccurrence); err != nil {
		return err
	}

	o, err := g.Storage.UpdateOccurrence(ctx, pID, oID, req.Occurrence, req.UpdateMask)
	if err != nil {
		return err
	}
	*resp = *o

	return nil
}

// DeleteOccurrence deletes the specified occurrence.
func (g *API) DeleteOccurrence(ctx context.Context, req *gpb.DeleteOccurrenceRequest, _ *emptypb.Empty) error {
	pID, oID, err := name.ParseOccurrence(req.Name)
	if err != nil {
		return err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := g.Auth.CheckAccessAndProject(ctx, pID, oID, OccurrencesDelete); err != nil {
		return err
	}

	o, err := g.Storage.GetOccurrence(ctx, pID, oID)
	if err != nil {
		return err
	}
	notePID, nID, err := name.ParseNote(o.NoteName)
	if err != nil {
		return err
	}
	if notePID != "" && nID != "" {
		if err := g.Auth.CheckAccessAndProject(ctx, notePID, nID, NotesAttachOccurrence); err != nil {
			return err
		}
	}

	if err := g.Storage.DeleteOccurrence(ctx, pID, oID); err != nil {
		return err
	}

	// Purge any IAM policies set on this entity.
	if err := g.Auth.PurgePolicy(ctx, pID, oID, Occurrences); err != nil {
		// This fails open, should not block on policy deletion failure.
		g.Logger.Warningf(ctx, "Error deleting policies for occurrence %q in project %q: %v", oID, pID, err)
	}

	return nil
}

// ListNoteOccurrences lists occurrences for the specified note.
func (g *API) ListNoteOccurrences(ctx context.Context, req *gpb.ListNoteOccurrencesRequest, resp *gpb.ListNoteOccurrencesResponse) error {
	pID, nID, err := name.ParseNote(req.Name)
	if err != nil {
		return err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := g.Auth.CheckAccessAndProject(ctx, pID, nID, NotesListOccurrences); err != nil {
		return err
	}

	if err := g.Filter.Validate(req.Filter); err != nil {
		return err
	}

	occs, npt, err := g.Storage.ListNoteOccurrences(ctx, pID, nID, req.Filter, req.PageToken, req.PageSize)
	if err != nil {
		return err
	}
	resp.Occurrences = occs
	resp.NextPageToken = npt

	return nil
}

// GetVulnerabilityOccurrencesSummary produces a summary of vulnerability
// occurrences grouped by severity that match the specified filter.
func (g *API) GetVulnerabilityOccurrencesSummary(ctx context.Context, req *gpb.GetVulnerabilityOccurrencesSummaryRequest, resp *gpb.VulnerabilityOccurrencesSummary) error {
	pID, err := name.ParseProject(req.Parent)
	if err != nil {
		return err
	}

	ctx = g.Logger.PrepareCtx(ctx, pID)

	if err := g.Auth.CheckAccessAndProject(ctx, pID, "", OccurrencesList); err != nil {
		return err
	}

	if err := g.Filter.Validate(req.Filter); err != nil {
		return err
	}

	summary, err := g.Storage.GetVulnerabilityOccurrencesSummary(ctx, pID, req.Filter)
	if err != nil {
		return err
	}
	*resp = *summary
	return nil
}
