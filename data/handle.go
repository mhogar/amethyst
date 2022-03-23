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
	rm := adapter.CreateReflectModel[T]()
	rm.SetModel(&model)

	return h.Adapter.Insert(rm)
}

func (h Handle[T]) Read(where *query.WhereClause) ([]*T, error) {
	rm := adapter.CreateReflectModel[T]()

	itr, err := h.Adapter.Select(rm, where)
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

		model, err := h.readModel(itr, rm)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (h Handle[T]) Update(model *T) (bool, error) {
	rm := adapter.CreateReflectModel[T]()
	rm.SetModel(&model)

	return h.Adapter.Update(rm)
}

func (h Handle[T]) Delete(model *T) (bool, error) {
	rm := adapter.CreateReflectModel[T]()
	rm.SetModel(&model)

	return h.Adapter.Delete(rm)
}

func (h Handle[T]) readModel(itr adapter.DataIterator, rm adapter.ReflectModel) (*T, error) {
	var model T
	rm.SetModel(&model)

	err := itr.Read(rm)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
