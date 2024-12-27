package validator

import (
	"reflect"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *Validation) Required() *Validation {
	v.addValidator(utils.Required)

	val := reflect.ValueOf(v.fieldValue)

	isZero, err := utils.IsZeroValue(val)
	if err != nil {
		v.addErrors(err)
		return v
	}
	if isZero {
		v.addValidation("field is required.")
	}

	return v
}
