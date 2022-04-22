package user

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Rank     int    `json:"rank"`
}

type UserResponse struct {
	Username string `json:"username"`
	Rank     int    `json:"rank"`
}

type User struct {
	Username     string `kiwi:"username,unique"`
	PasswordHash []byte `kiwi:"password_hash"`
	Rank         int    `kiwi:"rank"`
}

func NewUser(username string, hash []byte, rank int) *User {
	return &User{
		Username:     username,
		PasswordHash: hash,
		Rank:         rank,
	}
}

func newUserResponse(user *User) UserResponse {
	return UserResponse{
		Username: user.Username,
		Rank:     user.Rank,
	}
}
