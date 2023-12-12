package controllers

import (
	"account_app/entities"
	"database/sql"
	"fmt"
	"log"
)

func ReadAccount(db *sql.DB) {
	// SELECT data
	// Menyimpan data yang dibaca di query SELECT
	var accounts []entities.Account

	// menjalankan perintah query SELECT
	rows, errSelect := db.Query("SELECT id, full_name, address, phone, email, password, balance, created_at, updated_at FROM accounts")
	// handle error query SELECT
	if errSelect != nil {
		log.Fatal("error run query SELECT ", errSelect.Error())
	}

	//Proses membaca per baris/row
	for rows.Next() {
		var dataAccount entities.Account
		// proses scan dataAccount
		errScan := rows.Scan(&dataAccount.ID, &dataAccount.FullName, &dataAccount.Address, &dataAccount.Phone, &dataAccount.Email, &dataAccount.Password, &dataAccount.Balance, &dataAccount.CreatedAt, &dataAccount.UpdatedAt)
		if errScan != nil {
			log.Fatal("error scan SELECT ", errScan.Error())
		}
		// memasukkan dataAccount ke accounts
		accounts = append(accounts, dataAccount)
	}

	for _, v := range accounts {
		fmt.Printf("ID: %v, FullName: %v, Address: %v, Phone: %v, Email: %v, Password: %v, Balance: %v, CreatedAt: %v, UpdatedAt: %v\n", v.ID, v.FullName, v.Address, v.Phone, v.Email, v.Password, v.Balance, v.CreatedAt, v.UpdatedAt)
	}
}
