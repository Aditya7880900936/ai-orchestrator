package parser

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Parse[T any](raw string) (*T, error) {

	var result T

	// First attempt
	if err := json.Unmarshal([]byte(raw), &result); err == nil {
		return &result, nil
	}

	// Check if LLM wrapped JSON inside "answer"
	var wrapper map[string]any
	if err := json.Unmarshal([]byte(raw), &wrapper); err != nil {
		return nil, err
	}

	answer, ok := wrapper["answer"].(string)
	if !ok {
		return nil, fmt.Errorf("unable to parse response")
	}

	answer = strings.TrimSpace(answer)

	if !strings.HasPrefix(answer, "{") {
		return nil, fmt.Errorf("wrapped answer is not json")
	}

	// Try parsing wrapped JSON
	if err := json.Unmarshal([]byte(answer), &result); err == nil {
		return &result, nil
	}

	// Handle:
	// {"answer":["a","b","c"]}
	var nested struct {
		Answer []string `json:"answer"`
	}

	if err := json.Unmarshal([]byte(answer), &nested); err == nil {

		out := strings.Join(nested.Answer, "\n")

		fixed := fmt.Sprintf(`{"answer":%q}`, out)

		if err := json.Unmarshal([]byte(fixed), &result); err == nil {
			return &result, nil
		}
	}

	return nil, fmt.Errorf("failed to parse llm response")
}