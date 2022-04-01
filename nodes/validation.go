package nodes

import (
	"github.com/mhogar/kiwi/nodes/validator"
)

type ValidationNode[Context any] struct {
	Validator validator.Validator[Context]
}

func (f NodeFactory[Context, Model]) Validation(v validator.Validator[Context]) ValidationNode[Context] {
	return ValidationNode[Context]{
		Validator: v,
	}
}

func (n ValidationNode[Context]) Run(ctx Context, input interface{}) (interface{}, *Error) {
	verrs, err := n.Validator.Validate(ctx, input)

	if err != nil {
		return nil, InternalError(err)
	}

	if verrs.HasErrors() {
		return nil, ClientError(verrs.Errors()...)
	}

	return input, nil
}
