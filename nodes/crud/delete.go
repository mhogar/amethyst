package crud

import (
	"errors"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/nodes"
)

type DeleteModelNode[Model any] struct {
	NotFoundMessage string
}

func NewDeleteModelNode[Model any](notFoundMessage string) DeleteModelNode[Model] {
	return DeleteModelNode[Model]{
		NotFoundMessage: notFoundMessage,
	}
}

func (n DeleteModelNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	model := input.(*Model)
	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	exists, err := handle.Delete(model)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error deleting model", err))
	}

	if !exists {
		return nil, nodes.ClientError(errors.New(n.NotFoundMessage))
	}

	return model, nil
}
