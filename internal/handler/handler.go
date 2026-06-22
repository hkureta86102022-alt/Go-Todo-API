package handler

import (
	"fmt"
	"go-todo-api/internal/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetTodos(c echo.Context) error {
	Todos := []model.Todo{
		{ID: 1, Title: "Todo 1", Completed: false},
		{ID: 2, Title: "Todo 2", Completed: true},
	}
	return c.JSON(http.StatusOK, Todos)
}

func CreateTodo(c echo.Context) error {
	var todo model.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid request JSON")
	}
	return c.JSON(http.StatusCreated, todo)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, map[string]string{
		"message": fmt.Sprintf("ID %s: Todo deleted successfully", id),
	})

}
