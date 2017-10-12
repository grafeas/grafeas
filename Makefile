
API_VERSION ?= v1alpha1

# Prepend our _vendor directory to the system GOPATH
# # so that import path resolution will prioritize
# # our third party snapshots.
GOPATH := ${PWD}/vendor:${GOPATH}
export GOPATH

CODEGEN_JAR ?= /tmp/swagger-codegen-cli.jar

.PHONY: build fmt lint test vet

build:  vet fmt 
	go build -v ./...


# http://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt ./...

test:
	@go test -v ./...

# http://godoc.org/code.google.com/p/go.tools/cmd/vet
# go get code.google.com/p/go.tools/cmd/vet
vet:
	go vet ./...

$(CODEGEN_JAR):
	wget http://central.maven.org/maven2/io/swagger/swagger-codegen-cli/2.2.3/swagger-codegen-cli-2.2.3.jar -O $(CODEGEN_JAR)

# uses goimports from golang.org/x/tools/cmd/goimports
# since the upstream swagger generates go source that still needs to be formatted.
# https://github.com/swagger-api/swagger-codegen/issues/3518#issuecomment-336438908
server-api: ./$(API_VERSION)/grafeas.json $(CODEGEN_JAR)
	java -jar $(CODEGEN_JAR) generate \
		-l go \
		-i ./v1alpha1/grafeas.json \
		-o ./server-go/api/$(API_VERSION)/ \
		-D packageName=api \
		-D packageVersion=$(API_VERSION)
	goimports -w ./server-go/api
