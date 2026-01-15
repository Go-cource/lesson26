// Это хорошее тестовое API для демонстрации работы Gin
// # Обзор
// Это занятие №26 на курсе ЦДО МГТУ им.Баумана. Мы учимся документировать проекты
// Запуск go run main.go
// Запуск клиента --------
package ginModule

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// indexHandler - это ручка для вывода "Hello, World!!"
// Чтобы обратиться перейдите http://127.0.0.1:8080/
//
// @Summary      Базовое приветствие
// @Description  Возвращает стандартное приветственное сообщение "Hello, World!!"
// @Tags         Общие
// @Accept       json
// @Produce      plain
// @Success      200  {object}  MessageResponse  "Успешный ответ с приветствием"
// @Router       / [get]
func IndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!!")
}

// helloUser - это ручка для вывода имени.
// Обращаться: http://localhost:8080/user/:name, имя, тут будет выведено
// @Summary      Приветствие пользователя
// @Description  Возвращает персональное приветствие по имени
// @Tags         Пользователи
// @Accept       json
// @Produce      plain
// @Param        name   path      string  true  "Имя пользователя"
// @Success      200    {string}  string  "Приветственное сообщение"
// @Router       /user/{name} [get]
func HelloUser(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Hello, %s", c.Param("name")))
}

// Credo - это структура для аутентификации.
// Поля: Email (string), Password (string)
type Credo struct {
	Email    string
	Password string
}

// authHandler - это ручка, которая отвечает за аутентификацию пользователя по паролю.
// Работает POST метод. Принимает структуру "Credo"
// В случе неудачи: Status Bad Request.
// @Summary Аутентификация пользователя
// @Description Авторизует пользователя по email и паролю, возвращает токен доступа
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body Credo true "Данные для аутентификации"
// @Success 200 {object} map[string]string "Успешная аутентификация"
// @Failure 400 {object} map[string]string "Некорректные данные"
// @Router /auth [post]
func AuthHandler(c *gin.Context) {
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
	r.GET("/", IndexHandler)
	r.GET("/user/:name", HelloUser)
	r.POST("/auth", AuthHandler)

	r.Run(":8080")
}
