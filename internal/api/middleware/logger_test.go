package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	projectLogger "github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/gin-gonic/gin"
)

func TestLoggerMiddleware(t *testing.T) {

	gin.SetMode(gin.TestMode)

	if err := projectLogger.Init(); err != nil {
		t.Fatal(err)
	}

	r := gin.New()

	r.Use(Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", w.Code)
	}
}

func TestLoggerMiddleware_NextCalled(t *testing.T) {

	gin.SetMode(gin.TestMode)

	if err := projectLogger.Init(); err != nil {
		t.Fatal(err)
	}

	r := gin.New()

	called := false

	r.Use(Logger())

	r.GET("/", func(c *gin.Context) {
		called = true
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !called {
		t.Fatal("handler not executed")
	}
}
