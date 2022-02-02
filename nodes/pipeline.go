package nodes

type Pipeline []Node

func (p Pipeline) Run(input interface{}) interface{} {
	for _, node := range p {
		input = node.Run(input)
	}

	return input
}
