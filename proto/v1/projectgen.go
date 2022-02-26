// This generates the protocol buffer code in go for the v1 proto spec.

//go:generate rm -rf project_go_proto
//go:generate mkdir project_go_proto
//go:generate -command protoc ../../protoc/bin/protoc -I ../../ -I ./ -I ../../protodeps/grpc-gateway/third_party/googleapis -I ../../protodeps/grpc-gateway -I ../../protodeps/googleapis  --go_out=./project_go_proto --go-grpc_out=require_unimplemented_servers=false:./project_go_proto --grpc-gateway_out=logtostderr=true:./project_go_proto
//go:generate protoc ../../proto/v1/project.proto
//go:generate mv project_go_proto tmpp
//go:generate mv tmpp/github.com/grafeas/grafeas/proto/v1/project_go_proto .
//go:generate rm -rf tmpp
package v1
