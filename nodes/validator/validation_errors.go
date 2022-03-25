package validator

import (
	"fmt"
	"sort"
)

type ValidationError struct {
	Field   string
	Message string
}

func (verr ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", verr.Field, verr.Message)
}

type ValidationErrors struct {
	errors []ValidationError
}

func (verrs *ValidationErrors) Add(field string, messages ...string) {
	for _, message := range messages {
		verrs.errors = append(verrs.errors, ValidationError{
			Field:   field,
			Message: message,
		})
	}
}

func (verrs *ValidationErrors) Merge(other *ValidationErrors) {
	verrs.errors = append(verrs.errors, other.errors...)
}

func (verrs *ValidationErrors) HasErrors() bool {
	return len(verrs.errors) > 0
}

func (verrs *ValidationErrors) Errors() []error {
	sort.Slice(verrs.errors, func(i, j int) bool {
		return verrs.errors[i].Field < verrs.errors[j].Field
	})

	errs := make([]error, len(verrs.errors))
	for index, verr := range verrs.errors {
		errs[index] = verr
	}

	return errs
}
