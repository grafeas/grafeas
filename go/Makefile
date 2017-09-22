.PHONY: build fmt lint run test vendor_clean vendor_get vendor_update vet

+GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
+GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)

# Prepend our _vendor directory to the system GOPATH
# # so that import path resolution will prioritize
# # our third party snapshots.
GOPATH := ${PWD}/vendor:${GOPATH}
export GOPATH

default: build

build: vet
	go build -v ./... 

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/...

# https://github.com/golang/lint
# # go get github.com/golang/lint/golint
lint:
	golint ./src
#
run: build

test:
	@go test -v $(GOPACKAGES)

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./src/...

