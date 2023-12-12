package main

import (
	"account_app/controllers"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type EnvConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
}

func InitDB() (*sql.DB, error) {
	var config = EnvConfig{
		DB_USERNAME: os.Getenv("DB_USERNAME"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
	}
	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
	var db *sql.DB
	var err error
	// Cek konfigurasi database
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("error open connection to db: ", err)
		return nil, err
	}

	// Cek ping connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Println("error ping conenction: ", pingErr)
		return nil, pingErr
	}
	fmt.Println("success connect to db!")
	return db, nil
}

func main() {
	db, errInitDB := InitDB()
	if errInitDB != nil {
		log.Fatal("error connect to db ", errInitDB)
	}
	//Close connection to database
	defer db.Close()

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
			AddAccount()
		case 2:
			Login()
		case 3:
			controllers.ReadAccount(db)
		case 4:
			UpdateAccount()
		case 5:
			DeleteAccount()
		case 6:
			TopUp()
		case 7:
			Transfer()
		case 8:
			HistoryTopUp()
		case 9:
			HistoryTransfer()
		case 10:
			ViewOtherUserProfile()
		case 0:
			fmt.Println("Terima kasih telah bertransaksi.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
