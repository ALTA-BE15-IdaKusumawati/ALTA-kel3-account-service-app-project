package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
)

func TopupHistory(db *sql.DB, id string) {
	rows, err := db.Query("SELECT created_at, nominal FROM topup WHERE topup.user_id=?", id)
	if err != nil {
		panic(err)
	}
	var history []entities.Topup
	for rows.Next() {
		var datarow entities.Topup
		err := rows.Scan(&datarow.Tanggal_Transaksi, &datarow.Nominal)
		if err != nil {
			panic(err)
		}
		history = append(history, datarow)
	}
	if len(history) > 0 {
		for _, hist := range history {
			fmt.Println("Tanggal Transaksi   | Nominal Transaksi ")
			fmt.Println(hist.Tanggal_Transaksi, "|", hist.Nominal)
		}
	} else {
		fmt.Println("     Anda belum melakukan transaksi     ")
	}

}
