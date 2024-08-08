package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CorsPolicy(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, hx-target")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusOK)
		return
	}

	c.Next()
}
