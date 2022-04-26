package user

type User struct {
	Username string `kiwi:"username,unique"`
	Rank     int    `kiwi:"rank"`
}

func (u *User) GetUsername() string {
	return u.Username
}

type UserAuth struct {
	Username     string `kiwi:"username,unique"`
	PasswordHash []byte `kiwi:"password_hash"`
}

func (u UserAuth) GetPasswordHash() []byte {
	return u.PasswordHash
}

func NewUser(username string, rank int) *User {
	return &User{
		Username: username,
		Rank:     rank,
	}
}

func NewUserAuth(username string, hash []byte) *UserAuth {
	return &UserAuth{
		Username:     username,
		PasswordHash: hash,
	}
}
