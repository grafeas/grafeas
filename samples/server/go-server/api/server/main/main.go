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
	"log"

	"flag"
	"fmt"
	"net"

	"github.com/grafeas/grafeas/samples/server/go-server/api/server/storage"
	"github.com/grafeas/grafeas/samples/server/go-server/api/server/v1alpha1"
	pb "github.com/grafeas/grafeas/v1alpha1/proto"
	opspb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	log.Printf("Server started on port %v", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	g := v1alpha1.Grafeas{S: storage.NewMemStore()}
	pb.RegisterGrafeasServer(grpcServer, &g)
	pb.RegisterGrafeasProjectsServer(grpcServer, &g)
	opspb.RegisterOperationsServer(grpcServer, &g)
	grpcServer.Serve(lis)
}
