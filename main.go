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
func IndexHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!!")
}

// helloUser - это ручка для вывода имени.
// Обращаться: http://localhost:8080/user/:name, имя, тут будет выведено
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
