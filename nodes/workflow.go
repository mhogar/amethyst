package nodes

type Workflow []Node

func NewWorkflow(nodes ...Node) Workflow {
	return nodes
}

func (w Workflow) Run(ctx interface{}, input any) (any, *Error) {
	var err *Error

	for _, node := range w {
		input, err = node.Run(ctx, input)
		if err != nil {
			break
		}
	}

	return input, err
}
