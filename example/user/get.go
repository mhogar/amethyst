package user

import (
	"github.com/mhogar/kiwi/example/models"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/web"
)

func GetUsersWorkflow() nodes.Workflow {
	c := newUserConverter()

	return nodes.NewWorkflow(
		crud.NewReadModelsNode[models.User](),
		converter.NewConverterNode(c.UsersToResponse),
		web.NewDataResponseNode(),
	)
}

func GetUserEndpoint() nodes.Workflow {
	c := newUserConverter()

	return nodes.NewWorkflow(
		converter.NewConverterNode(c.EmptyUserFromParams),
		crud.NewReadUniqueModelNode[models.User]("user with username not found"),
		converter.NewConverterNode(c.UserToResponse),
		web.NewDataResponseNode(),
	)
}
