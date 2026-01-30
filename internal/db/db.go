package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(dsn string) *sql.DB {
	/*
		Function that recieves a database connection string and tries to connect with the database
		Returns a database connection
	*/
	// use the dsn to prepate the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// ping to verify that the connection works
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
