package session

import (
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/query"
)

func GetSessionWorkflow[Model Session](notFoundMessage string) nodes.Workflow {
	b := newSessionQueryBuilder()

	return nodes.NewWorkflow(
		query.NewBuildQueryNode(b.GetSessionByToken),
		crud.NewReadModelNode[Model](notFoundMessage),
	)
}
