package validator

import (
	"fmt"
	"reflect"
	"regexp"
)

type BaseValidator interface {
	ValidateLength(field string, val interface{}, minLen int, maxLen int) *ValidationErrors
	ValidatePassword(field string, val interface{}, minLen int, maxLen int, requireDigit bool, requireSymbol bool) *ValidationErrors
}

type BaseValidatorImpl struct{}

func (BaseValidatorImpl) ValidateLength(field string, val interface{}, minLen int, maxLen int) *ValidationErrors {
	len := reflect.ValueOf(val).Len()
	verrs := &ValidationErrors{}

	if minLen > 0 && len < minLen {
		verrs.Add(field, fmt.Sprintf("shorter than min length %d", minLen))
	} else if maxLen > 0 && len > maxLen {
		verrs.Add(field, fmt.Sprintf("longer than max length %d", maxLen))
	}

	return verrs
}

func (v BaseValidatorImpl) ValidatePassword(field string, val interface{}, minLen int, maxLen int, requireDigit bool, requireSymbol bool) *ValidationErrors {
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
