package model

type Converter interface {
	Convert(val interface{}) (*Field, error)
}
