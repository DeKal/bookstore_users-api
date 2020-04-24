package usersdb

import (
	"database/sql"
	"log"

	"github.com/DeKal/bookstore_users-api/src/datasources/mysql/source"
	// This includes a mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var (
	// Client is our db
	client *sql.DB
)

func openConnection() *sql.DB {
	db, err := sql.Open("mysql", source.GetDataSourceName())
	if err != nil {
		panic(err)
	}
	if err := db.Ping(); err != nil {
		panic(err)
	}
	return db
}

// GetNewClientConnection return new client connection
func GetNewClientConnection() *sql.DB {
	if client != nil {
		client.Close()
	}
	client = openConnection()
	log.Println("Database successfully connected.")
	return client
}
