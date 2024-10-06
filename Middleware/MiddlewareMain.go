package Middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	// start timer
	t := time.Now()

	// Process request
	c.Next()

	// calculate latency
	latency := time.Since(t)

	// Access the status we are sending
	status := c.Writer.Status()

	path := c.Request.URL.Path
	clientIP := c.ClientIP()
	method := c.Request.Method
	formatedTime := t.Format("02-01-2006 15:04:05")
	log.Printf("Middleware | Request | time start: %v | latency: %v | status: %v | path: %v | clientIP: %v | method:  %v | \n", formatedTime, latency, status, path, clientIP, method)
}
