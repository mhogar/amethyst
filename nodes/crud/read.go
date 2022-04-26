package crud

import (
	"errors"

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

type ReadModelNode[Model any] struct {
	NotFoundMessage string
}

func NewReadModelNode[Model any](notFoundMessage string) ReadModelNode[Model] {
	return ReadModelNode[Model]{
		NotFoundMessage: notFoundMessage,
	}
}

func (n ReadModelNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	where := input.(*query.WhereClause)
	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	models, err := handle.Read(where)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error reading model", err))
	}

	if len(models) > 0 {
		return models[0], nil
	}

	if n.NotFoundMessage != "" {
		return nil, nodes.ClientError(errors.New(n.NotFoundMessage))
	}

	return nil, nil
}
