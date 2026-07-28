package routes

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"
	"github.com/gin-gonic/gin"
)

func RegisterJobMatchRoutes(r *gin.Engine) {
	r.POST("/job/match", handler.MatchJob)
}