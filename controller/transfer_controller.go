package controller

import (
	"account-service-app/entity"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
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
	Uuid := uuid.New()
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
	_, err = tx.Exec("INSERT INTO transfer (id, user_id_sender, user_id_recipient, amount) VALUES(?, ?, ?, ?)", Uuid, senderID, recipientID, amount)
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

func TransferHistory(db *sql.DB, phoneNumber string) ([]entity.History, error) {
	var tfHistory []entity.History
	var createdAt []uint8

	senderQuery := `SELECT tf.id, recipient.username, recipient.phone_number, tf.amount, tf.created_at
	FROM transfer AS tf
	INNER JOIN users AS sender ON user_id_sender = sender.id
	INNER JOIN users AS recipient ON user_id_recipient = recipient.id
	WHERE sender.phone_number = ?
	AND sender.deleted_at IS NULL
	ORDER BY tf.id DESC;`

	recipientQuery := `SELECT tf.id, sender.username, sender.phone_number, tf.amount, tf.created_at
	FROM transfer AS tf
	INNER JOIN users AS sender ON user_id_sender = sender.id
	INNER JOIN users AS recipient ON user_id_recipient = recipient.id
	WHERE recipient.phone_number = ?
	AND recipient.deleted_at IS NULL
	ORDER BY tf.id DESC;`

	//mendapatkan riwayat transfeer sebagai SENDER
	senderRow, err := db.Query(senderQuery, phoneNumber)
	if err != nil {
		return nil, err
	}
	defer senderRow.Close()

	for senderRow.Next() {

		var tfHistories entity.History

		err := senderRow.Scan(&tfHistories.ID, &tfHistories.Username, &tfHistories.PhoneNumber, &tfHistories.Amount, &createdAt)
		if err != nil {
			return nil, err
		}
		//Konversi createdAt menjadi objek time.Time
		createdAtStr := string(createdAt)
		tfHistories.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, err
		}
		tfHistories.IsSender = true
		tfHistory = append(tfHistory, tfHistories)
	}

	//mendapatkan riwayat transfer sebagai penerima

	recipientRow, err := db.Query(recipientQuery, phoneNumber)
	if err != nil {
		return nil, err
	}
	defer recipientRow.Close()

	for recipientRow.Next() {
		var tfHistories entity.History
		err := recipientRow.Scan(&tfHistories.ID, &tfHistories.Username, &tfHistories.PhoneNumber, &tfHistories.Amount, &createdAt)
		if err != nil {
			return nil, err
		}

		//Konversi createdAt menjadi objek time.Time
		createdAtStr := string(createdAt)
		tfHistories.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, err
		}
		tfHistories.IsSender = false
		tfHistory = append(tfHistory, tfHistories)
	}
	return tfHistory, nil
}
