package session

import (
	"github.com/mhogar/kiwi/example/models"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/web"
)

func DeleteSessionWorkflow(notFoundMessage string) nodes.Workflow {
	return nodes.NewWorkflow(
		crud.NewDeleteUniqueModelNode[models.Session](notFoundMessage),
	)
}

func DeleteSessionEndpoint() nodes.Workflow {
	return nodes.NewWorkflow(
		web.NewParseTokenFromAuthorizationHeaderNode(),
		DeleteSessionWorkflow("bearer token invalid or expired"),
		web.NewSuccessResponseNode(),
	)
}
