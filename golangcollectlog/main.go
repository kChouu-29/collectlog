package main

import (
	"fmt"
	"golang-elk/router"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	// log.SetFormatter(&ecslogrus.Formatter{})
	log.SetLevel(log.TraceLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00", // Định dạng ISO8601
	})
	logFilePath := "logs/out.log"

	// Create logs directory if it doesn't exist
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Error creating logs directory:", err)
	}

	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	log.SetOutput(file)

	defer file.Close()

	fmt.Print("Start Service")

	log.Info("Start Service")
	router := router.InitRouter()

	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        router,
	}
	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}
}
