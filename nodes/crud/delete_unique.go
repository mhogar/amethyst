package crud

import (
	"errors"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/nodes"
)

type DeleteUniqueModelNode[Model any] struct {
	NotFoundMessage string
}

func NewDeleteUniqueModelNode[Model any](notFoundMessage string) DeleteUniqueModelNode[Model] {
	return DeleteUniqueModelNode[Model]{
		NotFoundMessage: notFoundMessage,
	}
}

func (n DeleteUniqueModelNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	exists, err := handle.DeleteUnique(input.(Identifier).GetIdentifier())
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error deleting model", err))
	}

	if !exists {
		return nil, nodes.ClientError(errors.New(n.NotFoundMessage))
	}

	return input, nil
}
