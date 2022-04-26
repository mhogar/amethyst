package user

import (
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/query"
	"github.com/mhogar/kiwi/nodes/web"
)

func GetUsersWorkflow() nodes.Workflow {
	c := newUserConverter()

	return nodes.NewWorkflow(
		crud.NewReadModelsNode[User](),
		converter.NewConverterNode(c.UsersToResponse),
		web.NewDataResponseNode(),
	)
}

func GetUserWorkflow(notFoundMessage string) nodes.Workflow {
	b := NewUserQueryBuilder()

	return nodes.NewWorkflow(
		query.NewBuildQueryNode(b.GetUserByUsername),
		crud.NewReadModelNode[User](notFoundMessage),
	)
}

func GetUserEndpoint() nodes.Workflow {
	c := newUserConverter()

	return nodes.NewWorkflow(
		converter.NewConverterNode(c.NewUserFromParams),
		GetUserWorkflow("user with username not found"),
		converter.NewConverterNode(c.UserToResponse),
		web.NewDataResponseNode(),
	)
}
