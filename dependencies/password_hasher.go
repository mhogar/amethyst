package dependencies

import "github.com/mhogar/kiwi/nodes/converter"

func createPasswordHasher() converter.PasswordHasher {
	return converter.BCryptPasswordHasher{}
}

var PasswordHasher = Dependency[converter.PasswordHasher]{
	createObject: createPasswordHasher,
}
