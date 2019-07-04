# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: proto test test-with-coverage

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . -I ${GOPATH}/src test.proto --go_out=${GOPATH}/src

test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -v ./...


test-with-coverage:
	${GOPATH}/bin/goveralls -service=travis-pro -ignore=testproto/*
