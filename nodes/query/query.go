package query

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data/query"
	"github.com/mhogar/kiwi/nodes"
)

type BuildQuery func(ctx interface{}, input any) (*query.WhereClause, error)

type BuildQueryNode struct {
	BuildQuery BuildQuery
}

func NewBuildQueryNode(q BuildQuery) BuildQueryNode {
	return BuildQueryNode{
		BuildQuery: q,
	}
}

func (n BuildQueryNode) Run(ctx interface{}, input any) (any, *nodes.Error) {
	where, err := n.BuildQuery(ctx, input)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error building query", err))
	}

	return where, nil
}
