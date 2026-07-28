package orchestrator

import (
	"encoding/json"
	"time"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	"github.com/Aditya7880900936/ai-orchestrator/internal/parser"
	"github.com/Aditya7880900936/ai-orchestrator/internal/retry"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

const (
	DefaultRetryAttempts = 3
	DefaultRetryDelay    = time.Second
)

func ExecutePipeline[T any](
	cacheKey string,
	input string,
	wf workflow.Workflow,
) (*T, error) {

	// Cache
	if data, err := cache.Get(cacheKey); err == nil && data != "" {
		var res T
		if err := json.Unmarshal([]byte(data), &res); err == nil {
			return &res, nil
		}
	}

	// Workflow
	raw, err := retry.Execute(
		DefaultRetryAttempts,
		DefaultRetryDelay,
		func() (string, error) {
			return wf.Run(input)
		},
	)
	if err != nil {
		return nil, err
	}

	// Parse
	jsonText := parser.ExtractJSON(raw)

	res, err := parser.Parse[T](jsonText)
	if err != nil {
		return nil, err
	}

	// Cache
	bytes, _ := json.Marshal(res)
	_ = cache.Set(cacheKey, string(bytes))

	return res, nil
}