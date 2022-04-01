package nodes

type Node interface {
	Run(ctx interface{}, input any) (any, *Error)
}
