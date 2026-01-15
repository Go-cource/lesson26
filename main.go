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

type Credo struct {
	Email    string
	Password string
}

func authHandler(c *gin.Context) {
	var credo Credo
	err := c.ShouldBindJSON(&credo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"authError": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access": "allowed",
		"token":  "wdjnWubdnkaIIw123",
	})

}

func main() {
	r := gin.Default()
	r.GET("/", indexHandler)
	r.GET("/user/:name", helloUser)
	r.POST("/auth", authHandler)

	r.Run(":8080")
}
