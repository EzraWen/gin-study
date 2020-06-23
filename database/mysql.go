package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Database *sql.DB

func init() {

	var err error
	Database, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/shop?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = Database.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}
