package user

import (
	"github.com/mhogar/kiwi/example/models"
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/auth"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/query"
	"github.com/mhogar/kiwi/nodes/validator"
	"github.com/mhogar/kiwi/nodes/web"
)

type updateUserAuthInput struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (u *updateUserAuthInput) GetUsername() string {
	return u.Username
}

func (u *updateUserAuthInput) SetUsername(val string) {
	u.Username = val
}

func (u *updateUserAuthInput) GetIdentifier() any {
	return u.Username
}

func (u *updateUserAuthInput) GetPassword() string {
	return u.OldPassword
}

func (u *updateUserAuthInput) GetNewPassword() string {
	return u.NewPassword
}

func UpdateUserAuthWorkflow() nodes.Workflow {
	b := NewUserQueryBuilder()
	c := newUserConverter()
	v := newUserValidator()

	updateAuth := nodes.NewWorkflow(
		auth.NewAuthenticateNode[models.UserAuth](),
		validator.NewValidatorNode(v.ValidatePasswordComplexity),
		converter.NewConverterNode(c.UserAuthFieldsToUserAuth),
		crud.NewUpdateModelNode[models.UserAuth](""),
	)

	deleteOtherSessions := nodes.NewWorkflow(
		query.NewBuildQueryNode(b.FindOtherUserSessions),
		crud.NewDeleteModelsNode[models.Session](),
	)

	return nodes.NewWorkflow(
		nodes.NewSplitWorkflowNode(
			updateAuth,
			deleteOtherSessions,
		),
	)
}

func UpdateUserAuthEndpoint() nodes.Workflow {
	c := newUserConverter()

	return nodes.NewWorkflow(
		web.SetSessionContextFromAuthorizationHeaderWorkflow[models.Session](),
		web.NewJSONBodyParserNode[updateUserAuthInput](),
		converter.NewConverterNode(c.SetUsernameFromSession),
		UpdateUserAuthWorkflow(),
		web.NewSuccessResponseNode(),
	)
}
