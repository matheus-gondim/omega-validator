package validator

import (
	"fmt"
	"strings"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *validator) EndsWith(suffix string) *validator {
	v.addValidator(utils.EndsWith)

	str, ok := v.value.(string)

	if !ok || v.value == nil {
		v.addInternalError(fmt.Errorf("error validating if fields ends with; field is not a string or is nil"))
		return v
	}

	required := utils.ContainsTypes(v.validators, utils.Required)

	if !strings.HasSuffix(str, suffix) && required {
		v.addValidationError(fmt.Sprintf("field does not end with %q", suffix))
	}

	return v
}
