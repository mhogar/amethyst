package sqladapter

import (
	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/data/query"
)

type ScriptBuilder interface {
	BuildSelectQuery(model adapter.ReflectModel, where *query.WhereClause) (string, []any)
	BuildInsertStatement(model adapter.ReflectModel) string
	BuildUpdateStatement(model adapter.ReflectModel) string
	BuildDeleteStatement(model adapter.ReflectModel) string
}
