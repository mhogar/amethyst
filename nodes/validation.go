package nodes

import (
	"fmt"

	"github.com/mhogar/kiwi/validator"
)

type ValidationNode struct {
	Validator validator.Validator
}

func (n ValidationNode) Run(input interface{}) interface{} {
	verrs := n.Validator.Validate(input)

	if verrs.HasErrors() {
		fmt.Println(verrs.Messages)
	}

	return input
}
