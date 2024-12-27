package validator

import (
	"fmt"
	"regexp"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *Validation) Regexp(expression string) *Validation {
	v.addValidator(utils.Regexp)

	str, ok := v.fieldValue.(string)
	if !ok {
		v.addErrors(fmt.Errorf("error validating field with regular expression; field is not a string"))
		return v
	}

	required := utils.ContainsTypes(v.validatorsAdded, utils.Required)

	if !required && str == "" {
	}

	re := regexp.MustCompile(expression)
	if re.MatchString(str) {
		return v
	}

	v.addValidation("field does not match")

	return v
}
