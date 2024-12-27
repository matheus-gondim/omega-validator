package validator

import (
	"fmt"
	"reflect"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *Validation) Min(value any) *Validation {
	v.addValidator(utils.Min)

	required := utils.ContainsTypes(v.validatorsAdded, utils.Required)

	val := reflect.ValueOf(v.fieldValue)
	valueToCompare := reflect.ValueOf(value)

	isMin, err := min(val, valueToCompare, required)

	if err != nil {
		v.addErrors(err)
		return v
	}

	if !isMin {
		v.addValidation(fmt.Sprintf("field cannot be less than %v", value))
	}
	return v
}

func min(value reflect.Value, valueToCompare reflect.Value, required bool) (bool, error) {
	t := value.Type()

	isZero, err := utils.IsZeroValue(value)
	if err != nil {
		return false, err
	}

	if isZero && !required {
		return true, nil
	}

	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() >= valueToCompare.Int(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return value.Uint() >= valueToCompare.Uint(), nil
	case reflect.Float32, reflect.Float64:
		return value.Float() >= valueToCompare.Float(), nil
	case reflect.String:
		return int64(len(value.String())) >= valueToCompare.Int(), nil
	case reflect.Array, reflect.Slice, reflect.Map:
		return int64(value.Len()) >= valueToCompare.Int(), nil
	default:
		return false, fmt.Errorf("error validating minimum value; type of %v unknown or unsupported: %v", value, valueToCompare.Kind())
	}
}
