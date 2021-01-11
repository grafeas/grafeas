package tools

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

// This file builds a number of proto specific binaries to get go protobufs working via plugins.

//go:generate env GO111MODULE=on go install -v    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
//go:generate env GO111MODULE=on go install -v    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
//go:generate env GO111MODULE=on go install -v    google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate env GO111MODULE=on go install -v    google.golang.org/grpc/cmd/protoc-gen-go-grpc
