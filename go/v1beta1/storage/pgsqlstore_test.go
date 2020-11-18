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
	"archive/zip"
	"crypto/tls"
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

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

func downloadFile(filepath string, url string) error {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
	}

	// Based off of http.DefaultTransport
	var transport http.RoundTripper = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		TLSClientConfig:       tlsConfig,
		ExpectContinueTimeout: 1 * time.Second,
	}

	httpClient := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func unzip(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

func startupPostgres(pgData *testPgHelper) error {
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
		binPath = filepath.Dir(string(output))
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

	workingDir, err := os.Getwd()
	if err != nil {
		return pgData, err
	}
	//The root is 3 levels up. Not an ideal means of getting this, but since there are multiple
	//ways of starting the tests, os.Executable and runtime.Caller don't seem like viable options
	basepath := filepath.Dir(filepath.Dir(filepath.Dir(workingDir)))
	basepath = filepath.Join(basepath, "test", "utilities")

	pgData.pgBinPath = ""
	testPgPath := filepath.Join(basepath, "pgsql")
	testPgBinPath := filepath.Join(testPgPath, "bin")
	if _, err := os.Stat(testPgBinPath); !os.IsNotExist(err) {
		pgData.pgBinPath = testPgBinPath
	} else {
		//Check for a global installation
		globalPgBinPath, err := getPostgresBinPathFromSystemPath()
		if err == nil && globalPgBinPath != "" {
			pgData.pgBinPath = globalPgBinPath
		}
	}

	if pgData.pgBinPath == "" {
		os.MkdirAll(basepath, 0755)

		fileNameBase := "postgresql-10.15.1"
		postgresZipFileName := filepath.Join(basepath, fmt.Sprintf("%s.zip", fileNameBase))

		postgresZipFound := false
		if _, err := os.Stat(postgresZipFileName); !os.IsNotExist(err) {
			postgresZipFound = true
		}

		if !postgresZipFound {
			//URLs from https://www.enterprisedb.com/download-postgresql-binaries
			postgresUrl := ""
			switch runtime.GOOS {
			case "darwin":
				postgresUrl = "https://sbp.enterprisedb.com/getfile.jsp?fileid=1257418&_ga=2.59164973.1284458936.1605634932-924989421.1605634932"
			case "linux":
				switch runtime.GOARCH {
				case "386":
					postgresUrl = "https://sbp.enterprisedb.com/getfile.jsp?fileid=1257416&_ga=2.255110536.1284458936.1605634932-924989421.1605634932"
				case "amd64":
					postgresUrl = "https://sbp.enterprisedb.com/getfile.jsp?fileid=1257417&_ga=2.59164973.1284458936.1605634932-924989421.1605634932"
				}
			case "windows":
				switch runtime.GOARCH {
				case "386":
					postgresUrl = "https://sbp.enterprisedb.com/getfile.jsp?fileid=1257419&_ga=2.255110536.1284458936.1605634932-924989421.1605634932"
				case "amd64":
					postgresUrl = "https://sbp.enterprisedb.com/getfile.jsp?fileid=1257420&_ga=2.255110536.1284458936.1605634932-924989421.1605634932"
				}
			}

			if postgresUrl == "" {
				return pgData, fmt.Errorf("Unable to test on %s with architecture %s", runtime.GOOS, runtime.GOARCH)
			}

			//Download postgres
			postgresTempZipFile, err := ioutil.TempFile("", fmt.Sprintf("%s-*.zip", fileNameBase))
			if err != nil {
				return pgData, err
			}

			fmt.Fprintln(os.Stderr, "testing: downloading", postgresUrl)
			if err := downloadFile(postgresTempZipFile.Name(), postgresUrl); err != nil {
				return pgData, err
			}

			//Move fully downloaded tempfile to relative location to prevent repeated downloads
			if err := os.Rename(postgresTempZipFile.Name(), postgresZipFileName); err != nil {
				return pgData, err
			}
		}

		//Extract the zip file
		fmt.Fprintln(os.Stderr, "testing: unzipping", postgresZipFileName)
		unzipPath, err := ioutil.TempDir("", fmt.Sprintf("%s-*", fileNameBase))
		if err != nil {
			return pgData, err
		}

		if _, err := unzip(postgresZipFileName, unzipPath); err != nil {
			return pgData, err
		}

		//Move fully extracted zip folder to relative location to prevent repeated unzipping
		if err := os.Rename(filepath.Join(unzipPath, "pgsql"), testPgPath); err != nil {
			return pgData, err
		}

		//Remove the zip file
		if err := os.Remove(postgresZipFileName); err != nil {
			return pgData, err
		}

		pgData.pgBinPath = testPgBinPath
	}

	//Startup postgres
	pgDataPath, err := ioutil.TempDir("", "pg-data-*")
	if err != nil {
		return nil, err
	}

	pgData.pgDataPath = pgDataPath

	if err := startupPostgres(pgData); err != nil {
		return pgData, err
	}

	return pgData, nil
}

func stopPostgres(pgData *testPgHelper) error {
	if pgData.startedPg {
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
