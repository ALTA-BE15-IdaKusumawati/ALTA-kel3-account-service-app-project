package main

import (
	"account-service-app-project/config"
	"account-service-app-project/entities"
	"account-service-app-project/feature"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.ConnectToDB()
	defer db.Close()
	var dataUserlogin entities.Users
	for {
		var input int
		fmt.Println("selamat datang")
		fmt.Println("silahkan pilih menu :")
		fmt.Println("menu 1 untuk mendaftar")
		fmt.Println("menu 2 untuk login jika sudah mempunyai akun")
		fmt.Println("menu 0 untuk membatalkan transaksi")
		fmt.Printf("masukan menu : ")
		fmt.Scanln(&input)

		switch input {

		case 1:

			newUsers := entities.Users{}
			fmt.Println("-----------------")
			fmt.Println("Masukkan ID User:")
			fmt.Scanln(&newUsers.ID)
			fmt.Println("-------------------")
			fmt.Println("Masukkan Name User:")
			fmt.Scanln(&newUsers.Name)
			fmt.Println("----------------------")
			fmt.Println("Masukkan Telepon User:")
			fmt.Scanln(&newUsers.Telepon)
			fmt.Println("--------------------")
			fmt.Println("Masukkan Email User:")
			fmt.Scanln(&newUsers.Email)
			fmt.Println("-----------------------")
			fmt.Println("Masukkan password User:")
			fmt.Scanln(&newUsers.Password)

			feature.InsertUser(db, newUsers)

		case 2:
			var connect string
			var telpon int
			var password string
			fmt.Println("----------------------")
			fmt.Print("Masukkan nomor telepon: ")
			fmt.Scanln(&telpon)
			fmt.Print("Masukkan password     : ")
			fmt.Scanln(&password)

			loginUsers := entities.Users{Telepon: telpon, Password: password}
			connect, dataUserlogin = feature.LoginUser(db, loginUsers)

			if connect == "login gagal" {
				fmt.Println("---------------------------")
				fmt.Println("Telepon atau password salah")
				fmt.Println("---------------------------")
			}
			if connect == "login berhasil" {
				var input1 int
				fmt.Println("---------------------")
				fmt.Println("silahkan pilih menu:")
				fmt.Println("menu 3 untuk melihat profil")
				fmt.Println("menu 4 untuk mengedit profil")
				fmt.Println("menu 5 untuk delete account")
				fmt.Println("menu 6 untuk melakukan topup")
				fmt.Println("menu 7 untuk melakukan transfer")
				fmt.Println("menu 8 untuk melihat riwayat topup")
				fmt.Println("menu 9 untuk melihat riwayat transfer")
				fmt.Println("menu 10 untuk melihat profil pengguna lain")
				fmt.Println("menu 0 untuk membatalkan transaksi")
				fmt.Println("masukan nomor menu : ")
				fmt.Scanln(&input1)

				switch input1 {
				case 3:
					{
						feature.GetUsers(db, dataUserlogin.ID)
					}
				case 4:
					{
						updatedataUser := entities.Users{}
						fmt.Println("--------------------------")
						fmt.Println("Masukkan Name User:")
						fmt.Scanln(&updatedataUser.Name)
						fmt.Println("--------------------------")
						fmt.Println("Masukkan Telepon User:")
						fmt.Scanln(&updatedataUser.Telepon)
						fmt.Println("--------------------------")
						fmt.Println("Masukkan Email User:")
						fmt.Scanln(&updatedataUser.Email)
						fmt.Println("--------------------------")
						feature.UpdateUser(db, dataUserlogin.ID, updatedataUser)
						fmt.Println("--------------------------")
					}
				case 5:
					{
						feature.DeleteUser(db, dataUserlogin.ID)
					}
				case 6:
					{
						fmt.Println("----------------------")
						fmt.Println("Masukkan nominal topup")
						var duit entities.Topup
						fmt.Scanln(&duit.Nominal)
						fmt.Println("----------------------")
						feature.EntryTopup(db, dataUserlogin.ID, duit)
						fmt.Println("----------------------")
					}
				case 7:
					{
						fmt.Println("-------------------------------")
						fmt.Println("Masukkan nomor telepon penerima:")
						var teleponPenerima int
						fmt.Scanln(&teleponPenerima)
						var nominal int
						fmt.Println("-------------------------------")
						fmt.Println("Masukkan nominal:")
						fmt.Scanln(&nominal)
						var saldo entities.Users
						fmt.Println("-------------------------------")
						feature.Transfer(db, dataUserlogin.ID, teleponPenerima, nominal, saldo)
						fmt.Println("-------------------------------")

					}
				case 8:
					{
						fmt.Println("----------------------------------------")
						fmt.Println("             RIWAYAT TOPUP")
						fmt.Println("----------------------------------------")
						feature.TopupHistory(db, dataUserlogin.ID)
						fmt.Println("----------------------------------------")
					}
				case 9:
					{
						fmt.Println("----------------------------------------")
						fmt.Println("           RIWAYAT TRANSFER")
						fmt.Println("----------------------------------------")
						feature.TransferHistory(db, dataUserlogin.ID)
						fmt.Println("----------------------------------------")
					}
				case 10:
					{
						fmt.Println("---------------------------")
						fmt.Println("Masukkan nomor telepon user")
						var profile entities.Users
						fmt.Scanln(&profile.Telepon)
						feature.OtherUser(db, dataUserlogin.ID, profile)
					}
				case 0:
					{
						fmt.Println("-------------------------------")
						fmt.Println("terima kasih telah bertransaksi")
						fmt.Println("-------------------------------")
						return
					}
				}
			}

		case 0:
			fmt.Println("terima kasih telah bertransaksi")
			return
		default:
			fmt.Println("menu yang anda masukan tidak ada")
		}

	}
}
