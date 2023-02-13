package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	var connectionString = os.Getenv("DB_CONNECTION")
	// var connectionString = "root:sgkp8ghd%@tcp(127.0.0.1:3306)/db_project1"
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		//return nil, err
		log.Fatal("error open connection", err.Error())
	}
	// See "Important settings" section.
	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db", errPing.Error())
	} else {
		fmt.Println("berhasil")
	}
	return db
}
