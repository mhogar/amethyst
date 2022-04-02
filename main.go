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

	"github.com/julienschmidt/httprouter"
)

func CreateUserWorkflow() nodes.Workflow {
	c := user.NewUserConverter()

	return nodes.NewWorkflow(
		web.NewJSONBodyParserNode[user.UserInput](),
		validator.NewValidatorNode(user.NewUserInputValidator()),
		converter.NewConverterNode(c.ConvertInputToUser),
		validator.NewValidatorNode(user.NewUserValidator()),
		crud.NewCreateModelNode[user.User](),
	)
}

func UpdateUserWorkflow() nodes.Workflow {
	c := user.NewUserConverter()

	return nodes.NewWorkflow(
		web.NewJSONBodyParserNode[user.UserInput](),
		converter.NewConverterNode(c.SetUsernameFromParams),
		validator.NewValidatorNode(user.NewUserInputValidator()),
		converter.NewConverterNode(c.ConvertInputToUser),
		crud.NewUpdateModelNode[user.User]("user with username not found"),
	)
}

func DeleteUserWorkflow() nodes.Workflow {
	c := user.NewUserConverter()

	return nodes.NewWorkflow(
		converter.NewConverterNode(c.NewUserFromParams),
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

	r := httprouter.New()

	r.POST("/user",
		web.NewHandler(adapter, CreateUserWorkflow()).ServeHTTPRouter,
	)
	r.PUT("/user/:username",
		web.NewHandler(adapter, CreateUserWorkflow()).ServeHTTPRouter,
	)
	r.DELETE("/user/:username",
		web.NewHandler(adapter, CreateUserWorkflow()).ServeHTTPRouter,
	)
}
