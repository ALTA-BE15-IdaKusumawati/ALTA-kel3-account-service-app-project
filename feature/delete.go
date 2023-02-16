package feature

import (
	"database/sql"
	"fmt"
	"log"
)

func DeleteUser(db *sql.DB, ID string) {
	statement, errdelete := db.Prepare("DELETE FROM users WHERE id = ?")
	if errdelete != nil {
		log.Fatal("error prepare delete", errdelete.Error())
	}
	defer statement.Close()

	hasil, errdelete := statement.Exec(ID)
	if errdelete != nil {
		log.Fatal("error Exec delete", errdelete.Error())
	} else {
		row, _ := hasil.RowsAffected()
		if row > 0 {
			fmt.Println("proses berhasil dijalankan")

		} else {
			fmt.Println("proses gagal")
		}
	}

}
