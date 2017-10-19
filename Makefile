.PHONY: build fmt test vet clean

# Prepend our _vendor directory to the system GOPATH
# # so that import path resolution will prioritize
# # our third party snapshots.
GOPATH := ${PWD}/vendor:${GOPATH}
export GOPATH

build:  vet fmt grafeas_go
	go build -v ./...


# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./... -./vendor/...

test:
	@go test -v ./... -./vendor/...

protoc_middleman_go: v1alpha1/proto/grafeas.proto
	protoc -I. -I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis -I vendor/github.com/googleapis/googleapis --go_out=. v1alpha1/proto/grafeas.proto
	@touch protoc_middleman_go

grafeas_go: protoc_middleman_go

vet:
	go vet ./... -./vendor/... 


clean:
	rm -f protoc_middleman_go v1alpha1/proto/*.pb.go