.PHONY: build fmt test vet clean

# Prepend our _vendor directory to the system GOPATH
# # so that import path resolution will prioritize
# # our third party snapshots.
GOPATH := ${PWD}/vendor:${GOPATH}
export GOPATH
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

install:
	go get -u -v github.com/golang/protobuf/protoc-gen-go

build:  vet fmt install grafeas_go
	go build -v ./...


# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@gofmt -l -w $(SRC)

test:
	@go test -v ./... 

protoc_middleman_go: v1alpha1/proto/grafeas.proto
	protoc -I. -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/googleapis/googleapis --go_out=. v1alpha1/proto/grafeas.proto
	@touch protoc_middleman_go

grafeas_go: protoc_middleman_go

vet:
	@go tool vet ${SRC}

clean:
	go clean ./...
	rm -f protoc_middleman_go v1alpha1/proto/*.pb.go

