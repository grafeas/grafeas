//go:generate -command protoc ../../protoc/bin/protoc -I ../../ -I ./ -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway -I ../../vendor/github.com/googleapis/googleapis --go_out=plugins=grpc,paths=source_relative:.  --grpc-gateway_out=logtostderr=true,paths=source_relative:.
//go:generate protoc attestation.proto
//go:generate rm -rf attestation_go_proto
//go:generate mkdir attestation_go_proto
//go:generate mv attestation.pb.go attestation_go_proto
//go:generate protoc common.proto
//go:generate rm -rf common_go_proto
//go:generate mkdir common_go_proto
//go:generate mv common.pb.go common_go_proto
//go:generate protoc deployment.proto
//go:generate rm -rf deployment_go_proto
//go:generate mkdir deployment_go_proto
//go:generate mv deployment.pb.go deployment_go_proto
//go:generate protoc grafeas.proto
//go:generate rm -rf grafeas_go_proto
//go:generate mkdir grafeas_go_proto
//go:generate mv grafeas.pb.go grafeas_go_proto
//go:generate mv grafeas.pb.gw.go grafeas_go_proto
//go:generate protoc package.proto
//go:generate rm -rf package_go_proto
//go:generate mkdir package_go_proto
//go:generate mv package.pb.go package_go_proto
//go:generate protoc provenance.proto
//go:generate rm -rf provenance_go_proto
//go:generate mkdir provenance_go_proto
//go:generate mv provenance.pb.go provenance_go_proto
//go:generate protoc build.proto
//go:generate rm -rf build_go_proto
//go:generate mkdir build_go_proto
//go:generate mv build.pb.go build_go_proto
//go:generate protoc cvss.proto
//go:generate rm -rf cvss_go_proto
//go:generate mkdir cvss_go_proto
//go:generate mv cvss.pb.go cvss_go_proto
//go:generate protoc discovery.proto
//go:generate rm -rf discovery_go_proto
//go:generate mkdir discovery_go_proto
//go:generate mv discovery.pb.go discovery_go_proto
//go:generate protoc image.proto
//go:generate rm -rf image_go_proto
//go:generate mkdir image_go_proto
//go:generate mv image.pb.go image_go_proto
//go:generate protoc vulnerability.proto
//go:generate rm -rf vulnerability_go_proto
//go:generate mkdir vulnerability_go_proto
//go:generate mv vulnerability.pb.go vulnerability_go_proto
package v1beta1
