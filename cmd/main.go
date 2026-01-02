package main

import (
	"log"

	"github.com/kasariks/notes_api/cmd/api"
	"github.com/kasariks/notes_api/config"
	"github.com/kasariks/notes_api/db"
)

func main() {
	db, err := db.NewSQLiteStorage(config.Envs.DBName)
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":"+config.Envs.Port, db)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
