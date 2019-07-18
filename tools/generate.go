// This file builds a number of proto specific binaries to get go protobufs working via plugins.

//go:generate env GO111MODULE=on go install -v github.com/golang/protobuf/protoc-gen-go
//go:generate env GO111MODULE=on go get -u -v github.com/golang/protobuf/protoc-gen-go
//go:generate env GO111MODULE=on go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
//go:generate env GO111MODULE=on go install -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
package tools
