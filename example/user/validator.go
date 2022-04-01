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

func NewUserInputValidator() UserInputValidator {
	return UserInputValidator{
		BaseValidator: dependencies.CreateBaseValidator[UserInput](),
	}
}

func (v UserInputValidator) Validate(_ interface{}, val any) (*validator.ValidationErrors, error) {
	user := val.(*UserInput)

	verrs := v.ValidateLength(user, "Username", 5, 30)
	verrs.Merge(v.ValidatePassword(user, "Password", 8, 0, true, true))

	return verrs, nil
}

type UserValidator struct {
	validator.BaseValidator[User]
}

func NewUserValidator() UserValidator {
	return UserValidator{
		BaseValidator: dependencies.CreateBaseValidator[User](),
	}
}

func (v UserValidator) Validate(ctx interface{}, val any) (*validator.ValidationErrors, error) {
	user := val.(*User)

	verrs, err := v.ValidateUniqueField(user, ctx.(nodes.Context).GetDataAdapter(), "already in use by another user")
	if err != nil {
		return verrs, common.ChainError("error validating user unique field", err)
	}

	return verrs, nil
}
