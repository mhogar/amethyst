package user

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/dependencies"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/validator"
)

type UserInputValidator struct {
	validator.BaseValidator[UserInput]
}

func CreateUserInputValidator() UserInputValidator {
	return UserInputValidator{
		BaseValidator: dependencies.CreateBaseValidator[UserInput](),
	}
}

func (v UserInputValidator) Validate(_ nodes.BaseContext, val interface{}) (*validator.ValidationErrors, error) {
	user := val.(*UserInput)

	verrs := v.ValidateLength(user, "Username", 5, 30)
	verrs.Merge(v.ValidatePassword(user, "Password", 8, 0, true, true))

	return verrs, nil
}

type UserValidator struct {
	validator.BaseValidator[User]
}

func CreateUserValidator() UserValidator {
	return UserValidator{
		BaseValidator: dependencies.CreateBaseValidator[User](),
	}
}

func (v UserValidator) Validate(ctx nodes.BaseContext, val interface{}) (*validator.ValidationErrors, error) {
	user := val.(*User)

	verrs, err := v.ValidateUniqueField(user, ctx.Adapter, "already in use by another user")
	if err != nil {
		return verrs, common.ChainError("error validating user unique field", err)
	}

	return verrs, nil
}
