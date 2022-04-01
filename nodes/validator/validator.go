package validator

import "github.com/mhogar/kiwi/nodes"

type Validator interface {
	Validate(ctx interface{}, val any) (*ValidationErrors, error)
}

type ValidatorNode struct {
	Validator Validator
}

func NewValidatorNode(v Validator) ValidatorNode {
	return ValidatorNode{
		Validator: v,
	}
}

func (n ValidatorNode) Run(ctx interface{}, input any) (any, *nodes.Error) {
	verrs, err := n.Validator.Validate(ctx, input)

	if err != nil {
		return nil, nodes.InternalError(err)
	}

	if verrs.HasErrors() {
		return nil, nodes.ClientError(verrs.Errors()...)
	}

	return input, nil
}
