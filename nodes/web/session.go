package web

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/session"
)

type SessionToken struct {
	ID uuid.UUID
}

func NewSessionToken(id uuid.UUID) SessionToken {
	return SessionToken{
		ID: id,
	}
}

func (t SessionToken) GetIdentifier() any {
	return t.ID
}

type ParseTokenFromAuthorizationHeaderNode struct{}

func NewParseTokenFromAuthorizationHeaderNode() ParseTokenFromAuthorizationHeaderNode {
	return ParseTokenFromAuthorizationHeaderNode{}
}

func (ParseTokenFromAuthorizationHeaderNode) Run(ctx interface{}, input any) (any, *nodes.Error) {
	req := ctx.(HTTPContext).GetRequest()

	splitTokens := strings.Split(req.Header.Get("Authorization"), "Bearer ")
	if len(splitTokens) != 2 {
		return nil, nodes.ClientError(errors.New("no bearer token provided"))
	}

	token, err := uuid.Parse(splitTokens[1])
	if err != nil {
		return nil, nodes.ClientError(errors.New("bearer token was in an invalid format"))
	}

	return NewSessionToken(token), nil
}

func SetSessionContextFromAuthorizationHeaderWorkflow[Model session.Session]() nodes.Workflow {
	return nodes.NewWorkflow(
		NewParseTokenFromAuthorizationHeaderNode(),
		crud.NewReadUniqueModelNode[Model]("bearer token invalid or expired"),
		session.NewSetSessionContextNode(),
	)
}
