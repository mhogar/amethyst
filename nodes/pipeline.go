package nodes

type Pipeline struct {
	Nodes []Node
}

func (p *Pipeline) Build(nodes ...Node) {
	p.Nodes = append(p.Nodes, nodes...)
}

func (p Pipeline) Run(input interface{}) interface{} {
	for _, node := range p.Nodes {
		input = node.Run(input)
	}

	return input
}
