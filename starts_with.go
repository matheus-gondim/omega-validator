package validator

import (
	"fmt"
	"strings"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *Validation) StartsWith(prefix string) *Validation {
	v.addValidator(utils.StartsWith)

	str, ok := v.fieldValue.(string)

	if !ok || v.fieldValue == nil {
		v.addErrors(fmt.Errorf("error validating if fields starts with; field is not a string or is nil"))
		return v
	}

	required := utils.ContainsTypes(v.validatorsAdded, utils.Required)

	if !strings.HasPrefix(str, prefix) && required {
		v.addValidation(fmt.Sprintf("field does not start with %q", prefix))
	}

	return v
}
