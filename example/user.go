package example

import "github.com/amethyst/validator"

type User struct {
	Username string
	Password string
}

func CreateNewUser(username string, password string) *User {
	return &User{
		Username: username,
		Password: password,
	}
}

type UserValidator struct {
	validator.BaseValidator
}

func (v UserValidator) Validate(val interface{}) *validator.ValidationErrors {
	user := val.(*User)

	verrs := v.ValidateLength("username", user.Username, 5, 30)
	verrs.Merge(v.ValidatePassword("password", user.Password, 8, 0, true, true))

	return verrs
}
