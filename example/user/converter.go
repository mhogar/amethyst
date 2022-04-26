package user

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/example/models"
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
	return models.NewUser(user.GetUsername(), user.GetRank()), nil
}

func (c userConverter) UserAuthFieldsToUserAuth(_ interface{}, val any) (any, error) {
	user := val.(UserAuthFields)

	hash, err := c.HashPassword(user.GetNewPassword())
	if err != nil {
		return nil, common.ChainError("error hashing password", err)
	}

	return models.NewUserAuth(user.GetUsername(), hash), nil
}

func (userConverter) SetUsernameFromParams(ctx interface{}, val any) (any, error) {
	user := val.(UsernameField)
	user.SetUsername(ctx.(web.HTTPRouterContext).GetParams().ByName("username"))
	return user, nil
}

func (userConverter) SetUsernameFromSession(ctx interface{}, val any) (any, error) {
	user := val.(UsernameField)
	session := ctx.(web.HTTPRouterContext).GetSession().(*models.Session)

	user.SetUsername(session.Username)
	return user, nil
}

func (userConverter) EmptyUserFromParams(ctx interface{}, _ any) (any, error) {
	username := ctx.(web.HTTPRouterContext).GetParams().ByName("username")
	return models.NewUser(username, 0), nil
}

func (userConverter) EmptyUserFromSession(ctx interface{}, _ any) (any, error) {
	session := ctx.(web.HTTPRouterContext).GetSession().(*models.Session)
	return models.NewUser(session.Username, 0), nil
}

func (userConverter) UserFieldsToResponse(_ interface{}, val any) (any, error) {
	user := val.(UserFields)
	return UserResponse{
		Username: user.GetUsername(),
		Rank:     user.GetRank(),
	}, nil
}

func (userConverter) UserToResponse(_ interface{}, val any) (any, error) {
	user := val.(*models.User)
	return newUserResponse(user), nil
}

func (userConverter) UsersToResponse(_ interface{}, val any) (any, error) {
	users := val.([]*models.User)
	res := make([]UserResponse, len(users))

	for i, user := range users {
		res[i] = newUserResponse(user)
	}

	return res, nil
}
