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

// Package provenance implements functions to validate that the fields of provenance entities being
// passed into the API meet our requirements.
package provenance

import (
	"errors"
	"fmt"

	ppb "github.com/grafeas/grafeas/proto/v1beta1/provenance_go_proto"
)

// ValidateBuildProvenance validates that a build provenance has all its required fields filled in.
func ValidateBuildProvenance(p *ppb.BuildProvenance) []error {
	errs := []error{}

	if p.GetId() == "" {
		errs = append(errs, errors.New("id is required"))
	}

	for i, c := range p.GetCommands() {
		if c == nil {
			errs = append(errs, fmt.Errorf("commands[%d] command cannot be null", i))
		} else {
			for _, err := range validateCommand(c) {
				errs = append(errs, fmt.Errorf("commands[%d].%s", i, err))
			}
		}
	}

	for i, a := range p.GetBuiltArtifacts() {
		if a == nil {
			errs = append(errs, fmt.Errorf("built_artifacts[%d] command cannot be null", i))
		} else {
			for _, err := range validateArtifact(a) {
				errs = append(errs, fmt.Errorf("built_artifacts[%d].%s", i, err))
			}
		}
	}

	if s := p.GetSourceProvenance(); s != nil {
		for _, err := range validateSource(s) {
			errs = append(errs, fmt.Errorf("source_provenance.%s", err))
		}
	}

	return errs
}

func validateCommand(c *ppb.Command) []error {
	errs := []error{}

	if c.GetName() == "" {
		errs = append(errs, errors.New("name is required"))
	}

	return errs
}

func validateArtifact(a *ppb.Artifact) []error {
	errs := []error{}
	return errs
}

func validateSource(s *ppb.Source) []error {
	errs := []error{}

	for filePath, fileHashes := range s.GetFileHashes() {
		if fileHashes == nil {
			errs = append(errs, fmt.Errorf("file_hashes[%q] file hashes cannot be null", filePath))
		} else {
			for _, err := range validateFileHashes(fileHashes) {
				errs = append(errs, fmt.Errorf("file_hashes[%q].%s", filePath, err))
			}
		}
	}

	return errs
}

func validateFileHashes(fileHashes *ppb.FileHashes) []error {
	errs := []error{}

	if fileHash := fileHashes.GetFileHash(); fileHash == nil {
		errs = append(errs, errors.New("file_hash is required"))
	} else if len(fileHash) == 0 {
		errs = append(errs, errors.New("file_hash requires at least 1 element"))
	} else {
		for i, h := range fileHash {
			if h == nil {
				errs = append(errs, fmt.Errorf("file_hash[%d] hash cannot be null", i))
			} else {
				for _, err := range validateHash(h) {
					errs = append(errs, fmt.Errorf("file_hash[%d].%s", i, err))
				}
			}
		}
	}

	return errs
}

func validateHash(h *ppb.Hash) []error {
	errs := []error{}

	if h.GetType() == ppb.Hash_HASH_TYPE_UNSPECIFIED {
		errs = append(errs, errors.New("type is required"))
	}

	if h.GetValue() == nil {
		errs = append(errs, errors.New("value is required"))
	}

	return errs
}
