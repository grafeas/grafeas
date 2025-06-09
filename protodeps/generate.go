//go:generate rm -rf grpc-gateway.zip
//go:generate rm -rf grpc-gateway
//go:generate rm -rf grpc-gateway-2.0.1
//go:generate curl https://github.com/grpc-ecosystem/grpc-gateway/archive/v2.0.1.zip -o grpc-gateway.zip -L
//go:generate unzip grpc-gateway.zip -d .
//sleep gives prior command time to release file handles
//go:generate sleep .25
//go:generate mv grpc-gateway-2.0.1 grpc-gateway
//go:generate rm -rf googleapis.zip
//go:generate rm -rf googleapis
//go:generate rm -rf googleapis-cc520460fa6b89750bc3578539f2f436c827d956
//go:generate curl https://github.com/googleapis/googleapis/archive/cc520460fa6b89750bc3578539f2f436c827d956.zip -o googleapis.zip -L
//go:generate unzip googleapis.zip
//sleep gives prior command time to release file handles
//go:generate sleep .25
//go:generate mv googleapis-cc520460fa6b89750bc3578539f2f436c827d956 googleapis

package protodeps
