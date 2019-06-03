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

package server

import (
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
)

// Storager is the interface that a Grafeas storage implementation would provide
type Storager interface {
	// CreateProject adds the specified project
	CreateProject(pID string) error

	// CreateNote adds the specified note
	CreateNote(n *pb.Note) error

	// CreateOccurrence adds the specified occurrence
	CreateOccurrence(o *pb.Occurrence) error

	// DeleteNote deletes the project with the given pID
	DeleteProject(pID string) error

	// DeleteNote deletes the note with the given pID and nID
	DeleteNote(pID, nID string) error

	// DeleteOccurrence deletes the occurrence with the given pID and oID
	DeleteOccurrence(pID, oID string) error

	// GetProject returns the project with the given pID
	GetProject(pID string) (*prpb.Project, error)

	// GetNote returns the note with project (pID) and note ID (nID)
	GetNote(pID, nID string) (*pb.Note, error)

	// GetNoteByOccurrence returns the note attached to occurrence with pID and oID
	GetNoteByOccurrence(pID, oID string) (*pb.Note, error)

	// GetOccurrence returns the occurrence with pID and oID
	GetOccurrence(pID, oID string) (*pb.Occurrence, error)

	// ListProjects returns up to pageSize number of projects beginning at pageToken (or from
	// start if pageToken is the empty string).
	ListProjects(filter string, pageSize int, pageToken string) ([]*prpb.Project, string, error)

	// ListNoteOccurrences returns up to pageSize number of occcurrences on the particular note (nID)
	// for this project (pID) projects beginning at pageToken (or from start if pageToken is the empty string).
	ListNoteOccurrences(pID, nID, filters string, pageSize int, pageToken string) ([]*pb.Occurrence, string, error)

	// ListNotes returns up to pageSize number of notes for this project (pID) beginning
	// at pageToken (or from start if pageToken is the empty string).
	ListNotes(pID, filters string, pageSize int, pageToken string) ([]*pb.Note, string, error)

	// ListOccurrences returns up to pageSize number of occurrences for this project (pID) beginning
	// at pageToken (or from start if pageToken is the empty string).
	ListOccurrences(pID, filters string, pageSize int, pageToken string) ([]*pb.Occurrence, string, error)

	// UpdateNote updates the existing note with the given pID and nID
	UpdateNote(pID, nID string, n *pb.Note) error

	// UpdateOccurrence updates the existing occurrence with the given projectID and occurrenceID
	UpdateOccurrence(pID, oID string, o *pb.Occurrence) error
}
