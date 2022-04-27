package session

import (
	"github.com/mhogar/kiwi/data/adapter"
	"github.com/mhogar/kiwi/nodes"
)

type SessionContext interface {
	nodes.Context

	GetSession() Session
	SetSession(val Session)
}

type SessionContextImpl struct {
	*nodes.ContextImpl

	Session Session
}

func NewSessionContext(adapter adapter.DataAdapter) *SessionContextImpl {
	return &SessionContextImpl{
		ContextImpl: nodes.NewContext(adapter),
	}
}

func (c *SessionContextImpl) GetSession() Session {
	return c.Session
}

func (c *SessionContextImpl) SetSession(val Session) {
	c.Session = val
}

type SetSessionContextNode struct{}

func NewSetSessionContextNode() SetSessionContextNode {
	return SetSessionContextNode{}
}

func (SetSessionContextNode) Run(ctx interface{}, input any) (any, *nodes.Error) {
	ctx.(SessionContext).SetSession(input.(Session))
	return input, nil
}
