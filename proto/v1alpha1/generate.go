// This file generates the v1alpha1 proto bindings for go.

//go:generate -command protoc ../../protoc/bin/protoc -I ../../ -I ./ -I ../../google/grpc-gateway/third_party/googleapis -I ../../google/grpc-gateway -I ../../google/googleapis --go_out=plugins=grpc,paths=source_relative:.  --grpc-gateway_out=logtostderr=true,paths=source_relative:.
//go:generate -command swagger ../../protoc/bin/protoc -I ../../ -I ./ -I ../../google/grpc-gateway/third_party/googleapis -I ../../google/grpc-gateway -I ../../google/googleapis --swagger_out=logtostderr=true:.
//go:generate rm -rf swagger
//go:generate mkdir swagger

//go:generate protoc grafeas.proto
//go:generate swagger grafeas.proto
//go:generate mv grafeas.swagger.json swagger
package grafeas
