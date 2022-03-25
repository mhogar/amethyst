package nodes

import (
	"fmt"

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

func (n ValidationNode[T]) Run(ctx T, input interface{}) interface{} {
	verrs := n.Validator.Validate(ctx, input)

	if verrs.HasErrors() {
		fmt.Println(verrs.Messages)
	}

	return input
}
