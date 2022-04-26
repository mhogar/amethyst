package crud

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/data/query"
	"github.com/mhogar/kiwi/nodes"
)

type DeleteModelsNode[Model any] struct{}

func NewDeleteModelsNode[Model any]() DeleteModelsNode[Model] {
	return DeleteModelsNode[Model]{}
}

func (n DeleteModelsNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	where := input.(*query.WhereClause)
	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	_, err := handle.Delete(where)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error deleting model", err))
	}

	return input, nil
}
