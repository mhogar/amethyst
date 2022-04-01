package nodes

import (
	"errors"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
)

type CreateModelNode[Context, Model any] struct{}

func (f NodeFactory[Context, Model]) CreateModel() CreateModelNode[Context, Model] {
	return CreateModelNode[Context, Model]{}
}

func (CreateModelNode[Context, Model]) Run(ctx BaseContext, input interface{}) (interface{}, *Error) {
	model := input.(*Model)
	handle := data.GetHandle[Model](ctx.Adapter)

	err := handle.Create(model)
	if err != nil {
		return nil, InternalError(common.ChainError("error creating model", err))
	}

	return model, nil
}

type UpdateModelNode[Context, Model any] struct {
	NotFoundMessage string
}

func (f NodeFactory[Context, Model]) UpdateModel(notFoundMessage string) UpdateModelNode[Context, Model] {
	return UpdateModelNode[Context, Model]{
		NotFoundMessage: notFoundMessage,
	}
}

func (n UpdateModelNode[Context, Model]) Run(ctx BaseContext, input interface{}) (interface{}, *Error) {
	model := input.(*Model)
	handle := data.GetHandle[Model](ctx.Adapter)

	exists, err := handle.Update(model)
	if err != nil {
		return nil, InternalError(common.ChainError("error updating model", err))
	}

	if !exists {
		return nil, ClientError(errors.New(n.NotFoundMessage))
	}

	return model, nil
}

type DeleteModelNode[Context, Model any] struct {
	NotFoundMessage string
}

func (f NodeFactory[Context, Model]) DeleteModel(notFoundMessage string) DeleteModelNode[Context, Model] {
	return DeleteModelNode[Context, Model]{
		NotFoundMessage: notFoundMessage,
	}
}

func (n DeleteModelNode[Context, Model]) Run(ctx BaseContext, input interface{}) (interface{}, *Error) {
	model := input.(*Model)
	handle := data.GetHandle[Model](ctx.Adapter)

	exists, err := handle.Delete(model)
	if err != nil {
		return nil, InternalError(common.ChainError("error deleting model", err))
	}

	if !exists {
		return nil, ClientError(errors.New(n.NotFoundMessage))
	}

	return model, nil
}
