package session

import (
	"github.com/mhogar/kiwi/example/user"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/auth"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/web"
)

type createSessionInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *createSessionInput) GetIdentifier() any {
	return s.Username
}

func (s *createSessionInput) GetUsername() string {
	return s.Username
}

func (s *createSessionInput) GetPassword() string {
	return s.Password
}

func CreateSessionWorkflow() nodes.Workflow {
	c := newSessionConverter()

	return nodes.NewWorkflow(
		auth.NewAuthenticateNode[user.UserAuth](),
		user.GetUserWorkflow(),
		converter.NewConverterNode(c.UserToSession),
		crud.NewCreateModelNode[Session](),
	)
}

func CreateSessionEndpoint() nodes.Workflow {
	c := newSessionConverter()

	return nodes.NewWorkflow(
		web.NewJSONBodyParserNode[createSessionInput](),
		CreateSessionWorkflow(),
		converter.NewConverterNode(c.SessionToResponse),
		web.NewDataResponseNode(),
	)
}
