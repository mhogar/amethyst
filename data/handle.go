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
	m := adapter.CreateReflectModel[T]()
	m.SetModel(model)

	return h.Adapter.Insert(m)
}

func (h Handle[T]) Read(where *query.WhereClause) ([]*T, error) {
	m := adapter.CreateReflectModel[T]()

	itr, err := h.Adapter.Select(m, where)
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

		model, err := h.readModel(itr, m)
		if err != nil {
			return nil, err
		}
		models = append(models, model)
	}

	return models, nil
}

func (h Handle[T]) ReadUnique(identifier any) (*T, error) {
	m := adapter.CreateReflectModel[T]()

	models, err := h.Read(query.Where(m.UniqueField(), "=", identifier))
	if err != nil {
		return nil, err
	}

	if len(models) > 0 {
		return models[0], nil
	}

	return nil, nil
}

func (h Handle[T]) Update(model *T) (bool, error) {
	m := adapter.CreateReflectModel[T]()
	m.SetModel(model)

	return h.Adapter.Update(m)
}

func (h Handle[T]) Delete(where *query.WhereClause) (bool, error) {
	m := adapter.CreateReflectModel[T]()
	return h.Adapter.Delete(m, where)
}

func (h Handle[T]) DeleteUnique(identifier any) (bool, error) {
	m := adapter.CreateReflectModel[T]()
	return h.Delete(query.Where(m.UniqueField(), "=", identifier))
}

func (h Handle[T]) readModel(itr adapter.DataIterator, m adapter.ReflectModel) (*T, error) {
	var model T
	m.SetModel(&model)

	err := itr.Read(m)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
