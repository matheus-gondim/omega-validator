package validator

import "github.com/matheus-gondim/omega-validator/utils"

func (v *validator) WithCustomRule(action func(value any) error) *validator {
	isRequired := utils.ContainsTypes(v.validators, utils.Required)

	if !isRequired && utils.IsEmpty(v.value) {
		return v
	}

	err := action(v.value)
	if err != nil {
		v.addValidationError(err.Error())
	}

	return v
}
