package converter

import (
	"golang.org/x/crypto/bcrypt"
)

type BaseConverter struct{}

func (BaseConverter) HashPassword(password string) []byte {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hash
}
