package auth

import (
	"errors"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
)

type AuthModel interface {
	GetPasswordHash() []byte
}

type AuthFields interface {
	GetUniqueValue() any
	GetPassword() string
}

type AuthenticateNode[Model AuthModel] struct {
	PasswordHasher converter.PasswordHasher
}

func NewAuthenticateNode[Model AuthModel]() AuthenticateNode[Model] {
	return AuthenticateNode[Model]{
		PasswordHasher: dependencies.PasswordHasher.Resolve(),
	}
}

func (n AuthenticateNode[Model]) Run(ctx interface{}, input any) (any, *nodes.Error) {
	fields := input.(AuthFields)

	handle := data.GetHandle[Model](ctx.(nodes.Context).GetDataAdapter())
	model, err := handle.ReadUnique(fields.GetUniqueValue())

	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error reading model", err))
	}

	if model == nil {
		return nil, nodes.ClientError(errors.New("invalid username and/or password"))
	}

	err = n.PasswordHasher.ComparePasswords((*model).GetPasswordHash(), fields.GetPassword())
	if err != nil {
		return nil, nodes.ClientError(errors.New("invalid username and/or password"))
	}

	return input, nil
}
