package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

func NewSQLiteStorage(dbName string) (*sql.DB, error) {
	var install bool

	_, err := os.Stat(dbName)
	if err != nil {
		install = true
	}

	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		return nil, err
	}

	if install {
		if err := createTables(db); err != nil {
			return nil, err
		}
	}

	return db, nil
}
