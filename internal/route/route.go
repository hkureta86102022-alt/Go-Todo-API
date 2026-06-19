package route

import (
	"go-todo-api/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handler.GetTodos)
}
