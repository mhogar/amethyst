package user

import (
	"github.com/mhogar/kiwi/example/models"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/validator"
	"github.com/mhogar/kiwi/nodes/web"
)

type updateUserInput struct {
	Username string `json:"username,omitempty"`
	Rank     int    `json:"rank"`
}

func (u *updateUserInput) GetUsername() string {
	return u.Username
}

func (u *updateUserInput) SetUsername(val string) {
	u.Username = val
}

func (u *updateUserInput) GetRank() int {
	return u.Rank
}

func UpdateUserWorkflow() nodes.Workflow {
	c := newUserConverter()
	v := newUserValidator()

	return nodes.NewWorkflow(
		converter.NewConverterNode(c.UserFieldsToUser),
		validator.NewValidatorNode(v.ValidateUser),
		crud.NewUpdateModelNode[models.User]("user with username not found"),
	)
}

func UpdateUserEndpoint() nodes.Workflow {
	c := newUserConverter()

	return nodes.NewWorkflow(
		web.SetSessionContextFromAuthorizationHeaderWorkflow[models.Session](),
		web.NewJSONBodyParserNode[updateUserInput](),
		converter.NewConverterNode(c.SetUsernameFromSession),
		UpdateUserWorkflow(),
		converter.NewConverterNode(c.UserToResponse),
		web.NewDataResponseNode(),
	)
}
