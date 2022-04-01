package nodes

type Workflow[Context any] []Node[Context]

func (f NodeFactory[Context, Model]) Workflow(nodes ...Node[Context]) Workflow[Context] {
	return nodes
}

func (w Workflow[Context]) Run(ctx Context, input interface{}) (interface{}, *Error) {
	var err *Error

	for _, node := range w {
		input, err = node.Run(ctx, input)
		if err != nil {
			break
		}
	}

	return input, err
}
