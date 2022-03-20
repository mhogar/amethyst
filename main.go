package main

import (
	"fmt"

	"github.com/mhogar/kiwi/example"
	"github.com/mhogar/kiwi/nodes"
)

func main() {
	p := nodes.Pipeline{}
	p.Build(
		nodes.ValidationNode{
			Validator: example.UserValidator{},
		},
		nodes.ConverterNode{
			Converter: example.UserConverter{},
		},
	)

	fmt.Println(p.Run(
		example.CreateNewUserInput("username", "Password123!"),
	))
}
