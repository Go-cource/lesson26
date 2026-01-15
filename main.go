package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!!")
}

func main() {
	r := gin.Default()
	r.GET("/", indexHandler)
	r.Run(":8080")
}
