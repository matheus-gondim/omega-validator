package validator

import (
	"fmt"
	"regexp"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *Validation) Email() *Validation {
	v.addValidator(utils.Email)
	str, ok := v.fieldValue.(string)

	if !ok || v.fieldValue == nil {
		v.addErrors(fmt.Errorf("error validating if fields is email; field is not a string or is nil"))
		return v
	}

	required := utils.ContainsTypes(v.validatorsAdded, utils.Required)
	if str == "" && !required {
		return v
	}

	emailRegex := `^(?i)([a-z0-9!#\$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-z0-9!#\$%&'*+/=?^_` + "`" + `{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x7f\x21\x23-\x5b\x5d-\x7e]|\\[\x01-\x09\x0b\x0c\x0d-\x7f])+")@([a-z0-9](?:[a-z0-9-]*[a-z0-9])?(?:\.[a-z0-9](?:[a-z0-9-]*[a-z0-9])?)*\.[a-z]{2,})$`
	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(str) {
		v.addValidation("email has format invalid")
	}

	return v
}
