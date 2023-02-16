package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
)

// func Transfer(db *sql.DB, id_pengirim string, telepon_penerima int, id_penerima entities.Users, nominal float32, amount entities.Users) {
func Transfer(db *sql.DB, id_pengirim string, telepon_penerima int, nominal float32, balance entities.Users) {
	querySaldo := "SELECT saldo from users where id=?"
	row := db.QueryRow(querySaldo, id_pengirim)
	switch err := row.Scan(&balance.Saldo); err {
	case sql.ErrNoRows:
		fmt.Println("Data tidak tersedia")
	case nil:
		switch {
		case balance.Saldo >= nominal:
			// fmt.Println("saldo akan ditransfer sebesar Rp.", nominal, ".00")
			fmt.Println("")
		case balance.Saldo < nominal:
			fmt.Println("Saldo Anda tidak mencukupi")
		}
	default:
		panic(err)
	}

	queryselect := "select id from users where telepon=?"
	queryTransfer := `INSERT INTO transfer(user_id_pengirim, user_id_penerima, nominal) 
						VALUES(?, ?, ?)`
	rows, err := db.Query(queryselect)
	rows.Scan(&telepon_penerima)
	if err != nil {
		panic(err)
	}
	// queryTransfer := "INSERT INTO transfer(user_id_pengirim, user_id_penerima, nominal) VALUES(?, ?, ?)"
	statement, errPrepare := db.Prepare(queryTransfer)
	if errPrepare != nil {
		panic(errPrepare.Error())
	}

	res, errInsert := statement.Exec(id_pengirim, id_penerima, nominal)
	if errInsert != nil {
		panic(errInsert.Error())
	} else {
		ro, _ := res.RowsAffected()
		if ro > 0 {
			fmt.Println("Transfer telah berhasil")
		} else {
			fmt.Println("Transfer gagal")
		}
	}

	//update saldo pengirim
	pengirim, errPengirim := db.Prepare("UPDATE users INNER JOIN transfer ON users.id=? SET saldo=saldo-?")
	if errPengirim != nil {
		panic(errPengirim.Error())
	}
	res, errUpdate := pengirim.Exec(id_pengirim, nominal)
	if errUpdate != nil {
		panic(errUpdate.Error())
	}

	a, errUpdate := res.RowsAffected()
	if errUpdate != nil {
		panic(errUpdate.Error())
	}
	if a > 0 {
		fmt.Println("")
	}

	//update saldo penerima
	penerima, errPenerima := db.Prepare("UPDATE users INNER JOIN transfer ON users.id=? SET saldo=saldo+?")
	if errPenerima != nil {
		panic(errPenerima.Error())
	}
	has, errUpdatePenerima := penerima.Exec(id_penerima, nominal)
	if errUpdate != nil {
		panic(errUpdatePenerima.Error())
	}

	b, errUpdatePenerima := has.RowsAffected()
	if errUpdate != nil {
		panic(errUpdatePenerima.Error())
	}
	if b > 0 {
		fmt.Println("")
	}
}
