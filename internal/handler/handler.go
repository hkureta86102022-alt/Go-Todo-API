package handler

import (
	"fmt"
	"go-todo-api/internal/model"
	"go-todo-api/internal/repository"

	"github.com/jackc/pgx/v5"

	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTodos(c echo.Context) error {
	Todos := []model.Todo{
		{ID: 1, Title: "Todo 1", Completed: false},
		{ID: 2, Title: "Todo 2", Completed: true},
	}
	return c.JSON(http.StatusOK, Todos)
}

func (h *Handler) CreateTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request JSON")
	}
	err := repository.CreateTodo(c.Request().Context(), h.conn, todo.Title, todo.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create todo")
	}

	return c.JSON(http.StatusCreated, todo)
}

func (h *Handler) DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("ID %s: Todo deleted successfully", id),
	})

}

func (h *Handler) UpdateTodo(c echo.Context) error {

	id := c.Param("id")
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request JSON")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"id":        id,
		"title":     todo.Title,
		"completed": todo.Completed,
	})
}

type Handler struct {
	conn *pgx.Conn
}

func NewHandler(conn *pgx.Conn) *Handler {
	return &Handler{conn: conn}
}
