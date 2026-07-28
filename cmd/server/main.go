package main

import (
	"log"

	handler "github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"
	"github.com/Aditya7880900936/ai-orchestrator/internal/api/routes"
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	llm "github.com/Aditya7880900936/ai-orchestrator/internal/llm"
	"github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/Aditya7880900936/ai-orchestrator/internal/metrics"
	"github.com/Aditya7880900936/ai-orchestrator/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	cache.InitRedis()
	metrics.Init()

	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}

	if err := llm.InitGemini(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(
		middleware.RequestID(),
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	r.POST("/analyze", handler.AnalyzeHandler)
	routes.RegisterResumeRoutes(r)
	routes.RegisterSkillRoutes(r)
	routes.RegisterATSRoutes(r)
	routes.RegisterJobMatchRoutes(r)
	routes.RegisterResumeImproveRoutes(r)
	routes.RegisterCoverLetterRoutes(r)
	routes.RegisterResumeChatRoutes(r)
	log.Println("Server running on :8080")
	r.Run(":8080")
}
