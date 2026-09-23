package main

import (
	"log"
	"mini-project/config"
	"mini-project/handler"
	"mini-project/pkg/db"
	"mini-project/storage"

	"github.com/gin-gonic/gin"
	"github.com/pressly/goose/v3"
)

func main() {
	cfg := config.Load()

	db, err := db.ConnectDB(*cfg)
	if err != nil {
		log.Fatalf("failed to connect db: %v\n", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("failed to goose set: %v\n", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal(err)
	}

	storage := storage.NewStorage(db)

	handler := handler.NewHandler(storage)

	route := gin.Default()

	route.POST("/users", handler.CreateUser)
	route.POST("/tasks", handler.CreateTask)

	if err := route.Run(":8180"); err != nil {
		log.Fatalf("failed to run port 8180: %v\n", err)
	}
}
