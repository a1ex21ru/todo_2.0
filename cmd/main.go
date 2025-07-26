package main

import (
	"log"

	todo "github.com/alex21ru/todo_2.0"
	"github.com/alex21ru/todo_2.0/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http")
	}
}
