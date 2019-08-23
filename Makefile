## The lines below can be uncommented for debugging the make rules
#
# OLD_SHELL := $(SHELL)
# SHELL = $(warning Building $@$(if $<, (from $<))$(if $?, ($? newer)))$(OLD_SHELL)
#
# print-%:
# 	@echo $* = $($*)

.PHONY: build fmt test vet clean generate

SRC = $(shell find . -type f -name '*.go' -not -path "./protodeps/*")
CLEAN := *~

.EXPORT_ALL_VARIABLES:

GO111MODULE=on

build: vet fmt generate
	go build -v ./...

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@gofmt -l -w $(SRC)

test: generate
	@go test -v ./...

vet: generate
	@go vet ./...

generate:
	# protoc and tools need to be run before all of the other generations.
	go generate ./protoc
	cd tools && go generate
	go generate ./protodeps
	go generate ./cel
	go generate ./proto/v1
	go generate ./proto/v1alpha1
	go generate ./proto/v1beta1

clean:
	go clean ./...
	rm -rf $(CLEAN)
