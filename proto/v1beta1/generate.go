// This generates the protocol buffer and swagger code in go for the v1beta1 proto spec.

//go:generate -command protoc ../../protoc/bin/protoc -I ../../ -I ./ -I ../../protodeps/grpc-gateway/third_party/googleapis -I ../../protodeps/grpc-gateway -I ../../protodeps/googleapis  --go_out=paths=source_relative:. --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:. --grpc-gateway_out=logtostderr=true,paths=source_relative:.
//go:generate -command swagger ../../protoc/bin/protoc -I ../../ -I ./ -I ../../protodeps/grpc-gateway/third_party/googleapis -I ../../protodeps/grpc-gateway -I ../../protodeps/googleapis --openapiv2_out=logtostderr=true:.
//go:generate rm -rf swagger
//go:generate mkdir swagger

// ==================
// Swagger Generation
// ==================
//
// We generate grafeas.swagger.json and project.swagger.json for backwards
// compatibility. A merged Swagger file is also generated in swagger/merged that
// merges both services into a single Swagger file for easier client generation.
//
// NOTE: You should only generate Swagger for a proto if it contains any
// `Service` definitions.

//go:generate swagger grafeas.proto
//go:generate mv grafeas.swagger.json swagger

//go:generate swagger project.proto
//go:generate mv project.swagger.json swagger

//go:generate mkdir swagger/merged
//go:generate swagger --openapiv2_opt=allow_merge=true,merge_file_name=grafeas grafeas.proto project.proto
//go:generate mv grafeas.swagger.json swagger/merged

// ===================
// Go Proto Generation
// ===================

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
//go:generate mv grafeas_grpc.pb.go grafeas_go_proto
//go:generate mv grafeas.pb.gw.go grafeas_go_proto

//go:generate protoc package.proto
//go:generate rm -rf package_go_proto
//go:generate mkdir package_go_proto
//go:generate mv package.pb.go package_go_proto

//go:generate protoc source.proto
//go:generate rm -rf source_go_proto
//go:generate mkdir source_go_proto
//go:generate mv source.pb.go source_go_proto

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

//go:generate protoc intoto.proto
//go:generate rm -rf intoto_go_proto
//go:generate mkdir intoto_go_proto
//go:generate mv intoto.pb.go intoto_go_proto

//go:generate protoc project.proto
//go:generate rm -rf project_go_proto
//go:generate mkdir project_go_proto
//go:generate mv project.pb.go project_go_proto
//go:generate mv project_grpc.pb.go project_go_proto
//go:generate mv project.pb.gw.go project_go_proto

//go:generate protoc spdx.proto
//go:generate rm -rf spdx_go_proto
//go:generate mkdir spdx_go_proto
//go:generate mv spdx.pb.go spdx_go_proto

//go:generate protoc vex.proto
//go:generate rm -rf vex_go_proto
//go:generate mkdir vex_go_proto
//go:generate mv vex.pb.go vex_go_proto

//go:generate protoc sbom_reference.proto
//go:generate rm -rf sbom_reference_go_proto
//go:generate mkdir sbom_reference_go_proto
//go:generate mv sbom_reference.pb.go sbom_reference_go_proto

package v1beta1
