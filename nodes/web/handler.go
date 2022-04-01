package web

import (
	"log"
	"net/http"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/nodes"
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

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := NewWebContext(h.Adapter, w, req)

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
