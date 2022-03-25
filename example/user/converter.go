package user

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
)

type UserConverter struct {
	converter.BaseConverter
}

func CreateUserConverter() UserConverter {
	return UserConverter{
		BaseConverter: dependencies.BaseConverter.Resolve(),
	}
}

func (c UserConverter) Convert(_ nodes.BaseContext, val interface{}) (interface{}, error) {
	user := val.(*UserInput)

	hash, err := c.HashPassword(user.Password)
	if err != nil {
		return nil, common.ChainError("error hashing password", err)
	}

	return CreateNewUser(user.Username, hash, user.Rank), nil
}
