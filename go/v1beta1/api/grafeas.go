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

// Package grafeas is an implementation of the v1beta1 https://github.com/grafeas/grafeas/ API.
package grafeas

import (
	"github.com/grafeas/grafeas/go/iam"
	gpb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	"golang.org/x/net/context"
	fieldmaskpb "google.golang.org/genproto/protobuf/field_mask"
)

const (
	defaultPageSize = 20
	maxPageSize     = 1000
	maxBatchSize    = 1000

	// NotesGet is the permission to get a note.
	NotesGet = iam.Permission("notes.get")
	// NotesList is the permission to list notes.
	NotesList = iam.Permission("notes.list")
	// NotesCreate is the permission to create a note.
	NotesCreate = iam.Permission("notes.create")
	// NotesUpdate is the permission to update a note.
	NotesUpdate = iam.Permission("notes.update")
	// NotesDelete is the permission to delete a note.
	NotesDelete = iam.Permission("notes.delete")

	// OccurrencesGet is the permission to get an occurrence.
	OccurrencesGet = iam.Permission("occurrences.get")
	// OccurrencesList is the permission to list occurrences.
	OccurrencesList = iam.Permission("occurrences.list")
	// OccurrencesCreate is the permission to create an occurrence.
	OccurrencesCreate = iam.Permission("occurrences.create")
	// OccurrencesUpdate is the permission to update an occurrence.
	OccurrencesUpdate = iam.Permission("occurrences.update")
	// OccurrencesDelete is the permission to delete an occurrence.
	OccurrencesDelete = iam.Permission("occurrences.delete")

	// NotesListOccurrences is the permission to list occurrences associated for a note you own.
	NotesListOccurrences = iam.Permission("notes.listOccurrences")
	// NotesAttachOccurrence is the permission to attach occurrences for a note you own.
	NotesAttachOccurrence = iam.Permission("notes.attachOccurrence")

	// Notes is the resource type for notes.
	Notes = iam.Resource("notes")
	// Occurrences is the resource type for occurrences.
	Occurrences = iam.Resource("occurrences")
)

// Storage provides storage functions for this API.
type Storage interface {
	// GetOccurrence gets the specified occurrence from storage.
	GetOccurrence(ctx context.Context, projectID, occID string) (*gpb.Occurrence, error)
	// ListOccurrences lists occurrences for the specified project from storage.
	ListOccurrences(ctx context.Context, projectID, filter, pageToken string, pageSize int32) ([]*gpb.Occurrence, string, error)
	// CreateOccurrence creates the specified occurrence in storage.
	CreateOccurrence(ctx context.Context, projectID, userID string, o *gpb.Occurrence) (*gpb.Occurrence, error)
	// BatchCreateOccurrences batch creates the specified occurrences in storage.
	BatchCreateOccurrences(ctx context.Context, projectID string, userID string, occs []*gpb.Occurrence) ([]*gpb.Occurrence, []error)
	// UpdateOccurrence updates the specified occurrence in storage.
	UpdateOccurrence(ctx context.Context, projectID, occID string, o *gpb.Occurrence, mask *fieldmaskpb.FieldMask) (*gpb.Occurrence, error)
	// DeleteOccurrence deletes the specified occurrence in storage.
	DeleteOccurrence(ctx context.Context, projectID, occID string) error

	// GetNote gets the specified note from storage.
	GetNote(ctx context.Context, projectID, nID string) (*gpb.Note, error)
	// ListNotes lists notes for the specified project from storage.
	ListNotes(ctx context.Context, projectID, filter, pageToken string, pageSize int32) ([]*gpb.Note, string, error)
	// CreateNote creates the specified note in storage.
	CreateNote(ctx context.Context, projectID, nID string, userID string, n *gpb.Note) (*gpb.Note, error)
	// BatchCreateNotes batch creates the specified notes in storage.
	BatchCreateNotes(ctx context.Context, projectID string, userID string, notes map[string]*gpb.Note) ([]*gpb.Note, []error)
	// UpdateNote updates the specified note in storage.
	UpdateNote(ctx context.Context, projectID, nID string, n *gpb.Note, mask *fieldmaskpb.FieldMask) (*gpb.Note, error)
	// DeleteNote deletes the specified note in storage.
	DeleteNote(ctx context.Context, projectID, nID string) error

	// GetOccurrenceNote gets the note for the specified occurrence from storage.
	GetOccurrenceNote(ctx context.Context, projectID, oID string) (*gpb.Note, error)
	// ListNoteOccurrences lists occurrences for the specified note from storage.
	ListNoteOccurrences(ctx context.Context, projectID, nID, filter, pageToken string, pageSize int32) ([]*gpb.Occurrence, string, error)
	// GetVulnerabilityOccurrencesSummary gets a summary of vulnerability occurrences from storage.
	GetVulnerabilityOccurrencesSummary(ctx context.Context, projectID, filter string) (*gpb.VulnerabilityOccurrencesSummary, error)
}

// Auth provides authorization functions for this API.
type Auth interface {
	// CheckAccessAndProject checks to see whether an API call is allowed. It can check things like
	// whether the project and entity exists, and whether the user has access to the project and the
	// entity, This method should generally be the first thing an API method calls before taking any
	// action. To prevent leaking information, the only kind of error this should return is a
	// permission denied error.
	CheckAccessAndProject(ctx context.Context, projectID string, entityID string, p iam.Permission) error
	// EndUserID returns the ID of the user making an API call.
	EndUserID(ctx context.Context) (string, error)
	// PurgePolicy purges any auth policies that may have been set on the specified entity if using an
	// IAM service.
	PurgePolicy(ctx context.Context, projectID string, entityID string, r iam.Resource) error
}

// Filter provides functions for parsing filter strings.
type Filter interface {
	// Validate determines whether the specified filter string is a valid filter.
	Validate(f string) error
}

// Logger provides functions for logging at various levels.
type Logger interface {
	// PrepareCtx adds values to the context for logging if necessary.
	PrepareCtx(ctx context.Context, projectID string) context.Context
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})
	Warning(ctx context.Context, args ...interface{})
	Warningf(ctx context.Context, format string, args ...interface{})
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
}

// API implements the methods in the v1beta1 Grafeas API.
type API struct {
	Storage           Storage
	Auth              Auth
	Filter            Filter
	Logger            Logger
	EnforceValidation bool
}
