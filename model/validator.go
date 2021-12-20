package model

const (
	ValidateFieldValid       = 0x0
	ValidateFieldInvalidType = 0x1
)

type Validator interface {
	Validate(val interface{}) (int, string)
}
