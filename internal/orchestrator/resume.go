package orchestrator

import (
	"encoding/json"
	"time"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	"github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/parser"
	"github.com/Aditya7880900936/ai-orchestrator/internal/retry"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func AnalyzeResume(req model.ResumeAnalyzeRequest) (*model.ResumeAnalyzeResponse, error) {

	// Cache Key
	key := cache.GenerateKey("resume:" + req.ResumeText)

	// Check Cache
	if data, err := cache.Get(key); err == nil && data != "" {
		var res model.ResumeAnalyzeResponse
		if err := json.Unmarshal([]byte(data), &res); err == nil {
			return &res, nil
		}
	}

	// Workflow
	wf := workflow.NewResumeWorkflow()

	raw, err := retry.Execute(
		3,
		time.Second,
		func() (string, error) {
			return wf.Run(req.ResumeText)
		},
	)
	if err != nil {
		return nil, err
	}

	// Extract JSON
	jsonText := parser.ExtractJSON(raw)

	// Parse
	res, err := parser.Parse[model.ResumeAnalyzeResponse](jsonText)
	if err != nil {
		return nil, err
	}

	// Cache Result
	bytes, _ := json.Marshal(res)
	_ = cache.Set(key, string(bytes))

	return res, nil
}
