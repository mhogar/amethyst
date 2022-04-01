package converter

import (
	"github.com/mhogar/kiwi/common"
	"github.com/mhogar/kiwi/nodes"
)

type Converter interface {
	Convert(ctx interface{}, val any) (any, error)
}

type ConverterNode struct {
	Converter Converter
}

func NewConverterNode(c Converter) ConverterNode {
	return ConverterNode{
		Converter: c,
	}
}

func (n ConverterNode) Run(ctx interface{}, input any) (any, *nodes.Error) {
	output, err := n.Converter.Convert(ctx, input)
	if err != nil {
		return nil, nodes.InternalError(common.ChainError("error converting model", err))
	}

	return output, nil
}
