package session

import (
	"github.com/google/uuid"
	"github.com/mhogar/kiwi/example/user"
)

type sessionConverter struct{}

func newSessionConverter() sessionConverter {
	return sessionConverter{}
}

func (sessionConverter) UserToSession(_ interface{}, val any) (any, error) {
	user := val.(*user.User)
	return CreateNewSession(user.Username, user.Rank), nil
}

func (sessionConverter) SessionToResponse(_ interface{}, val any) (any, error) {
	session := val.(*Session)
	return newSessionResponse(session), nil
}

func (sessionConverter) NewSessionFromToken(_ interface{}, val any) (any, error) {
	return CreateSession(val.(uuid.UUID), "", 0), nil
}
