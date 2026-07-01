package route

import (
	"go-todo-api/internal/handler"
	"go-todo-api/internal/middleware"
	"go-todo-api/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *pgxpool.Pool, userHandler *handler.UserHandler) {

	api := e.Group("/api")

	api.POST("/signup", userHandler.Signup)
	api.POST("/login", userHandler.Login)

	auth := api.Group("")
	auth.Use(middleware.JWTMiddleware())

	auth.GET("/me", userHandler.Me)

	repo := repository.NewTodoRepository(db)
	h := handler.NewHandler(repo)

	auth.GET("/todos", h.GetTodos)
	auth.POST("/todos", h.CreateTodo)
	auth.DELETE("/todos/:id", h.DeleteTodo)
	auth.PUT("/todos/:id", h.UpdateTodo)
}
