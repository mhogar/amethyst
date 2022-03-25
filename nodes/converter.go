package nodes

import (
	"github.com/mhogar/kiwi/nodes/converter"
)

type ConverterNode[T any] struct {
	Converter converter.Converter[T]
}

func (f NodeFactory[T]) Converter(c converter.Converter[T]) ConverterNode[T] {
	return ConverterNode[T]{
		Converter: c,
	}
}

func (n ConverterNode[T]) Run(ctx T, input interface{}) (interface{}, *Error) {
	return n.Converter.Convert(ctx, input), nil
}
