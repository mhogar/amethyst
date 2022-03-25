package converter

type Converter[T any] interface {
	Convert(ctx T, val interface{}) interface{}
}
