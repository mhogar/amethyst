package validator

type Validator interface {
	HasErrors() bool
	GetMessages() []string
	ClearErrors()
	Validate(val interface{})
}
