package main

import (
	"account-service-app-project/config"
<<<<<<< HEAD
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type users struct {
	ID       int
	name     string
	telepon  int
	email    string
	password string
}

=======
	_ "github.com/go-sql-driver/mysql"
)

>>>>>>> e1da229391c00ccfdb813282f589dbde480b596a
func main() {
	// var connectionString = os.Getenv("DB_CONNECTION")
	// // var connectionString = "root:sgkp8ghd%@tcp(127.0.0.1:3306)/db_project1"
	// db, err := sql.Open("mysql", connectionString)
	// if err != nil {
	// 	//return nil, err
	// 	log.Fatal("error open connection", err.Error())
	// }
	// // See "Important settings" section.
	// db.SetConnMaxIdleTime(time.Minute * 3)
	// db.SetMaxOpenConns(10)
	// db.SetMaxIdleConns(10)

	// errPing := db.Ping()
	// if errPing != nil {
	// 	log.Fatal("error connect to db", errPing.Error())
	// } else {
	// 	fmt.Println("berhasil")
	// }
	db := config.ConnectToDB()
	defer db.Close()

<<<<<<< HEAD
	var input int
	fmt.Println("selamat datang")
	fmt.Println("silahkan pilih menu :")
	fmt.Println("menu 1 untuk mendaftar")
	fmt.Println("menu 2 untuk login jika sudah mempunyai akun")
	fmt.Printf("masukan menu : ")
	fmt.Scan(&input)

	switch input {

	case 1:
		{

		}
	}

=======
>>>>>>> e1da229391c00ccfdb813282f589dbde480b596a
}
