package adapter

import (
	"reflect"
)

type ReflectModel struct {
	fields map[string]int

	Name      string
	Fields    []string
	Values    []interface{}
	Addresses []interface{}
}

func CreateReflectModel[T any]() ReflectModel {
	rm := ReflectModel{
		fields: map[string]int{},
	}

	var model T
	st := reflect.TypeOf(model)
	rm.Name = st.Name()

	for i := 0; i < st.NumField(); i++ {
		tag := st.Field(i).Tag.Get("kiwi")
		if tag != "" {
			rm.fields[tag] = i
			rm.Fields = append(rm.Fields, tag)
		}
	}

	return rm
}

func (rm *ReflectModel) UniqueField() string {
	return rm.Fields[0]
}

func (rm *ReflectModel) UniqueValue() interface{} {
	return rm.Values[0]
}

func (rm *ReflectModel) SetModel(model interface{}) {
	sv := reflect.ValueOf(model).Elem()

	rm.Values = []interface{}{}
	rm.Addresses = []interface{}{}

	for _, index := range rm.fields {
		field := sv.Field(index)
		rm.Values = append(rm.Values, field.Interface())
		rm.Addresses = append(rm.Addresses, field.Addr())
	}
}
