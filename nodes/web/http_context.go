package web

import (
	"net/http"

	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/nodes/session"
)

type HTTPContext interface {
	session.SessionContext

	GetRequest() *http.Request
	GetResponseWriter() http.ResponseWriter
}

type HTTPContextImpl struct {
	*session.SessionContextImpl

	Request        *http.Request
	ResponseWriter http.ResponseWriter
}

func NewHTTPContext(adapter adapter.DataAdapter, w http.ResponseWriter, req *http.Request) *HTTPContextImpl {
	return &HTTPContextImpl{
		SessionContextImpl: session.NewSessionContext(adapter),
		ResponseWriter:     w,
		Request:            req,
	}
}

func (c *HTTPContextImpl) GetRequest() *http.Request {
	return c.Request
}

func (c *HTTPContextImpl) GetResponseWriter() http.ResponseWriter {
	return c.ResponseWriter
}
