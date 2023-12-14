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
	fmt.Scan(&newAccount.FullName)
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
	_, err := db.Exec("INSERT INTO accounts (full_name, address, phone, email, password, balance, created_at) VALUES (?, ?,?, ?,?, ?,?)", newAccount.FullName, newAccount.Address, newAccount.Phone, newAccount.Email, newAccount.Password, newAccount.Balance, newAccount.CreatedAt)
	if err != nil {
		log.Println("Error adding account:", err)
		return
	}

	fmt.Println("Account successfully registered!")

}

func DeleteAccount(db *sql.DB, sessionLogin *entities.Account) {
	// Perform the SQL UPDATE operation
	_, err := db.Exec("UPDATE accounts SET deleted_at = ? WHERE phone = ?", time.Now(), sessionLogin.Phone)
	if err != nil {
		log.Println("Error soft deleting account:", err)
		return
	}

	log.Println("Account successfully soft deleted!")
	fmt.Println("Exit menu.....")

}

func ReadAccount(db *sql.DB, sessionLogin *entities.Account) {
	// Menyimpan data yang dibaca di query SELECT
	var accounts []entities.Account

	// menjalankan perintah query SELECT
	rows, errSelect := db.Query("SELECT id, full_name, address, phone, email, password, balance, created_at, updated_at, deleted_at FROM accounts where phone = ?", sessionLogin.Phone)
	// handle error query SELECT
	if errSelect != nil {
		log.Fatal("error run query SELECT ", errSelect.Error())
	}

	//Proses membaca per baris/row
	for rows.Next() {
		// proses scan data pada sessionLogin
		errScan := rows.Scan(&sessionLogin.ID, &sessionLogin.FullName, &sessionLogin.Address, &sessionLogin.Phone, &sessionLogin.Email, &sessionLogin.Password, &sessionLogin.Balance, &sessionLogin.CreatedAt, &sessionLogin.UpdatedAt, &sessionLogin.DeletedAt)
		if errScan != nil {
			log.Fatal("error scan SELECT ", errScan.Error())
		}
		// memasukkan sessionLogin ke accounts
		accounts = append(accounts, *sessionLogin)
	}

	for _, v := range accounts {
		//Cek apakah deleted_at nya memiliki nilai atau tidak
		if v.DeletedAt.Valid == true {
			fmt.Printf("ID: %v\nFullName: %v\nAddress: %v\nPhone: %v\nEmail: %v\nPassword: %v\nBalance: %v\nCreatedAt: %v\nUpdatedAt: %v\nDeletedAt: %v\n", v.ID, v.FullName, v.Address, v.Phone, v.Email, v.Password, v.Balance, v.CreatedAt, v.UpdatedAt, v.DeletedAt.Time)
		} else {
			fmt.Printf("ID: %v\nFullName: %v\nAddress: %v\nPhone: %v\nEmail: %v\nPassword: %v\nBalance: %v\nCreatedAt: %v\nUpdatedAt: %v\nDeletedAt: Null\n", v.ID, v.FullName, v.Address, v.Phone, v.Email, v.Password, v.Balance, v.CreatedAt, v.UpdatedAt)
		}
	}
}

func UpdateAccount(db *sql.DB, sessionLogin *entities.Account) {
	// mengisi data yang ingin di UPDATE
	fmt.Println("***NOTE: jika tidak ingin merubah data pada kolom tertentu gunakan (-)***")
	fmt.Print("Masukan nama baru: ")
	var fullName string
	fmt.Scan(&fullName)
	if fullName == "-" {
		fullName = sessionLogin.FullName
	} else {
		sessionLogin.FullName = fullName
	}
	fmt.Print("Masukan alamat baru: ")
	var address string
	fmt.Scan(&address)
	if address == "-" {
		address = sessionLogin.Address
	} else {
		sessionLogin.Address = address
	}
	fmt.Print("Masukan phone baru: ")
	var phone string
	fmt.Scan(&phone)
	if phone == "-" {
		phone = sessionLogin.Phone
	} else {
		sessionLogin.Phone = phone
	}
	fmt.Print("Masukan email baru: ")
	var email string
	fmt.Scan(&email)
	if email == "-" {
		email = sessionLogin.Email
	} else {
		sessionLogin.Email = email
	}
	fmt.Print("Masukan Password baru: ")
	var password string
	fmt.Scan(&password)
	if password == "-" {
		password = sessionLogin.Password
	} else {
		sessionLogin.Password = password
	}

	//Menjalankan query UPDATE
	result, errUpd := db.Exec("UPDATE accounts SET full_name = ?, address = ?, phone = ?, email = ?, password = ? WHERE id = ?", fullName, address, phone, email, password, sessionLogin.ID)
	if errUpd != nil {
		log.Fatal("error run query UPDATE", errUpd.Error())
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			log.Println("Success update data!")
		} else {
			fmt.Println("Tidak ada data yang di update!")
		}
	}
}

func Login(db *sql.DB) (*entities.Account, error) {
	//mendapatkan data dari variabel yang di input
	fmt.Print("Enter your phone number: ")
	var phone string
	fmt.Scan(&phone)
	fmt.Print("Enter your password: ")
	var password string
	fmt.Scan(&password)

	//mendapatkan data dari data account dan memasukannya ke dalam struct Account
	var dataLogin entities.Account

	// Mengambil satu baris data dari tabel accounts berdasarkan nomor telepon dan kata sandi
	err := db.QueryRow("SELECT id, full_name, address, phone, email, password, balance, created_at, updated_at, deleted_at FROM accounts WHERE phone = ? AND password = ? AND deleted_at IS NULL", phone, password).
		Scan(&dataLogin.ID, &dataLogin.FullName, &dataLogin.Address, &dataLogin.Phone, &dataLogin.Email, &dataLogin.Password, &dataLogin.Balance, &dataLogin.CreatedAt, &dataLogin.UpdatedAt, &dataLogin.DeletedAt)

	//error handling
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("login failed: account not found")
		}
		// Terjadi error lain
		return nil, fmt.Errorf("login failed: %v", err)
	}
	return &dataLogin, nil
}

func ViewOtherUserProfile(db *sql.DB, Phone string) error {
	var Account entities.Account

	query := "SELECT id, full_name, address, phone, email, password, balance, created_at, deleted_at FROM accounts WHERE phone = ? AND deleted_at is null "
	err := db.QueryRow(query, Phone).Scan(
		&Account.ID, &Account.FullName, &Account.Address, &Account.Phone, &Account.Email,
		&Account.Password, &Account.Balance, &Account.CreatedAt, &Account.DeletedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("Error viewing account profile:", err)
		}
		return fmt.Errorf("error fetching user details: %v", err)
	}

	fmt.Println("Account Details:")
	fmt.Printf("ID: *****\nFullName: %v\nAddress: %v\nPhone: %v\nEmail: %v\nPassword: ****\nBalance: ****\n",
		Account.FullName, Account.Address, Account.Phone, Account.Email)

	// Check if DeletedAt is valid before printing
	// if Account.DeletedAt.Valid {
	// 	fmt.Printf("\nDeletedAt: %v", Account.DeletedAt.Time.Format("2006-01-02 15:04:05"))
	// } else {
	// 	fmt.Println("\nDeletedAt: Null")
	// }

	return nil
}
