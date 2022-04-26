package user

type UsernameGetter interface {
	GetUsername() string
}

type UsernameField interface {
	GetUsername() string
	SetUsername(val string)
}

type RankField interface {
	GetRank() int
}

type UserFields interface {
	UsernameField
	RankField
}

type PasswordField interface {
	GetNewPassword() string
}

type UserAuthFields interface {
	UsernameField
	PasswordField
}
