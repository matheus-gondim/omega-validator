package validator

import (
	"fmt"
	"regexp"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *validator) Regexp(expression string) *validator {
	v.addValidator(utils.Regexp)

	str, ok := v.value.(string)
	if !ok {
		v.addInternalError(fmt.Errorf("error validating field with regular expression; field is not a string"))
		return v
	}

	required := utils.ContainsTypes(v.validators, utils.Required)

	if !required && str == "" {
		return v
	}

	re := regexp.MustCompile(expression)
	if re.MatchString(str) {
		return v
	}

	v.addValidationError("field does not match")

	return v
}
