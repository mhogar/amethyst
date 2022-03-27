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

type BaseValidator[T any] interface {
	ValidateLength(model *T, field string, minLen int, maxLen int) *ValidationErrors
	ValidatePassword(model *T, field string, minLen int, maxLen int, requireDigit bool, requireSymbol bool) *ValidationErrors
	ValidateUniqueField(model *T, db adapter.DataAdapter, message string) (*ValidationErrors, error)
}

type BaseValidatorImpl[T any] struct{}

func (BaseValidatorImpl[T]) ValidateLength(model *T, field string, minLen int, maxLen int) *ValidationErrors {
	len := reflect.ValueOf(model).Elem().FieldByName(field).Len()
	verrs := &ValidationErrors{}

	if minLen > 0 && len < minLen {
		verrs.Add(field, fmt.Sprintf("shorter than min length %d", minLen))
	} else if maxLen > 0 && len > maxLen {
		verrs.Add(field, fmt.Sprintf("longer than max length %d", maxLen))
	}

	return verrs
}

func (v BaseValidatorImpl[T]) ValidatePassword(model *T, field string, minLen int, maxLen int, requireDigit bool, requireSymbol bool) *ValidationErrors {
	str := reflect.ValueOf(model).Elem().FieldByName(field).String()
	verrs := v.ValidateLength(model, field, minLen, maxLen)

	if requireDigit && regexp.MustCompile(`[0-9]`).FindString(str) == "" {
		verrs.Add(field, "must contain digit")
	}

	if requireSymbol && regexp.MustCompile(`[^0-9a-zA-Z]`).FindString(str) == "" {
		verrs.Add(field, "must contain symbol")
	}

	return verrs
}

func (v BaseValidatorImpl[T]) ValidateUniqueField(model *T, db adapter.DataAdapter, message string) (*ValidationErrors, error) {
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
