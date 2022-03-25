package user

import (
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/validator"
)

type UserValidator struct {
	validator.BaseValidator
}

func CreateUserValidator() UserValidator {
	return UserValidator{
		BaseValidator: dependencies.BaseValidator.Resolve(),
	}
}

func (v UserValidator) Validate(_ nodes.BaseContext, val interface{}) *validator.ValidationErrors {
	user := val.(*UserInput)

	verrs := v.ValidateLength("username", user.Username, 5, 30)
	verrs.Merge(v.ValidatePassword("password", user.Password, 8, 0, true, true))

	return verrs
}
