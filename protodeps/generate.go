//go:generate rm -rf grpc-gateway.zip
//go:generate rm -rf grpc-gateway
//go:generate rm -rf grpc-gateway-1.9.6
//go:generate curl https://github.com/grpc-ecosystem/grpc-gateway/archive/v1.9.6.zip -o grpc-gateway.zip -L
//go:generate unzip grpc-gateway.zip -d .
//go:generate mv grpc-gateway-1.9.6 grpc-gateway
//go:generate rm -rf googleapis.zip
//go:generate rm -rf googleapis
//go:generate rm -rf googleapis-fb6fa4cfb16917da8dc5d23c2494d422dd3e9cd4
//go:generate curl https://github.com/googleapis/googleapis/archive/fb6fa4cfb16917da8dc5d23c2494d422dd3e9cd4.zip -o googleapis.zip -L
//go:generate unzip googleapis.zip
//go:generate mv googleapis-fb6fa4cfb16917da8dc5d23c2494d422dd3e9cd4 googleapis

package protodeps
