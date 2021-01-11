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
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"testing"

	"github.com/grafeas/grafeas/go/config"
	grafeas "github.com/grafeas/grafeas/go/v1beta1/api"
	"github.com/grafeas/grafeas/go/v1beta1/project"
	"github.com/grafeas/grafeas/go/v1beta1/storage"
)

type testPgHelper struct {
	pgDataPath string
	pgBinPath  string
	startedPg  bool
	pgConfig   *config.PgSQLConfig
}

var (
	//Unfortunately, not a good way to pass this information around to tests except via a globally scoped var
	pgsqlstoreTestPgConfig *testPgHelper
)

func startupPostgres(pgData *testPgHelper) error {
	//Create a test database instance directory
	if pgDataPath, err := ioutil.TempDir("", "pg-data-*"); err != nil {
		return err
	} else {
		pgData.pgDataPath = filepath.ToSlash(pgDataPath)
	}

	//Make password file
	passwordTempFile, err := ioutil.TempFile("", "pgpassword-*")
	if err != nil {
		return err
	}
	defer os.Remove(passwordTempFile.Name())

	if _, err = io.WriteString(passwordTempFile, pgData.pgConfig.Password); err != nil {
		return err
	}

	if err := passwordTempFile.Sync(); err != nil {
		return err
	}

	port, err := findAvailablePort()
	if err != nil {
		return err
	}
	pgData.pgConfig.Host = fmt.Sprintf("127.0.0.1:%d", port)

	//Init db
	pgCtl := filepath.Join(pgData.pgBinPath, "pg_ctl")
	fmt.Fprintln(os.Stderr, "testing: intializing test postgres instance under", pgData.pgDataPath)
	pgCtlInitDBOptions := fmt.Sprintf("--username %s --pwfile %s", pgData.pgConfig.User, passwordTempFile.Name())
	cmd := exec.Command(pgCtl, "--pgdata", pgData.pgDataPath, "-o", pgCtlInitDBOptions, "initdb")
	if err := cmd.Run(); err != nil {
		return err
	}

	//Start postgres
	fmt.Fprintln(os.Stderr, "testing: starting test postgres instance on port", port)
	pgCtlStartOptions := fmt.Sprintf("-p %d", port)
	cmd = exec.Command(pgCtl, "--pgdata", pgData.pgDataPath, "-o", pgCtlStartOptions, "start")
	if err := cmd.Run(); err != nil {
		return err
	}

	pgData.startedPg = true

	return nil
}

func findAvailablePort() (availablePort int, err error) {
	for availablePort = 5432; availablePort < 6000; availablePort++ {
		l, err := net.Listen("tcp", fmt.Sprintf(":%d", availablePort))
		defer l.Close()
		if err == nil {
			return availablePort, nil
		}
	}

	return -1, fmt.Errorf("Unable to find an open port")
}

func isPostgresRunning(config *config.PgSQLConfig) bool {
	source := storage.CreateSourceString(config.User, config.Password, config.Host, "postgres", config.SSLMode)
	db, err := sql.Open("postgres", source)
	if err != nil {
		return false
	}
	defer db.Close()

	if db.Ping() != nil {
		return false
	}
	return true
}

func getPostgresBinPathFromSystemPath() (binPath string, err error) {
	cmd := exec.Command("which", "pg_ctl")
	output, err := cmd.Output()
	if output != nil && err == nil {
		binPath = filepath.ToSlash(filepath.Dir(string(output)))
	}

	//Deal with "which" Linux-style output on Windows, a bit of a corner case
	regex := regexp.MustCompile("^/([a-z])/(.*)$")
	regexMatches := regex.FindStringSubmatch(binPath)
	if runtime.GOOS == "windows" && regexMatches != nil && len(regexMatches) == 3 {
		binPath = fmt.Sprintf("%s:/%s", regexMatches[1], regexMatches[2])
	}

	return
}

func setup() (pgData *testPgHelper, err error) {
	pgConfig := &config.PgSQLConfig{
		Host:     "127.0.0.1:5432",
		User:     "postgres",
		Password: "password",
		SSLMode:  "disable",
	}

	pgData = &testPgHelper{
		startedPg: false,
		pgConfig:  pgConfig,
	}

	//See if postgres is already available and running
	if isPostgresRunning(pgConfig) {
		return
	}

	//Check for a global installation
	if pgData.pgBinPath, err = getPostgresBinPathFromSystemPath(); err != nil {
		err = fmt.Errorf("Unable to find a running Postgres instance or Postgres binaries necessary for testing on the system PATH: %v", err)
		return
	}

	//Startup postgres
	if err = startupPostgres(pgData); err != nil {
		return
	}

	return pgData, nil
}

func stopPostgres(pgData *testPgHelper) error {
	if pgData != nil && pgData.startedPg {
		//Stop postgres
		pgCtl := filepath.Join(pgData.pgBinPath, "pg_ctl")

		fmt.Fprintln(os.Stderr, "testing: stopping test postgres instance")
		cmd := exec.Command(pgCtl, "--pgdata", pgData.pgDataPath, "stop")
		if err := cmd.Run(); err != nil {
			return err
		}

		//Cleanup
		if err := os.RemoveAll(pgData.pgDataPath); err != nil {
			return err
		}
	}

	return nil
}

func teardown(pgData *testPgHelper) error {
	return stopPostgres(pgData)
}

func dropDatabase(t *testing.T, config *config.PgSQLConfig) {
	t.Helper()
	// Open database
	source := storage.CreateSourceString(config.User, config.Password, config.Host, "postgres", config.SSLMode)
	db, err := sql.Open("postgres", source)
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	// Kill opened connection
	if _, err := db.Exec(`
		SELECT pg_terminate_backend(pid)
		FROM pg_stat_activity
		WHERE datname = $1`, config.DbName); err != nil {
		t.Fatalf("Failed to drop database: %v", err)
	}
	// Drop database
	if _, err := db.Exec("DROP DATABASE " + config.DbName); err != nil {
		t.Fatalf("Failed to drop database: %v", err)
	}
}

func TestMain(m *testing.M) {
	var err error
	pgsqlstoreTestPgConfig, err = setup()
	if err != nil {
		log.Fatal(err)
	}

	exitVal := m.Run()

	if err := teardown(pgsqlstoreTestPgConfig); err != nil {
		log.Fatal(err)
	}

	// os.Exit() does not respect defer statements
	os.Exit(exitVal)
}

func TestBetaPgSQLStore(t *testing.T) {
	createPgSQLStore := func(t *testing.T) (grafeas.Storage, project.Storage, func()) {
		t.Helper()
		config := &config.PgSQLConfig{
			Host:          pgsqlstoreTestPgConfig.pgConfig.Host,
			DbName:        "test_db",
			User:          pgsqlstoreTestPgConfig.pgConfig.User,
			Password:      pgsqlstoreTestPgConfig.pgConfig.Password,
			SSLMode:       pgsqlstoreTestPgConfig.pgConfig.SSLMode,
			PaginationKey: "XxoPtCUzrUv4JV5dS+yQ+MdW7yLEJnRMwigVY/bpgtQ=",
		}
		pg, err := storage.NewPgSQLStore(config)
		if err != nil {
			t.Errorf("Error creating PgSQLStore, %s", err)
		}
		var g grafeas.Storage = pg
		var gp project.Storage = pg
		return g, gp, func() { dropDatabase(t, config); pg.Close() }
	}

	storage.DoTestStorage(t, createPgSQLStore)
}

func TestPgSQLStoreWithUserAsEnv(t *testing.T) {
	createPgSQLStore := func(t *testing.T) (grafeas.Storage, project.Storage, func()) {
		t.Helper()
		config := &config.PgSQLConfig{
			Host:          pgsqlstoreTestPgConfig.pgConfig.Host,
			DbName:        "test_db",
			User:          "",
			Password:      "",
			SSLMode:       pgsqlstoreTestPgConfig.pgConfig.SSLMode,
			PaginationKey: "XxoPtCUzrUv4JV5dS+yQ+MdW7yLEJnRMwigVY/bpgtQ=",
		}
		_ = os.Setenv("PGUSER", pgsqlstoreTestPgConfig.pgConfig.User)
		_ = os.Setenv("PGPASSWORD", pgsqlstoreTestPgConfig.pgConfig.Password)
		pg, err := storage.NewPgSQLStore(config)
		if err != nil {
			t.Errorf("Error creating PgSQLStore, %s", err)
		}
		var g grafeas.Storage = pg
		var gp project.Storage = pg
		return g, gp, func() { dropDatabase(t, config); pg.Close() }
	}

	storage.DoTestStorage(t, createPgSQLStore)
}

func TestBetaPgSQLStoreWithNoPaginationKey(t *testing.T) {
	createPgSQLStore := func(t *testing.T) (grafeas.Storage, project.Storage, func()) {
		t.Helper()
		config := &config.PgSQLConfig{
			Host:          pgsqlstoreTestPgConfig.pgConfig.Host,
			DbName:        "test_db",
			User:          pgsqlstoreTestPgConfig.pgConfig.User,
			Password:      pgsqlstoreTestPgConfig.pgConfig.Password,
			SSLMode:       pgsqlstoreTestPgConfig.pgConfig.SSLMode,
			PaginationKey: "",
		}
		pg, err := storage.NewPgSQLStore(config)
		if err != nil {
			t.Errorf("Error creating PgSQLStore, %s", err)
		}
		var g grafeas.Storage = pg
		var gp project.Storage = pg
		return g, gp, func() { dropDatabase(t, config); pg.Close() }
	}

	storage.DoTestStorage(t, createPgSQLStore)
}

func TestBetaPgSQLStoreWithInvalidPaginationKey(t *testing.T) {
	config := &config.PgSQLConfig{
		Host:          pgsqlstoreTestPgConfig.pgConfig.Host,
		DbName:        "test_db",
		User:          pgsqlstoreTestPgConfig.pgConfig.User,
		Password:      pgsqlstoreTestPgConfig.pgConfig.Password,
		SSLMode:       pgsqlstoreTestPgConfig.pgConfig.SSLMode,
		PaginationKey: "INVALID_VALUE",
	}
	pg, err := storage.NewPgSQLStore(config)
	if pg != nil {
		pg.Close()
	}
	if err == nil {
		t.Errorf("expected error for invalid pagination key; got none")
	}
	if err.Error() != "invalid pagination key; must be 256-bit URL-safe base64" {
		t.Errorf("expected error message about invalid pagination key; got: %s", err.Error())
	}
}
