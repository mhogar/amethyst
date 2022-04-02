package web

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mhogar/kiwi/data/adapter"
)

type HTTPRouterContext interface {
	HTTPContext

	GetParams() httprouter.Params
}

type HTTPRouterContextImpl struct {
	HTTPContextImpl

	Params httprouter.Params
}

func NewHTTPRouterContext(adapter adapter.DataAdapter, w http.ResponseWriter, req *http.Request, params httprouter.Params) HTTPRouterContextImpl {
	return HTTPRouterContextImpl{
		HTTPContextImpl: NewHTTPContext(adapter, w, req),
		Params:          params,
	}
}

func (c HTTPRouterContextImpl) GetParams() httprouter.Params {
	return c.Params
}
