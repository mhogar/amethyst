package user

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/nodes/converter"
)

type UserConverter struct {
	converter.BaseConverter
}

func NewUserConverter() UserConverter {
	return UserConverter{
		BaseConverter: dependencies.BaseConverter.Resolve(),
	}
}

func (c UserConverter) Convert(_ interface{}, val any) (any, error) {
	user := val.(*UserInput)

	hash, err := c.HashPassword(user.Password)
	if err != nil {
		return nil, common.ChainError("error hashing password", err)
	}

	return CreateNewUser(user.Username, hash, user.Rank), nil
}
