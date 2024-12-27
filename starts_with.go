package validator

import (
	"fmt"
	"strings"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *validator) StartsWith(prefix string) *validator {
	v.addValidator(utils.StartsWith)

	str, ok := v.value.(string)

	if !ok || v.value == nil {
		v.addInternalError(fmt.Errorf("error validating if fields starts with; field is not a string or is nil"))
		return v
	}

	required := utils.ContainsTypes(v.validators, utils.Required)

	if !strings.HasPrefix(str, prefix) && required {
		v.addValidationError(fmt.Sprintf("field does not start with %q", prefix))
	}

	return v
}
