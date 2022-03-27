package adapter

import (
	"github.com/mhogar/kiwi/data/query"
)

type DataIterator interface {
	Next() (bool, error)
	Read(model ReflectModel) error
	Close()
}

type DataAdapter interface {
	// Setup sets up the adapter and returns any errors.
	Setup() error

	// CleanUp cleans up the adapter and returns any errors.
	CleanUp() error

	Select(model ReflectModel, where *query.WhereClause) (DataIterator, error)
	Insert(model ReflectModel) error
	Update(model ReflectModel) (bool, error)
	Delete(model ReflectModel) (bool, error)
}
