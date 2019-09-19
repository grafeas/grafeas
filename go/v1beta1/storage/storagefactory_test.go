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
	"testing"
)

func TestRegisterStorageTypeProvider_AddsProviders(t *testing.T) {
	// clear global variable before test
	for k := range registeredStorageTypeProviders {
		delete(registeredStorageTypeProviders, k)
	}

	numProviders := len(registeredStorageTypeProviders)
	if numProviders != 0 {
		t.Errorf("Expected 0 storage providers at start of test, got %d", numProviders)
	}

	err := RegisterStorageTypeProvider("p1", func(storageType string, storageConfig *interface{}) (storage *Storage, e error) {
		return nil, nil
	})

	if err != nil {
		t.Errorf("Error adding provider, %s", err)
	}

	err = RegisterStorageTypeProvider("p2", func(storageType string, storageConfig *interface{}) (storage *Storage, e error) {
		return nil, nil
	})

	if err != nil {
		t.Errorf("Error adding provider, %s", err)
	}

	numProviders = len(registeredStorageTypeProviders)
	if numProviders != 2 {
		t.Errorf("Expected 2 storage providers at start of test, got %d", numProviders)
	}

}

// test variable used to confirm that storage provider function is called
var providerExecutionTestVariable int

func TestCreateStorageOfType_CorrectProviderIsCalled(t *testing.T) {
	// clear global variables before test
	for k := range registeredStorageTypeProviders {
		delete(registeredStorageTypeProviders, k)
	}
	providerExecutionTestVariable = 0

	err := RegisterStorageTypeProvider("p1", func(storageType string, storageConfig *interface{}) (storage *Storage, e error) {
		providerExecutionTestVariable = 1
		return nil, nil
	})

	if err != nil {
		t.Errorf("Error adding provider, %s", err)
	}

	err = RegisterStorageTypeProvider("p2", func(storageType string, storageConfig *interface{}) (storage *Storage, e error) {
		providerExecutionTestVariable = 2
		return nil, nil
	})

	if err != nil {
		t.Errorf("Error adding provider, %s", err)
	}

	_, err = CreateStorageOfType("p1", nil)
	if providerExecutionTestVariable != 1 {
		t.Errorf("Provider 'p1' not called")
	}

	_, err = CreateStorageOfType("p2", nil)
	if providerExecutionTestVariable != 2 {
		t.Errorf("Provider 'p2' not called")
	}

	_, err = CreateStorageOfType("p3", nil)
	if err == nil {
		t.Errorf("Called unsupported storage, expected error, got none")
	}

}
