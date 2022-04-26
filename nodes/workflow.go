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

type SplitWorkflowNode struct {
	Branches []Workflow
}

func NewSplitWorkflowNode(branches ...Workflow) SplitWorkflowNode {
	return SplitWorkflowNode{
		Branches: branches,
	}
}

func (w SplitWorkflowNode) Run(ctx interface{}, input any) (any, *Error) {
	for _, branch := range w.Branches {
		_, err := branch.Run(ctx, input)
		if err != nil {
			return nil, err
		}
	}

	return input, nil
}
