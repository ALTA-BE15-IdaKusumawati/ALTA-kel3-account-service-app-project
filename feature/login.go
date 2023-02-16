package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"log"
)

func LoginUser(db *sql.DB, loginUsers entities.Users) (statement string, dataUser entities.Users) {

	var telepon entities.Users
	row := db.QueryRow("SELECT id,name,Telepon,email,saldo,password FROM users WHERE telepon = ?", loginUsers.Telepon)
	err := row.Scan(&telepon.ID, &telepon.Name, &telepon.Telepon, &telepon.Email, &telepon.Saldo, &telepon.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			statement = "login gagal"
			return statement, dataUser
		} else {
			log.Fatal(err)
		}

	}
	if telepon.Password != loginUsers.Password {
		statement = "login gagal"
	} else {
		statement = "login berhasil"
		dataUser = telepon
	}
	return statement, dataUser
}
