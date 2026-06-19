package main

import (
	"go-todo-api/internal/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	route.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
