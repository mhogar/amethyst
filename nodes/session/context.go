package session

import "github.com/mhogar/kiwi/nodes"

type SessionContext interface {
	GetSession() Session
	SetSession(val Session)
}

type SessionContextImpl struct {
	Session Session
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
