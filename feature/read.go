package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

func GetallUsers(db *sql.DB) {

	rows, errSelect := db.Query("SELECT id, name, telepon, email, saldo, password FROM users")
	if errSelect != nil {
		log.Fatal("error query select", errSelect.Error())
	}

	var allUsers []entities.Users
	for rows.Next() {
		var datarow entities.Users
		errScan := rows.Scan(&datarow.ID, &datarow.Name, &datarow.Telepon, &datarow.Email, &datarow.Saldo, &datarow.Password)
		if errScan != nil {
			log.Fatal("error scan select", errScan.Error())
		}
		allUsers = append(allUsers, datarow)

	}

	for _, v := range allUsers {
		fmt.Print("ID : ", v.ID, "\n", "Nama : ", v.Name, "\n", "Telepon : ", v.Telepon, "\n", "Email : ", v.Email, "\n", "Saldo : ", v.Saldo)
	}
}
