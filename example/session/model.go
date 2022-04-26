package session

import "github.com/google/uuid"

type Session struct {
	Token    uuid.UUID `kiwi:"token,unique"`
	Username string    `kiwi:"username"`
	Rank     int       `kiwi:"rank"`
}

func CreateSession(token uuid.UUID, username string, rank int) *Session {
	return &Session{
		Token:    token,
		Username: username,
		Rank:     rank,
	}
}

func CreateNewSession(username string, rank int) *Session {
	return CreateSession(uuid.New(), username, rank)
}
