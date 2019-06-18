//go:generate curl -O https://raw.githubusercontent.com/googleapis/googleapis/e266a1ca806aac193ee3ed5d6d2031bea351546e/google/api/expr/v1alpha1/syntax.proto
//go:generate ../protoc/bin/protoc syntax.proto --go_out=:.
//go:generate mv google.golang.org/genproto/googleapis/api/expr/v1alpha1/syntax.pb.go syntax.pb.go
//go:generate rm -r google.golang.org
package expr
