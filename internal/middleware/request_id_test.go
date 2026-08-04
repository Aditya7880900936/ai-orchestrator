package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRequestIDMiddleware(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.Use(RequestID())

	r.GET("/", func(c *gin.Context) {

		id, exists := c.Get("request_id")

		if !exists {
			t.Fatal("request_id not found in context")
		}

		requestID, ok := id.(string)
		if !ok {
			t.Fatal("request_id is not string")
		}

		if requestID == "" {
			t.Fatal("request_id is empty")
		}

		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", w.Code)
	}

	header := w.Header().Get("X-Request-ID")

	if header == "" {
		t.Fatal("X-Request-ID header missing")
	}
}

func TestRequestIDMiddleware_UniqueIDs(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.Use(RequestID())

	r.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req1 := httptest.NewRequest(http.MethodGet, "/", nil)
	res1 := httptest.NewRecorder()

	req2 := httptest.NewRequest(http.MethodGet, "/", nil)
	res2 := httptest.NewRecorder()

	r.ServeHTTP(res1, req1)
	r.ServeHTTP(res2, req2)

	id1 := res1.Header().Get("X-Request-ID")
	id2 := res2.Header().Get("X-Request-ID")

	if id1 == "" || id2 == "" {
		t.Fatal("missing request ids")
	}

	if id1 == id2 {
		t.Fatal("expected unique request ids")
	}
}
