package web

import (
	"log"
	"net/http"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/nodes"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	Adapter  adapter.DataAdapter
	Workflow nodes.Workflow
}

func NewHandler(adapter adapter.DataAdapter, workflow nodes.Workflow) Handler {
	return Handler{
		Adapter:  adapter,
		Workflow: workflow,
	}
}

func (h Handler) serve(ctx interface{}, w http.ResponseWriter) {
	_, err := h.Workflow.Run(ctx, nil)
	if err != nil {
		if err.Type == nodes.ERROR_TYPE_CLIENT {
			sendClientErrorResponse(w, http.StatusBadRequest, err.Errors)
		} else {
			log.Println(err.Errors)
			sendInternalErrorResponse(w)
		}
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	h.serve(NewHTTPContext(h.Adapter, w, req), w)
}

func (h Handler) ServeHTTPRouter(w http.ResponseWriter, req *http.Request, params httprouter.Params) {
	h.serve(NewHTTPRouterContext(h.Adapter, w, req, params), w)
}
