package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!!")
}

func helloUser(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Hello, %s", c.Param("name")))
}



func main() {
	r := gin.Default()
	r.GET("/", indexHandler)
	r.GET("/user/:name", helloUser)
	
	r.Run(":8080")
}
