package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"log"
)

func LoginUser(db *sql.DB, loginUsers entities.Users) (berhasil string) {

	var telepon entities.Users
	row := db.QueryRow("SELECT Telepon,Password FROM users WHERE telepon = ?", loginUsers.Telepon)
	err := row.Scan(&telepon.Telepon, &telepon.Password)
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
	}
	return berhasil
}
