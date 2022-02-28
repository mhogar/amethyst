package validator

type ValidationErrors struct {
	Messages map[string][]string
}

func CreateNewValidationErrors() *ValidationErrors {
	return &ValidationErrors{
		Messages: map[string][]string{},
	}
}

func (v *ValidationErrors) Add(field string, messages ...string) {
	_, ok := v.Messages[field]
	if !ok {
		v.Messages[field] = messages
	} else {
		v.Messages[field] = append(v.Messages[field], messages...)
	}
}

func (v *ValidationErrors) Merge(other *ValidationErrors) {
	for field, messages := range other.Messages {
		v.Add(field, messages...)
	}
}

func (v *ValidationErrors) HasErrors() bool {
	return len(v.Messages) > 0
}
