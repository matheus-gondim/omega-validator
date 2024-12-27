package validator

import (
	"fmt"
	"reflect"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *Validation) Contains(element any) *Validation {
	v.addValidator(utils.Contains)

	val := reflect.ValueOf(v.fieldValue)
	t := val.Type()

	if t.Kind() != reflect.Array && t.Kind() != reflect.Slice && t.Kind() != reflect.Map {
		v.addErrors(fmt.Errorf("error validating if field contains; unsupported type: %v", t.Kind()))
		return v
	}

	required := utils.ContainsTypes(v.validatorsAdded, utils.Required)

	if !required && val.Len() == 0 {
		return v
	}

	contains := false
	switch t.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(val.Index(i).Interface(), element) {
				contains = true
				break
			}
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			if reflect.DeepEqual(val.MapIndex(key).Interface(), element) {
				contains = true
				break
			}
		}
	default:
		v.addErrors(fmt.Errorf("error validating if field contains; unsupported type: %v", t.Kind()))
		return v
	}

	if contains {
		return v
	}

	v.addValidation(fmt.Sprintf("field does not contain the specified element: %v", element))

	return v
}
