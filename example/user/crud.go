package user

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/nodes"
)

type CreateUserNode struct{}

func (CreateUserNode) Run(ctx nodes.BaseContext, input interface{}) (interface{}, *nodes.Error) {
	user := input.(*User)
	handle := data.GetHandle[User](ctx.Adapter)

	err := handle.Create(user)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error creating user", err))
	}

	return user, nil
}

type UpdateUserNode struct{}

func (UpdateUserNode) Run(ctx nodes.BaseContext, input interface{}) (interface{}, *nodes.Error) {
	user := input.(*User)
	handle := data.GetHandle[User](ctx.Adapter)

	exists, err := handle.Update(user)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error updating user", err))
	}

	if !exists {
		return nil, nodes.ClientError(common.NewError(`user "%s" not found`, user.Username))
	}

	return user, nil
}

type DeleteUserNode struct{}

func (DeleteUserNode) Run(ctx nodes.BaseContext, input interface{}) (interface{}, *nodes.Error) {
	user := input.(*User)
	handle := data.GetHandle[User](ctx.Adapter)

	exists, err := handle.Delete(user)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error deleting user", err))
	}

	if !exists {
		return nil, nodes.ClientError(common.NewError(`user "%s" not found`, user.Username))
	}

	return user, nil
}
