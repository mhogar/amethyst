package nodes

import "github.com/mhogar/kiwi/data/adapter"

type Context interface {
	GetDataAdapter() adapter.DataAdapter
}

type ContextImpl struct {
	Adapter adapter.DataAdapter
}

func NewContext(adapter adapter.DataAdapter) ContextImpl {
	return ContextImpl{
		Adapter: adapter,
	}
}

func (c ContextImpl) GetDataAdapter() adapter.DataAdapter {
	return c.Adapter
}
