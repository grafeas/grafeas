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
	server "github.com/grafeas/grafeas/server-go"

	"io/ioutil"
	"path/filepath"
	"strconv"
	"sync/atomic"
	"testing"
)

func TestEmbeddedStore(t *testing.T) {
	dir, err := ioutil.TempDir("", "embeddedstore")
	if err != nil {
		t.Fatalf("ioutil.TempDir failed %v", err)
	}
	var instance int32
	doTestStorager(t, func(t *testing.T) (server.Storager, func()) {
		testDir := filepath.Join(dir, strconv.Itoa(int(atomic.AddInt32(&instance, 1))))
		return NewEmbeddedStore(&EmbeddedStoreConfig{Path: testDir}), func() {}
	})
}
