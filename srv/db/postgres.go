package db

import (
	"database/sql"
	"fmt"
	"os"

	// needed
	_ "github.com/lib/pq"
)

const (
	dbhost = "DBHOST"
	dbport = "DBPORT"
	dbuser = "DBUSER"
	dbpass = "DBPASS"
	dbname = "DBNAME"
)

var db *sql.DB

func InitDb() {
	config := dbConfig()
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config[dbhost], config[dbport],
		config[dbuser], config[dbpass], config[dbname])

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	installDbs()
}

func dbConfig() map[string]string {
	conf := make(map[string]string)
	host, ok := os.LookupEnv(dbhost)
	if !ok {
		panic("DBHOST environment variable required but not set")
	}
	port, ok := os.LookupEnv(dbport)
	if !ok {
		panic("DBPORT environment variable required but not set")
	}
	user, ok := os.LookupEnv(dbuser)
	if !ok {
		panic("DBUSER environment variable required but not set")
	}
	password, ok := os.LookupEnv(dbpass)
	if !ok {
		panic("DBPASS environment variable required but not set")
	}
	name, ok := os.LookupEnv(dbname)
	if !ok {
		panic("DBNAME environment variable required but not set")
	}
	conf[dbhost] = host
	conf[dbport] = port
	conf[dbuser] = user
	conf[dbpass] = password
	conf[dbname] = name
	return conf
}

func Close() {
	db.Close()
}

func Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	return db.Query(query, args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args...)
}

func Prepare(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	return stmt.Query(args...)
}

func PrepareRow(query string, args ...interface{}) *sql.Row {
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	return stmt.QueryRow(args...)
}

func GetDb() *sql.DB {
	if db == nil {
		InitDb()
	}
	return db
}
