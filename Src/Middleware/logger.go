package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Logg creates and returns a logger for a specific log place
func Logg(LogPlace string) (*log.Logger, *os.File) {
	logDir := "Logs"

	if err := os.MkdirAll(filepath.Join(logDir, LogPlace), 0755); err != nil {
		log.Fatalf("Could not create log directories: %v", err)
	}

	logFilePath := filepath.Join(logDir, LogPlace, fmt.Sprintf("%s.log", LogPlace))
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger, file
}

// LoggerMiddleware logs requests using the provided logger
func LoggerMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)
		logger.Printf("Request: %s %s | Status: %d | Duration: %s",
			c.Request.Method, c.Request.URL.Path, c.Writer.Status(), duration)
	}
}
