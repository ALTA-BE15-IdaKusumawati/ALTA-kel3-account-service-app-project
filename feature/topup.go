package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
	"log"
)

func EntryTopup(db *sql.DB, id string, duit entities.Topup) {

	var query = "INSERT INTO topup(user_id, nominal) VALUES(?,?)"
	statement, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal("error", errPrepare.Error())
	}

	res, errInsert := statement.Exec(id, duit.Nominal)
	if errInsert != nil {
		log.Fatal("error", errInsert.Error())
	} else {
		ro, _ := res.RowsAffected()
		if ro > 0 {
			fmt.Println("Topup telah berhasil")
		} else {
			fmt.Println("Topup gagal")
		}
	}

	updateSaldo(db, int(duit.Nominal), id)

	// ya, err := db.Prepare("UPDATE users SET saldo=saldo+? WHERE users.id=?")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// // var n string = duit.User_ID
	// res, errUpdate := ya.Exec(duit.Nominal, id)
	// if errUpdate != nil {
	// 	panic(errUpdate.Error())
	// }

	// a, errUpdate := res.RowsAffected()
	// if errUpdate != nil {
	// 	panic(errUpdate.Error())
	// }
	// if a > 0 {
	// 	fmt.Printf("Selamat, saldo anda telah bertambah")
	// }

}
