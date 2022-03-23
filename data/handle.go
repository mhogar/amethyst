package data

import (
	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/data/query"
)

type Handle[T any] struct {
	Adapter adapter.DataAdapter
}

func GetHandle[T any](da adapter.DataAdapter) Handle[T] {
	return Handle[T]{
		Adapter: da,
	}
}

func (h Handle[T]) Create(model *T) error {
	return h.Adapter.Insert(model)
}

func (h Handle[T]) Read(where *query.WhereClause) ([]*T, error) {
	itr, err := h.Adapter.Select(where)
	if err != nil {
		return nil, err
	}
	defer itr.Close()

	//read the data
	models := []*T{}
	for {
		hasNext, err := itr.Next()
		if err != nil {
			return nil, err
		}

		if !hasNext {
			break
		}

		model, err := h.readModel(itr)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (h Handle[T]) Update(model *T) (bool, error) {
	return h.Adapter.Update(model)
}

func (h Handle[T]) Delete(model *T) (bool, error) {
	return h.Adapter.Delete(model)
}

func (h Handle[T]) readModel(itr adapter.DataIterator) (*T, error) {
	var model T

	err := itr.Read(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
