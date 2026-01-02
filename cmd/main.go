package main

import (
	"log"

	"github.com/kasariks/notes_api/cmd/api"
	"github.com/kasariks/notes_api/config"
)

func main() {
	server := api.NewAPIServer(":"+config.Envs.Port, nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
