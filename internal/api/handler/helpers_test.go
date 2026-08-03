package handler

import (
	"testing"

	"github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/gin-gonic/gin"
)

func setupHandlerTest(t *testing.T) {

	t.Helper()

	gin.SetMode(gin.TestMode)

	if err := logger.Init(); err != nil {
		t.Fatalf("failed to init logger: %v", err)
	}
}