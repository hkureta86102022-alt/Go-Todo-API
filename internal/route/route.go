package route

import (
	"go-todo-api/internal/handler"

	"go-todo-api/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *pgxpool.Pool) {
	repo := repository.NewTodoRepository(db)
	h := handler.NewHandler(repo)

	e.GET("/todos", h.GetTodos)
	e.POST("/todos", h.CreateTodo)
	e.DELETE("/todos/:id", h.DeleteTodo)
	e.PUT("/todos/:id", h.UpdateTodo)

}

func UserInitRoutes(e *echo.Echo, h *handler.UserHandler) {

	e.POST("/signup", h.Signup)
}
