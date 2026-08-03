package main

import (
	"log"

	_ "github.com/Aditya7880900936/ai-orchestrator/docs"

	handler "github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"
	middle "github.com/Aditya7880900936/ai-orchestrator/internal/api/middleware"
	"github.com/Aditya7880900936/ai-orchestrator/internal/api/routes"
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	"github.com/Aditya7880900936/ai-orchestrator/internal/database"
	llm "github.com/Aditya7880900936/ai-orchestrator/internal/llm"
	"github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/Aditya7880900936/ai-orchestrator/internal/metrics"
	"github.com/Aditya7880900936/ai-orchestrator/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title AI Orchestrator API
// @version 1.0
// @description Production-grade AI Orchestration Backend with Resume Analysis, ATS Scoring, Job Matching, Resume Chat, Cover Letter Generation and AI Workflows.
// @host localhost:8080
// @BasePath /
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	cache.InitRedis()
	database.Init()
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
		middle.Logger(),
		middle.RateLimiter(),
	)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

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

	log.Println("Server running on :8080")
	log.Println("Swagger Docs: http://localhost:8080/swagger/index.html")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
