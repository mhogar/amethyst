package validator

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/data/query"
)

func ValidateLength(field string, value any, minLen int, maxLen int) *ValidationErrors {
	len := reflect.ValueOf(value).Len()
	verrs := &ValidationErrors{}

	if minLen > 0 && len < minLen {
		verrs.Add(field, fmt.Sprintf("shorter than min length %d", minLen))
	} else if maxLen > 0 && len > maxLen {
		verrs.Add(field, fmt.Sprintf("longer than max length %d", maxLen))
	}

	return verrs
}

func ValidateMin(field string, value any, minValue int64) *ValidationErrors {
	num := reflect.ValueOf(value).Int()
	verrs := &ValidationErrors{}

	if num < minValue {
		verrs.Add(field, fmt.Sprintf("smaller than min value %d", minValue))
	}

	return verrs
}

func ValidateMax(field string, value any, maxValue int64) *ValidationErrors {
	num := reflect.ValueOf(value).Int()
	verrs := &ValidationErrors{}

	if num > maxValue {
		verrs.Add(field, fmt.Sprintf("larger than max value %d", maxValue))
	}

	return verrs
}

func ValidatePassword(field string, value string, minLen int, maxLen int, requireDigit bool, requireSymbol bool) *ValidationErrors {
	verrs := ValidateLength(field, value, minLen, maxLen)

	if requireDigit && regexp.MustCompile(`[0-9]`).FindString(value) == "" {
		verrs.Add(field, "must contain digit")
	}

	if requireSymbol && regexp.MustCompile(`[^0-9a-zA-Z]`).FindString(value) == "" {
		verrs.Add(field, "must contain symbol")
	}

	return verrs
}

func ValidateUniqueField[T any](model *T, db adapter.DataAdapter, message string) (*ValidationErrors, error) {
	verrs := &ValidationErrors{}

	m := adapter.CreateReflectModel[T]()
	m.SetModel(model)

	handle := data.GetHandle[T](db)
	other, err := handle.Read(
		query.Where(m.UniqueField(), "=", m.UniqueValue()),
	)

	if err != nil {
		return verrs, common.ChainError("error getting model by unique field", err)
	}

	if len(other) > 0 {
		verrs.Add(m.UniqueField(), message)
	}

	return verrs, nil
}
