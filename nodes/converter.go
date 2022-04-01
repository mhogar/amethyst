package nodes

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/nodes/converter"
)

type ConverterNode[Context any] struct {
	Converter converter.Converter[Context]
}

func (f NodeFactory[Context, Model]) Converter(c converter.Converter[Context]) ConverterNode[Context] {
	return ConverterNode[Context]{
		Converter: c,
	}
}

func (n ConverterNode[Context]) Run(ctx Context, input interface{}) (interface{}, *Error) {
	output, err := n.Converter.Convert(ctx, input)
	if err != nil {
		return nil, InternalError(common.ChainError("error converting model", err))
	}

	return output, nil
}
