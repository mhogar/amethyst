package validator

import (
	"fmt"
	"strings"
)

type ValidationErrors struct {
	Messages []string
}

func CreateNewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		Messages: []string{},
	}
}

func (v *ValidationErrors) Add(field string, message string) {
	v.Messages = append(v.Messages, fmt.Sprintf("%s: %s", field, message))
}

func (v *ValidationErrors) Merge(other ValidationErrors) {
	v.Messages = append(v.Messages, other.Messages...)
}

func (v *ValidationErrors) HasErrors() bool {
	return len(v.Messages) > 0
}

func (v *ValidationErrors) FormatMessages() string {
	return strings.Join(v.Messages, ", ")
}
