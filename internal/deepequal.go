package deepequal

import (
	"fmt"
	"reflect"
)

type Diff struct {
	Key         string
	OldValue    string
	NewValue    string
	Description string
}

type Comparison struct {
	Removed  []Diff
	Added    []Diff
	Modified []Diff
}

func Compare(old interface{}, new interface{}) (Comparison, error) {

	if old == nil && new == nil {
		return Comparison{}, nil
	}
	oldValue := reflect.ValueOf(old)
	newValue := reflect.ValueOf(new)

	if !oldValue.IsValid() {
		return Comparison{}, fmt.Errorf("%q is invalid for comparison", oldValue)
	}

	if !newValue.IsValid() {
		return Comparison{}, fmt.Errorf("%q is invalid for comparison", newValue)
	}

	if expectedValue.Type() != actualValue.Type() {
		return false, diff.PrimitiveTypeMismatch{
			ExpectedType: expectedValue.Type(),
			ActualValue:  actualValue.Interface(),
		}
	}
	//
	// switch actualValue.Kind() {
	// case reflect.Slice:
	// 	return Slice(expectedValue, actualValue)
	//
	// case reflect.Map:
	// 	return Map(expectedValue, actualValue)
	//
	// default:
	// 	return Primitive(expected, actual)
	// }

	return Comparison{
		Removed: []Diff{
			{
				Key:      "test",
				OldValue: "test",
				NewValue: "test",
			},
		},
	}
}
