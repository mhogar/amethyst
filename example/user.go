package example

import (
	"github.com/amethyst/converter"
	"github.com/amethyst/validator"
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
	Username     string
	PasswordHash []byte
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

func (v UserValidator) Validate(val interface{}) *validator.ValidationErrors {
	user := val.(*UserInput)

	verrs := v.ValidateLength("username", user.Username, 5, 30)
	verrs.Merge(v.ValidatePassword("password", user.Password, 8, 0, true, true))

	return verrs
}

type UserConverter struct {
	converter.BaseConverter
}

func (c UserConverter) Convert(val interface{}) interface{} {
	user := val.(*UserInput)
	return CreateNewUser(user.Username, c.HashPassword(user.Password))
}
