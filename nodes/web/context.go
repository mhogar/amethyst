package web

import (
	"net/http"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/nodes"
)

type WebContext interface {
	nodes.Context

	GetRequest() *http.Request
}

type WebContextImpl struct {
	nodes.ContextImpl

	Writer  http.ResponseWriter
	Request *http.Request
}

func NewWebContext(adapter adapter.DataAdapter, w http.ResponseWriter, req *http.Request) WebContext {
	return WebContextImpl{
		ContextImpl: nodes.ContextImpl{
			Adapter: adapter,
		},
		Writer:  w,
		Request: req,
	}
}

func (c WebContextImpl) GetRequest() *http.Request {
	return c.Request
}
