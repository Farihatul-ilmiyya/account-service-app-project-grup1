package controller

import (
	"account-service-app/entity"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
)

func TopUp(db *sql.DB, phoneNumber string, amount float64) (string, error) {
	// membuat permintaan transaksi.
	transaction, err := db.Begin()
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %v", err)
	}
	// Menunda transaksi jika ada yang gagal
	defer transaction.Rollback()

	Uuid := uuid.New()

	// Memeriksa saldo user
	var startingBalance float64
	sqlQuery1 := `SELECT balance FROM users WHERE phone_number = ?`
	err = db.QueryRow(sqlQuery1, phoneNumber).Scan(&startingBalance)
	if err != nil {
		return "", fmt.Errorf("failed to fetch starting balance: %v", err)
	}

	//update user balance
	var newBalance = startingBalance + amount
	sqlQuery2 := `UPDATE users SET balance = balance + ? WHERE phone_number = ?`
	_, err = db.Exec(sqlQuery2, newBalance, phoneNumber)
	if err != nil {
		return "", fmt.Errorf("failed to update user balance: %v", err)
	}

	// Get user ID
	var userId string
	sqlQuery3 := `SELECT id FROM users WHERE phone_number = ?`
	err = db.QueryRow(sqlQuery3, phoneNumber).Scan(&userId)
	if err != nil {
		return "", err
	}

	// Menambahkan baris baru di tabel topup_histories

	sqlQuery4 := `INSERT INTO top_up (id,user_id, amount, created_at) VALUES (?,?, ?, NOW())`
	_, err = db.Exec(sqlQuery4, Uuid, userId, amount)
	if err != nil {
		return "", err
	}

	// commit transaction
	if err = transaction.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %v", err)
	}

	outputStr := "\n[SUCCESS] Top Up was successfull.\n"
	return outputStr, nil

}

func HistoryTopUp(db *sql.DB, phoneNumber string) ([]entity.HistoryTopUp, error) {

	// Query top-up histories for a specific user
	sqlQuery := `
	SELECT th.user_id,u.username ,th.amount, th.created_at
	FROM top_up AS th
	JOIN users AS u ON th.user_id = u.id
	WHERE u.phone_number = ?
	ORDER BY th.created_at DESC
	`
	// Uuid := uuid.New()
	rows, err := db.Query(sqlQuery, phoneNumber)
	if err != nil {
		return []entity.HistoryTopUp{}, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	histories := []entity.HistoryTopUp{}
	for rows.Next() {
		var history entity.HistoryTopUp
		var createdAt []uint8 // Use []byte to store the raw value

		err := rows.Scan(&history.Id, &history.Username, &history.Amount, &createdAt)
		if err != nil {
			return []entity.HistoryTopUp{}, fmt.Errorf("failed to scan row: %v", err)
		}
		// Parse the createdAt value into a time.Time variable
		createdAtStr := string(createdAt)
		history.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			log.Printf("failed to parse created_at value: %v", err)
			continue
		}
		histories = append(histories, history)
	}

	if err = rows.Err(); err != nil {
		return []entity.HistoryTopUp{}, fmt.Errorf("an error occurred while retrieving rows: %v", err)
	}

	return histories, nil
}
