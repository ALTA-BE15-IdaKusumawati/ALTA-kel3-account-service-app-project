package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"log"
)

func LoginUser(db *sql.DB, loginUsers entities.Users) (berhasil string, dataUser entities.Users) {

	var telepon entities.Users
	row := db.QueryRow("SELECT id,name,Telepon,email,Password,saldo FROM users WHERE telepon = ?", loginUsers.Telepon)
	err := row.Scan(&telepon.ID, &telepon.Name, &telepon.Telepon, &telepon.Email, &telepon.Password, &telepon.Saldo)
	if err != nil {
		if err == sql.ErrNoRows {
			berhasil = "login gagal"
		} else {
			log.Fatal(err)
		}
	}
	if telepon.Password != loginUsers.Password {
		berhasil = "login gagal"
	} else {
		berhasil = "login berhasil"
		dataUser = telepon
	}
	return berhasil, dataUser
}
