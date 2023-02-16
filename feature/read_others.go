package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

// func OtherUser(db *sql.DB, id string, profile entities.Users ) {
func OtherUser(db *sql.DB, id string, profile entities.Users) {
	ro, err := db.Query("SELECT telepon, name FROM users where telepon=?", profile.Telepon)
	if err != nil {
		log.Fatal("error select", err.Error())
	}

	var User []entities.Users
	for ro.Next() {
		var datarow entities.Users
		errScan := ro.Scan(&datarow.Telepon, &datarow.Name)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		User = append(User, datarow)

	}
	if len(User) > 0 {
		for _, v := range User {
			// fmt.Println(v.Name, "    |", v.Telepon)
			fmt.Println("---------------")
			fmt.Println("Data Pengguna |")
			fmt.Println("-------------------------")
			fmt.Println("Nama          :", v.Name)
			fmt.Println("Nomor Telepon :", v.Telepon)
		}
	} else {
		fmt.Println("-------------------------")
		fmt.Println("   Data tidak tersedia")
		fmt.Println("-------------------------")
	}
}
