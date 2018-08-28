.PHONY: build fmt test vet clean grafeas_go

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
CLEAN := *~

default: build

install.tools: .install.protoc-gen-go .install.grpc-gateway

CLEAN += .install.protoc-gen-go .install.grpc-gateway
.install.protoc-gen-go:
	go get -u -v github.com/golang/protobuf/protoc-gen-go && touch $@

.install.grpc-gateway:
	go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && touch $@

build:  vet fmt grafeas_go
	go build -v ./...

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@gofmt -l -w $(SRC)

test:
	@go test -v ./... 

vet:
	@go tool vet ${SRC}

grafeas_go_v1alpha1: .install.protoc-gen-go .install.grpc-gateway v1alpha1/proto/grafeas.proto
	protoc \
		-I ./ \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I vendor/github.com/googleapis/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		v1alpha1/proto/grafeas.proto

define gen_go_proto
	protoc \
		-I ./ \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I vendor/github.com/googleapis/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		proto/v1beta1/$(1).proto && \
	mv proto/v1beta1/$(1).pb.go proto/v1beta1/$(1)_go_proto && \
	mv proto/v1beta1/$(1).swagger.json proto/v1beta1/swagger
endef

attestation_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/attestation.proto proto/v1beta1/attestation_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,attestation)

build_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/build.proto proto/v1beta1/build_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,build)

common_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/common.proto proto/v1beta1/common_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,common)

deployment_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/deployment.proto proto/v1beta1/deployment_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,deployment)

discovery_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/discovery.proto proto/v1beta1/discovery_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,discovery)

grafeas_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/grafeas.proto proto/v1beta1/grafeas_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,grafeas)

image_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/image.proto proto/v1beta1/image_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,image)

package_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/package.proto proto/v1beta1/package_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,package)

project_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/project.proto proto/v1beta1/project_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,project)

provenance_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/provenance.proto proto/v1beta1/provenance_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,provenance)

source_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/source.proto proto/v1beta1/source_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,source)

vulnerability_go_v1beta1: .install.protoc-gen-go .install.grpc-gateway proto/v1beta1/vulnerability.proto proto/v1beta1/vulnerability_go_proto proto/v1beta1/swagger
	$(call gen_go_proto,vulnerability)

clean:
	go clean ./...
	rm -f $(CLEAN)
