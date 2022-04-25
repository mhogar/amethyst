package user

import (
	"github.com/mhogar/kiwi/nodes"
	"github.com/mhogar/kiwi/nodes/converter"
	"github.com/mhogar/kiwi/nodes/crud"
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

func (u *updateUserAuthInput) GetPassword() string {
	return u.NewPassword
}

func (u *updateUserAuthInput) GetRank() int {
	return -1
}

func UpdateUserAuthWorkflow() nodes.Workflow {
	c := newUserConverter()
	v := newUserValidator()

	return nodes.NewWorkflow(
		web.NewJSONBodyParserNode[updateUserAuthInput](),
		converter.NewConverterNode(c.SetUsernameFromParams),
		validator.NewValidatorNode(v.ValidatePasswordComplexity),
		//TODO: validate old password
		converter.NewConverterNode(c.UserAuthFieldsToUserAuth),
		crud.NewUpdateModelNode[UserAuth](""),
		web.NewSuccessResponseNode(),
	)
}
