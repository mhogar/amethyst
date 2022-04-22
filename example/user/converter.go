package user

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/web"
)

type UserConverter struct {
	converter.BaseConverter
}

func NewUserConverter() UserConverter {
	return UserConverter{
		BaseConverter: dependencies.BaseConverter.Resolve(),
	}
}

func (c UserConverter) ConvertInputToUser(_ interface{}, val any) (any, error) {
	user := val.(*UserInput)

	hash, err := c.HashPassword(user.Password)
	if err != nil {
		return nil, common.ChainError("error hashing password", err)
	}

	return NewUser(user.Username, hash, user.Rank), nil
}

func (UserConverter) SetUsernameFromParams(ctx interface{}, val any) (any, error) {
	user := val.(*UserInput)
	user.Username = ctx.(web.HTTPRouterContext).GetParams().ByName("username")
	return user, nil
}

func (UserConverter) NewUserFromParams(ctx interface{}, _ any) (any, error) {
	username := ctx.(web.HTTPRouterContext).GetParams().ByName("username")
	return NewUser(username, nil, 0), nil
}

func (UserConverter) ConvertUserToResponse(_ interface{}, val any) (any, error) {
	user := val.(*User)
	return newUserResponse(user), nil
}

func (UserConverter) ConvertUsersToResponse(_ interface{}, val any) (any, error) {
	users := val.([]*User)
	res := make([]UserResponse, len(users))

	for i, user := range users {
		res[i] = newUserResponse(user)
	}

	return res, nil
}
