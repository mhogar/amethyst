package dependencies

import "github.com/mhogar/kiwi/nodes/validator"

func createBaseValidator() validator.BaseValidator {
	return validator.BaseValidatorImpl{}
}

var BaseValidator = Dependency[validator.BaseValidator]{
	createObject: createBaseValidator,
}
