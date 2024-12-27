package utils

import (
	"fmt"
	"reflect"
)

func IsZeroValue(value reflect.Value) (bool, error) {
	t := value.Type()

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0, nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() == 0, nil
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0.0, nil
	case reflect.String:
		return value.String() == "", nil
	case reflect.Array, reflect.Slice, reflect.Map:
		return value.Len() == 0, nil
	case reflect.Struct:
		return value.IsNil(), nil
	case reflect.Ptr:
		return value.IsNil(), nil
	default:
		return false, fmt.Errorf("error validating required fields; unknown or unsupported type: %v", t.Kind())
	}
}
