package main

import (
	"log"

	todo "github.com/maximka200/NotesWebApp"
	handler "github.com/maximka200/NotesWebApp/pkg/handler"
)

func main() {
	/// *gin.Engine реализует интерфейс http.Handler
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8800", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err)
	}
}
