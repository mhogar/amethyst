package main

import (
	"fmt"

	sqladapter "github.com/mhogar/kiwi/data/adapter/sql_adapter"
	"github.com/mhogar/kiwi/example"
	"github.com/mhogar/kiwi/nodes"
)

func main() {
	f := nodes.NodeFactory[nodes.BaseContext]{}

	w := f.Workflow(
		f.Validation(example.UserValidator{}),
		f.Converter(example.UserConverter{}),
		example.CreateUserNode{},
	)

	ctx := nodes.BaseContext{
		Adapter: &sqladapter.SqlAdapter{},
	}

	fmt.Println(w.Run(ctx,
		example.CreateNewUserInput("username", "Password123!"),
	))
}
