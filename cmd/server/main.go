package main

import (
	"log"

	handler "github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	llm "github.com/Aditya7880900936/ai-orchestrator/internal/llm"
	"github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/Aditya7880900936/ai-orchestrator/internal/middleware"
	"github.com/Aditya7880900936/ai-orchestrator/internal/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	log.Println("Server running on :8080")
	r.Run(":8080")
}
