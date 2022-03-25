package main

import (
	"fmt"

	sqladapter "github.com/mhogar/kiwi/data/adapter/sql_adapter"
	"github.com/mhogar/kiwi/example/user"
	"github.com/mhogar/kiwi/nodes"
)

func main() {
	f := nodes.NodeFactory[nodes.BaseContext]{}

	w := f.Workflow(
		f.Validation(user.CreateUserValidator()),
		f.Converter(user.CreateUserConverter()),
		user.CreateUserNode{},
	)

	ctx := nodes.BaseContext{
		Adapter: &sqladapter.SqlAdapter{},
	}

	user, err := w.Run(
		ctx, user.CreateNewUserInput("username", "Password123!", 3),
	)

	if err != nil {
		fmt.Println(err.Errors)
	} else {
		fmt.Println(user)
	}
}
