package routes

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"

	"github.com/gin-gonic/gin"
)

func RegisterResumeRoutes(r *gin.Engine) {
	r.POST("/resume/analyze", handler.AnalyzeResume)
}