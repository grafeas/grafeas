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

package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/grafeas/grafeas/samples/server/go-server/api/server/cert"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/storage"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/v1alpha1"
	pb "github.com/grafeas/grafeas/v1alpha1/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	opspb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

// grpcHandlerFunc returns an http.Handler that delegates to grpcServer on incoming gRPC
// connections or otherHandler otherwise. Copied from cockroachdb via philips
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// This is a partial recreation of gRPC's internal checks https://github.com/grpc/grpc-go/pull/514/files#diff-95e9a25b738459a2d3030e1e6fa2a718R61
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	})
}

func main() {
	flag.Parse()
	grafeasEndpoint := fmt.Sprintf("localhost:%d", *port)

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewClientTLSFromCert(cert.Pool(), grafeasEndpoint))}
	grpcServer := grpc.NewServer(opts...)
	g := v1alpha1.Grafeas{S: storage.NewMemStore()}
	pb.RegisterGrafeasServer(grpcServer, &g)
	opspb.RegisterOperationsServer(grpcServer, &g)
	log.Printf("Server started on port %v", *port)

	ctx := context.Background()

	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		ServerName: "localhost",
		RootCAs:    cert.Pool(),
	}))}
	err := pb.RegisterGrafeasHandlerFromEndpoint(ctx, mux, grafeasEndpoint, dialOpts)
	if err != nil {
		log.Fatalf("Failed to resigister handler: %v", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err != nil {
		log.Fatal(err)
		return
	}
	pair := cert.Pair()
	srv := &http.Server{
		Addr:    "localhost",
		Handler: grpcHandlerFunc(grpcServer, mux),
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{*pair},
			NextProtos:   []string{"h2"},
		},
	}
	log.Printf("Server started on port %v", *port)
	err = srv.Serve(tls.NewListener(lis, srv.TLSConfig))
	if err != nil {
		log.Fatalf("Unable to serve: %v", err)
	}
}
