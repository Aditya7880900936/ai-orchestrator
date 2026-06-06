package orchestrator

type WorkflowFunc func() (string, error)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

func (e *Executor) Execute(fn WorkflowFunc) (string, error) {
	return fn()
}