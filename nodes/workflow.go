package nodes

type Workflow[T any] []Node[T]

func (f NodeFactory[T]) Workflow(nodes ...Node[T]) Workflow[T] {
	return nodes
}

func (w Workflow[T]) Run(ctx T, input interface{}) interface{} {
	for _, node := range w {
		input = node.Run(ctx, input)
	}

	return input
}
