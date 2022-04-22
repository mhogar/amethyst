package web

import (
	"net/http"

	"github.com/mhogar/kiwi/nodes"
)

type SuccessResponseNode struct{}

func NewSuccessResponseNode() SuccessResponseNode {
	return SuccessResponseNode{}
}

func (SuccessResponseNode) Run(ctx interface{}, _ any) (any, *nodes.Error) {
	sendJSONResponse(ctx.(HTTPContext).GetResponseWriter(), http.StatusOK, NewSuccessResponse())
	return nil, nil
}

type DataResponseNode struct{}

func NewDataResponseNode() DataResponseNode {
	return DataResponseNode{}
}

func (DataResponseNode) Run(ctx interface{}, input any) (any, *nodes.Error) {
	sendJSONResponse(ctx.(HTTPContext).GetResponseWriter(), http.StatusOK, NewSuccessDataResponse(input))
	return nil, nil
}
