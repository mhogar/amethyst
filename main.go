package main

import (
	"fmt"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/example/user"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/validator"
)

func Run(adapter adapter.DataAdapter, workflow nodes.Workflow, input any) *nodes.Error {
	ctx := nodes.ContextImpl{
		Adapter: adapter,
	}

	output, err := workflow.Run(ctx, input)
	if err != nil {
		return err
	}

	fmt.Println(output)
	return nil
}

func CreateUserWorkflow() nodes.Workflow {
	return nodes.NewWorkflow(
		validator.NewValidatorNode(user.NewUserInputValidator()),
		converter.NewConverterNode(user.NewUserConverter()),
		validator.NewValidatorNode(user.NewUserValidator()),
		crud.NewCreateModelNode[user.User](),
	)
}

func UpdateUserWorkflow() nodes.Workflow {
	return nodes.NewWorkflow(
		validator.NewValidatorNode(user.NewUserInputValidator()),
		converter.NewConverterNode(user.NewUserConverter()),
		crud.NewUpdateModelNode[user.User]("user with username not found"),
	)
}

func DeleteUserWorkflow() nodes.Workflow {
	return nodes.NewWorkflow(
		crud.NewDeleteModelNode[user.User]("user with username not found"),
	)
}

func main() {
	adapter := dependencies.DataAdapter.Resolve()

	w := CreateUserWorkflow()
	//w := UpdateUserWorkflow()
	//w := DeleteUserWorkflow()

	user := user.CreateNewUserInput("user2", "Password123!", 1)
	//user := user.CreateNewUser("user2", nil, 0)

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
