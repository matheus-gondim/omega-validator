package validator

import (
	"fmt"

	"github.com/matheus-gondim/omega-validator/utils"
)

type ValidationError struct {
	Message string
	Errors  map[string][]string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%v: %v", e.Message, e.Errors)
}

type validator struct {
	value            any
	name             string
	validationErrors []string
	validators       []utils.ValidatorTypes
	internalErrors   []error
}

func Compose(builders ...*validator) (isSuccessful bool, err *ValidationError) {
	m := make(map[string][]string)

	for _, builder := range builders {
		valid, values := builder.Validate()
		if valid || values == nil {
			continue
		}
		for key, value := range values.Errors {
			m[key] = value
		}
	}

	for _, v := range m {
		if len(v) > 0 {
			return false, &ValidationError{
				Message: "fields is invalids",
				Errors:  m,
			}
		}
	}

	return true, nil
}

func New(fieldName string, fiedlValue any) *validator {
	return &validator{
		value:            fiedlValue,
		name:             fieldName,
		validationErrors: []string{},
		validators:       []utils.ValidatorTypes{},
	}
}

func (v *validator) addValidationError(errMsg string) {
	v.validationErrors = append(v.validationErrors, errMsg)
}

func (v *validator) addValidator(validatorTypes utils.ValidatorTypes) {
	v.validators = append(v.validators, validatorTypes)
}

func (v *validator) addInternalError(err error) {
	v.internalErrors = append(v.internalErrors, err)
}

func (v *validator) Validate() (isSuccessful bool, err *ValidationError) {
	errorsMap := make(map[string][]string)
	for _, err := range v.internalErrors {
		errorsMap["internal"] = append(errorsMap["internal"], fmt.Sprintf("%q: %q", v.name, err.Error()))
	}

	if len(v.validationErrors) > 0 {
		errorsMap[v.name] = v.validationErrors
	}

	for _, val := range errorsMap {
		if len(val) > 0 {
			return false, &ValidationError{
				Message: fmt.Sprintf("field %q is invalid", v.name),
				Errors:  errorsMap,
			}
		}
	}

	return true, nil
}
