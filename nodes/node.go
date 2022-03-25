package nodes

type Node[T any] interface {
	Run(ctx T, input interface{}) (interface{}, *Error)
}

type NodeFactory[T any] struct{}
