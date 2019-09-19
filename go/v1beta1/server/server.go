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

package server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/cockroachdb/cmux"
	"github.com/grafeas/grafeas/go/config"
	"github.com/grafeas/grafeas/go/v1beta1/api"
	"github.com/grafeas/grafeas/go/v1beta1/project"
	"github.com/grafeas/grafeas/go/v1beta1/storage"
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	prpb "github.com/grafeas/grafeas/proto/v1beta1/project_go_proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	configFile = flag.String("config", "", "Path to a config file")
)

func StartGrafeas() {
	flag.Parse()
	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		log.Fatalf("Failed to load cfg file: %s", err)
	}
	var db grafeas.Storage
	var proj project.Storage

	s, err := storage.CreateStorageOfType(cfg.StorageType, cfg.StorageConfig)
	if err != nil {
		log.Fatalf("Failed to create storage: %s", err)
	}
	db = s
	proj = s
	run(cfg.API, &db, &proj)
}

// run initializes grpc and grpc gateway api services on the same address
func run(config *config.ServerConfig, db *grafeas.Storage, proj *project.Storage) {
	network, address := "tcp", config.Address
	if strings.HasPrefix(config.Address, "unix://") {
		network = "unix"
		address = strings.TrimPrefix(config.Address, "unix://")
		// Remove existing socket if found
		os.Remove(address)
	}
	l, err := net.Listen(network, address)
	if err != nil {
		log.Fatalln("could not listen to address", config.Address)
	}
	log.Printf("starting grpc server on %s", address)

	var (
		apiHandler  http.Handler
		apiListener net.Listener
		srv         *http.Server
		ctx         = context.Background()
		httpMux     = http.NewServeMux()
		tcpMux      = cmux.New(l)
	)

	tlsConfig, err := tlsClientConfig(config.CAFile)
	if err != nil {
		log.Fatal("Failed to create tls config", err)
	}

	if tlsConfig != nil {
		cert, err := tls.LoadX509KeyPair(config.CertFile, config.KeyFile)
		if err != nil {
			log.Fatalln("Failed to load certificate files", err)
		}
		tlsConfig.Certificates = []tls.Certificate{cert}
		tlsConfig.NextProtos = []string{"h2"}

		apiListener = tls.NewListener(tcpMux.Match(cmux.Any()), tlsConfig)
		go func() { handleShutdown(tcpMux.Serve()) }()

		grpcServer := newGrpcServer(tlsConfig, db, proj)
		gwmux := newGrpcGatewayServer(ctx, apiListener.Addr().String(), tlsConfig)

		httpMux.Handle("/", gwmux)
		apiHandler = grpcHandlerFunc(grpcServer, httpMux)

		log.Println("grpc server is configured with client certificate authentication")
	} else {
		grpcL := tcpMux.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
		apiListener = tcpMux.Match(cmux.Any())
		go func() { handleShutdown(tcpMux.Serve()) }()

		grpcServer := newGrpcServer(nil, db, proj)
		go func() { handleShutdown(grpcServer.Serve(grpcL)) }()

		gwmux := newGrpcGatewayServer(ctx, apiListener.Addr().String(), nil)

		httpMux.Handle("/", gwmux)
		apiHandler = httpMux

		log.Println("grpc server is configured without client certificate authentication")
	}

	// Setup the CORS middleware. If `config.CORSAllowedOrigins` is empty, no CORS
	// Origins will be allowed through.
	cors := cors.New(cors.Options{
		AllowedOrigins: config.CORSAllowedOrigins,
	})

	srv = &http.Server{
		Handler:   cors.Handler(apiHandler),
		TLSConfig: tlsConfig,
	}

	// blocking call
	handleShutdown(srv.Serve(apiListener))
	log.Println("Grpc API stopped")
}

// handleShutdown handles the server shut down error.
func handleShutdown(err error) {
	if err != nil {
		if opErr, ok := err.(*net.OpError); !ok || (ok && opErr.Op != "accept") {
			log.Fatal(err)
		}
	}
}

func newGrpcServer(tlsConfig *tls.Config, db *grafeas.Storage, proj *project.Storage) *grpc.Server {
	grpcOpts := []grpc.ServerOption{}

	if tlsConfig != nil {
		grpcOpts = append(grpcOpts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}

	grpcServer := grpc.NewServer(grpcOpts...)
	g := grafeas.API{
		Storage:           *db,
		Auth:              &grafeas.NoOpAuth{},
		Filter:            &grafeas.NoOpFilter{},
		Logger:            &grafeas.NoOpLogger{},
		EnforceValidation: true,
	}
	pb.RegisterGrafeasV1Beta1Server(grpcServer, &g)

	gp := project.API{Storage: *proj}
	prpb.RegisterProjectsServer(grpcServer, &gp)

	return grpcServer
}

func newGrpcGatewayServer(ctx context.Context, listenerAddr string, tlsConfig *tls.Config) http.Handler {
	var (
		gwTLSConfig *tls.Config
		gwOpts      []grpc.DialOption
	)

	if tlsConfig != nil {
		gwTLSConfig = tlsConfig.Clone()
		gwTLSConfig.InsecureSkipVerify = true
		gwOpts = append(gwOpts, grpc.WithTransportCredentials(credentials.NewTLS(gwTLSConfig)))
	} else {
		gwOpts = append(gwOpts, grpc.WithInsecure())
	}

	// changes json serializer to include empty fields with default values
	jsonOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{EmitDefaults: true})
	gwmux := runtime.NewServeMux(jsonOpt)

	conn, err := grpc.DialContext(ctx, listenerAddr, gwOpts...)
	if err != nil {
		log.Fatal("could not initialize grpc gateway connection")
	}
	err = pb.RegisterGrafeasV1Beta1Handler(ctx, gwmux, conn)
	if err != nil {
		log.Fatal("could not initialize ancestry grpc gateway")
	}

	err = prpb.RegisterProjectsHandler(ctx, gwmux, conn)
	if err != nil {
		log.Fatal("could not initialize notification grpc gateway")
	}

	return http.Handler(gwmux)
}

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb.
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

// tlsClientConfig initializes a *tls.Config using the given CA. The resulting
// *tls.Config is meant to be used to configure an HTTP server to do client
// certificate authentication.
//
// If no CA is given, a nil *tls.Config is returned; no client certificate will
// be required and verified. In other words, authentication will be disabled.
func tlsClientConfig(caPath string) (*tls.Config, error) {
	if caPath == "" {
		return nil, nil
	}

	caCert, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}

	return tlsConfig, nil
}
