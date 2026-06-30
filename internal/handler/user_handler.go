package handler

import (
	"go-todo-api/internal/model"
	"go-todo-api/internal/repository"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) Signup(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return c.JSON(400, "invalid request")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(500, "hash error")
	}

	user.Password = string(hashedPassword)

	if err := h.Repo.CreateUser(c.Request().Context(), user); err != nil {
		return c.JSON(500, err.Error())
	}

	return c.JSON(201, "user created")
}
