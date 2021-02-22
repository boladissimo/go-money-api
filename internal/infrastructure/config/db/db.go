package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"sync"

	"github.com/boladissimo/go-money-api/internal/infrastructure/util"
	// importing mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var once sync.Once
var db *sql.DB
var migrationsPath = "/scripts/migrations/"

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

//getDBParameterFromEnv return the database connection parameters from the env CLEARDB_DATABASE_URL
//it needs to be in the format user@password.host:port/dbname and the port needs to be made of 4 digits
func getDBParameterFromEnv() (host, user, password, dbname, port string) {
	databaseURL := os.Getenv("CLEARDB_DATABASE_URL")
	if databaseURL == "" {
		panic("invalid env CLEARDB_DATABASE_URL")
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

//RunMigrations scans the /scripts/migrations folder to check what tables does it needs migrate
func RunMigrations() {
	_, _, _, dbname, _ := getDBParameterFromEnv()
	tables := getTableNames()
	for _, table := range tables {
		if !tableExists(dbname, table) {
			util.LogInfo(fmt.Sprintf("running %s migration", table))
			db.Exec(getMigrationScript(table))
		}
	}
}

//tableExists check if the given table is present on the database
func tableExists(dbname, table string) bool {
	db := GetDB()
	result, err := db.Query("SELECT * FROM information_schema.tables WHERE table_schema = ? AND table_name = ? LIMIT 1", dbname, table)
	if err != nil {
		util.LogError(err)
	}
	return result.Next()
}

//getTableNames return all tables names insed the migrations folder
func getTableNames() (tables []string) {

	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		util.LogError(err)
	}

	for _, file := range files {
		fileFullName := file.Name()
		tableName := fileFullName[:len(fileFullName)-4]
		tables = append(tables, tableName)
		util.LogInfo(tableName)
	}

	return
}

//getMigrationScript return
func getMigrationScript(table string) string {
	script, err := ioutil.ReadFile(fmt.Sprintf("%s%s.sql", migrationsPath, table))
	if err != nil {
		util.LogError(err)
	}
	return string(script)
}
