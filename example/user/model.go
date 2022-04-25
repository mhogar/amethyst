package user

type User struct {
	Username string `kiwi:"username,unique"`
	Rank     int    `kiwi:"rank"`
}

type UserAuth struct {
	User         *User  `kiwi:"user,unique,ref"`
	PasswordHash []byte `kiwi:"password_hash"`
}

func NewUser(username string, rank int) *User {
	return &User{
		Username: username,
		Rank:     rank,
	}
}

func NewUserAuth(username string, hash []byte, rank int) *UserAuth {
	return &UserAuth{
		User:         NewUser(username, rank),
		PasswordHash: hash,
	}
}
