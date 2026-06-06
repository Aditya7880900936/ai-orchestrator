package orchestrator

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/parser"
	"github.com/Aditya7880900936/ai-orchestrator/internal/retry"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func Analyze(req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {

	cacheKey := cache.GenerateKey(req.Prompt)

	cachedResponse, err := cache.Get(cacheKey)
	if err == nil {

		var cached models.AnalyzeResponse

		if err := json.Unmarshal([]byte(cachedResponse), &cached); err == nil {

			fmt.Println("CACHE HIT")

			return &cached, nil
		}
	}

	fmt.Println("CACHE MISS")

	wf := workflow.NewAnalyzeWorkflow()

	resp, err := retry.Execute(
		3,
		2*time.Second,
		func() (string, error) {
			return wf.Run(req.Prompt)
		},
	)

	if err != nil {
		return nil, err
	}

	cleaned := parser.ExtractJSON(resp)

	parsed, err := parser.Parse[models.AnalyzeResponse](cleaned)
	if err != nil {
		return nil, err
	}

	serialized, _ := json.Marshal(parsed)

	fmt.Println("WRITING TO CACHE")

	_ = cache.Set(
		cacheKey,
		string(serialized),
	)

	return parsed, nil
}