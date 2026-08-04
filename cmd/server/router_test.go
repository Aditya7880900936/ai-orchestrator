package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	projectLogger "github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	if err := projectLogger.Init(); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func hasRoute(routes gin.RoutesInfo, method, path string) bool {
	for _, r := range routes {
		if r.Method == method && r.Path == path {
			return true
		}
	}
	return false
}

func TestSetupRouter_NotNil(t *testing.T) {

	r := setupRouter()

	if r == nil {
		t.Fatal("router is nil")
	}
}

func TestHealthRoute(t *testing.T) {

	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", w.Code)
	}
}

func TestMetricsRoute(t *testing.T) {

	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", w.Code)
	}
}

func TestSwaggerRouteRegistered(t *testing.T) {

	r := setupRouter()

	if !hasRoute(r.Routes(), "GET", "/swagger/*any") {
		t.Fatal("swagger route missing")
	}
}

func TestAnalyzeRouteRegistered(t *testing.T) {

	r := setupRouter()

	if !hasRoute(r.Routes(), "POST", "/analyze") {
		t.Fatal("analyze route missing")
	}
}

func TestResumeRoutesRegistered(t *testing.T) {

	r := setupRouter()

	routes := r.Routes()

	checks := []struct {
		Method string
		Path   string
	}{
		{"POST", "/resume/analyze"},
		{"POST", "/resume/upload"},
		{"POST", "/resume/chat"},
		{"POST", "/resume/improve"},
		{"POST", "/skills/extract"},
		{"POST", "/ats/score"},
		{"POST", "/job/match"},
		{"POST", "/cover-letter/generate"},
	}

	for _, c := range checks {
		if !hasRoute(routes, c.Method, c.Path) {
			t.Fatalf("missing route %s %s", c.Method, c.Path)
		}
	}
}

func TestMiddlewaresAttached(t *testing.T) {

	r := setupRouter()

	if len(r.Handlers) == 0 {
		t.Fatal("expected middlewares to be attached")
	}
}
