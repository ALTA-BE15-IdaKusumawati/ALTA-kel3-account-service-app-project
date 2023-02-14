package main

import (
	"account-service-app-project/config"
	"account-service-app-project/entities"
	"account-service-app-project/feature"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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

	var input int
	fmt.Println("selamat datang")
	fmt.Println("silahkan pilih menu :")
	fmt.Println("menu 1 untuk mendaftar")
	fmt.Println("menu 2 untuk login jika sudah mempunyai akun")
	fmt.Println("menu 0 untuk membatalkan transaksi")
	fmt.Printf("masukan menu : ")
	fmt.Scanln(&input)

	switch input {

	case 1:

		newUsers := entities.Users{}
		fmt.Println("Masukkan ID User:")
		fmt.Scanln(&newUsers.ID)
		fmt.Println("Masukkan Name User:")
		fmt.Scanln(&newUsers.Name)
		fmt.Println("Masukkan Telepon User:")
		fmt.Scanln(&newUsers.Telepon)
		fmt.Println("Masukkan Email User:")
		fmt.Scanln(&newUsers.Email)
		fmt.Println("Masukkan password User:")
		fmt.Scanln(&newUsers.Password)

		feature.InsertUser(db, newUsers)

	case 0:
		fmt.Println("terima kasih telah bertransaksi")
		return
	default:
		fmt.Println("menu yang anda masukan tidak ada")
	}

}
