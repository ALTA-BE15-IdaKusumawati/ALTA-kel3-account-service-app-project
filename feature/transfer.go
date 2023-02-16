package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

// func Transfer(db *sql.DB, id_pengirim string, telepon_penerima int, nominal float32, balance entities.Users) (dbsub *sql.Rows) {
func Transfer(db *sql.DB, id_pengirim string, telepon_penerima int, nominal float32, balance entities.Users) {
	//pengecekan saldo pengirim apakah cukup atau tidak
	querySaldo := "SELECT saldo from users where id=?"
	row := db.QueryRow(querySaldo, id_pengirim)
	switch err := row.Scan(&balance.Saldo); err {
	case sql.ErrNoRows:
		fmt.Println("Data tidak tersedia")
	case nil:
		switch {
		case balance.Saldo >= nominal:
			// fmt.Println("boleh")
		case balance.Saldo < nominal:
			fmt.Println("Saldo Anda tidak mencukupi")
		}
	default:
		// panic(err)
		log.Fatal("error, gan", err.Error())
	}

	//deklarasi query
	var queryselect = "select id from users where telepon=?"
	var queryTransfer = "insert into transfer(user_id_pengirim, user_id_penerima, nominal) values(?, ?, ?)"
	//mengecek adanya telepon penerima di database
	rows, errTransfer := db.Query(queryselect, telepon_penerima)
	if errTransfer != nil {
		log.Fatal("error trf", errTransfer.Error())
	}
	var prf []entities.Users //membuat penampung untuk data yang di-generate dari database
	for rows.Next() {
		var datarow entities.Users
		//mengecek apakah id penerima (berdasarkan nomor telepon yang diinput) ada di database
		errTransfer := rows.Scan(&datarow.ID)
		if errTransfer != nil {
			panic(errTransfer.Error())
		}
		prf = append(prf, datarow) //memasukkan data ke var prf

	}
	var id_penerima string //membuat variabel penampung untuk id penerima
	for _, v := range prf {
		id_penerima = v.ID //memasukkan data id penerima ke variabel id_penerima
	}

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
