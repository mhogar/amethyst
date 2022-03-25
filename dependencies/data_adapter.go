package dependencies

import (
	"github.com/mhogar/kiwi/data/adapter"
	sqladapter "github.com/mhogar/kiwi/data/adapter/sql_adapter"
)

func createDataAdapter() adapter.DataAdapter {
	return sqladapter.SqlAdapter{}
}

var DataAdapter = Dependency[adapter.DataAdapter]{
	createObject: createDataAdapter,
}
