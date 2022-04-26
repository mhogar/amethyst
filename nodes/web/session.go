package web

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/session"
)

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

	return token, nil
}

func SetSessionContextFromAuthorizationHeaderWorkflow[Model session.Session]() nodes.Workflow {
	return nodes.NewWorkflow(
		NewParseTokenFromAuthorizationHeaderNode(),
		session.GetSessionWorkflow[Model]("bearer token invalid or expired"),
		session.NewSetSessionContextNode(),
	)
}
