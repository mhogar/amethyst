package nodes

type Workflow[T any] []Node[T]

func (f NodeFactory[T]) Workflow(nodes ...Node[T]) Workflow[T] {
	return nodes
}

func (w Workflow[T]) Run(ctx T, input interface{}) (interface{}, *Error) {
	var err *Error

	for _, node := range w {
		input, err = node.Run(ctx, input)
		if err != nil {
			break
		}
	}

	return input, err
}
