package example

import "github.com/amethyst/validator"

type User struct {
	Username string
}

func CreateNewUser(username string) *User {
	return &User{
		Username: username,
	}
}

type UserValidator struct {
	validator.BaseValidator
}

func (v UserValidator) Validate(val interface{}) *validator.ValidationErrors {
	user := val.(*User)

	return v.ValidateLength("username", user.Username, 5, 30)
}
