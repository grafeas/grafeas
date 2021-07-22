// This generates the protocol buffer code in go for the v1 proto spec.

//go:generate rm -rf grafeas_go_proto
//go:generate mkdir grafeas_go_proto
//go:generate -command protoc ../../protoc/bin/protoc -I ../../ -I ./ -I ../../protodeps/grpc-gateway/third_party/googleapis -I ../../protodeps/grpc-gateway -I ../../protodeps/googleapis  --go_out=./grafeas_go_proto --go-grpc_out=require_unimplemented_servers=false:./grafeas_go_proto --grpc-gateway_out=logtostderr=true:./grafeas_go_proto
//go:generate protoc ../../proto/v1/attestation.proto
//go:generate protoc ../../proto/v1/common.proto
//go:generate protoc ../../proto/v1/deployment.proto
//go:generate protoc ../../proto/v1/intoto_provenance.proto
//go:generate protoc ../../proto/v1/dsse_attestation.proto
//go:generate protoc ../../proto/v1/grafeas.proto
//go:generate protoc ../../proto/v1/package.proto
//go:generate protoc ../../proto/v1/provenance.proto
//go:generate protoc ../../proto/v1/build.proto
//go:generate protoc ../../proto/v1/cvss.proto
//go:generate protoc ../../proto/v1/discovery.proto
//go:generate protoc ../../proto/v1/image.proto
//go:generate protoc ../../proto/v1/vulnerability.proto
//go:generate protoc ../../proto/v1/upgrade.proto
//go:generate protoc ../../proto/v1/compliance.proto
//go:generate mv grafeas_go_proto tmp
//go:generate mv tmp/github.com/grafeas/grafeas/proto/v1/grafeas_go_proto .
//go:generate rm -rf tmp
package v1
