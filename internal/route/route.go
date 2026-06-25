package route

import (
	"go-todo-api/internal/handler"

	"github.com/jackc/pgx/v5"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, conn *pgx.Conn) {
	h := handler.NewHandler(conn)
	e.GET("/todos", h.GetTodos)
	e.POST("/todos", h.CreateTodo)
	e.DELETE("/todos/:id", h.DeleteTodo)
	e.PUT("/todos/:id", h.UpdateTodo)
}
