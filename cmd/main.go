package main

import (
	"context"

	"go-todo-api/internal/database"
	"go-todo-api/internal/handler"
	"go-todo-api/internal/repository"
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
	db, err := database.ConnectDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	log.Println("Connected to the database successfully")

	e := echo.New()
	userHandler := handler.NewUserHandler(repository.NewUserRepository(db))
	route.InitRoutes(e, db, userHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
