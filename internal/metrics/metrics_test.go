package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
)

func TestInit(t *testing.T) {

	oldRegister := mustRegister

	defer func() {
		mustRegister = oldRegister
	}()

	called := false

	mustRegister = func(cs ...prometheus.Collector) {

		called = true

		if len(cs) != 1 {
			t.Fatalf("expected 1 collector got %d", len(cs))
		}

		if cs[0] != AnalyzeRequests {
			t.Fatal("unexpected collector")
		}
	}

	Init()

	if !called {
		t.Fatal("MustRegister not called")
	}
}

func TestAnalyzeRequests_NotNil(t *testing.T) {

	if AnalyzeRequests == nil {
		t.Fatal("AnalyzeRequests should not be nil")
	}
}
