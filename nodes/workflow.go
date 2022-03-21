package nodes

type Workflow []Node

func (f NodeFactory) Workflow(nodes ...Node) Workflow {
	return nodes
}

func (w Workflow) Run(input interface{}) interface{} {
	for _, node := range w {
		input = node.Run(input)
	}

	return input
}
