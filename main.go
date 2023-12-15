package main

import (
	"account_app/controllers"
	"account_app/entities"
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

	var sessionLogin entities.Account
	var choice int
	for {
		// Menampilkan menu
		fmt.Println("Menu:")
		fmt.Println("[1] Add Account (Register)")
		fmt.Println("[2] Login")
		fmt.Println("[0] Keluar dari Sistem")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&choice)
		fmt.Println("========================================================OUTPUT=======================================================")

		// Menampilkan menu ke-1
		switch choice {
		case 1:
			controllers.AddAccount(db)
		case 2:
			dataLogin, err := controllers.Login(db)
			if err != nil {
				//menampilkan error dengan value error di controllerAccount dan menghentikan program program
				log.Fatal(err)
			}
			//menampung data account yang login kedalam variabel sessionLogin
			sessionLogin = *dataLogin
			// Berhasil login
			fmt.Println("--------------------------------------------------------")
			fmt.Println("Login successful!")
			fmt.Println("Selamat datang", dataLogin.FullName)
			fmt.Println("--------------------------------------------------------")
			if err == nil {
				// jika login berhasil tanpa error maka akan menampilkan menu ke-2
				var pilih int
				var exit bool
				for !exit {
					fmt.Println("[3] Read Account")
					fmt.Println("[4] Update Account")
					fmt.Println("[5] Delete Account")
					fmt.Println("[6] Top-Up")
					fmt.Println("[7] Transfer")
					fmt.Println("[8] History Top-Up")
					fmt.Println("[9] History Transfer")
					fmt.Println("[10] View Other User Profile")
					fmt.Println("[0] Logout")
					fmt.Print("Pilihan Anda: ")
					fmt.Scan(&pilih)
					fmt.Println("========================================================OUTPUT=======================================================")
					switch pilih {
					case 3:
						controllers.ReadAccount(db, &sessionLogin)
					case 4:
						controllers.UpdateAccount(db, &sessionLogin)
					case 5:
						controllers.DeleteAccount(db, &sessionLogin)
						exit = true
					case 6:
						controllers.TopUp(db, &sessionLogin)
					case 7:
						// Transfer Funds
						// var receiverIdentifier string
						// var amount int

						// fmt.Print("Enter your phone (sender): ")
						// fmt.Scan(&sessionLogin)

						_, err := controllers.Transfer(db, &sessionLogin)
						if err != nil {
							log.Println("Error transferring funds:", err)
						}

					case 8:
						controllers.HistoryTopUp(db, &sessionLogin)
					case 9:
						// var phoneNumber string
						// fmt.Print("Enter the account phone number: ")
						// fmt.Scan(&phoneNumber)

						_, err := controllers.HistoryTransfer(db, &sessionLogin)
						if err != nil {
							log.Println("Error displaying transfer history for account:", err)
						}

						// fmt.Println("Transfer History:")
						// for _, transfer := range transfers {
						// 	fmt.Printf("TransferID: %d, SenderID: %d, ReceiverID: %d, Amount: %d, CreatedAt: %s\n", transfer.ID, transfer.AccountIdSender, transfer.AccountIdReceiver, transfer.Amount, transfer.CreatedAt.Format("2006-01-02 15:04:05"))
						// }

					case 10:
						var Phone string
						fmt.Print("Masukan phone yang profilenya ingin dilihat: ")
						fmt.Scan(&Phone)

						if err := controllers.ViewOtherUserProfile(db, Phone); err != nil {
							log.Println("Error viewing user profile:", err)
						}
					case 0:
						log.Println("Berhasil Logout.")
						exit = true
						fmt.Println("=====================================================================================================================")
					default:
						fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
					}
					fmt.Println("=====================================================================================================================")
					if exit != true {
						var konfirmasi string
						fmt.Print("Apakah anda ingin memilih menu lagi? (y/n): ")
						fmt.Scan(&konfirmasi)
						if konfirmasi != "y" {
							log.Println("Logout Account")
							fmt.Println("=====================================================================================================================")
							exit = true
						}
					}
				}
			}
		case 0:
			log.Println("Terima kasih telah bertransaksi, Exit program...")
			fmt.Println("=====================================================================================================================")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			fmt.Println("=====================================================================================================================")
		}
	}
}
