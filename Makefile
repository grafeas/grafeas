## The lines below can be uncommented for debugging the make rules
#
# OLD_SHELL := $(SHELL)
# SHELL = $(warning Building $@$(if $<, (from $<))$(if $?, ($? newer)))$(OLD_SHELL)
#
# print-%:
# 	@echo $* = $($*)

.PHONY: build fmt test vet clean generate

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
CLEAN := *~

default: .check_makefile_in_gopath build

EXPECTED_MAKE := $(shell go env GOPATH)/src/github.com/grafeas/grafeas/Makefile

.check_makefile_in_gopath:
	if [ "$(realpath ${EXPECTED_MAKE})" != "$(realpath $(lastword $(MAKEFILE_LIST)))" ]; \
	then  \
	echo "Makefile is not in GOPATH root"; \
	false; \
	fi

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
	go generate ./tools
	go generate ./...

clean:
	go clean ./...
	rm -rf $(CLEAN)
