.PHONY: build fmt test vet clean

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
CLEAN := *~

default: build

install.tools: .install.protoc-gen-go

CLEAN += .install.protoc-gen-go
.install.protoc-gen-go:
	go get -u -v github.com/golang/protobuf/protoc-gen-go && touch $@

build:  vet fmt grafeas_go
	go build -v ./...

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@gofmt -l -w $(SRC)

test:
	@go test -v ./... 

vet:
	@go tool vet ${SRC}

model-go/v1alpha1/v1alpha1/proto/grafeas.pb.go: .install.protoc-gen-go v1alpha1/proto/grafeas.proto
	@mkdir -p ./model-go/v1alpha1/
	protoc \
		-I. \
		-I ./vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I ./vendor/github.com/googleapis/googleapis \
		--go_out=./v1alpha1/ \
		v1alpha1/proto/grafeas.proto

.PHONY: grafeas_go
grafeas_go: model-go/v1alpha1/v1alpha1/proto/grafeas.pb.go

clean:
	go clean ./...
	rm -f $(CLEAN)

