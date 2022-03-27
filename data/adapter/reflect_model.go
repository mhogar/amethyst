package adapter

import (
	"reflect"
	"strings"
)

type ReflectModel struct {
	fieldIndices     []int
	uniqueFieldIndex int

	Name      string
	Fields    []string
	Values    []interface{}
	Addresses []interface{}
}

func CreateReflectModel[T any]() ReflectModel {
	m := ReflectModel{}

	var model T
	st := reflect.TypeOf(model)
	m.Name = st.Name()

	for index := 0; index < st.NumField(); index++ {
		tokens := strings.Split(st.Field(index).Tag.Get("kiwi"), ",")

		if len(tokens) >= 1 {
			m.fieldIndices = append(m.fieldIndices, index)
			m.Fields = append(m.Fields, tokens[0])
		}

		if len(tokens) >= 2 && tokens[1] == "unique" {
			m.uniqueFieldIndex = index
		}
	}

	return m
}

func (m *ReflectModel) UniqueField() string {
	return m.Fields[m.uniqueFieldIndex]
}

func (m *ReflectModel) UniqueValue() interface{} {
	return m.Values[m.uniqueFieldIndex]
}

func (m *ReflectModel) SetModel(model interface{}) {
	sv := reflect.ValueOf(model).Elem()

	m.Values = []interface{}{}
	m.Addresses = []interface{}{}

	for index := range m.fieldIndices {
		field := sv.Field(index)
		m.Values = append(m.Values, field.Interface())
		m.Addresses = append(m.Addresses, field.Addr().Interface())
	}
}
