package feature

import (
	"database/sql"
	"fmt"
)

func updateSaldo(db *sql.DB, nominal int, id_penerima string) {
	stmt, err := db.Prepare("UPDATE users SET saldo=saldo+? WHERE users.id=?")
	if err != nil {
		panic(err.Error())
	}

	result, err := stmt.Exec(nominal, id_penerima)
	if err != nil {
		panic(err.Error())
	}

	RowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}

	if RowsAffected > 0 {
		fmt.Println("Saldo updated")
	}
}
