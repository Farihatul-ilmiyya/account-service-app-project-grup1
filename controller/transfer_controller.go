package controller

import (
	"database/sql"
	"fmt"
	"log"
)

func Transfer(db *sql.DB, phoneSender, phoneRecipient string, amount float64) (string, error) {
	//0.memulai transaksi
	//1. cek uang pengirim dulu
	//2. jika uang lebih besar atau sama dengan amount
	//
	//4. lalu kirim dengan cara update uang pengirim yaitu balance - ...
	//5. update uang penerima bertambah balance + ...
	//6. masukkan transaksinya ke tabel topup agar tercatat

	tx, err := db.Begin()
	if err != nil {
		log.Fatal("failed to begin transaction", err)
	}

	//cek uang sender
	var senderBalance float64
	err = tx.QueryRow("SELECT balance FROM users WHERE phone_number = ? AND deleted_at IS NULL", phoneSender).Scan(&senderBalance)
	if err != nil {
		return "", err
	}

	if senderBalance < amount {
		return "insufficient balance", nil
	}

	//uang sender berkurang
	_, err = tx.Exec("UPDATE users SET balance= balance - ? WHERE phone_number = ? AND deleted_at IS NULL", amount, phoneSender)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	//select userID dari sender
	QueryID := ("SELECT id FROM users WHERE phone_number = ? AND deleted_at IS NULL")

	var senderID string
	err = tx.QueryRow(QueryID, phoneSender).Scan(&senderID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("sender's account not found")
		}
		return "", fmt.Errorf("error querying sender's account: %v", err)
	}

	//uang penerima bertambah
	_, err = tx.Exec("UPDATE users SET balance= balance + ? WHERE phone_number = ? AND deleted_at IS NULL", amount, phoneRecipient)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	//select userID dari recipient
	var recipientID string
	err = tx.QueryRow(QueryID, phoneRecipient).Scan(&recipientID)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("recipient's account not found")
		}
		return "", fmt.Errorf("error querying recipient's account: %v", err)
	}

	//jika transfer berhasil maka masukkan data ke transfer history
	_, err = tx.Exec("INSERT INTO transfer (user_id_sender, user_id_recipient, amount) VALUES(?, ?, ?)", senderID, recipientID, amount)
	if err != nil {
		return "", err
	}

	//commit transaksi
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal()
	}
	outputStr := "\n[SUCCESS] Transfer successfully.\n\n"
	return outputStr, nil
}
