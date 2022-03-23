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
	rm := ReflectModel{}

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
		rm.Values = append(rm.Values, sv.Field(index))
		rm.Addresses = append(rm.Addresses, sv.Field(index).Addr())
	}
}
