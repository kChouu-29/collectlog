package main

import (
	"fmt"
	"golang-elk/router"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers" // <-- THÊM IMPORT NÀY
	log "github.com/sirupsen/logrus"

)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.000Z07:00", // Định dạng text
	})
	log.SetLevel(log.TraceLevel)

	logFilePath := "logs/out.log"

	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Fatal("Error creating logs directory:", err)
	}

	// Mở file log để cả Logrus và Gorilla Handlers cùng ghi vào
	file, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening log file:", err)
	}
	log.SetOutput(file) // Logrus sẽ ghi vào file này

	defer file.Close()

	fmt.Print("Start Service")

	log.Info("Start Service") // Log ứng dụng dạng text
	router := router.InitRouter()

	// TẠO MIDDLEWARE BỌC ROUTER CỦA BẠN
	// Nó sẽ ghi log truy cập (định dạng Apache) vào "file" (logs/out.log)
	loggedRouter := handlers.CombinedLoggingHandler(file, router)

	server := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		// Handler:        router, // <-- THAY DÒNG NÀY
		Handler: loggedRouter, // <-- BẰNG DÒNG NÀY
	}
	server_err := server.ListenAndServe()
	if server_err != nil {
		panic(server_err)
	}
}
