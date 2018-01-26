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

package config

import (
	"io/ioutil"

	"github.com/grafeas/grafeas/samples/server/go-server/api/server/api"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/storage"
	"gopkg.in/yaml.v2"
)

// File is the grafeas config file.
type file struct {
	Grafeas *config `yaml:"grafeas"`
}

// Config is the global configuration for an instance of Grafeas.
type config struct {
	API         *api.Config          `yaml:"api"`
	StorageType string               `yaml:"storage_type"` // Supported storage types are "memstore" and "postgres"
	PgSQLConfig *storage.PgSQLConfig `yaml:"postgres"`
}

// DefaultConfig is a configuration that can be used as a fallback value.
func defaultConfig() *config {
	return &config{
		API: &api.Config{
			Address:  "localhost:10000",
			CertFile: "",
			KeyFile:  "",
			CAFile:   "",
		},
		StorageType: "memstore",
		PgSQLConfig: &storage.PgSQLConfig{},
	}
}

// Creates a config from a YAML-file. If fileName is an empty
// string a default config will be returned.
func LoadConfig(fileName string) (*config, error) {
	if fileName == "" {
		return defaultConfig(), nil
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var configFile file
	err = yaml.Unmarshal(data, &configFile)
	if err != nil {
		return nil, err
	}
	return configFile.Grafeas, nil
}
