package sqladapter

import (
	"database/sql"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data"
	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/data/query"
)

type SqlAdapter struct {
	DB             *sql.DB
	ContextFactory data.ContextFactory
	ScriptBuilder  ScriptBuilder
}

func (a SqlAdapter) Select(model adapter.ReflectModel, where *query.WhereClause) (adapter.DataIterator, error) {
	ctx, cancel := a.ContextFactory.CreateStandardTimeoutContext()
	rows, err := a.DB.QueryContext(ctx, a.ScriptBuilder.BuildSelectQuery(model, where))
	defer cancel()

	if err != nil {
		return nil, common.ChainError("error executing query", err)
	}

	return &RowsIterator{
		Rows: rows,
	}, nil
}

func (a SqlAdapter) Insert(model adapter.ReflectModel) error {
	ctx, cancel := a.ContextFactory.CreateStandardTimeoutContext()
	_, err := a.DB.ExecContext(ctx, a.ScriptBuilder.BuildInsertStatement(model), model.Values...)
	defer cancel()

	if err != nil {
		return common.ChainError("error executing insert statement", err)
	}
	return nil
}

func (a SqlAdapter) Update(model adapter.ReflectModel) (bool, error) {
	ctx, cancel := a.ContextFactory.CreateStandardTimeoutContext()
	res, err := a.DB.ExecContext(ctx, a.ScriptBuilder.BuildUpdateStatement(model), model.Values...)
	defer cancel()

	if err != nil {
		return false, common.ChainError("error executing update statement", err)
	}

	count, _ := res.RowsAffected()
	return count > 0, nil
}

func (a SqlAdapter) Delete(model adapter.ReflectModel) (bool, error) {
	ctx, cancel := a.ContextFactory.CreateStandardTimeoutContext()
	res, err := a.DB.ExecContext(ctx, a.ScriptBuilder.BuildDeleteStatement(model), model.UniqueValue())
	defer cancel()

	if err != nil {
		return false, common.ChainError("error executing delete statement", err)
	}

	count, _ := res.RowsAffected()
	return count > 0, nil
}
