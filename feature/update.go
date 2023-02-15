package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

func UpdateUser(db *sql.DB, ID string, updatedataUser entities.Users) {

	var query = "UPDATE users SET name=?, telepon=?, email=? WHERE id=?"
	statment, errupdate := db.Prepare(query)
	if errupdate != nil {
		log.Fatal("error update", errupdate)
	}

	result, errupdate := statment.Exec(updatedataUser.Name, updatedataUser.Telepon, updatedataUser.Email, ID)
	if errupdate != nil {
		log.Fatal("error exec insert", errupdate.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("proses berhasil dijalankan")

		} else {
			fmt.Println("proses gagal")
		}
	}
	fmt.Println(updatedataUser)
}
