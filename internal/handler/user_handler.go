package handler

import (
	"go-todo-api/internal/auth"
	"go-todo-api/internal/model"
	"go-todo-api/internal/repository"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
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

func (h *UserHandler) Login(c echo.Context) error {
	var req model.User

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "invalid request")
	}

	user, err := h.Repo.GetUserByEmail(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(401, "invalid Email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(401, "invalid Email or password")
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		return c.JSON(500, "failed to generate token")
	}

	return c.JSON(200, map[string]string{"token": token})

}

func (h *UserHandler) Me(c echo.Context) error {
	userIDFloat, ok := c.Get("user_id").(float64)
	if !ok {
		return c.JSON(401, "invalid user id")
	}

	user, err := h.Repo.GetUserByID(c.Request().Context(), int(userIDFloat))

	if err != nil {
		return c.JSON(500, "failed to get user")
	}

	return c.JSON(200, user)
}
