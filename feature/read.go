package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

func GetUsers(db *sql.DB, ID string) {

	rows, errSelect := db.Query("SELECT id, name, telepon, email, saldo, password FROM users where id=?", ID)
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var Users []entities.Users
	for rows.Next() {
		var datarow entities.Users
		errScan := rows.Scan(&datarow.ID, &datarow.Name, &datarow.Telepon, &datarow.Email, &datarow.Saldo, &datarow.Password)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		Users = append(Users, datarow)

	}

	for _, v := range Users {
		fmt.Println("ID 	 : ", v.ID)
		fmt.Println("Nama	 : ", v.Name)
		fmt.Println("Telepon  : ", v.Telepon)
		fmt.Println("Email 	 : ", v.Email)
		fmt.Println("Saldo 	 : ", v.Saldo)
	}
}
