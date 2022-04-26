package crud

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/data/query"
	"github.com/mhogar/kiwi/nodes"
)

type ReadModelsNode[Model any] struct{}

func NewReadModelsNode[Model any]() ReadModelsNode[Model] {
	return ReadModelsNode[Model]{}
}

func (ReadModelsNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	var where *query.WhereClause
	if input != nil {
		where = input.(*query.WhereClause)
	}

	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	models, err := handle.Read(where)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error reading models", err))
	}

	return models, nil
}
