package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/khilmi-aminudin/golang_restful_api/helper"
)

const (
	driver   string = "mysql"
	user     string = "root"
	password string = "1234q5678"
	host     string = "localhost"
	port     int    = 3306
	dbname   string = "belajar_golang_restful_api"
)

var stringConnection string = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbname)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", stringConnection)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
