package validator

import (
	"reflect"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *validator) Required() *validator {
	v.addValidator(utils.Required)

	val := reflect.ValueOf(v.value)

	isZero, err := utils.IsZeroValue(val)
	if err != nil {
		v.addInternalError(err)
		return v
	}
	if isZero {
		v.addValidationError("field is required.")
	}

	return v
}
