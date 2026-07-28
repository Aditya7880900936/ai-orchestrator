package orchestrator

import "github.com/Aditya7880900936/ai-orchestrator/internal/workflow"

type WorkflowFunc func() (string, error)

type Executor struct {
	Workflow workflow.Workflow
}

func NewExecutor(w workflow.Workflow) *Executor {
	return &Executor{
		Workflow: w,
	}
}

func (e *Executor) Execute(fn WorkflowFunc) (string, error) {
	return fn()
}
