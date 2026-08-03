package workflow

type Workflow interface {
	Run(input string) (string, error)
}
