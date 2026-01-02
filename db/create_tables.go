package db

import "database/sql"

const userTable = "CREATE TABLE IF NOT EXISTS users (" +
	"id INTEGER PRIMARY KEY AUTOINCREMENT," +
	"nickname VARCHAR(255) NOT NULL," +
	"email VARCHAR(255) NOT NULL UNIQUE," +
	"password VARCHAR(255) NOT NULL" +
	");"

const notesTable = "CREATE TABLE IF NOT EXISTS notes (" +
	"id INTEGER PRIMARY KEY AUTOINCREMENT," +
	"userId INT NOT NULL," +
	"name VARCHAR(255) NOT NULL," +
	"value TEXT NOT NULL," +
	"createdAt TEXT NOT NULL," +
	"FOREIGN KEY (userId) REFERENCES users(id)" +
	");"

func createTables(db *sql.DB) error {
	if _, err := db.Exec(userTable); err != nil {
		return err
	}

	if _, err := db.Exec(notesTable); err != nil {
		return err
	}

	return nil
}
