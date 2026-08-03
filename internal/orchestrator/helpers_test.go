package orchestrator

import (
	"encoding/json"
	"testing"
)

func MustMarshal(t *testing.T, v any) string {
	t.Helper()

	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal: %v", err)
	}

	return string(data)
}
