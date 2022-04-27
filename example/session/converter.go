package session

import (
	"github.com/mhogar/kiwi/example/models"
)

type sessionConverter struct{}

func newSessionConverter() sessionConverter {
	return sessionConverter{}
}

func (sessionConverter) UserToSession(_ interface{}, val any) (any, error) {
	user := val.(*models.User)
	return models.CreateNewSession(user.Username, user.Rank), nil
}

func (sessionConverter) SessionToResponse(_ interface{}, val any) (any, error) {
	session := val.(*models.Session)
	return newSessionResponse(session), nil
}
