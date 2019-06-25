// This generates the protocol buffer code in go for the v1beta1 proto spec.

//go:generate rm -rf grafeas_go_proto
//go:generate mkdir grafeas_go_proto
//go:generate -command protoc ../../protoc/bin/protoc -I ../../ -I ./ -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway -I ../../vendor/github.com/googleapis/googleapis --go_out=plugins=grpc,paths=source_relative:.  --grpc-gateway_out=logtostderr=true,paths=source_relative:.
//go:generate protoc attestation.proto
//go:generate mv attestation.pb.go grafeas_go_proto
//go:generate protoc common.proto
//go:generate mv common.pb.go grafeas_go_proto
//go:generate protoc deployment.proto
//go:generate mv deployment.pb.go grafeas_go_proto
//go:generate protoc grafeas.proto
//go:generate mv grafeas.pb.go grafeas_go_proto
//go:generate mv grafeas.pb.gw.go grafeas_go_proto
//go:generate protoc package.proto
//go:generate mv package.pb.go grafeas_go_proto
//go:generate protoc provenance.proto
//go:generate mv provenance.pb.go grafeas_go_proto
//go:generate protoc build.proto
//go:generate mv build.pb.go grafeas_go_proto
//go:generate protoc cvss.proto
//go:generate mv cvss.pb.go grafeas_go_proto
//go:generate protoc discovery.proto
//go:generate mv discovery.pb.go grafeas_go_proto
//go:generate protoc image.proto
//go:generate mv image.pb.go grafeas_go_proto
//go:generate protoc vulnerability.proto
//go:generate mv vulnerability.pb.go grafeas_go_proto
package v1
