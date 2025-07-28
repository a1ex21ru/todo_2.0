package main

import (
	"log"

	todo "github.com/alex21ru/todo_2.0"
	"github.com/alex21ru/todo_2.0/pkg/handler"
	"github.com/alex21ru/todo_2.0/pkg/repository"
	"github.com/alex21ru/todo_2.0/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	if err := InitConfig(); err != nil {
		log.Fatalf("error init config: %s", err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http")
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
