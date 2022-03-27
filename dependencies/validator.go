package dependencies

import "github.com/mhogar/kiwi/nodes/validator"

func CreateBaseValidator[T any]() validator.BaseValidator[T] {
	return &validator.BaseValidatorImpl[T]{}
}
