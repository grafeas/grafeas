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

package config

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func defaultServerConfig(t *testing.T) *ServerConfig {
	return &ServerConfig{
		Address:            "0.0.0.0:8080",
		CertFile:           "",
		KeyFile:            "",
		CAFile:             "",
		CORSAllowedOrigins: nil,
	}
}

func userServerConfig(t *testing.T) *ServerConfig {
	return &ServerConfig{
		Address:  "0.0.0.0:8081",
		CertFile: "abc",
		KeyFile:  "def",
		CAFile:   "ghi",
		CORSAllowedOrigins: []string{
			"http://example.com",
			"https://somewhere.else.com",
		},
	}
}

func userConfig_memstore_yaml(t *testing.T) []byte {
	return []byte(`
grafeas:
  api:
    address: "0.0.0.0:8081"
    certfile: abc
    keyfile: def
    cafile:  ghi
    cors_allowed_origins:
      - "http://example.com"
      - "https://somewhere.else.com"
  storage_type: "memstore"
`)
}

func userEmbeddedConfig(t *testing.T) *EmbeddedStoreConfig {
	return &EmbeddedStoreConfig{
		Path: "/some/path",
	}
}

func userConfig_embedded_yaml(t *testing.T) []byte {
	return []byte(`
grafeas:
  api:
    address: "0.0.0.0:8081"
    certfile: abc
    keyfile: def
    cafile:  ghi
    cors_allowed_origins:
      - "http://example.com"
      - "https://somewhere.else.com"
  storage_type: "embedded"
  embedded: 
    path: /some/path
`)
}

func TestLoadConfig_ReturnsDefaultConfig_NoInput(t *testing.T) {
	cfg, err := LoadConfig("")
	if err != nil {
		t.Error(err)
	}

	if cfg.API == nil {
		t.Errorf("config API is nil")
	}

	if !cmp.Equal(cfg.API, defaultServerConfig(t)) {
		t.Errorf("Values in cfg.API are not correct\n%s", cmp.Diff(cfg.API, defaultServerConfig(t)))
	}

	if cfg.StorageType != "memstore" {
		t.Errorf("Storage type %s is not memstore", cfg.StorageType)
	}

	if cfg.StorageConfig != nil {
		t.Errorf("storage configuration is not nil")
	}
}

func TestLoadConfig_ReturnsConfig_UserSuppliedValues_Memstore(t *testing.T) {
	file, err := ioutil.TempFile("", "config.*.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = os.Remove(file.Name())
	}()

	_, err = file.Write(userConfig_memstore_yaml(t))
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := LoadConfig(file.Name())
	if err != nil {
		t.Error(err)
	}

	if cfg.API == nil {
		t.Errorf("config API is nil")
	}

	if !cmp.Equal(cfg.API, userServerConfig(t)) {
		t.Errorf("Values in cfg.API are not correct\n%s", cmp.Diff(cfg.API, userServerConfig(t)))
	}

	if cfg.StorageType != "memstore" {
		t.Errorf("Storage type %s is not memstore", cfg.StorageType)
	}

	if cfg.StorageConfig != nil {
		t.Errorf("storage configuration is not nil")
	}
}

func TestLoadConfig_ReturnsConfig_UserSuppliedValues_Embedded(t *testing.T) {
	file, err := ioutil.TempFile("", "config.*.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = os.Remove(file.Name())
	}()

	_, err = file.Write(userConfig_embedded_yaml(t))
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := LoadConfig(file.Name())
	if err != nil {
		t.Error(err)
	}

	if cfg.API == nil {
		t.Errorf("config API is nil")
	}

	if !cmp.Equal(cfg.API, userServerConfig(t)) {
		t.Errorf("Values in cfg.API are not correct\n%s", cmp.Diff(cfg.API, userServerConfig(t)))
	}

	if cfg.StorageType != "embedded" {
		t.Errorf("Storage type %s is not embedded", cfg.StorageType)
	}

	if cfg.StorageConfig == nil {
		t.Errorf("storage configuration is nil")
	}

	var storeConfig EmbeddedStoreConfig

	err = ConvertGenericConfigToSpecificType(*cfg.StorageConfig, &storeConfig)
	if err != nil {
		log.Fatal(err)
	}

	if !cmp.Equal(storeConfig, *userEmbeddedConfig(t)) {
		t.Errorf("Values in storage configuration are not correct\n%s", cmp.Diff(storeConfig, *userEmbeddedConfig(t)))
	}
}

// TODO(#341) move these 2 supporting functions and the test case to the new project
func userPostgresConfig(t *testing.T) *PgSQLConfig {
	return &PgSQLConfig{
		Host:          "127.0.0.1:5432",
		DbName:        "postgres",
		User:          "postgres",
		Password:      "password",
		SSLMode:       "require",
		PaginationKey: "",
	}
}

func userConfig_postgres_yaml(t *testing.T) []byte {
	return []byte(`
grafeas:
  api:
    address: "0.0.0.0:8081"
    certfile: abc
    keyfile: def
    cafile:  ghi
    cors_allowed_origins:
      - "http://example.com"
      - "https://somewhere.else.com"
  storage_type: "postgres"
  postgres:
    host: "127.0.0.1:5432"
    dbname: "postgres"
    user: "postgres"
    password: "password"
    sslmode: "require"
    paginationkey:
`)
}

func TestLoadConfig_ReturnsConfig_UserSuppliedValues_Postgres(t *testing.T) {
	file, err := ioutil.TempFile("", "config.*.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = os.Remove(file.Name())
	}()

	_, err = file.Write(userConfig_postgres_yaml(t))
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := LoadConfig(file.Name())
	if err != nil {
		t.Error(err)
	}

	if cfg.API == nil {
		t.Errorf("config API is nil")
	}

	if !cmp.Equal(cfg.API, userServerConfig(t)) {
		t.Errorf("Values in cfg.API are not correct\n%s", cmp.Diff(cfg.API, userServerConfig(t)))
	}

	if cfg.StorageType != "postgres" {
		t.Errorf("Storage type %s is not postgres", cfg.StorageType)
	}

	if cfg.StorageConfig == nil {
		t.Errorf("storage configuration is nil")
	}

	var storeConfig PgSQLConfig

	err = ConvertGenericConfigToSpecificType(*cfg.StorageConfig, &storeConfig)
	if err != nil {
		log.Fatal(err)
	}

	if !cmp.Equal(storeConfig, *userPostgresConfig(t)) {
		t.Errorf("Values in storage configuration are not correct\n%s", cmp.Diff(storeConfig, *userPostgresConfig(t)))
	}
}

// The project defines an example of configuration in go/v1beta1/config.yaml, so this test
// confirms that this file remains parseable.
// It does not check the config file values to avoid the test becoming brittle.
func TestLoadConfig_LoadsConfigWithoutError_ForExampleProjectConfiguration(t *testing.T) {
	cfg, err := LoadConfig("../v1beta1/config.yaml")
	if err != nil {
		t.Error(err)
	}

	if cfg.API == nil {
		t.Errorf("config API is nil")
	}
}
