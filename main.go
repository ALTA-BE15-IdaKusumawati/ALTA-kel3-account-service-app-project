package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var connectionString = os.Getenv("DB_CONNECTION")
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
}
