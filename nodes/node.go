package nodes

type Node[T any] interface {
	Run(ctx T, input interface{}) interface{}
}

type NodeFactory[T any] struct{}
