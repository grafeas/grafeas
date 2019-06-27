// This generates the protocol buffer and swagger code in go for the v1beta1 proto spec.

//go:generate -command protoc ../../protoc/bin/protoc -I ../../ -I ./ -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway -I ../../vendor/github.com/googleapis/googleapis --go_out=plugins=grpc,paths=source_relative:.  --grpc-gateway_out=logtostderr=true,paths=source_relative:.
//go:generate -command swagger ../../protoc/bin/protoc -I ../../ -I ./ -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I ../../vendor/github.com/grpc-ecosystem/grpc-gateway -I ../../vendor/github.com/googleapis/googleapis --swagger_out=logtostderr=true:.
//go:generate rm -rf swagger
//go:generate mkdir swagger

//go:generate protoc attestation.proto
//go:generate rm -rf attestation_go_proto
//go:generate mkdir attestation_go_proto
//go:generate mv attestation.pb.go attestation_go_proto
//go:generate swagger attestation.proto
//go:generate mv attestation.swagger.json swagger

//go:generate protoc common.proto
//go:generate rm -rf common_go_proto
//go:generate mkdir common_go_proto
//go:generate mv common.pb.go common_go_proto
//go:generate swagger common.proto
//go:generate mv common.swagger.json swagger

//go:generate protoc deployment.proto
//go:generate rm -rf deployment_go_proto
//go:generate mkdir deployment_go_proto
//go:generate mv deployment.pb.go deployment_go_proto
//go:generate swagger deployment.proto
//go:generate mv deployment.swagger.json swagger

//go:generate protoc grafeas.proto
//go:generate rm -rf grafeas_go_proto
//go:generate mkdir grafeas_go_proto
//go:generate mv grafeas.pb.go grafeas_go_proto
//go:generate mv grafeas.pb.gw.go grafeas_go_proto
//go:generate swagger grafeas.proto
//go:generate mv grafeas.swagger.json swagger

//go:generate protoc package.proto
//go:generate rm -rf package_go_proto
//go:generate mkdir package_go_proto
//go:generate mv package.pb.go package_go_proto
//go:generate swagger package.proto
//go:generate mv package.swagger.json swagger

//go:generate protoc source.proto
//go:generate rm -rf source_go_proto
//go:generate mkdir source_go_proto
//go:generate mv source.pb.go source_go_proto
//go:generate swagger source.proto
//go:generate mv source.swagger.json swagger

//go:generate protoc provenance.proto
//go:generate rm -rf provenance_go_proto
//go:generate mkdir provenance_go_proto
//go:generate mv provenance.pb.go provenance_go_proto
//go:generate swagger provenance.proto
//go:generate mv provenance.swagger.json swagger

//go:generate protoc build.proto
//go:generate rm -rf build_go_proto
//go:generate mkdir build_go_proto
//go:generate mv build.pb.go build_go_proto
//go:generate swagger build.proto
//go:generate mv build.swagger.json swagger

//go:generate protoc cvss.proto
//go:generate rm -rf cvss_go_proto
//go:generate mkdir cvss_go_proto
//go:generate mv cvss.pb.go cvss_go_proto
//go:generate swagger cvss.proto
//go:generate mv cvss.swagger.json swagger

//go:generate protoc discovery.proto
//go:generate rm -rf discovery_go_proto
//go:generate mkdir discovery_go_proto
//go:generate mv discovery.pb.go discovery_go_proto
//go:generate swagger discovery.proto
//go:generate mv discovery.swagger.json swagger

//go:generate protoc image.proto
//go:generate rm -rf image_go_proto
//go:generate mkdir image_go_proto
//go:generate mv image.pb.go image_go_proto
//go:generate swagger image.proto
//go:generate mv image.swagger.json swagger

//go:generate protoc vulnerability.proto
//go:generate rm -rf vulnerability_go_proto
//go:generate mkdir vulnerability_go_proto
//go:generate mv vulnerability.pb.go vulnerability_go_proto
//go:generate swagger vulnerability.proto
//go:generate mv vulnerability.swagger.json swagger

//go:generate protoc project.proto
//go:generate rm -rf project_go_proto
//go:generate mkdir project_go_proto
//go:generate mv project.pb.go project_go_proto
//go:generate mv project.pb.gw.go project_go_proto
//go:generate swagger project.proto
//go:generate mv project.swagger.json swagger
package v1beta1
