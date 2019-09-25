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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

// file is the Grafeas configuration file.
type file struct {
	Grafeas GrafeasConfig `mapstructure:"grafeas"`
}

type GrafeasConfig struct {
	API           *ServerConfig `mapstructure:"api"`
	StorageType   string        `mapstructure:"storage_type"` // Natively supported storage types are "memstore" and "embedded"
	StorageConfig *interface{}
}

// ServerConfig is the Grafeas server configuration.
type ServerConfig struct {
	Address            string   `mapstructure:"address"`              // Endpoint address, e.g. localhost:8080 or unix:///var/run/grafeas.sock
	CertFile           string   `mapstructure:"certfile"`             // A PEM encoded certificate file
	KeyFile            string   `mapstructure:"keyfile"`              // A PEM encoded private key file
	CAFile             string   `mapstructure:"cafile"`               // A PEM encoded CA's certificate file
	CORSAllowedOrigins []string `mapstructure:"cors_allowed_origins"` // Permitted CORS origins.
}

// EmbeddedStoreConfig is the configuration for embedded store.
type EmbeddedStoreConfig struct {
	Path string `mapstructure:"path"` // Path is the folder path to storage files
}

// TODO(#341) Move this to its own project
// PgSQLConfig is the configuration for PostgreSQL store.
type PgSQLConfig struct {
	Host     string `mapstructure:"host"`
	DbName   string `mapstructure:"dbname"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	// Valid sslmodes: disable, allow, prefer, require, verify-ca, verify-full.
	// See https://www.postgresql.org/docs/current/static/libpq-connect.html for details
	SSLMode       string `mapstructure:"sslmode"`
	PaginationKey string `mapstructure:"paginationkey"`
}

// defaultConfig is a configuration that can be used as a fallback value.
var defaultConfig = []byte(`
grafeas:
  # Grafeas api server config
  api:
    # Endpoint address
    address: "0.0.0.0:8080"
    # PKI configuration (optional)
    certfile:
    keyfile:
    cafile: 
    # CORS configuration (optional)
    cors_allowed_origins:
      # - "http://example.net"
  # Supported storage types are "memstore" and "postgres"
  storage_type: "memstore"
`)

// LoadConfig creates a config from a YAML-file. If fileName is an empty
// string a default config will be returned.
func LoadConfig(fileName string) (*GrafeasConfig, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	var err error
	data := defaultConfig
	// now read from config cfg if required
	if fileName != "" {
		data, err = ioutil.ReadFile(fileName)
		if err != nil {
			return nil, err
		}
	}
	if err = v.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return nil, err
	}

	var config GrafeasConfig

	// parse server config
	serverCfg := ServerConfig{}
	if err = v.UnmarshalKey("grafeas.api", &serverCfg); err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to decode into struct, %v", err))
	}
	config.API = &serverCfg

	// parse storage type
	config.StorageType = v.GetString("grafeas.storage_type")

	// parse storage type-specific configuration
	genericConfig := v.Get(fmt.Sprintf("grafeas.%s", config.StorageType))

	if config.StorageType != "memstore" && genericConfig != nil {
		config.StorageConfig = &genericConfig
	}

	return &config, nil
}

// ConvertGenericConfigToSpecificType will attempt to copy generic configuration within source
// to a target struct that represents the specific storage configuration, represented as an interface{}.
// see config_test.go for example usage.
func ConvertGenericConfigToSpecificType(source interface{}, target interface{}) error {
	b, err := json.Marshal(source)
	if err != nil {
		return errors.New(fmt.Sprintf("Error parsing configuration, %v", err))
	}

	return json.Unmarshal(b, &target)
}
