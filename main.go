package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from Gin Gonic Go!")
	})
	r.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello from user!")
	})

	r.Run(":8080")
}
