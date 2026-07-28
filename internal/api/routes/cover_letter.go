package routes

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterCoverLetterRoutes(r *gin.Engine) {
	r.POST("/cover-letter/generate", handler.GenerateCoverLetter)
}