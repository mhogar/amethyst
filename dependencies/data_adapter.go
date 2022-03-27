package dependencies

import (
	"github.com/mhogar/kiwi/data/adapter"
	sqladapter "github.com/mhogar/kiwi/data/adapter/database/sql_adapter"
)

func createDataAdapter() adapter.DataAdapter {
	return sqladapter.CreateSQLAdapter(SQLDriver.Resolve())
}

var DataAdapter = Dependency[adapter.DataAdapter]{
	createObject: createDataAdapter,
}
