package crud

import (
	"errors"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/nodes"
)

type Identifier interface {
	GetIdentifier() any
}

type ReadUniqueModelNode[Model any] struct {
	NotFoundMessage string
}

func NewReadUniqueModelNode[Model any](notFoundMessage string) ReadUniqueModelNode[Model] {
	return ReadUniqueModelNode[Model]{
		NotFoundMessage: notFoundMessage,
	}
}

func (n ReadUniqueModelNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())

	model, err := handle.ReadUnique(input.(Identifier).GetIdentifier())
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error reading model", err))
	}

	if model != nil {
		return model, nil
	}

	if n.NotFoundMessage != "" {
		return nil, nodes.ClientError(errors.New(n.NotFoundMessage))
	}

	return nil, nil
}
