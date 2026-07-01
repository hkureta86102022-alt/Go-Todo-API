package middleware

import (
	"go-todo-api/internal/auth"
	"strings"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(401, "missing token")
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			claims, err := auth.ValidateJWT(tokenStr)
			if err != nil {
				return c.JSON(401, "invalid token")
			}

			// contextにuser_id入れる
			c.Set("user_id", claims["user_id"])

			return next(c)
		}
	}
}
