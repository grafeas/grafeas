## Protobuf Field Mask utils for Go

[![Build Status](https://travis-ci.org/mennanov/fieldmask-utils.svg?branch=master)](https://travis-ci.org/mennanov/fieldmask-utils)
[![Coverage Status](https://coveralls.io/repos/github/mennanov/fieldmask-utils/badge.svg?branch=master)](https://coveralls.io/github/mennanov/fieldmask-utils?branch=master)

Features:

* Copy from any Go struct to any compatible Go struct with a field mask applied
* Copy from any Go struct to a `map[string]interface{}` with a field mask applied
* Extensible masks (e.g. inverse mask: copy all except those mentioned, etc.)
* Supports [Protobuf Any](https://developers.google.com/protocol-buffers/docs/proto3#any) message types.

### Examples

Copy from a protobuf message to a protobuf message:

```proto
// test.proto

message UpdateUserRequest {
    User user = 1;
    google.protobuf.FieldMask field_mask = 2;
}
```

```go
import "github.com/golang/protobuf/protoc-gen-go/generator"

var request UpdateUserRequest
userDst := &testproto.User{} // a struct to copy to
mask, err := fieldmask_utils.MaskFromPaths(request.FieldMask.Paths, generator.CamelCase)
// handle err...
fieldmask_utils.StructToStruct(mask, request.User, userDst)
// Only the fields mentioned in the field mask will be copied to userDst, other fields are left intact
```

Copy from a protobuf message to a `map[string]interface{}`:

```go
import "github.com/golang/protobuf/protoc-gen-go/generator"

var request UpdateUserRequest
userDst := make(map[string]interface{}) // a map to copy to
mask, err := fieldmask_utils.MaskFromProtoFieldMask(request.FieldMask, generator.CamelCase)
// handle err...
err := fieldmask_utils.StructToMap(mask, request.User, userDst)
// handle err..
// Only the fields mentioned in the field mask will be copied to userDst, other fields are left intact
```

Copy with an inverse mask:

```go
import "github.com/golang/protobuf/protoc-gen-go/generator"

var request UpdateUserRequest
userDst := &testproto.User{} // a struct to copy to
mask := fieldmask_utils.MaskInverse{"Id": nil, "Friends": fieldmask_utils.MaskInverse{"Username": nil}}
fieldmask_utils.StructToStruct(mask, request.User, userDst)
// Only the fields that are not mentioned in the field mask will be copied to userDst, other fields are left intact.
```

### Limitations

1.  Larger scope field masks have no effect and are not considered invalid:

    field mask strings `"a", "a.b", "a.b.c"` will result in a mask `a{b{c}}`, which is the same as `"a.b.c"`.

2.  Masks inside a protobuf `Map` are not supported.
3.  When copying from a struct to struct the destination struct must have the same fields (or a subset)
    as the source struct. Pointers must also be coherent: if a field is a pointer in the source struct, then
    it also must be a pointer (not a value field) in the destination struct.
