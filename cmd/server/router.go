package main

import (
	handler "github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"
	middle "github.com/Aditya7880900936/ai-orchestrator/internal/api/middleware"
	"github.com/Aditya7880900936/ai-orchestrator/internal/api/routes"
	"github.com/Aditya7880900936/ai-orchestrator/internal/middleware"

	_ "github.com/Aditya7880900936/ai-orchestrator/docs"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.Use(
		middleware.RequestID(),
		middle.Logger(),
		middle.RateLimiter(),
	)

	// Health Check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Metrics
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	r.POST("/analyze", handler.AnalyzeHandler)

	routes.RegisterResumeRoutes(r)
	routes.RegisterSkillRoutes(r)
	routes.RegisterATSRoutes(r)
	routes.RegisterJobMatchRoutes(r)
	routes.RegisterResumeImproveRoutes(r)
	routes.RegisterCoverLetterRoutes(r)
	routes.RegisterResumeChatRoutes(r)

	return r
}
