package routes

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func hasRoute(routes gin.RoutesInfo, method, path string) bool {
	for _, r := range routes {
		if r.Method == method && r.Path == path {
			return true
		}
	}
	return false
}

func TestRegisterATSRoutes(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	RegisterATSRoutes(r)

	if !hasRoute(r.Routes(), "POST", "/ats/score") {
		t.Fatal("ATS route not registered")
	}
}

func TestRegisterCoverLetterRoutes(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	RegisterCoverLetterRoutes(r)

	if !hasRoute(r.Routes(), "POST", "/cover-letter/generate") {
		t.Fatal("cover letter route not registered")
	}
}

func TestRegisterJobMatchRoutes(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	RegisterJobMatchRoutes(r)

	if !hasRoute(r.Routes(), "POST", "/job/match") {
		t.Fatal("job match route not registered")
	}
}

func TestRegisterResumeRoutes(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	RegisterResumeRoutes(r)

	routes := r.Routes()

	if !hasRoute(routes, "POST", "/resume/analyze") {
		t.Fatal("resume analyze route missing")
	}

	if !hasRoute(routes, "POST", "/resume/upload") {
		t.Fatal("resume upload route missing")
	}
}

func TestRegisterResumeChatRoutes(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	RegisterResumeChatRoutes(r)

	if !hasRoute(r.Routes(), "POST", "/resume/chat") {
		t.Fatal("resume chat route not registered")
	}
}

func TestRegisterResumeImproveRoutes(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	RegisterResumeImproveRoutes(r)

	if !hasRoute(r.Routes(), "POST", "/resume/improve") {
		t.Fatal("resume improve route not registered")
	}
}

func TestRegisterSkillRoutes(t *testing.T) {

	gin.SetMode(gin.TestMode)

	r := gin.New()

	RegisterSkillRoutes(r)

	if !hasRoute(r.Routes(), "POST", "/skills/extract") {
		t.Fatal("skill extraction route not registered")
	}
}
