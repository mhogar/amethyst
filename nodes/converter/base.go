package converter

type BaseConverter interface {
	HashPassword(password string) ([]byte, error)
}

type BaseConverterImpl struct {
	PasswordHasher PasswordHasher
}

func (c BaseConverterImpl) HashPassword(password string) ([]byte, error) {
	hash, err := c.PasswordHasher.HashPassword(password)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
