package nodes

import (
	"fmt"
	"strings"

	"github.com/amethyst/validator"
)

type ValidationNode struct {
	Validator validator.Validator
}

func (n ValidationNode) Run(input interface{}) interface{} {
	n.Validator.Validate(input)

	if n.Validator.HasErrors() {
		fmt.Println(strings.Join(n.Validator.GetMessages(), ", "))
	}

	n.Validator.ClearErrors()
	return input
}
