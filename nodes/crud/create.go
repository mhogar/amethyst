package crud

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/nodes"
)

type CreateModelNode[Model any] struct{}

func NewCreateModelNode[Model any]() CreateModelNode[Model] {
	return CreateModelNode[Model]{}
}

func (CreateModelNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	model := input.(*Model)
	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	err := handle.Create(model)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error creating model", err))
	}

	return model, nil
}
