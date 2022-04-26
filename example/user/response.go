package user

import "github.com/mhogar/kiwi/example/models"

type UserResponse struct {
	Username string `json:"username"`
	Rank     int    `json:"rank"`
}

func newUserResponse(user *models.User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Rank:     user.Rank,
	}
}
