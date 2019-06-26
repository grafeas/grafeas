// This file generates the v1alpha1 proto bindings for go.

//go:generate ../../protoc/bin/protoc -I ./ -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway -I ../../vendor/github.com/googleapis/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true:. --swagger_out=logtostderr=true:. grafeas.proto
package v1alpha1
