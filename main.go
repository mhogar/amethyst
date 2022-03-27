package main

import (
	"fmt"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/example/user"
	"github.com/mhogar/kiwi/nodes"
)

func Run(adapter adapter.DataAdapter, workflow nodes.Workflow[nodes.BaseContext]) *nodes.Error {
	ctx := nodes.BaseContext{
		Adapter: adapter,
	}

	user, err := workflow.Run(ctx, user.CreateNewUserInput("username", "Password123!", 3))
	if err != nil {
		return err
	}

	fmt.Println(user)
	return nil
}

func main() {
	f := nodes.NodeFactory[nodes.BaseContext]{}
	w := f.Workflow(
		f.Validation(user.CreateUserValidator()),
		f.Converter(user.CreateUserConverter()),
		user.CreateUserNode{},
	)

	adapter := dependencies.DataAdapter.Resolve()

	err := adapter.Setup()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer adapter.CleanUp()

	nErr := Run(adapter, w)
	if nErr != nil {
		fmt.Println(nErr.Errors)
	}
}
