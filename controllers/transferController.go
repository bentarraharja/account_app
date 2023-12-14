package controllers

import (
	"account_app/entities"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// GetUser mengambil detail pengguna berdasarkan ID atau nomor telepon
func GetUser(db *sql.DB, identifier interface{}) (*entities.Account, error) {
	var akun entities.Account

	switch identifier := identifier.(type) {
	case int:
		// Identifier ID
		err := db.QueryRow("SELECT id, full_name, address, phone, email, password, balance, created_at, updated_at FROM accounts WHERE id = ?", identifier).
			Scan(&akun.ID, &akun.FullName, &akun.Address, &akun.Phone, &akun.Email, &akun.Password, &akun.Balance, &akun.CreatedAt, &akun.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error fetching user details: %v", err)
		}
	case string:
		// Identifier Phone number
		err := db.QueryRow("SELECT id, full_name, address, phone, email, password, balance, created_at, updated_at FROM accounts WHERE phone = ?", identifier).
			Scan(&akun.ID, &akun.FullName, &akun.Address, &akun.Phone, &akun.Email, &akun.Password, &akun.Balance, &akun.CreatedAt, &akun.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error fetching user details: %v", err)
		}
	default:
		return nil, fmt.Errorf("invalid identifier type")
	}

	return &akun, nil
}

// Transfer menggunakan ID atau nomor telepon
func Transfer(db *sql.DB, senderIdentifier, receiverIdentifier string, amount int) (*entities.Transfer, error) {
	var sender *entities.Account
	var receiver *entities.Account

	// Check if sender and receiver exist
	//sender
	sender, err := GetUser(db, senderIdentifier)
	if err != nil {
		return nil, fmt.Errorf("error fetching sender details: %v", err)
	}

	receiver, err = GetUser(db, receiverIdentifier)
	if err != nil {
		return nil, fmt.Errorf("error fetching receiver details: %v", err)
	}

	// Check if sender has sufficient balance
	if sender.Balance < amount {
		return nil, fmt.Errorf("insufficient funds for transfer")
	}

	// Begin a transaction
	tx, err := db.Begin()
	if err != nil {
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}
	defer func() {
		if err := recover(); err != nil {
			_ = tx.Rollback()
		}
	}()

	// Update sender's balance
	_, err = tx.Exec("UPDATE accounts SET balance =  ? WHERE id = ?", sender.Balance-amount, sender.ID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error updating sender's balance: %v", err)
	}

	// Update receiver's balance
	_, err = tx.Exec("UPDATE accounts SET balance =   ? WHERE id = ?", receiver.Balance+amount, receiver.ID)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error updating receiver's balance: %v", err)
	}

	transfer := &entities.Transfer{
		AccountIdSender:   sender.ID,
		AccountIdReceiver: receiver.ID,
		Amount:            amount,
		CreatedAt:         time.Now(),
	}

	// Insert transfer record
	result, err := tx.Exec("INSERT INTO transfers (account_id_sender, account_id_receiver, amount, created_at) VALUES (?, ?, ?, ?)",
		transfer.AccountIdSender, transfer.AccountIdReceiver, transfer.Amount, transfer.CreatedAt)
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error inserting transfer record: %v", err)
	}

	// Get the last inserted ID
	transferID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("error getting last insert ID: %v", err)
	}
	transfer.ID = int(transferID)

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("error committing transaction: %v", err)
	}

	fmt.Printf("Transfer of $%v from %s to %s successful.\n", amount, sender.FullName, receiver.FullName)
	return transfer, nil
}

func HistoryTransfer(db *sql.DB, phoneNumber string) ([]entities.Transfer, error) {
	var accountID int

	// Assuming you have a table 'accounts' with columns 'ID' and 'Phone'
	err := db.QueryRow("SELECT id FROM accounts WHERE phone = ?", phoneNumber).Scan(&accountID) // phoneNumber).Scan(&accountID)
	if err != nil {
		return nil, fmt.Errorf("error fetching account ID: %v", err)
	}

	rows, err := db.Query(`
		SELECT id, account_id_sender, account_id_receiver, amount, created_at
		FROM transfers
		WHERE account_id_sender = ? OR account_id_receiver = ?
	`, accountID, accountID)
	if err != nil {
		return nil, fmt.Errorf("error fetching transfer history: %v", err)
	}
	defer rows.Close()

	var transferHistory []entities.Transfer
	for rows.Next() {
		var transfer entities.Transfer
		err := rows.Scan(&transfer.ID, &transfer.AccountIdSender, &transfer.AccountIdReceiver, &transfer.Amount, &transfer.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning transfer record: %v", err)
		}
		transferHistory = append(transferHistory, transfer)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over transfer records: %v", err)
	}

	fmt.Printf("Transfer History for Account with Phone Number %s:\n", phoneNumber)
	for _, transfer := range transferHistory {
		fmt.Printf("Transfer ID: %d\nSender Account ID: %d\nReceiver Account ID: %d\nAmount: %d\nCreated At: %s\n\n",
			transfer.ID, transfer.AccountIdSender, transfer.AccountIdReceiver, transfer.Amount, transfer.CreatedAt.Format("15:05 02-Jan-2006"))
	}

	return transferHistory, nil
}
