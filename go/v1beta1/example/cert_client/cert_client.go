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

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// TODO: Update the paths below.
	certFile = "/path/to/client.crt"
	keyFile  = "/path/to/client.key"
	caFile   = "/path/to/ca.crt"
)

func main() {
	// Load client cert
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	creds := credentials.NewTLS(tlsConfig)
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds))
	defer conn.Close()
	client := pb.NewGrafeasV1Beta1Client(conn)

	// List notes
	resp, err := client.ListNotes(context.Background(),
		&pb.ListNotesRequest{
			Parent: "projects/myproject",
		})
	if err != nil {
		log.Fatal(err)
	}

	if len(resp.Notes) != 0 {
		log.Println(resp.Notes)
	} else {
		log.Println("Project does not contain any notes")
	}
}
