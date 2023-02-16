package feature

import (
	"account-service-app-project/entities"
	"database/sql"
	"fmt"
)

func TransferHistory(db *sql.DB, id string) {
	rowsTransfer, errTransfer := db.Query("SELECT created_at, nominal FROM transfer WHERE YEARWEEK(created_at)=YEARWEEK(NOW()) AND transfer.user_id_pengirim=?", id)
	// rowsTransfer, errTransfer := db.Query("SELECT created_at, nominal FROM transfer WHERE transfer.user_id_pengirim=?", id)
	if errTransfer != nil {
		panic(errTransfer.Error())
	}
	var historyTransfer []entities.Transfer
	for rowsTransfer.Next() {
		var datarow entities.Transfer
		errTransfer := rowsTransfer.Scan(&datarow.Tanggal_Transaksi, &datarow.Nominal)
		if errTransfer != nil {
			panic(errTransfer.Error())
		}
		historyTransfer = append(historyTransfer, datarow)
	}
	if len(historyTransfer) > 0 {
		for _, histTransfer := range historyTransfer {
			fmt.Println("Tanggal Transaksi   | Nominal Transaksi ")
			fmt.Println(histTransfer.Tanggal_Transaksi, "|", histTransfer.Nominal)
		}
	} else {
		fmt.Println("     Anda belum melakukan transaksi     ")
	}
}
