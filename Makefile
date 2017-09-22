GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
GOPACKAGES = $(shell go list ./...  | grep -v /vendor/)


test: test-all

test-all: 
	@go test -v $(GOPACKAGES)
