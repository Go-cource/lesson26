package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoggingAllRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("New Request: ", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(LoggingAllRequests())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Name": "Sergey",
			"Age":  66,
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})

	r.POST("/user", func(c *gin.Context) {
		type User struct {
			Name  string
			Email string
		}
		var user User

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"Status": "Ok!",
		})
	})

	r.PUT("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Changed, %s!", name)
	})
	r.DELETE("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Deleted, %s!", name)
	})

	r.Run(":8080")
}
