package dependencies

import "github.com/mhogar/kiwi/nodes/converter"

func createBaseConverter() converter.BaseConverter {
	return converter.BaseConverterImpl{
		PasswordHasher: PasswordHasher.Resolve(),
	}
}

var BaseConverter = Dependency[converter.BaseConverter]{
	createObject: createBaseConverter,
}
