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
	"gopkg.in/yaml.v2"
)

// file is the Grafeas configuration file.
type file struct {
	Grafeas *GrafeasConfig `yaml:"grafeas"`
}

// ServerConfig is the Grafeas server configuration.
type ServerConfig struct {
	Address            string   `yaml:"address"`              // Endpoint address, e.g. localhost:8080 or unix:///var/run/grafeas.sock
	CertFile           string   `yaml:"certfile"`             // A PEM encoded certificate file
	KeyFile            string   `yaml:"keyfile"`              // A PEM encoded private key file
	CAFile             string   `yaml:"cafile"`               // A PEM encoded CA's certificate file
	CORSAllowedOrigins []string `yaml:"cors_allowed_origins"` // Permitted CORS origins.
}

// EmbeddedStoreConfig is the configuration for embedded store.
type EmbeddedStoreConfig struct {
	Path string `yaml:"path"` // Path is the folder path to storage files
}

// PgSQLConfig is the configuration for PostgreSQL store.
type PgSQLConfig struct {
	Host     string `yaml:"host"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	// Valid sslmodes: disable, allow, prefer, require, verify-ca, verify-full.
	// See https://www.postgresql.org/docs/current/static/libpq-connect.html for details
	SSLMode       string `yaml:"sslmode"`
	PaginationKey string `yaml:"paginationkey"`
}

// GrafeasConfig is the global configuration for an instance of Grafeas.
type GrafeasConfig struct {
	API            *ServerConfig        `yaml:"api"`
	StorageType    string               `yaml:"storage_type"` // Supported storage types are "memstore", "postgres" and "embedded"
	PgSQLConfig    *PgSQLConfig         `yaml:"postgres"`
	EmbeddedConfig *EmbeddedStoreConfig `yaml:"embedded"` // EmbeddedConfig is the embedded store config
}

// defaultConfig is a configuration that can be used as a fallback value.
func defaultConfig() *GrafeasConfig {
	return &GrafeasConfig{
		API: &ServerConfig{
			Address:  "0.0.0.0:8080",
			CertFile: "",
			KeyFile:  "",
			CAFile:   "",
		},
		StorageType: "memstore",
		PgSQLConfig: &PgSQLConfig{},
	}
}

// LoadConfig creates a config from a YAML-file. If fileName is an empty
// string a default config will be returned.
func LoadConfig(fileName string) (*GrafeasConfig, error) {
	if fileName == "" {
		return defaultConfig(), nil
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var configFile file
	if err := yaml.Unmarshal(data, &configFile); err != nil {
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
