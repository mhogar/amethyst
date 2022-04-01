package crud

import (
	"errors"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/nodes"
)

type UpdateModelNode[Model any] struct {
	NotFoundMessage string
}

func NewUpdateModelNode[Model any](notFoundMessage string) UpdateModelNode[Model] {
	return UpdateModelNode[Model]{
		NotFoundMessage: notFoundMessage,
	}
}

func (n UpdateModelNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	model := input.(*Model)
	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	exists, err := handle.Update(model)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error updating model", err))
	}

	if !exists {
		return nil, nodes.ClientError(errors.New(n.NotFoundMessage))
	}

	return model, nil
}
