package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

func InsertUser(db *sql.DB, newUsers entities.Users) {
	var query = "INSERT INTO users(id, name, telepon, email, password) VALUES(?,?,?,?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error prepare insert", errPrepare.Error())
	}

	result, errInsert := statement.Exec(newUsers.ID, newUsers.Name, newUsers.Telepon, newUsers.Email, newUsers.Password)
	if errInsert != nil {
		log.Fatal("error exec insert", errInsert.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("--------------------------")
			fmt.Println("proses berhasil dijalankan")
			fmt.Println("--------------------------")
		} else {
			fmt.Println("------------")
			fmt.Println("proses gagal")
			fmt.Println("------------")
		}
	}
}
