package dependencies

import (
	sqladapter "github.com/mhogar/kiwi/data/adapter/database/sql_adapter"
	"github.com/mhogar/kiwi/data/adapter/database/sql_adapter/postgres"
)

func createSQLDriver() sqladapter.SQLDriver {
	return postgres.Driver{}
}

var SQLDriver = Dependency[sqladapter.SQLDriver]{
	createObject: createSQLDriver,
}
