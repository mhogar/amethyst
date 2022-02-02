package validator

import "fmt"

type BaseValidator struct {
	Messages []string
}

func (v BaseValidator) HasErrors() bool {
	return len(v.Messages) > 0
}

func (v BaseValidator) GetMessages() []string {
	return v.Messages
}

func (v *BaseValidator) ClearErrors() {
	v.Messages = []string{}
}

func (v *BaseValidator) addMessage(field string, message string) {
	v.Messages = append(v.Messages, fmt.Sprintf("%s: %s", field, message))
}
