package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Пример API на Gin с Swagger
// @version 1.0
// @description Это пример API, созданный с использованием Gin и Swagger-аннотаций.
// @host localhost:8080
// @BasePath /

// Структура для ответа
type MessageResponse struct {
	Message string
}

// @Summary Получить приветственное сообщение
// @Description Возвращает простое приветствие в формате JSON
// @Tags пример
// @Produce json
// @Success 200 {object} MessageResponse
// @Router /hello [get]
func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, MessageResponse{Message: "Привет, мир!"})
}

// @Summary Получить список пользователей
// @Description Возвращает список пользователей
// @Tags пример
// @Produce json
// @Success 200 {array} string
// @Router /users [get]
func usersHandler(c *gin.Context) {
	users := []string{"Алиса", "Боб", "Чарли"}
	c.JSON(http.StatusOK, users)
}

func main() {
	r := gin.Default()

	r.GET("/hello", helloHandler)
	r.GET("/users", usersHandler)

	r.Run() // по умолчанию на :8080
}
