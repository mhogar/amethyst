package user

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
	GetPassword() string
}

type UserAuthFields interface {
	UserFields
	PasswordField
}
