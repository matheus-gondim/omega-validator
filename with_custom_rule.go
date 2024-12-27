package validator

import (
	"reflect"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *validator) WithCustomRule(action func(value any) error) *validator {
	isRequired := utils.ContainsTypes(v.validators, utils.Required)

	val := reflect.ValueOf(v.value)
	isZero, err := utils.IsZeroValue(val)
	if err != nil {
		v.addInternalError(err)
		return v
	}

	if !isRequired && isZero {
		return v
	}

	err = action(v.value)
	if err != nil {
		v.addValidationError(err.Error())
	}

	return v
}
