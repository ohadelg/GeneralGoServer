package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ohadelg/GeneralGoServer/Middleware"
)

func main() {
	// middlewareMain := Middleware.MiddlewareMain
	router := gin.Default()
	router.Use(Middleware.LoggerMiddleware)
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8080")
}
