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
	"errors"
	"fmt"

	"github.com/grafeas/grafeas/go/config"
	"github.com/grafeas/grafeas/go/v1beta1/api"
	"github.com/grafeas/grafeas/go/v1beta1/project"
)

// use type aliasing to get multiple inheritance even though both interfaces are called Storage
type ps = project.Storage
type gs = grafeas.Storage

// unified interface that basically gives us a single interface called Storage that implements
// both anonymous interfaces.
type Storage struct {
	ps
	gs
}

var registeredStorageTypeProviders = map[string]func(storageType string, storageConfig *interface{}) (*Storage, error){}

// RegisterStorageTypeProvider registers a new provider to create a specific type of Storage
func RegisterStorageTypeProvider(storageType string, provider func(storageType string, storageConfig *interface{}) (*Storage, error)) error {
	if _, present := registeredStorageTypeProviders[storageType]; !present {
		registeredStorageTypeProviders[storageType] = provider
		return nil
	} else {
		return errors.New(fmt.Sprintf("Storage provider %s already exists", storageType))
	}
}

// CreateStorageOfType will create an instance of Storage by name or an error if that type is unsupported.
func CreateStorageOfType(storageType string, storageConfig *interface{}) (*Storage, error) {
	if provider, present := registeredStorageTypeProviders[storageType]; present {
		return provider(storageType, storageConfig)
	} else {
		return nil, errors.New(fmt.Sprintf("Unsupported storage type %s", storageType))
	}

}

// MemstoreStorageTypeProvider returns a memstore storage instance
func memstoreStorageTypeProvider(storageType string, storageConfig *interface{}) (*Storage, error) {
	if storageType != "memstore" {
		return nil, errors.New(fmt.Sprintf("Unknown storage type %s, must be 'memstore'", storageType))
	}

	s := NewMemStore()
	storage := &Storage{
		ps: s,
		gs: s,
	}

	return storage, nil
}

// EmbeddedStorageTypeProvider returns an embedded storage instance
func embeddedStorageTypeProvider(storageType string, storageConfig *interface{}) (*Storage, error) {
	if storageType != "embedded" {
		return nil, errors.New(fmt.Sprintf("Unknown storage type %s, must be 'embedded'", storageType))
	}

	var storeConfig config.EmbeddedStoreConfig

	err := config.ConvertGenericConfigToSpecificType(storageConfig, &storeConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to create EmbeddedStoreConfig, %s", err))
	}

	s := NewEmbeddedStore(&storeConfig)
	storage := &Storage{
		ps: s,
		gs: s,
	}

	return storage, nil
}

// PostgresStorageTypeProvider returns a postgres storage instance
// TODO(#341) move this function to a separate project
func postgresStorageTypeProvider(storageType string, storageConfig *interface{}) (*Storage, error) {
	if storageType != "postgres" {
		return nil, errors.New(fmt.Sprintf("Unknown storage type %s, must be 'postgres'", storageType))
	}

	var storeConfig config.PgSQLConfig

	err := config.ConvertGenericConfigToSpecificType(storageConfig, &storeConfig)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to create PgSQLConfig, %s", err))
	}

	s := NewPgSQLStore(&storeConfig)
	storage := &Storage{
		ps: s,
		gs: s,
	}

	return storage, nil
}

// RegisterDefaultStorageTypeProviders adds support for memstore, embedded and Postgres storage types
// TODO(#341) remove support for Postgres and move to a separate Register...() implementation in a separate project
func RegisterDefaultStorageTypeProviders() error {
	err := RegisterStorageTypeProvider("memstore", memstoreStorageTypeProvider)
	if err != nil {
		return err
	}

	err = RegisterStorageTypeProvider("embedded", embeddedStorageTypeProvider)
	if err != nil {
		return err
	}

	// TODO(#341) move this function invocation to a separate function within a separate project
	err = RegisterStorageTypeProvider("postgres", postgresStorageTypeProvider)
	if err != nil {
		return err
	}

	return nil
}
