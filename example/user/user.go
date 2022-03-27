package user

type UserInput struct {
	Username string
	Password string
	Rank     int
}

func CreateNewUserInput(username string, password string, rank int) *UserInput {
	return &UserInput{
		Username: username,
		Password: password,
		Rank:     rank,
	}
}

type User struct {
	Username     string `kiwi:"username,unique"`
	PasswordHash []byte `kiwi:"password_hash"`
	Rank         int    `kiwi:"rank"`
}

func CreateNewUser(username string, hash []byte, rank int) *User {
	return &User{
		Username:     username,
		PasswordHash: hash,
		Rank:         rank,
	}
}
