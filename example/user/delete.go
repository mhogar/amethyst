package user

import (
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/web"
)

func DeleteUserWorkflow() nodes.Workflow {
	c := newUserConverter()

	deleteUserAuth := nodes.NewWorkflow(
		converter.NewConverterNode(c.NewUserAuthFromParams),
		crud.NewDeleteModelNode[UserAuth]("user with username not found"),
	)

	deleteUser := nodes.NewWorkflow(
		converter.NewConverterNode(c.NewUserFromParams),
		crud.NewDeleteModelNode[User]("user with username not found"),
	)

	return nodes.NewWorkflow(
		nodes.NewSplitWorkflowNode(
			deleteUserAuth,
			deleteUser,
		),
		web.NewSuccessResponseNode(),
	)
}
