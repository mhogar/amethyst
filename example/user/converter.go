package user

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/web"
)

type userConverter struct {
	converter.BaseConverter
}

func newUserConverter() userConverter {
	return userConverter{
		BaseConverter: dependencies.BaseConverter.Resolve(),
	}
}

func (c userConverter) UserFieldsToUser(_ interface{}, val any) (any, error) {
	user := val.(UserFields)
	return NewUser(user.GetUsername(), user.GetRank()), nil
}

func (c userConverter) UserAuthFieldsToUserAuth(_ interface{}, val any) (any, error) {
	user := val.(UserAuthFields)

	hash, err := c.HashPassword(user.GetPassword())
	if err != nil {
		return nil, common.ChainError("error hashing password", err)
	}

	return NewUserAuth(user.GetUsername(), hash, user.GetRank()), nil
}

func (userConverter) SetUsernameFromParams(ctx interface{}, val any) (any, error) {
	user := val.(UsernameField)
	user.SetUsername(ctx.(web.HTTPRouterContext).GetParams().ByName("username"))
	return user, nil
}

func (userConverter) NewUserFromParams(ctx interface{}, _ any) (any, error) {
	username := ctx.(web.HTTPRouterContext).GetParams().ByName("username")
	return NewUser(username, 0), nil
}

func (userConverter) NewUserAuthFromParams(ctx interface{}, _ any) (any, error) {
	username := ctx.(web.HTTPRouterContext).GetParams().ByName("username")
	return NewUserAuth(username, nil, 0), nil
}

func (userConverter) UserFieldsToResponse(_ interface{}, val any) (any, error) {
	user := val.(UserFields)
	return UserResponse{
		Username: user.GetUsername(),
		Rank:     user.GetRank(),
	}, nil
}

func (userConverter) UserToResponse(_ interface{}, val any) (any, error) {
	user := val.(*User)
	return newUserResponse(user), nil
}

func (userConverter) UsersToResponse(_ interface{}, val any) (any, error) {
	users := val.([]*User)
	res := make([]UserResponse, len(users))

	for i, user := range users {
		res[i] = newUserResponse(user)
	}

	return res, nil
}
