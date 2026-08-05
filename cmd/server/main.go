// @title AI Orchestrator API
// @version 1.0
// @description Production-grade AI Orchestration Backend with Resume Analysis, ATS Scoring, Job Matching, Resume Chat, Cover Letter Generation and AI Workflows.
// @termsOfService https://github.com/Aditya7880900936/ai-orchestrator
//
// @contact.name Aditya Sanskar Srivastav
// @contact.url https://github.com/Aditya7880900936
// @contact.email adityasanskarsrivastav788@gmail.com
//
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
//
// @host localhost:8080
// @BasePath /
package main

import "log"

var (
	infoLog = func(v ...any) {
		log.Println(v...)
	}

	fatalLog = func(v ...any) {
		log.Fatal(v...)
	}

	runServer = func() error {
		return setupRouter().Run(":8080")
	}
)

func main() {

	if err := initialize(); err != nil {
		fatalLog(err)
		return
	}

	infoLog("Server running on :8080")
	infoLog("Swagger Docs: http://localhost:8080/swagger/index.html")

	if err := runServer(); err != nil {
		fatalLog(err)
	}
}
