package handler

import (
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
