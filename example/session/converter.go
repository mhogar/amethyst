package session

import (
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
