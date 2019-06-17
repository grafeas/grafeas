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

package storage_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"sync/atomic"
	"testing"

	"github.com/grafeas/grafeas/go/v1beta1/api"
	"github.com/grafeas/grafeas/go/v1beta1/project"
	"github.com/grafeas/grafeas/go/v1beta1/storage"
)

func TestBetaEmbeddedStore(t *testing.T) {
	dir, err := ioutil.TempDir("", "embeddedstore")
	if err != nil {
		t.Fatalf("ioutil.TempDir failed %v", err)
	}
	// clean up
	defer os.RemoveAll(dir)

	var instance int32
	doTestStorage(t, func(t *testing.T) (grafeas.Storage, project.Storage, func()) {
		testDir := filepath.Join(dir, strconv.Itoa(int(atomic.AddInt32(&instance, 1))))
		s := storage.NewEmbeddedStore(&storage.EmbeddedStoreConfig{Path: testDir})
		var g grafeas.Storage = s
		var gp project.Storage = s
		return g, gp, func() {}
	})
}
