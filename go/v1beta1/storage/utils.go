package storage

import (
	pb "github.com/grafeas/grafeas/proto/v1beta1/grafeas_go_proto"
	utils "github.com/mennanov/fieldmask-utils"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/generator"
	"google.golang.org/genproto/protobuf/field_mask"
)

// CreateFieldMask creates a new field mask from a given array of paths
func CreateFieldMask(paths []string) *field_mask.FieldMask {
	if paths == nil {
		paths = []string{}
	}
	return &field_mask.FieldMask{
		Paths: paths,
	}
}

// ApplyUpdateOnOccurrence applies an update (src) on an existing Occurrence (dst) indicated by a field mask (updates)
func ApplyUpdateOnOccurrence(dst, src *pb.Occurrence, updates *field_mask.FieldMask) (*pb.Occurrence, error) {
	v := proto.Clone(dst).(*pb.Occurrence)
	mask, err := utils.MaskFromPaths(updates.GetPaths(), generator.CamelCase)
	if err != nil {
		return nil, err
	}

	err = utils.StructToStruct(mask, src, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// ApplyUpdateOnNote applies an update (src) on an existing Note (dst) indicated by a field mask (updates)
func ApplyUpdateOnNote(dst, src *pb.Note, updates *field_mask.FieldMask) (*pb.Note, error) {
	v := proto.Clone(dst).(*pb.Note)
	mask, err := utils.MaskFromPaths(updates.GetPaths(), generator.CamelCase)
	if err != nil {
		return nil, err
	}

	err = utils.StructToStruct(mask, src, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
