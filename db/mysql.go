package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysql() *sql.DB {
	// Initialize the database
	db, err := sql.Open("mysql", getConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	// Check database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
}
