package parser

import "encoding/json"

func Parse[T any](raw string) (*T, error) {

	var result T

	err := json.Unmarshal([]byte(raw), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}