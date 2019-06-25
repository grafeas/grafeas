## The lines below can be uncommented for debugging the make rules
#
# OLD_SHELL := $(SHELL)
# SHELL = $(warning Building $@$(if $<, (from $<))$(if $?, ($? newer)))$(OLD_SHELL)
#
# print-%:
# 	@echo $* = $($*)

.PHONY: build fmt test vet clean generate grafeas_go_v1alpha1 swagger_docs

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
build: vet fmt generate swagger_docs
	go build -v ./...

# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	@gofmt -l -w $(SRC)

test: generate
	@go test -v ./...

vet: generate
	@go vet ./...

generate:
	go generate ./protoc
	go generate ./tools
	go generate ./...

PROTOC_CMD=protoc/bin/protoc -I ./ \
	-I vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	-I vendor/github.com/grpc-ecosystem/grpc-gateway \
	-I vendor/github.com/googleapis/googleapis

swagger_docs: proto/v1beta1/swagger/*.swagger.json

proto/v1beta1/swagger/%.swagger.json: proto/v1beta1/%.proto generate
	$(PROTOC_CMD) --swagger_out=logtostderr=true:. $<
	mv $(<D)/*.swagger.json $@

clean:
	go clean ./...
	rm -rf $(CLEAN)
