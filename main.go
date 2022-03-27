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

	user, err := workflow.Run(ctx, user.CreateNewUserInput("user2", "Password123!", 3))
	if err != nil {
		return err
	}

	fmt.Println(user)
	return nil
}

func main() {
	f := nodes.NodeFactory[nodes.BaseContext]{}
	w := f.Workflow(
		f.Validation(user.CreateUserInputValidator()),
		f.Converter(user.CreateUserConverter()),
		f.Validation(user.CreateUserValidator()),
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
