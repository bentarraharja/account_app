package controllers

import (
	"account_app/entities"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func AddAccount(db *sql.DB) {
	var newAccount entities.Account

	fmt.Print("Enter nama: ")
	fmt.Scan(&newAccount.Fullname)
	fmt.Print("Enter alamat: ")
	fmt.Scan(&newAccount.Address)
	fmt.Print("Enter phone: ")
	fmt.Scan(&newAccount.Phone)
	fmt.Print("Enter email: ")
	fmt.Scan(&newAccount.Email)
	fmt.Print("Enter Password: ")
	fmt.Scan(&newAccount.Password)
	newAccount.Balance = 0.0
	newAccount.CreatedAt = time.Now()

	// Perform the SQL INSERT operation
	_, err := db.Exec("INSERT INTO accounts (Fullname,Address,Phone,Email, Password,Balance,Created_at) VALUES (?, ?,?, ?,?, ?,?)", newAccount.Fullname, newAccount.Address, newAccount.Phone, newAccount.Email, newAccount.Password, newAccount.Balance, newAccount.CreatedAt)
	if err != nil {
		log.Println("Error adding account:", err)
		return
	}

	fmt.Println("Account successfully registered!")

}

func DeleteAccount(db *sql.DB, accountID int) {
	// Perform the SQL UPDATE operation
	_, err := db.Exec("UPDATE accounts SET deleted_at = ? WHERE id = ?", time.Now(), accountID)
	if err != nil {
		log.Println("Error soft deleting account:", err)
		return
	}

	fmt.Println("Account successfully soft deleted!")
}
