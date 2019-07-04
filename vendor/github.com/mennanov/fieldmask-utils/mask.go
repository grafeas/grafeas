package fieldmask_utils

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/genproto/protobuf/field_mask"
)

// FieldFilter is an interface used by the copying function to filter fields that are needed to be copied.
type FieldFilter interface {
	// Filter should return a corresponding FieldFilter for the given fieldName and a boolean result. If result is true
	// then the field is copied, skipped otherwise.
	Filter(fieldName string) (FieldFilter, bool)

	// Returns true if the FieldFilter is empty. In this case all fields are copied.
	IsEmpty() bool
}

// Mask is a tree-based implementation of a FieldFilter.
type Mask map[string]FieldFilter

// Compile time interface check.
var _ FieldFilter = Mask{}

// Filter returns true for those fieldNames that exist in the underlying map.
// Field names that start with "XXX_" are ignored as unexported.
func (m Mask) Filter(fieldName string) (FieldFilter, bool) {
	if len(m) == 0 {
		// If the mask is empty choose all the exported fields.
		return Mask{}, !strings.HasPrefix(fieldName, "XXX_")
	}
	subFilter, ok := m[fieldName]
	if !ok {
		subFilter = Mask{}
	}
	return subFilter, ok
}

func (m Mask) IsEmpty() bool {
	return len(m) == 0
}

func mapToString(m map[string]FieldFilter) string {
	if len(m) == 0 {
		return ""
	}
	var result []string
	for fieldName, maskNode := range m {
		r := fieldName
		var sub string
		if stringer, ok := maskNode.(fmt.Stringer); ok {
			sub = stringer.String()
		} else {
			sub = fmt.Sprint(maskNode)
		}
		if sub != "" {
			r += "{" + sub + "}"
		}
		result = append(result, r)
	}
	return strings.Join(result, ",")
}

func (m Mask) String() string {
	return mapToString(m)
}

// MaskInverse is an inversed version of a Mask (will copy all the fields except those mentioned in the mask).
type MaskInverse Mask

// Filter returns true for those fieldNames that do NOT exist in the underlying map.
// Field names that start with "XXX_" are ignored as unexported.
func (m MaskInverse) Filter(fieldName string) (FieldFilter, bool) {
	subFilter, ok := m[fieldName]
	if !ok {
		return MaskInverse{}, !strings.HasPrefix(fieldName, "XXX_")
	}
	return subFilter, subFilter != nil
}

func (m MaskInverse) IsEmpty() bool {
	return len(m) == 0
}

func (m MaskInverse) String() string {
	return mapToString(m)
}

// MaskFromProtoFieldMask creates a Mask from the given FieldMask.
func MaskFromProtoFieldMask(fm *field_mask.FieldMask, naming func(string) string) (Mask, error) {
	return MaskFromPaths(fm.Paths, naming)
}

// MaskFromPaths creates a new Mask from the given paths.
func MaskFromPaths(paths []string, naming func(string) string) (Mask, error) {
	root := make(Mask)
	for _, path := range paths {
		mask := root
		for _, fieldName := range strings.Split(path, ".") {
			if fieldName == "" {
				return nil, errors.Errorf("invalid fieldName FieldFilter format: \"%s\"", path)
			}
			newFieldName := naming(fieldName)
			subNode, ok := mask[newFieldName]
			if !ok {
				mask[newFieldName] = make(Mask)
				subNode = mask[newFieldName]
			}
			mask = subNode.(Mask)
		}
	}
	return root, nil
}

// MaskFromString creates a `Mask` from a string `s`.
// `s` is supposed to be a valid string representation of a FieldFilter like "a,b,c{d,e{f,g}},d".
// This is the same string format as in FieldFilter.String(). This function should only be used in tests as it does not
// validate the given string and is only convenient to easily create DefaultMasks.
func MaskFromString(s string) Mask {
	mask, _ := maskFromRunes([]rune(s))
	return mask
}

func maskFromRunes(runes []rune) (Mask, int) {
	mask := make(Mask)
	var fieldName []string
	runes = append(runes, []rune(",")...)
	pos := 0
	for pos < len(runes) {
		char := fmt.Sprintf("%c", runes[pos])
		switch char {
		case " ", "\n", "\t":
			// Ignore white spaces.

		case ",", "{", "}":
			if len(fieldName) == 0 {
				switch char {
				case "}":
					return mask, pos
				case ",":
					pos += 1
					continue
				default:
					panic("invalid mask string format")
				}
			}

			var subMask FieldFilter
			if char == "{" {
				var jump int
				// Parse nested tree.
				subMask, jump = maskFromRunes(runes[pos+1:])
				pos += jump + 1
			} else {
				subMask = make(Mask)
			}
			f := strings.Join(fieldName, "")
			mask[f] = subMask
			// Reset FieldName.
			fieldName = []string{}

			if char == "}" {
				return mask, pos
			}

		default:
			fieldName = append(fieldName, char)
		}
		pos += 1
	}
	return mask, pos
}
