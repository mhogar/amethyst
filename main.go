package main

import (
	"fmt"

	"github.com/amethyst/example"
	"github.com/amethyst/nodes"
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
