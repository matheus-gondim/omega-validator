package validator

import (
	"fmt"
	"strings"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *Validation) EndsWith(suffix string) *Validation {
	v.addValidator(utils.EndsWith)

	str, ok := v.fieldValue.(string)

	if !ok || v.fieldValue == nil {
		v.addErrors(fmt.Errorf("error validating if fields ends with; field is not a string or is nil"))
		return v
	}

	required := utils.ContainsTypes(v.validatorsAdded, utils.Required)

	if !strings.HasSuffix(str, suffix) && required {
		v.addValidation(fmt.Sprintf("field does not end with %q", suffix))
	}

	return v
}
