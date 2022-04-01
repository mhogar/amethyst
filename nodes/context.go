package nodes

import "github.com/mhogar/kiwi/data/adapter"

type Context interface {
	DataAdapter() adapter.DataAdapter
}

type ContextImpl struct {
	Adapter adapter.DataAdapter
}

func (c ContextImpl) DataAdapter() adapter.DataAdapter {
	return c.Adapter
}
