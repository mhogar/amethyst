package nodes

type Node[Context any] interface {
	Run(ctx Context, input interface{}) (interface{}, *Error)
}

type NodeFactory[Context any, Model any] struct{}
