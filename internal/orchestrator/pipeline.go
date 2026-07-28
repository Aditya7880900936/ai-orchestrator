package orchestrator

import (
	"encoding/json"
	"fmt"
	"strings"
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

type unwrap struct {
	Answer string `json:"answer"`
}

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

		fmt.Println("Cache unmarshal failed, regenerating...")
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

	// Extract JSON
	jsonText := parser.ExtractJSON(raw)

	var u unwrap
	if err := json.Unmarshal([]byte(jsonText), &u); err == nil {
		if strings.HasPrefix(strings.TrimSpace(u.Answer), "{") {
			jsonText = u.Answer
		}
	}

	// Parse
	res, err := parser.Parse[T](jsonText)
	if err != nil {
		return nil, err
	}

	// Cache
	bytes, _ := json.Marshal(res)
	_ = cache.Set(cacheKey, string(bytes))


	return res, nil
}
