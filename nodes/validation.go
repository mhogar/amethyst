package nodes

import (
	"fmt"

	"github.com/amethyst/validator"
)

type ValidationNode struct {
	Validator validator.Validator
}

func (n ValidationNode) Run(input interface{}) interface{} {
	verrs := n.Validator.Validate(input)

	if verrs.HasErrors() {
		fmt.Println(verrs.FormatMessages())
	}

	return input
}
