package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetLimiter(t *testing.T) {

	clients = make(map[string]*client)

	l1 := getLimiter("127.0.0.1")
	l2 := getLimiter("127.0.0.1")

	if l1 == nil {
		t.Fatal("limiter is nil")
	}

	if l1 != l2 {
		t.Fatal("expected same limiter instance")
	}
}

func TestGetLimiter_NewClient(t *testing.T) {

	clients = make(map[string]*client)

	l1 := getLimiter("127.0.0.1")
	l2 := getLimiter("127.0.0.2")

	if l1 == nil || l2 == nil {
		t.Fatal("nil limiter")
	}

	if l1 == l2 {
		t.Fatal("expected different limiters")
	}
}

func TestRateLimiter_Allowed(t *testing.T) {

	gin.SetMode(gin.TestMode)

	clients = make(map[string]*client)

	r := gin.New()

	r.Use(RateLimiter())

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "127.0.0.1:12345"

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", w.Code)
	}
}

func TestRateLimiter_Exceeded(t *testing.T) {

	gin.SetMode(gin.TestMode)

	clients = make(map[string]*client)

	r := gin.New()

	r.Use(RateLimiter())

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Consume initial burst (5 requests)
	for i := 0; i < 5; i++ {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.RemoteAddr = "127.0.0.1:12345"

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
	}

	// 6th request should be blocked
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.RemoteAddr = "127.0.0.1:12345"

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusTooManyRequests {
		t.Fatalf("expected 429 got %d", w.Code)
	}
}

func TestRateLimiter_DifferentIPs(t *testing.T) {

	gin.SetMode(gin.TestMode)

	clients = make(map[string]*client)

	r := gin.New()

	r.Use(RateLimiter())

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	req1.RemoteAddr = "127.0.0.1:12345"

	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	req2.RemoteAddr = "127.0.0.2:12345"

	w1 := httptest.NewRecorder()
	w2 := httptest.NewRecorder()

	r.ServeHTTP(w1, req1)
	r.ServeHTTP(w2, req2)

	if w1.Code != http.StatusOK {
		t.Fatal("first request failed")
	}

	if w2.Code != http.StatusOK {
		t.Fatal("second request failed")
	}
}
