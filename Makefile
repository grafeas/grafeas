.PHONY: build fmt test vet clean go_protos grafeas_go_v1alpha1 swagger_docs 

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
CLEAN := *~

default: build

install.tools: .install.protoc-gen-go .install.grpc-gateway

CLEAN += .install.protoc-gen-go .install.grpc-gateway
.install.protoc-gen-go:
	go get -u -v github.com/golang/protobuf/protoc-gen-go && touch $@

.install.grpc-gateway:
	go get -u -v github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && touch $@

build: vet fmt go_protos swagger_docs
	go build -v ./...

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@gofmt -l -w $(SRC)

test:
	@go test -v ./... 

vet:
	@go tool vet ${SRC}

protoc/bin/protoc:
	mkdir -p protoc
	curl https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip -o protoc/protoc.zip -L
	unzip protoc/protoc -d protoc

CLEAN += protoc

go_protos: grafeas_go_v1alpha1 proto/v1beta1/*_go_proto proto/v1/*_go_proto

PROTOC_CMD=protoc/bin/protoc -I ./ \
	-I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I vendor/github.com/grpc-ecosystem/grpc-gateway \
	-I vendor/github.com/googleapis/googleapis

grafeas_go_v1alpha1: .install.protoc-gen-go .install.grpc-gateway v1alpha1/proto/grafeas.proto protoc/bin/protoc
	$(PROTOC_CMD) \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
		v1alpha1/proto/grafeas.proto

# Builds go proto packages (e.g. grafeas_go_proto/grafeas.pb.go and grafeas_go_proto/grafeas.pb.gw.go)
%_go_proto: %.proto protoc/bin/protoc install.tools
	$(PROTOC_CMD) \
		--go_out=plugins=grpc,paths=source_relative:. \
		--grpc-gateway_out=logtostderr=true,paths=source_relative:. \
		$<
	mv $*.pb.go $@
	if [ -f $*.pb.gw.go ]; then mv $*.pb.gw.go $@; fi

swagger_docs: proto/v1beta1/swagger/*.swagger.json

proto/v1beta1/swagger/%.swagger.json: proto/v1beta1/%.proto protoc/bin/protoc install.tools
	$(PROTOC_CMD) --swagger_out=logtostderr=true:. $<
	mv $(<D)/*.swagger.json $@

clean:
	go clean ./...
	rm -rf $(CLEAN)
