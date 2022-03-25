package example

import (
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/validator"
)

type UserInput struct {
	Username string
	Password string
}

func CreateNewUserInput(username string, password string) *UserInput {
	return &UserInput{
		Username: username,
		Password: password,
	}
}

type User struct {
	Username     string `kiwi:"username"`
	Rank         int    `kiwi:"rank"`
	PasswordHash []byte `kiwi:"password_hash"`
}

func CreateNewUser(username string, hash []byte) *User {
	return &User{
		Username:     username,
		PasswordHash: hash,
	}
}

type UserValidator struct {
	validator.BaseValidator
}

func (v UserValidator) Validate(_ nodes.BaseContext, val interface{}) *validator.ValidationErrors {
	user := val.(*UserInput)

	verrs := v.ValidateLength("username", user.Username, 5, 30)
	verrs.Merge(v.ValidatePassword("password", user.Password, 8, 0, true, true))

	return verrs
}

type UserConverter struct {
	converter.BaseConverter
}

func (c UserConverter) Convert(_ nodes.BaseContext, val interface{}) interface{} {
	user := val.(*UserInput)
	return CreateNewUser(user.Username, c.HashPassword(user.Password))
}

type CreateUserNode struct{}

func (CreateUserNode) Run(ctx nodes.BaseContext, input interface{}) (interface{}, *nodes.Error) {
	user := input.(*User)
	handle := data.GetHandle[User](ctx.Adapter)

	err := handle.Create(user)
	if err != nil {
		return nil, nodes.InternalError(err)
	}

	return user, nil
}
