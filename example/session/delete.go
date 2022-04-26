package session

import (
	"github.com/mhogar/kiwi/example/models"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/web"
)

func DeleteSessionEndpoint() nodes.Workflow {
	c := newSessionConverter()

	return nodes.NewWorkflow(
		web.NewParseTokenFromAuthorizationHeaderNode(),
		converter.NewConverterNode(c.NewSessionFromToken),
		crud.NewDeleteUniqueModelNode[models.Session]("bearer token invalid or expired"),
		web.NewSuccessResponseNode(),
	)
}
