package user

import (
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
	"github.com/mhogar/kiwi/nodes/validator"
	"github.com/mhogar/kiwi/nodes/web"
)

type createUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Rank     int    `json:"rank"`
}

func (u *createUserInput) GetUsername() string {
	return u.Username
}

func (u *createUserInput) SetUsername(val string) {
	u.Username = val
}

func (u *createUserInput) GetNewPassword() string {
	return u.Password
}

func (u *createUserInput) GetRank() int {
	return u.Rank
}

func CreateUserWorkflow() nodes.Workflow {
	c := newUserConverter()
	v := newUserValidator()

	createUser := nodes.NewWorkflow(
		converter.NewConverterNode(c.UserFieldsToUser),
		validator.NewValidatorNode(v.ValidateUser),
		validator.NewValidatorNode(v.ValidateUserUnique),
		crud.NewCreateModelNode[User](),
	)

	createUserAuth := nodes.NewWorkflow(
		validator.NewValidatorNode(v.ValidatePasswordComplexity),
		converter.NewConverterNode(c.UserAuthFieldsToUserAuth),
		crud.NewCreateModelNode[UserAuth](),
	)

	sendResponse := nodes.NewWorkflow(
		converter.NewConverterNode(c.UserFieldsToResponse),
		web.NewDataResponseNode(),
	)

	return nodes.NewWorkflow(
		web.NewJSONBodyParserNode[createUserInput](),
		nodes.NewSplitWorkflowNode(
			createUser,
			createUserAuth,
			sendResponse,
		),
	)
}
