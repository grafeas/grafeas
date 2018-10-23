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

package storage

import (
	"fmt"
	"sort"
	"strings"

	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/name"
	"github.com/grafeas/grafeas/server-go"
	opspb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

type EmbeddedStoreConfig struct {
	Path string `yaml:"path"` // Path is the folder path to storage files
}

// embeddedStore is a storage solution for Grafeas based on boltdb
type embeddedStore struct {
	db *bolt.DB
}

// NewEmbeddedStore creates a embeddedS store with initialized filesystem
func NewEmbeddedStore(config *EmbeddedStoreConfig) server.Storager {
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
	return &embeddedStore{db: db}
}

// CreateProject adds the specified project to the embedded store
func (m *embeddedStore) CreateProject(pID string) error {
	err := m.update(bucketProjects, pID, true, &prpb.Project{Name: name.FormatProject(pID)})
	if err == errKeyExists {
		return status.Errorf(codes.AlreadyExists, "Project with name %q already exists", pID)
	}
	return err
}

// DeleteProject deletes the project with the given pID from the embedded store
func (m *embeddedStore) DeleteProject(pID string) error {
	err := m.delete(bucketProjects, pID)
	if err == errNoKey {
		return status.Errorf(codes.NotFound, "Project with name %q does not Exist", pID)
	}
	return err
}

// GetProject returns the project with the given pID from the embedded store
func (m *embeddedStore) GetProject(pID string) (*prpb.Project, error) {
	var project prpb.Project
	err := m.get(bucketProjects, pID, &project)
	if err == errNoKey {
		return nil, status.Errorf(codes.NotFound, "Project with name %q does not Exist", pID)
	}
	return &project, err
}

// ListProjects returns up to pageSize number of projects beginning at pageToken (or from
// start if pageToken is the empty string).
func (m *embeddedStore) ListProjects(filter string, pageSize int, pageToken string) ([]*prpb.Project, string, error) {
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

// CreateOccurrence adds the specified occurrence to the embedded store
func (m *embeddedStore) CreateOccurrence(o *pb.Occurrence) error {
	err := m.update(bucketOccurrences, o.Name, true, o)
	if err == errKeyExists {
		return status.Errorf(codes.AlreadyExists, "Occurrence with name %q already exists", o.Name)
	}
	return err

}

// DeleteOccurrence deletes the occurrence with the given pID and oID from the embedded store
func (m *embeddedStore) DeleteOccurrence(pID, oID string) error {
	oName := name.OccurrenceName(pID, oID)
	err := m.delete(bucketOccurrences, oName)
	if err == errNoKey {
		return status.Errorf(codes.NotFound, "Occurrence with oName %q does not Exist", oName)
	}
	return err
}

// UpdateOccurrence updates the existing occurrence with the given projectID and occurrenceID
func (m *embeddedStore) UpdateOccurrence(pID, oID string, o *pb.Occurrence) error {
	oName := name.OccurrenceName(pID, oID)
	err := m.update(bucketOccurrences, oName, false, o)
	if err == errNoKey {
		return status.Errorf(codes.NotFound, "Occurrence with name %q does not Exist", oName)
	}
	return err
}

// GetOccurrence returns the occurrence with pID and oID
func (m *embeddedStore) GetOccurrence(pID, oID string) (*pb.Occurrence, error) {
	oName := name.OccurrenceName(pID, oID)
	var o pb.Occurrence
	err := m.get(bucketOccurrences, oName, &o)
	if err == errNoKey {
		return nil, status.Errorf(codes.NotFound, "Occurrence with name %q does not exist", oName)
	}
	return &o, err
}

// ListOccurrences returns up to pageSize number of occurrences for this project (pID) beginning
// at pageToken (or from start if pageToken is the empty string).
func (m *embeddedStore) ListOccurrences(pID, filters string, pageSize int, pageToken string) ([]*pb.Occurrence, string, error) {
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
	endPos := min(startPos+pageSize, len(os))
	return os[startPos:endPos], nextPageToken(endPos, len(os)), nil
}

// CreateNote adds the specified note to the embedded store
func (m *embeddedStore) CreateNote(n *pb.Note) error {
	err := m.update(bucketNotes, n.Name, true, n)
	if err == errKeyExists {
		return status.Errorf(codes.AlreadyExists, "Note with name %q already exists", n.Name)
	}
	return err
}

// DeleteNote deletes the note with the given pID and nID from the embedded store
func (m *embeddedStore) DeleteNote(pID, nID string) error {
	nName := name.NoteName(pID, nID)
	err := m.delete(bucketNotes, nName)
	if err == errNoKey {
		return status.Errorf(codes.NotFound, "Note with name %q does not Exist", nName)
	}
	return err
}

// UpdateNote updates the existing note with the given pID and nID
func (m *embeddedStore) UpdateNote(pID, nID string, n *pb.Note) error {
	nName := name.NoteName(pID, nID)
	err := m.update(bucketNotes, nName, false, n)
	if err == errNoKey {
		return status.Errorf(codes.NotFound, "Note with name %q does not Exist", nName)
	}
	return err
}

// GetNote returns the note with pID and nID
func (m *embeddedStore) GetNote(pID, nID string) (*pb.Note, error) {
	nName := name.NoteName(pID, nID)
	var n pb.Note
	err := m.get(bucketNotes, nName, &n)
	if err == errNoKey {
		return nil, status.Errorf(codes.NotFound, "Note with name %q does not Exist", nName)
	}
	return &n, err
}

// GetNoteByOccurrence returns the note attached to occurrence with pID and oID
func (m *embeddedStore) GetNoteByOccurrence(pID, oID string) (*pb.Note, error) {
	o, err := m.GetOccurrence(pID, oID)
	if err != nil {
		return nil, err
	}
	var n pb.Note
	err = m.get(bucketNotes, o.NoteName, &n)
	if err == errNoKey {
		return nil, status.Errorf(codes.NotFound, "Note with name %q does not Exist", o.NoteName)
	}
	return &n, err
}

// ListNotes returns up to pageSize number of notes for this project (pID) beginning
// at pageToken (or from start if pageToken is the empty string).
func (m *embeddedStore) ListNotes(pID, filters string, pageSize int, pageToken string) ([]*pb.Note, string, error) {
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
	endPos := min(startPos+pageSize, len(ns))
	return ns[startPos:endPos], nextPageToken(endPos, len(ns)), nil
}

// ListNoteOccurrences returns up to pageSize number of occcurrences on the particular note (nID)
// for this project (pID) projects beginning at pageToken (or from start if pageToken is the empty string).
func (m *embeddedStore) ListNoteOccurrences(pID, nID, filters string, pageSize int, pageToken string) ([]*pb.Occurrence, string, error) {
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
	endPos := min(startPos+pageSize, len(os))
	return os[startPos:endPos], nextPageToken(endPos, len(os)), nil
}

// GetOperation returns the operation with pID and oID
func (m *embeddedStore) GetOperation(pID, opID string) (*opspb.Operation, error) {
	oName := name.OperationName(pID, opID)
	var o opspb.Operation
	err := m.get(bucketOperations, oName, &o)
	if err == errNoKey {
		return nil, status.Errorf(codes.NotFound, "Operation with name %q does not Exist", oName)
	}
	return &o, err
}

// CreateOperation adds the specified operation to the embedded store
func (m *embeddedStore) CreateOperation(o *opspb.Operation) error {
	err := m.update(bucketOperations, o.Name, true, o)
	if err == errKeyExists {
		return status.Errorf(codes.AlreadyExists, "Operation with name %q already exists", o.Name)
	}
	return err
}

// DeleteOperation deletes the operation with the given pID and oID from the embeddedStore
func (m *embeddedStore) DeleteOperation(pID, opID string) error {
	opName := name.OperationName(pID, opID)
	err := m.delete(bucketOperations, opName)
	if err == errNoKey {
		return status.Errorf(codes.NotFound, "Operation with name %q does not Exist", opName)
	}
	return err
}

// UpdateOperation updates the existing operation with the given pID and nID
func (m *embeddedStore) UpdateOperation(pID, opID string, op *opspb.Operation) error {
	opName := name.OperationName(pID, opID)
	err := m.update(bucketOperations, opName, false, op)
	if err == errNoKey {
		return status.Errorf(codes.NotFound, "Operation with name %q does not Exist", opName)
	}
	return err
}

// ListOperations returns up to pageSize number of operations for this project (pID) beginning
// at pageToken (or from start if pageToken is the empty string).
func (m *embeddedStore) ListOperations(pID, filters string, pageSize int, pageToken string) ([]*opspb.Operation, string, error) {
	var os []*opspb.Operation
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketOperations))
		err := b.ForEach(func(k, v []byte) error {
			var o opspb.Operation
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
	endPos := min(startPos+pageSize, len(os))
	return os[startPos:endPos], nextPageToken(endPos, len(os)), nil
}

func (m *embeddedStore) update(bucket string, key string, new bool, pb proto.Message) error {
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

func (m *embeddedStore) get(bucket string, key string, pb proto.Message) error {
	return m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value := b.Get([]byte(key))
		if value == nil {
			return errNoKey
		}
		return proto.Unmarshal(value, pb)
	})
}

func (m *embeddedStore) delete(bucket string, key string) error {
	return m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		value := b.Get([]byte(key))
		if value == nil {
			return errNoKey
		}
		return b.Delete([]byte(key))
	})
}
