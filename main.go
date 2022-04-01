package main

import (
	"fmt"

	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/example/user"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/validator"
	"github.com/mhogar/kiwi/nodes/web"
)

func CreateUserWorkflow() nodes.Workflow {
	return nodes.NewWorkflow(
		web.NewJSONBodyParserNode[user.UserInput](),
		validator.NewValidatorNode(user.NewUserInputValidator()),
		converter.NewConverterNode(user.NewUserConverter()),
		validator.NewValidatorNode(user.NewUserValidator()),
		crud.NewCreateModelNode[user.User](),
	)
}

func UpdateUserWorkflow() nodes.Workflow {
	return nodes.NewWorkflow(
		web.NewJSONBodyParserNode[user.UserInput](),
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

	err := adapter.Setup()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer adapter.CleanUp()

	//TODO: add web server
}
