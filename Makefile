.PHONY: build fmt test vet clean

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
CLEAN := *~

default: build

install.tools: .install.protoc-gen-go .install.grpc-gateway .install.googleapis

CLEAN += .install.protoc-gen-go .install.grpc-gateway .install.googleapis

.install.protoc-gen-go:
	cd vendor/github.com/golang/protobuf/protoc-gen-go && go install .

.install.googleapis:
	git clone https://github.com/googleapis/googleapis
	touch $@

.install.grpc-gateway:
	cd vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && go install .
	cd vendor/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger && go install .

prepare: ensure install.tools

ensure:
	dep ensure

build: vet fmt grafeas_go
	go build -v ./...

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@gofmt -l -w $(SRC)

test:
	@go test -v ./...

vet:
	@go tool vet ${SRC}

v1alpha1/proto/grafeas.pb.go: .install.protoc-gen-go .install.grpc-gateway .install.googleapis v1alpha1/proto/grafeas.proto
	protoc \
		-I ./ \
		-I ./include \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I googleapis \
		--go_out=plugins=grpc:. \
	    --grpc-gateway_out=logtostderr=true:. \
        --swagger_out=logtostderr=true:. \
	    v1alpha1/proto/grafeas.proto


.PHONY: grafeas_go
grafeas_go: v1alpha1/proto/grafeas.pb.go

clean:
	go clean ./...
	rm -f $(CLEAN)
