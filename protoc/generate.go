// This downloads and installs the protobuf compiler

//go:generate rm -rf include
//go:generate rm -rf bin
//go:generate rm -rf readme.txt
//go:generate rm -rf protoc.zip
//go:generate curl https://github.com/protocolbuffers/protobuf/releases/download/v3.7.1/protoc-3.7.1-linux-x86_64.zip -o protoc/protoc.zip -L
//go:generate unzip protoc.zip -d .
package protoc
