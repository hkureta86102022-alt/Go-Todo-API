package handler

import (
	"fmt"
	"go-todo-api/internal/model"
	"go-todo-api/internal/repository"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetTodos(c echo.Context) error {
	todos, err := h.repo.GetTodos(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to fetch todos")
	}
	return c.JSON(http.StatusOK, todos)
}

func (h *Handler) CreateTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request JSON")
	}
	err := h.repo.CreateTodo(c.Request().Context(), todo.Title, todo.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to create todo")
	}

	return c.JSON(http.StatusCreated, todo)
}
func (h *Handler) DeleteTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	err = h.repo.DeleteTodo(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to delete todo")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("ID %d: Todo deleted successfully", id),
	})
}
func (h *Handler) UpdateTodo(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request JSON")
	}

	err = h.repo.UpdateTodo(c.Request().Context(), id, todo.Title, todo.Completed)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Failed to update todo")
	}

	return c.JSON(http.StatusOK, map[string]any{
		"id":        id,
		"title":     todo.Title,
		"completed": todo.Completed,
	})
}

type Handler struct {
	repo *repository.TodoRepository
}

func NewHandler(repo *repository.TodoRepository) *Handler {
	return &Handler{repo: repo}
}
