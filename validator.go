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

type Validation struct {
	fieldValue      any
	fieldName       string
	results         []string
	validatorsAdded []utils.ValidatorTypes
	errs            []error
}

func Compose(builders ...*Validation) (isSuccessful bool, err *ValidationError) {
	m := make(map[string][]string)

	for _, builder := range builders {
		valid, values := builder.Builder()
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

func New(fieldName string, fiedlValue any) *Validation {
	return &Validation{
		fieldValue:      fiedlValue,
		fieldName:       fieldName,
		results:         []string{},
		validatorsAdded: []utils.ValidatorTypes{},
	}
}

func (v *Validation) addValidation(errMsg string) {
	v.results = append(v.results, errMsg)
}

func (v *Validation) addValidator(validatorTypes utils.ValidatorTypes) {
	v.validatorsAdded = append(v.validatorsAdded, validatorTypes)
}

func (v *Validation) addErrors(err error) {
	v.errs = append(v.errs, err)
}

func (v *Validation) Builder() (isSuccessful bool, err *ValidationError) {
	m := make(map[string][]string)
	for _, err := range v.errs {
		m["internal"] = append(m["internal"], fmt.Sprintf("%q: %q", v.fieldName, err.Error()))
	}

	if len(v.results) > 0 {
		m[v.fieldName] = v.results
	}

	for _, val := range m {
		if len(val) > 0 {
			return false, &ValidationError{
				Message: fmt.Sprintf("field %q is invalid", v.fieldName),
				Errors:  m,
			}
		}
	}

	return true, nil
}
