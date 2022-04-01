package main

import (
	"fmt"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/example/user"
	"github.com/mhogar/kiwi/nodes"
)

func Run(adapter adapter.DataAdapter, workflow nodes.Workflow[nodes.BaseContext], input any) *nodes.Error {
	ctx := nodes.BaseContext{
		Adapter: adapter,
	}

	output, err := workflow.Run(ctx, input)
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}

func CreateUserWorkflow() nodes.Workflow[nodes.BaseContext] {
	f := nodes.NodeFactory[nodes.BaseContext, user.User]{}
	return f.Workflow(
		f.Validation(user.CreateUserInputValidator()),
		f.Converter(user.CreateUserConverter()),
		f.Validation(user.CreateUserValidator()),
		f.CreateModel(),
	)
}

func UpdateUserWorkflow() nodes.Workflow[nodes.BaseContext] {
	f := nodes.NodeFactory[nodes.BaseContext, user.User]{}
	return f.Workflow(
		f.Validation(user.CreateUserInputValidator()),
		f.Converter(user.CreateUserConverter()),
		f.UpdateModel("user with username not found"),
	)
}

func DeleteUserWorkflow() nodes.Workflow[nodes.BaseContext] {
	f := nodes.NodeFactory[nodes.BaseContext, user.User]{}
	return f.Workflow(
		f.DeleteModel("user with username not found"),
	)
}

func main() {
	adapter := dependencies.DataAdapter.Resolve()

	//w := CreateUserWorkflow()
	//w := UpdateUserWorkflow()
	w := DeleteUserWorkflow()

	//user := user.CreateNewUserInput("user2", "Password123!", 3)
	user := user.CreateNewUser("user2", nil, 0)

	err := adapter.Setup()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer adapter.CleanUp()

	nErr := Run(adapter, w, user)
	if nErr != nil {
		fmt.Println(nErr.Errors)
	}
}
