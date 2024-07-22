package main

import (
	"log"

	todo "github.com/maximka200/NotesWebApp"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8800"); err != nil {
		log.Fatalf("Error occured while running http server: %s", err)
	}
}
