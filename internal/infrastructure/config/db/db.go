package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"sync"

	"github.com/boladissimo/go-money-api/internal/infrastructure/util"
	// importing mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once
var db *sql.DB

//GetDB return the sql.DB instance as a singleton
func GetDB() *sql.DB {
	if db == nil {
		once.Do(func() {
			host, user, password, dbname, port := getDBParameterFromEnv()
			dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
			util.LogInfo(dbConnection)
			sqlDB, err := sql.Open("mysql", dbConnection)
			if err != nil {
				util.LogError(err)
			}
			db = sqlDB
		})
	}
	return db
}

//getDBParameterFromEnv return the database connection parameters from the env DATABASE_URL
//it needs to be in the format user@password.host:port/dbname and the port needs to be made of 4 digits
func getDBParameterFromEnv() (host, user, password, dbname, port string) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "postgres://money:pass123@localhost:3306/money"
		// panic("invalid env DATABASE_URL")
	}

	url, err := url.Parse(databaseURL)
	if err != nil {
		util.LogError("Error extracting data from database url")
		panic(err)
	}

	host = url.Host[0 : len(url.Host)-5]
	user = url.User.Username()
	password, _ = url.User.Password()
	dbname = url.Path[1:len(url.Path)]
	port = url.Host[len(url.Host)-4 : len(url.Host)]

	return
}
