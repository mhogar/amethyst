package nodes

type Node interface {
	Run(input interface{}) interface{}
}
