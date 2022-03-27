package validator

type Validator[T any] interface {
	Validate(ctx T, val interface{}) (*ValidationErrors, error)
}
