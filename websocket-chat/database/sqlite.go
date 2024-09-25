package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/chat.db")

	if err != nil {
		return nil, err
	}

	DB = db

	sqlFile, err := os.ReadFile("./database/init.sql")
	if err != nil {
		return nil, fmt.Errorf("error reading 'init.sql' file : %v", err)
	}

	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return nil, fmt.Errorf("error executing 'init.sql' code : %v", err)
	}

	log.Println("Database successfully initialized.")

	return db, nil
}
