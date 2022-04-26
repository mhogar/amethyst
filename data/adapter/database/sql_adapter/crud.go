package sqladapter

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/data/query"
)

func (a *SQLAdapter) Select(model adapter.ReflectModel, where *query.WhereClause) (adapter.DataIterator, error) {
	script, values := a.SQLDriver.BuildSelectQuery(model, where)

	ctx, cancel := a.ContextFactory.CreateStandardTimeoutContext()
	rows, err := a.DB.QueryContext(ctx, script, values...)

	if err != nil {
		return nil, common.ChainError("error executing query", err)
	}

	return &RowsIterator{
		cancel: cancel,
		Rows:   rows,
	}, nil
}

func (a *SQLAdapter) Insert(model adapter.ReflectModel) error {
	script := a.SQLDriver.BuildInsertStatement(model)

	ctx, cancel := a.ContextFactory.CreateStandardTimeoutContext()
	_, err := a.DB.ExecContext(ctx, script, model.Values...)
	defer cancel()

	if err != nil {
		return common.ChainError("error executing insert statement", err)
	}
	return nil
}

func (a *SQLAdapter) Update(model adapter.ReflectModel) (bool, error) {
	script := a.SQLDriver.BuildUpdateStatement(model)

	ctx, cancel := a.ContextFactory.CreateStandardTimeoutContext()
	res, err := a.DB.ExecContext(ctx, script, model.Values...)
	defer cancel()

	if err != nil {
		return false, common.ChainError("error executing update statement", err)
	}

	count, _ := res.RowsAffected()
	return count > 0, nil
}

func (a *SQLAdapter) Delete(model adapter.ReflectModel, where *query.WhereClause) (bool, error) {
	script, values := a.SQLDriver.BuildDeleteStatement(model, where)

	ctx, cancel := a.ContextFactory.CreateStandardTimeoutContext()
	res, err := a.DB.ExecContext(ctx, script, values...)
	defer cancel()

	if err != nil {
		return false, common.ChainError("error executing delete statement", err)
	}

	count, _ := res.RowsAffected()
	return count > 0, nil
}
