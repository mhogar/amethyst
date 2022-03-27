package sqladapter

import (
	"context"
	"database/sql"

	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/data/adapter"
)

type RowsIterator struct {
	cancel context.CancelFunc
	Rows   *sql.Rows
}

func (itr *RowsIterator) Next() (bool, error) {
	if !itr.Rows.Next() {
		err := itr.Rows.Err()
		if err != nil {
			return false, common.ChainError("error preparing next row", err)
		}

		return false, nil
	}

	return true, nil
}

func (itr *RowsIterator) Read(model adapter.ReflectModel) error {
	err := itr.Rows.Scan(model.Addresses...)

	if err != nil {
		return common.ChainError("error reading row", err)
	}
	return nil
}

func (itr *RowsIterator) Close() {
	itr.cancel()
	itr.Rows.Close()
}
