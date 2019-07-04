package fieldmask_utils

import (
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/pkg/errors"
)

// StructToStruct copies `src` struct to `dst` struct using the given FieldFilter.
// Only the fields where FieldFilter returns true will be copied to `dst`.
// `src` and `dst` must be coherent in terms of the field names, but it is not required for them to be of the same type.
// Unexported fields are copied only if the corresponding struct filter is empty and `dst` is assignable to `src`.
func StructToStruct(filter FieldFilter, src, dst interface{}) error {
	dstVal := reflect.ValueOf(dst)
	if dstVal.Kind() != reflect.Ptr {
		return errors.Errorf("dst must be a pointer, %s given", dstVal.Kind())
	}
	srcVal := indirect(reflect.ValueOf(src))
	if srcVal.Kind() != reflect.Struct {
		return errors.Errorf("src kind must be a struct, %s given", srcVal.Kind())
	}
	dstVal = indirect(dstVal)
	if dstVal.Kind() != reflect.Struct {
		return errors.Errorf("dst kind must be a struct, %s given", dstVal.Kind())
	}
	return structToStruct(filter, &srcVal, &dstVal)
}

func structToStruct(filter FieldFilter, src, dst *reflect.Value) error {
	if src.Kind() != dst.Kind() {
		return errors.Errorf("src kind %s differs from dst kind %s", src.Kind(), dst.Kind())
	}

	switch src.Kind() {
	case reflect.Struct:
		if srcAny, ok := src.Interface().(any.Any); ok {
			dstAny, ok := src.Interface().(any.Any)
			if !ok {
				return errors.Errorf("dst type is %s, expected: %s ", dst.Type(), "any.Any")
			}

			newSrcProto, err := ptypes.Empty(&srcAny)
			if err != nil {
				return errors.WithStack(err)
			}
			if err := ptypes.UnmarshalAny(&srcAny, newSrcProto); err != nil {
				return errors.WithStack(err)
			}
			newSrc := reflect.ValueOf(newSrcProto)

			newDstProto, err := ptypes.Empty(&dstAny)
			if err != nil {
				return errors.WithStack(err)
			}
			if err := ptypes.UnmarshalAny(&dstAny, newDstProto); err != nil {
				return errors.WithStack(err)
			}
			newDst := reflect.ValueOf(newDstProto)

			if err := structToStruct(filter, &newSrc, &newDst); err != nil {
				return err
			}

			newDstAny, err := ptypes.MarshalAny(newDst.Interface().(proto.Message))
			if err != nil {
				return errors.WithStack(err)
			}

			dst.Set(reflect.ValueOf(*newDstAny))
			return nil
		}

		if dst.CanSet() && dst.Type().AssignableTo(src.Type()) && filter.IsEmpty() {
			dst.Set(*src)
			return nil
		}

		for i := 0; i < src.NumField(); i++ {
			fieldName := src.Type().Field(i).Name

			subFilter, ok := filter.Filter(fieldName)
			if !ok {
				// Skip this field.
				continue
			}

			dstField := dst.FieldByName(fieldName)
			if !dstField.CanSet() {
				return errors.Errorf("Can't set a value on a destination field %s", fieldName)
			}

			srcField := src.FieldByName(fieldName)
			if err := structToStruct(subFilter, &srcField, &dstField); err != nil {
				return err
			}
		}

	case reflect.Ptr:
		if src.IsNil() {
			// If src is nil set dst to nil too.
			dst.Set(reflect.Zero(dst.Type()))
			break
		}
		if dst.IsNil() {
			// If dst is nil create a new instance of the underlying type and set dst to the pointer of that instance.
			dst.Set(reflect.New(dst.Type().Elem()))
		}

		srcElem, dstElem := src.Elem(), dst.Elem()
		if err := structToStruct(filter, &srcElem, &dstElem); err != nil {
			return err
		}

	case reflect.Interface:
		if src.IsNil() {
			// If src is nil set dst to nil too.
			dst.Set(reflect.Zero(dst.Type()))
			break
		}
		if dst.IsNil() {
			if src.Elem().Kind() != reflect.Ptr {
				// Non-pointer interface implementations are not addressable.
				return errors.Errorf("expected a pointer for an interface value, got %s instead", src.Elem().Kind())
			}
			dst.Set(reflect.New(src.Elem().Elem().Type()))
		}

		srcElem, dstElem := src.Elem(), dst.Elem()
		if err := structToStruct(filter, &srcElem, &dstElem); err != nil {
			return err
		}

	case reflect.Slice:
		dstLen := dst.Len()
		for i := 0; i < src.Len(); i++ {
			srcItem := src.Index(i)
			var dstItem reflect.Value
			if i < dstLen {
				// Use an existing item.
				dstItem = dst.Index(i)
			} else {
				// Create a new item if needed.
				dstItem = reflect.New(dst.Type().Elem()).Elem()
			}

			if err := structToStruct(filter, &srcItem, &dstItem); err != nil {
				return err
			}

			if i >= dstLen {
				// Append newly created items to the slice.
				dst.Set(reflect.Append(*dst, dstItem))
			}
		}

	case reflect.Array:
		dstLen := dst.Len()
		if dstLen < src.Len() {
			return errors.Errorf("dst array size %d is less than src size %d", dstLen, src.Len())
		}
		for i := 0; i < src.Len(); i++ {
			srcItem := src.Index(i)
			dstItem := dst.Index(i)
			if err := structToStruct(filter, &srcItem, &dstItem); err != nil {
				return errors.WithStack(err)
			}
		}

	default:
		if !dst.CanSet() {
			return errors.Errorf("dst %s, %s is not settable", dst, dst.Type())
		}
		dst.Set(*src)
	}

	return nil
}

// StructToMap copies `src` struct to the `dst` map.
// Behavior is similar to `StructToStruct`.
func StructToMap(filter FieldFilter, src interface{}, dst map[string]interface{}) error {
	srcVal := indirect(reflect.ValueOf(src))
	srcType := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		fieldName := srcType.Field(i).Name
		subFilter, ok := filter.Filter(fieldName)
		if !ok {
			// Skip this field.
			continue
		}
		srcField := srcVal.FieldByName(fieldName)

		switch srcField.Kind() {
		case reflect.Ptr, reflect.Interface:
			if srcField.IsNil() {
				dst[fieldName] = nil
				continue
			}

			var newValue map[string]interface{}
			existingValue, ok := dst[fieldName]
			if ok {
				newValue = existingValue.(map[string]interface{})
			} else {
				newValue = make(map[string]interface{})
			}
			if err := StructToMap(subFilter, srcField.Interface(), newValue); err != nil {
				return err
			}
			dst[fieldName] = newValue

		case reflect.Array, reflect.Slice:
			// Check if it is a slice of primitive values.
			itemKind := srcField.Type().Elem().Kind()
			if itemKind != reflect.Ptr && itemKind != reflect.Struct && itemKind != reflect.Interface {
				// Handle this array/slice as a regular non-nested data structure: copy it entirely to dst.
				if srcField.Len() > 0 {
					dst[fieldName] = srcField.Interface()
				} else {
					dst[fieldName] = []interface{}(nil)
				}
				continue
			}

			var newValue []map[string]interface{}
			existingValue, ok := dst[fieldName]
			if ok {
				newValue = existingValue.([]map[string]interface{})
			} else {
				newValue = make([]map[string]interface{}, srcField.Len())
			}

			// Iterate over items of the slice/array.
			dstLen := len(newValue)
			srcLen := srcField.Len()
			if dstLen < srcLen {
				return errors.Errorf("dst slice len %d is less than src slice len %d", dstLen, srcLen)
			}
			for i := 0; i < srcLen; i++ {
				subValue := srcField.Index(i)
				if newValue[i] == nil {
					newValue[i] = make(map[string]interface{})
				}
				newDst := newValue[i]
				if err := StructToMap(subFilter, subValue.Interface(), newDst); err != nil {
					return err
				}
				if i < dstLen {
					newValue[i] = newDst
				} else {
					newValue = append(newValue, newDst)
				}
			}
			dst[fieldName] = newValue

		case reflect.Struct:
			var newValue map[string]interface{}
			existingValue, ok := dst[fieldName]
			if ok {
				newValue = existingValue.(map[string]interface{})
			} else {
				newValue = make(map[string]interface{})
			}
			if err := StructToMap(subFilter, srcField.Interface(), newValue); err != nil {
				return err
			}
			dst[fieldName] = newValue

		default:
			// Set a value on a map.
			dst[fieldName] = srcField.Interface()
		}
	}
	return nil
}

func indirect(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	return v
}
