package web

import (
	"net/http"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/nodes"
)

type HTTPContext interface {
	nodes.Context

	GetRequest() *http.Request
}

type HTTPContextImpl struct {
	nodes.ContextImpl

	Writer  http.ResponseWriter
	Request *http.Request
}

func NewHTTPContext(adapter adapter.DataAdapter, w http.ResponseWriter, req *http.Request) HTTPContextImpl {
	return HTTPContextImpl{
		ContextImpl: nodes.NewContext(adapter),
		Writer:      w,
		Request:     req,
	}
}

func (c HTTPContextImpl) GetRequest() *http.Request {
	return c.Request
}
