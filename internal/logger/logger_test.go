package logger

import (
	"testing"

	"go.uber.org/zap"
)

func TestInit(t *testing.T) {

	// Reset logger
	Log = nil

	err := Init()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if Log == nil {
		t.Fatal("logger was not initialized")
	}
}

func TestInit_MultipleTimes(t *testing.T) {

	err := Init()
	if err != nil {
		t.Fatal(err)
	}

	first := Log

	err = Init()
	if err != nil {
		t.Fatal(err)
	}

	if Log == nil {
		t.Fatal("logger is nil")
	}

	if first == nil {
		t.Fatal("first logger is nil")
	}
}

func TestLoggerUsage(t *testing.T) {

	err := Init()
	if err != nil {
		t.Fatal(err)
	}

	// Verify logger can actually log without panic.
	Log.Info("test info log")
	Log.Error("test error log")

	if Log.Core() == nil {
		t.Fatal("logger core is nil")
	}

	if _, ok := interface{}(Log).(*zap.Logger); !ok {
		t.Fatal("unexpected logger type")
	}
}
