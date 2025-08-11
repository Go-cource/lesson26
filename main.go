// Package newGin реализует демонстрационный веб-сервер на Gin.
//
// Этот проект создан для учебных целей, чтобы показать слушателям
// основы работы с фреймворком Gin.
package newGin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PORT определяет порт, на котором запускается приложение.
const PORT = ":8080"

// User описывает структуру пользователя.
type User struct {
	Name  string // Имя пользователя
	Email string // Email пользователя
}

// LoggingAllRequests возвра middleware для логирования всех входящих запросов.
//
// Middleware выводит в консоль метод и путь каждого запроса.
//
// BUG(Дмитрий Веревкин): Эта функция избыточна, так как Gin предоставляет встроенное логирование.
// BUG(Сергей): Пожалуйста, не используйте мое имя в примерах ошибок.
func LoggingAllRequests() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("New Request: ", c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

// PrintHello выводит приветствие в консоль.
func PrintHello() {
	fmt.Println("Hello!")
}

// newGin - главная функция.
func newGin() {
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

	r.Run(PORT)
}
