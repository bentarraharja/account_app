package controllers

import (
	"account_app/entities"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func TopUp(db *sql.DB, sessionLogin *entities.Account) {
	// Memulai proses topup
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// Input topup balance
	var topupBalance int
	fmt.Print("Masukan amount top-up: ")
	fmt.Scan(&topupBalance)

	// Pada proses transaksi bank sebaiknya INSERT topup history dulu baru UPDATE saldo di akunnya, gunannya untuk menghindari keluhan user akibat value saldo yang tiba tiba berubah ketika terjadi kesalahan
	// Melakukan proses INSERT topup history
	_, errInsTopup := tx.Exec("INSERT INTO top_ups (account_id, amount) VALUES (?, ?)", sessionLogin.ID, topupBalance)
	if errInsTopup != nil {
		// Rollback topup jika ada kesalahan dalam INSERT data ke table topup
		tx.Rollback()
		log.Fatal("error run query INSERT topup ", errInsTopup.Error())
	}

	// Melakukan proses update saldo
	result, errTopup := tx.Exec("UPDATE accounts SET balance = balance + ? WHERE id = ?", topupBalance, sessionLogin.ID)
	if errTopup != nil {
		// Rollback topup jika ada kesalahan dalam menambahkan balance ke table account
		tx.Rollback()
		log.Fatal("error run query UPDATE topup balance ", errTopup.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			log.Println("Success update data balance!")
		} else {
			fmt.Println("Tidak ada data balance yang di update!")
		}
	}

	// Simpan perubahan dengan commit jika semua proses berhasil
	errCom := tx.Commit()
	if errCom != nil {
		log.Fatal(errCom)
	}

	fmt.Println("Transaksi top-up berhasil!")
}

func HistoryTopUp(db *sql.DB, sessionLogin *entities.Account) {
	// Menyimpan data yang dibaca di query SELECT
	var topups []entities.TopUp

	// menjalankan perintah query SELECT
	rows, errSelect := db.Query("SELECT id, account_id, amount, created_at FROM top_ups where account_id = ?", sessionLogin.ID)
	// handle error query SELECT
	if errSelect != nil {
		log.Fatal("error run query SELECT ", errSelect.Error())
	}

	//Proses membaca per baris/row
	for rows.Next() {
		// proses scan data pada struct TopUp
		var dataTopup entities.TopUp
		errScan := rows.Scan(&dataTopup.ID, &dataTopup.AccountID, &dataTopup.Amount, &dataTopup.CreatedAt)
		if errScan != nil {
			log.Fatal("error scan SELECT ", errScan.Error())
		}
		// memasukkan data ke topups
		topups = append(topups, dataTopup)
	}

	//Proses menampilkan data history top-up
	for _, v := range topups {
		fmt.Printf("ID: %v\nAccountID: %v\nAmount: %v\nCreatedAt: %v\n", v.ID, v.AccountID, v.Amount, v.CreatedAt)
	}
}
