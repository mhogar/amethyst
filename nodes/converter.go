package nodes

import (
	"github.com/amethyst/converter"
)

type ConverterNode struct {
	Converter converter.Converter
}

func (n ConverterNode) Run(input interface{}) interface{} {
	return n.Converter.Convert(input)
}
