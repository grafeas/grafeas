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
	"strings"

	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/grafeas/grafeas/go/config"
	"github.com/grafeas/grafeas/go/errors"
	"github.com/grafeas/grafeas/go/name"
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	"golang.org/x/net/context"
	fieldmaskpb "google.golang.org/genproto/protobuf/field_mask"
	"google.golang.org/grpc/codes"
	"log"
	"os"
	"path/filepath"
)

const (
	bucketOccurrences = "occurrences"
	bucketProjects    = "projects"
	bucketNotes       = "notes"
	bucketOperations  = "operations"
)

var (
	errKeyExists = fmt.Errorf("key exists")
	errNoKey     = fmt.Errorf("key missing")
)

// EmbeddedStore is a storage solution for Grafeas based on boltdb
type EmbeddedStore struct {
	db *bolt.DB
}

// NewEmbeddedStore creates a embeddedS store with initialized filesystem
func NewEmbeddedStore(config *config.EmbeddedStoreConfig) *EmbeddedStore {
	if err := os.MkdirAll(config.Path, 0700); err != nil {
		log.Fatalf("Failed to create config directory %v", err)
	}
	db, err := bolt.Open(filepath.Join(config.Path, "grafeas.db"), 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketOccurrences)); err != nil {
			return err
		}
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketProjects)); err != nil {
			return err
		}
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketNotes)); err != nil {
			return err
		}
		if _, err := tx.CreateBucketIfNotExists([]byte(bucketOperations)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}
	return &EmbeddedStore{db: db}
}

// CreateProject creates the specified project in embedded store.
func (m *EmbeddedStore) CreateProject(ctx context.Context, pID string, p *prpb.Project) (*prpb.Project, error) {
	err := m.update(bucketProjects, pID, true, p)
	if err == errKeyExists {
		return nil, errors.Newf(codes.AlreadyExists, "Project with name %q already exists", pID)
	}
	return p, err
}

// GetProject gets the specified project from embedded store.
func (m *EmbeddedStore) GetProject(ctx context.Context, pID string) (*prpb.Project, error) {
	var project prpb.Project
	err := m.get(bucketProjects, pID, &project)
	if err == errNoKey {
		return nil, errors.Newf(codes.NotFound, "Project with name %q does not Exist", pID)
	}
	return &project, err
}

// ListProjects returns up to pageSize number of projects beginning at pageToken, or from
// start if pageToken is the empty string.
func (m *EmbeddedStore) ListProjects(ctx context.Context, filter string, pageSize int, pageToken string) ([]*prpb.Project, string, error) {
	var projects []*prpb.Project
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketProjects))
		err := b.ForEach(func(k, v []byte) error {
			var project prpb.Project
			if err := proto.Unmarshal(v, &project); err != nil {
				return err
			}
			projects = append(projects, &project)
			return nil
		})
		return err
	})
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].Name < projects[j].Name
	})
	startPos := parsePageToken(pageToken, 0)
	endPos := min(startPos+pageSize, len(projects))
	return projects[startPos:endPos], nextPageToken(endPos, len(projects)), nil
}

// DeleteProject deletes the specified project from embedded store.
func (m *EmbeddedStore) DeleteProject(ctx context.Context, pID string) error {
	err := m.delete(bucketProjects, pID)
	if err == errNoKey {
		return errors.Newf(codes.NotFound, "Project with name %q does not Exist", pID)
	}
	return err
}

// GetOccurrence gets the specified occurrence from embedded store.
func (m *EmbeddedStore) GetOccurrence(ctx context.Context, pID, oID string) (*pb.Occurrence, error) {
	oName := name.FormatOccurrence(pID, oID)
	var o pb.Occurrence
	err := m.get(bucketOccurrences, oName, &o)
	if err == errNoKey {
		return nil, errors.Newf(codes.NotFound, "Occurrence with name %q does not exist", oName)
	}

	// Set the output-only field before returning
	o.Name = name.FormatOccurrence(pID, oID)
	return &o, err
}

// ListOccurrences returns up to pageSize number of occurrences for this project (pID) beginning
// at pageToken (or from start if pageToken is the empty string).
func (m *EmbeddedStore) ListOccurrences(ctx context.Context, pID, filters, pageToken string, pageSize int32) ([]*pb.Occurrence, string, error) {
	var os []*pb.Occurrence
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketOccurrences))
		err := b.ForEach(func(k, v []byte) error {
			var o pb.Occurrence
			if err := proto.Unmarshal(v, &o); err != nil {
				return err
			}
			if strings.HasPrefix(o.Name, fmt.Sprintf("projects/%v", pID)) {
				os = append(os, &o)
			}
			return nil
		})
		return err
	})
	sort.Slice(os, func(i, j int) bool {
		return os[i].Name < os[j].Name
	})
	startPos := parsePageToken(pageToken, 0)
	endPos := min(startPos+int(pageSize), len(os))
	return os[startPos:endPos], nextPageToken(endPos, len(os)), nil
}

// CreateOccurrence creates the specified occurrence in embedded store.
func (m *EmbeddedStore) CreateOccurrence(ctx context.Context, pID, uID string, o *pb.Occurrence) (*pb.Occurrence, error) {
	o = proto.Clone(o).(*pb.Occurrence)

	if err := m.get(bucketOccurrences, o.Name, &pb.Occurrence{}); err == errNoKey {
		o.CreateTime = ptypes.TimestampNow()
		err := m.update(bucketOccurrences, o.Name, true, o)
		return o, err
	}

	return nil, errors.Newf(codes.AlreadyExists, "Occurrence with name %q already exists", o.Name)
}

// BatchCreateOccurrence batch creates the specified occurrences in embedded store.
func (m *EmbeddedStore) BatchCreateOccurrences(ctx context.Context, pID string, uID string, occs []*pb.Occurrence) ([]*pb.Occurrence, []error) {
	clonedOccs := []*pb.Occurrence{}
	for _, o := range occs {
		clonedOccs = append(clonedOccs, proto.Clone(o).(*pb.Occurrence))
	}
	occs = clonedOccs

	errs := []error{}
	created := []*pb.Occurrence{}
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

// UpdateOccurrence updates the specified occurrence in embedded store.
func (m *EmbeddedStore) UpdateOccurrence(ctx context.Context, pID, oID string, o *pb.Occurrence, mask *fieldmaskpb.FieldMask) (*pb.Occurrence, error) {
	o = proto.Clone(o).(*pb.Occurrence)
	oName := name.FormatOccurrence(pID, oID)
	// TODO(#312): implement the update operation
	o.UpdateTime = ptypes.TimestampNow()

	err := m.update(bucketOccurrences, oName, false, o)
	if err == errNoKey {
		return nil, errors.Newf(codes.NotFound, "Occurrence with name %q does not Exist", oName)
	}
	return o, err
}

// DeleteOccurrence deletes the specified occurrence in embedded store.
func (m *EmbeddedStore) DeleteOccurrence(ctx context.Context, pID, oID string) error {
	oName := name.FormatOccurrence(pID, oID)
	err := m.delete(bucketOccurrences, oName)
	if err == errNoKey {
		return errors.Newf(codes.NotFound, "Occurrence with oName %q does not Exist", oName)
	}
	return err
}

// GetNote gets the specified note from embedded store.
func (m *EmbeddedStore) GetNote(ctx context.Context, pID, nID string) (*pb.Note, error) {
	nName := name.FormatNote(pID, nID)
	var n pb.Note
	err := m.get(bucketNotes, nName, &n)
	if err == errNoKey {
		return nil, errors.Newf(codes.NotFound, "Note with name %q does not Exist", nName)
	}

	// Set the output-only field before returning
	n.Name = name.FormatNote(pID, nID)
	return &n, err
}

// ListNotes returns up to pageSize number of notes for the project beginning
// at pageToken, or from start if pageToken is the empty string.
func (m *EmbeddedStore) ListNotes(ctx context.Context, pID, filter, pageToken string, pageSize int32) ([]*pb.Note, string, error) {
	var ns []*pb.Note
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketNotes))
		err := b.ForEach(func(k, v []byte) error {
			var n pb.Note
			if err := proto.Unmarshal(v, &n); err != nil {
				return err
			}
			if strings.HasPrefix(n.Name, fmt.Sprintf("projects/%v", pID)) {
				ns = append(ns, &n)
			}
			return nil
		})
		return err
	})
	sort.Slice(ns, func(i, j int) bool {
		return ns[i].Name < ns[j].Name
	})
	startPos := parsePageToken(pageToken, 0)
	endPos := min(startPos+int(pageSize), len(ns))
	return ns[startPos:endPos], nextPageToken(endPos, len(ns)), nil
}

// CreateNote creates the specified note in embedded store.
func (m *EmbeddedStore) CreateNote(ctx context.Context, pID, nID, uID string, n *pb.Note) (*pb.Note, error) {
	n = proto.Clone(n).(*pb.Note)

	if err := m.get(bucketNotes, n.Name, &pb.Note{}); err == errNoKey {
		n.CreateTime = ptypes.TimestampNow()
		err := m.update(bucketNotes, n.Name, true, n)
		return n, err
	}
	return nil, errors.Newf(codes.AlreadyExists, "Note with name %q already exists", n.Name)
}

// BatchCreateNotes batch creates the specified notes in embedded store.
func (m *EmbeddedStore) BatchCreateNotes(ctx context.Context, pID, uID string, notes map[string]*pb.Note) ([]*pb.Note, []error) {
	clonedNotes := map[string]*pb.Note{}
	for nID, n := range notes {
		clonedNotes[nID] = proto.Clone(n).(*pb.Note)
	}
	notes = clonedNotes

	errs := []error{}
	created := []*pb.Note{}
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

// UpdateNote updates the specified note in embedded store.
func (m *EmbeddedStore) UpdateNote(ctx context.Context, pID, nID string, n *pb.Note, mask *fieldmaskpb.FieldMask) (*pb.Note, error) {
	n = proto.Clone(n).(*pb.Note)
	// TODO(#312): implement the update operation
	n.UpdateTime = ptypes.TimestampNow()
	nName := name.FormatNote(pID, nID)
	n.Name = nName

	err := m.update(bucketNotes, nName, false, n)
	if err == errNoKey {
		return nil, errors.Newf(codes.NotFound, "Note with name %q does not Exist", nName)
	}
	return n, err
}

// DeleteNote deletes the specified note in embedded store.
func (m *EmbeddedStore) DeleteNote(ctx context.Context, pID, nID string) error {
	nName := name.FormatNote(pID, nID)
	err := m.delete(bucketNotes, nName)
	if err == errNoKey {
		return errors.Newf(codes.NotFound, "Note with name %q does not Exist", nName)
	}
	return err
}

// GetOccurrenceNote gets the note for the specified occurrence from embedded store.
func (m *EmbeddedStore) GetOccurrenceNote(ctx context.Context, pID, oID string) (*pb.Note, error) {
	o, err := m.GetOccurrence(ctx, pID, oID)
	if err != nil {
		return nil, err
	}
	var n pb.Note
	err = m.get(bucketNotes, o.NoteName, &n)
	if err == errNoKey {
		return nil, errors.Newf(codes.NotFound, "Note with name %q does not Exist", o.NoteName)
	}
	return &n, err
}

// ListNoteOccurrences returns up to pageSize number of occurrences on the note
// for the project beginning at pageToken, or from start if pageToken is empty.
func (m *EmbeddedStore) ListNoteOccurrences(ctx context.Context, pID, nID, filter, pageToken string, pageSize int32) ([]*pb.Occurrence, string, error) {
	// TODO: use filters
	nName := name.FormatNote(pID, nID)
	var os []*pb.Occurrence
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketOccurrences))
		err := b.ForEach(func(k, v []byte) error {
			var o pb.Occurrence
			if err := proto.Unmarshal(v, &o); err != nil {
				return err
			}
			if o.NoteName == nName {
				os = append(os, &o)
			}
			return nil
		})
		return err
	})
	sort.Slice(os, func(i, j int) bool {
		return os[i].Name < os[j].Name
	})
	startPos := parsePageToken(pageToken, 0)
	endPos := min(startPos+int(pageSize), len(os))
	return os[startPos:endPos], nextPageToken(endPos, len(os)), nil
}

// GetVulnerabilityOccurrencesSummary gets a summary of vulnerability occurrences from storage.
func (m *EmbeddedStore) GetVulnerabilityOccurrencesSummary(ctx context.Context, projectID, filter string) (*pb.VulnerabilityOccurrencesSummary, error) {
	return &pb.VulnerabilityOccurrencesSummary{}, nil
}

func (m *EmbeddedStore) update(bucket string, key string, new bool, pb proto.Message) error {
	return m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value := b.Get([]byte(key))
		if new && value != nil {
			return errKeyExists
		} else if !new && value == nil {
			return errNoKey
		}
		buf, err := proto.Marshal(pb)
		if err != nil {
			return err
		}
		return b.Put([]byte(key), buf)
	})
}

func (m *EmbeddedStore) get(bucket string, key string, pb proto.Message) error {
	return m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value := b.Get([]byte(key))
		if value == nil {
			return errNoKey
		}
		return proto.Unmarshal(value, pb)
	})
}

func (m *EmbeddedStore) delete(bucket string, key string) error {
	return m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value := b.Get([]byte(key))
		if value == nil {
			return errNoKey
		}
		return b.Delete([]byte(key))
	})
}
