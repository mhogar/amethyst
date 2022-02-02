package validator

import (
	"fmt"
	"reflect"
)

func (v *BaseValidator) ValidateLength(field string, val interface{}, minLen int, maxLen int) {
	len := reflect.ValueOf(val).Len()

	if len < minLen {
		v.addMessage(field, fmt.Sprintf("shorter than min length %d", minLen))
	} else if len > maxLen {
		v.addMessage(field, fmt.Sprintf("longer than max length %d", maxLen))
	}
}
