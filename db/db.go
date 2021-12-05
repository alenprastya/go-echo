package db

import (
	"database/sql"

	"github.com/alen/echo-framework/config"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfirations()
	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAMER

	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		panic("Connection String error")
	}
	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateConf() *sql.DB {
	return db
}
