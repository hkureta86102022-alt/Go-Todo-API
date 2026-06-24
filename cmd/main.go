package main

import (
	"context"
	"go-todo-api/internal/database"
	"go-todo-api/internal/route"
	"log"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	conn, err := database.ConnectDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(ctx)

	log.Println("Connected to the database successfully")

	e := echo.New()
	route.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
