package adapter

import "github.com/mhogar/kiwi/data/query"

type DataIterator interface {
	Next() (bool, error)
	Read(model interface{}) error
	Close()
}

type DataAdapter interface {
	Select(where *query.WhereClause) (DataIterator, error)
	Insert(model interface{}) error
	Update(model interface{}) (bool, error)
	Delete(model interface{}) (bool, error)
}
