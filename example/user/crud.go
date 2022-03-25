package user

import (
	"github.com/mhogar/kiwi/nodes"
)

type CreateUserNode struct{}

func (CreateUserNode) Run(ctx nodes.BaseContext, input interface{}) (interface{}, *nodes.Error) {
	user := input.(*User)
	// handle := data.GetHandle[User](ctx.Adapter)

	// err := handle.Create(user)
	// if err != nil {
	// 	return nil, nodes.InternalError(common.ChainError("error creating user", err))
	// }

	return user, nil
}
