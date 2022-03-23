package sqladapter

import (
	"database/sql"

	"github.com/mhogar/kiwi/common"
)

type RowsIterator struct {
	Rows *sql.Rows
}

func (itr *RowsIterator) Next() (bool, error) {
	if !itr.Rows.Next() {
		err := itr.Rows.Err()
		if err != nil {
			return false, common.ChainError("error preparing next row", err)
		}
	}

	return true, nil
}

func (itr *RowsIterator) Read(model interface{}) error {
	//TODO: scan into model
	err := itr.Rows.Scan()

	if err != nil {
		return common.ChainError("error reading row", err)
	}
	return nil
}

func (itr *RowsIterator) Close() {
	itr.Rows.Close()
}
