package session

type SessionResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Rank     int    `json:"rank"`
}

func newSessionResponse(s *Session) SessionResponse {
	return SessionResponse{
		Token:    s.Token.String(),
		Username: s.Username,
		Rank:     s.Rank,
	}
}
