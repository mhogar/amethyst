package validator

import (
	"regexp"
)

func (v BaseValidator) ValidatePassword(field string, val interface{}, minLen int, maxLen int, requireDigit bool, requireSymbol bool) *ValidationErrors {
	str := val.(string)
	verrs := v.ValidateLength(field, str, minLen, maxLen)

	if requireDigit && regexp.MustCompile(`[0-9]`).FindString(str) == "" {
		verrs.Add(field, "must contain digit")
	}

	if requireSymbol && regexp.MustCompile(`[^0-9a-zA-Z]`).FindString(str) == "" {
		verrs.Add(field, "must contain symbol")
	}

	return verrs
}
