package source

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	dataSourceFormat   = "%s:%s@tcp(%s)/%s?charset=utf8"
	mysqlUsersUsername = "mysql_users_username"
	mysqlUsersPass     = "mysql_users_pass"
	mysqlUsersHost     = "mysql_users_host"
	mysqlUsersDb       = "mysql_users_db"
)

var (
	user   string
	pass   string
	host   string
	dbName string
)

func init() {
	godotenv.Load()
	user = os.Getenv(mysqlUsersUsername)
	pass = os.Getenv(mysqlUsersPass)
	host = os.Getenv(mysqlUsersHost)
	dbName = os.Getenv(mysqlUsersDb)
}

// GetDataSourceName Return datasource name
func GetDataSourceName() string {
	return fmt.Sprintf(dataSourceFormat, user, pass, host, dbName)
}
