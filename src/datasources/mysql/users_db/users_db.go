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
	Client *sql.DB
)

func init() {
	Client = openConnection()
	log.Println("Database successfully connected.")
}

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
