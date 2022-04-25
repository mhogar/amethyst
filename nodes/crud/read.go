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

type ReadModelNode[Model any] struct{}

func NewReadModelNode[Model any]() ReadModelNode[Model] {
	return ReadModelNode[Model]{}
}

func (ReadModelNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	where := input.(*query.WhereClause)
	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	models, err := handle.Read(where)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error reading model", err))
	}

	if len(models) < 0 {
		return nil, nil
	}

	return models[0], nil
}
