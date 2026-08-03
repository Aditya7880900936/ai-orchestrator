package orchestrator

import "github.com/Aditya7880900936/ai-orchestrator/internal/workflow"

type MockWorkflow struct {
	Response string
	Err      error
}

func (m *MockWorkflow) Run(input string) (string, error) {
	return m.Response, m.Err
}

// Compile-time interface check
var _ workflow.Workflow = (*MockWorkflow)(nil)