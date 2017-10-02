.PHONY: build fmt lint test vet

# Prepend our _vendor directory to the system GOPATH
# # so that import path resolution will prioritize
# # our third party snapshots.
GOPATH := ${PWD}/vendor:${GOPATH}
export GOPATH


build:  vet fmt 
	go build -v ./src/github.com/grafeas/...


# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./src/github.com/grafeas/...

test:
	@go test -v ./src/github.com/grafeas/...


# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./src/github.com/grafeas/...

