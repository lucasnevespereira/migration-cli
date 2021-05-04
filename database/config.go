package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Open() *sql.DB {
	db, err := sql.Open("sqlite3", "myDB")

	if err != nil {
		fmt.Printf("Error Opening DB: %v \n", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("Error Pinging DB: %v \n", err)
	}

	fmt.Println("Connected to db!")

	return db
}
