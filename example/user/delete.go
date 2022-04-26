package user

import (
	"github.com/mhogar/kiwi/example/models"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/query"
	"github.com/mhogar/kiwi/nodes/web"
)

func DeleteUserWorkflow() nodes.Workflow {
	b := NewUserQueryBuilder()
	c := newUserConverter()

	deleteSessions := nodes.NewWorkflow(
		query.NewBuildQueryNode(b.FindUserSessions),
		crud.NewDeleteModelsNode[models.Session](),
	)

	deleteUserAuth := nodes.NewWorkflow(
		converter.NewConverterNode(c.EmptyUserFromSession),
		crud.NewDeleteUniqueModelNode[models.UserAuth]("user with username not found"),
	)

	deleteUser := nodes.NewWorkflow(
		converter.NewConverterNode(c.EmptyUserFromSession),
		crud.NewDeleteUniqueModelNode[models.User]("user with username not found"),
	)

	return nodes.NewWorkflow(
		nodes.NewSplitWorkflowNode(
			deleteSessions,
			deleteUserAuth,
			deleteUser,
		),
	)
}

func DeleteUserEndpoint() nodes.Workflow {
	return nodes.NewWorkflow(
		web.SetSessionContextFromAuthorizationHeaderWorkflow[models.Session](),
		DeleteUserWorkflow(),
		web.NewSuccessResponseNode(),
	)
}
