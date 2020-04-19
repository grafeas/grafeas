// This downloads and installs the protobuf compiler

//go:generate rm -rf include
//go:generate rm -rf bin
//go:generate rm -rf readme.txt
//go:generate rm -rf protoc.zip
//go:generate ./downloadProtoc.sh
//go:generate unzip protoc.zip -d .
package protoc
