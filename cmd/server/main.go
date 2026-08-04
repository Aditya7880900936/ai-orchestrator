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
