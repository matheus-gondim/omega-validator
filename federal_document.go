package validator

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/matheus-gondim/omega-validator/utils"
)

func (v *validator) FederalDocument() *validator {
	v.addValidator(utils.FederalDocument)

	document, ok := v.value.(string)
	if !ok {
		v.addInternalError(fmt.Errorf("error validating field with federal document; field is not a string"))
		return v
	}

	required := utils.ContainsTypes(v.validators, utils.Required)
	if document == "" && !required {
		return v
	}

	cleanedDoc := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(document, ".", ""), "-", ""), "/", "")

	if len(cleanedDoc) == 0 || strings.Count(cleanedDoc, string(cleanedDoc[0])) == len(cleanedDoc) {
		v.addValidationError("federal document is invalid")
		return v
	}

	isValid := func() bool {
		switch len(document) {
		case 14:
			return isCnpjValid(document)
		case 11:
			return isCpfValid(document)
		default:
			return false
		}
	}

	if isValid() {
		return v
	}

	v.addValidationError("federal document is invalid")
	return v
}

func isCnpjValid(document string) bool {
	digit1, err := calculateValidatorDigit(document, []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	if err != nil {
		return false
	}

	digit2, err := calculateValidatorDigit(document, []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	if err != nil {
		return false
	}

	return strings.HasSuffix(document, fmt.Sprintf("%d%d", digit1, digit2))
}

func isCpfValid(document string) bool {
	digit1, err := calculateValidatorDigit(document, []int{10, 9, 8, 7, 6, 5, 4, 3, 2})
	if err != nil {
		return false
	}
	digit2, err := calculateValidatorDigit(document, []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2})
	if err != nil {
		return false
	}

	return strings.HasSuffix(document, fmt.Sprintf("%d%d", digit1, digit2))
}

func calculateValidatorDigit(doc string, validationIndexes []int) (int, error) {
	sum := 0
	for ix := 0; ix < len(validationIndexes); ix++ {
		digit, err := strconv.Atoi(string(doc[ix]))
		if err != nil {
			return 0, fmt.Errorf("error converting character at position %d to integer: %v", ix, err)
		}
		sum += digit * validationIndexes[ix]
	}

	rest := sum % 11
	if rest < 2 {
		return 0, nil
	}
	return 11 - rest, nil
}
