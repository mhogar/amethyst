package nodes

import (
	"github.com/mhogar/kiwi/common"
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
	output, err := n.Converter.Convert(ctx, input)
	if err != nil {
		return nil, InternalError(common.ChainError("error converting model", err))
	}

	return output, nil
}
