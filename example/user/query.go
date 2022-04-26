package user

import (
	"github.com/mhogar/kiwi/data/query"
	"github.com/mhogar/kiwi/example/models"
	"github.com/mhogar/kiwi/nodes/web"
)

type UserQueryBuilder struct{}

func NewUserQueryBuilder() UserQueryBuilder {
	return UserQueryBuilder{}
}

func (UserQueryBuilder) FindUserSessions(ctx interface{}, _ any) (*query.WhereClause, error) {
	session := ctx.(web.HTTPContext).GetSession().(*models.Session)
	return query.Where("username", "=", session.Username), nil
}

func (UserQueryBuilder) FindOtherUserSessions(ctx interface{}, _ any) (*query.WhereClause, error) {
	session := ctx.(web.HTTPContext).GetSession().(*models.Session)
	return query.Where("username", "=", session.Username).And(query.Where("token", "!=", session.Token)), nil
}
