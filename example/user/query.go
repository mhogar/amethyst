package user

import "github.com/mhogar/kiwi/data/query"

type UserQueryBuilder struct{}

func NewUserQueryBuilder() UserQueryBuilder {
	return UserQueryBuilder{}
}

func (UserQueryBuilder) GetUserByUsername(_ interface{}, input any) (*query.WhereClause, error) {
	user := input.(*User)
	return query.Where("username", "=", user.Username), nil
}
