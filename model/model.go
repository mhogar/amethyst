package model

type Field struct {
	DataType   string
	Value      interface{}
	Validators []Validator
	Converter  Converter
}

type Model map[string]Field
