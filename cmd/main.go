package main

import (
	"log"

	todo "github.com/maximka200/NotesWebApp"
	handler "github.com/maximka200/NotesWebApp/pkg/handler"
	"github.com/maximka200/NotesWebApp/pkg/repository"
	"github.com/maximka200/NotesWebApp/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	/// *gin.Engine реализует интерфейс http.Handler
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8800"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err)
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
