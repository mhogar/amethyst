package converter

type Converter interface {
	Convert(val interface{}) interface{}
}
