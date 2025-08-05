package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	logrus "github.com/sirupsen/logrus"

	todo "github.com/alex21ru/todo_2.0"
	"github.com/alex21ru/todo_2.0/pkg/handler"
	"github.com/alex21ru/todo_2.0/pkg/repository"
	"github.com/alex21ru/todo_2.0/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

// @title Todo App API
// @version 1.0
// @description API Server for todolist app

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error init config: %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Port:     viper.GetString("db.port"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed connect to db, error: %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http")
		}
	}()

	logrus.Print("Todoapp started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Todoapp started")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error server shutting: %s", err.Error())
	}

	if err = db.Close(); err != nil {
		logrus.Errorf("error db connect close")
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
