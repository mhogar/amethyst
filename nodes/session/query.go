package session

import (
	"github.com/google/uuid"
	"github.com/mhogar/kiwi/data/query"
)

type sessionQueryBuilder struct{}

func newSessionQueryBuilder() sessionQueryBuilder {
	return sessionQueryBuilder{}
}

func (sessionQueryBuilder) GetSessionByToken(_ interface{}, input any) (*query.WhereClause, error) {
	token := input.(uuid.UUID)
	return query.Where("token", "=", token), nil
}
