package session

import "github.com/mhogar/kiwi/example/models"

type SessionResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Rank     int    `json:"rank"`
}

func newSessionResponse(s *models.Session) SessionResponse {
	return SessionResponse{
		Token:    s.Token.String(),
		Username: s.Username,
		Rank:     s.Rank,
	}
}
