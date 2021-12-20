package model

import "reflect"

const (
	ValidateLengthTooShort = 0x10
	ValidateLengthTooLong  = 0x11
)

type LengthValidator struct {
	Min int
	Max int
}

func (v LengthValidator) Validate(val interface{}) (int, string) {
	var len int

	switch reflect.TypeOf(val).Kind() {
	case reflect.String, reflect.Array, reflect.Slice:
		len = reflect.ValueOf(val).Len()
	default:
		return ValidateFieldInvalidType, "invalid type"
	}

	if len < v.Min {
		return ValidateLengthTooShort, "length too short"
	}

	if len < v.Max {
		return ValidateLengthTooLong, "length too long"
	}

	return ValidateFieldValid, ""
}
