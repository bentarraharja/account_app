package main

import (
	"account_app"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var choice int

	for {
		// Menampilkan menu
		fmt.Println("Menu:")
		fmt.Println("1. Add Account (Register)")
		fmt.Println("2. Login")
		fmt.Println("3. Read Account")
		fmt.Println("4. Update Account")
		fmt.Println("5. Delete Account")
		fmt.Println("6. Top-Up")
		fmt.Println("7. Transfer")
		fmt.Println("8. History Top-Up")
		fmt.Println("9. History Transfer")
		fmt.Println("10. View Other User Profile")
		fmt.Println("0. Keluar dari Sistem")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&choice)

		// Melakukan aksi berdasarkan pilihan menu
		switch choice {
		case 1:
			addAccount()
		case 2:
			login()
		case 3:
			readAccount()
		case 4:
			updateAccount()
		case 5:
			deleteAccount()
		case 6:
			topUp()
		case 7:
			transfer()
		case 8:
			historyTopUp()
		case 9:
			historyTransfer()
		case 10:
			viewOtherUserProfile()
		case 0:
			fmt.Println("Terima kasih telah bertransaksi.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
