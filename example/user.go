package example

import (
	"errors"

	"github.com/amethyst/model"
)

func User() model.Model {
	return model.Model{
		"username": {
			DataType: "varstr(30)",
			Validators: []model.Validator{
				model.LengthValidator{
					Min: 5,
					Max: 30,
				},
			},
		},
		"password_hash": {
			Validators: []model.Validator{
				model.LengthValidator{
					Min: 8,
				},
			},
			Converter: PasswordHashConverter{},
		},
	}
}

type PasswordHashConverter struct{}

func (PasswordHashConverter) Convert(val interface{}) (*model.Field, error) {
	str, ok := val.(string)
	if !ok {
		return nil, errors.New("value is not a string")
	}

	//TODO: actually hash password
	hash := []byte(str)

	return &model.Field{
		DataType: "bytes",
		Value:    hash,
	}, nil
}
