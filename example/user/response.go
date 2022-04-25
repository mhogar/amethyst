package user

type UserResponse struct {
	Username string `json:"username"`
	Rank     int    `json:"rank"`
}

func newUserResponse(user *User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Rank:     user.Rank,
	}
}
