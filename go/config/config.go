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
	"errors"
	"io/ioutil"
	"log"

	fernet "github.com/fernet/fernet-go"
	"github.com/grafeas/grafeas/go/v1beta1/storage"
	"gopkg.in/yaml.v2"
)

// File is the grafeas config file.
type file struct {
	Grafeas *config `yaml:"grafeas"`
}

type Config struct {
	Address            string   `yaml:"address"`              // Endpoint address, e.g. localhost:8080 or unix:///var/run/grafeas.sock
	CertFile           string   `yaml:"certfile"`             // A PEM eoncoded certificate file
	KeyFile            string   `yaml:"keyfile"`              // A PEM encoded private key file
	CAFile             string   `yaml:"cafile"`               // A PEM eoncoded CA's certificate file
	CORSAllowedOrigins []string `yaml:"cors_allowed_origins"` // Permitted CORS origins.
}

// Config is the global configuration for an instance of Grafeas.
type config struct {
	API            *Config                  `yaml:"api"`
	StorageType    string                       `yaml:"storage_type"` // Supported storage types are "memstore", "postgres" and "embedded"
	PgSQLConfig    *storage.PgSQLConfig         `yaml:"postgres"`
	EmbeddedConfig *storage.EmbeddedStoreConfig `yaml:"embedded"` // EmbeddedConfig is the embedded store config
}

// DefaultConfig is a configuration that can be used as a fallback value.
func defaultConfig() *config {
	return &config{
		API: &Config{
			Address:  "0.0.0.0:8080",
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
	config := configFile.Grafeas

	if config.StorageType == "postgres" {
		// Generate a pagination key if none is provided.
		if config.PgSQLConfig.PaginationKey == "" {
			log.Println("pagination key is empty, generating...")
			var key fernet.Key
			if err = key.Generate(); err != nil {
				return nil, err
			}
			config.PgSQLConfig.PaginationKey = key.Encode()
		} else {
			_, err = fernet.DecodeKey(config.PgSQLConfig.PaginationKey)
			if err != nil {
				err = errors.New("Invalid Pagination key; must be 32-bit URL-safe base64")
				return nil, err
			}
		}
	}
	return config, nil
}
