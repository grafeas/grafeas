// This downloads and installs the protobuf compiler

//go:generate rm -rf include
//go:generate rm -rf bin
//go:generate rm -rf readme.txt
//go:generate rm -rf protoc.zip
//go:generate curl https://github.com/google/protobuf/releases/download/v3.3.0/protoc-3.3.0-linux-x86_64.zip -o protoc.zip -L
//go:generate unzip protoc.zip -d .
package protoc
