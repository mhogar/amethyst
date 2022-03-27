package nodes

import (
	"github.com/mhogar/kiwi/nodes/validator"
)

type ValidationNode[T any] struct {
	Validator validator.Validator[T]
}

func (f NodeFactory[T]) Validation(v validator.Validator[T]) ValidationNode[T] {
	return ValidationNode[T]{
		Validator: v,
	}
}

func (n ValidationNode[T]) Run(ctx T, input interface{}) (interface{}, *Error) {
	verrs, err := n.Validator.Validate(ctx, input)

	if err != nil {
		return nil, InternalError(err)
	}

	if verrs.HasErrors() {
		return nil, ClientError(verrs.Errors()...)
	}

	return input, nil
}
