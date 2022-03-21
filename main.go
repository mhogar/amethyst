package main

import (
	"fmt"

	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/example"
	"github.com/mhogar/kiwi/nodes"
)

func main() {
	f := nodes.NodeFactory{}

	w := f.Workflow(
		f.Validation(example.UserValidator{}),
		f.Converter(example.UserConverter{}),
	)

	fmt.Println(w.Run(
		example.CreateNewUserInput("username", "Password123!"),
	))

	user := data.Select[example.User](
		data.Where("username", "=", "username"),
	)

	println(user.Username)
}
