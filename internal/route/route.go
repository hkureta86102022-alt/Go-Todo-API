package route

import (
	"go-todo-api/internal/handler"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/todos", handler.GetTodos)
	e.POST("/todos", handler.CreateTodo)
	e.DELETE("/todos/:id", handler.DeleteTodo)
	e.PUT("/todos/:id", handler.UpdateTodo)
}