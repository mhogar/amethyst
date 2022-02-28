package validator

import (
	"fmt"
	"reflect"
)

func (BaseValidator) ValidateLength(field string, val interface{}, minLen int, maxLen int) *ValidationErrors {
	len := reflect.ValueOf(val).Len()
	verrs := CreateNewValidationErrors()

	if len < minLen {
		verrs.Add(field, fmt.Sprintf("shorter than min length %d", minLen))
	} else if len > maxLen {
		verrs.Add(field, fmt.Sprintf("longer than max length %d", maxLen))
	}

	return verrs
}
