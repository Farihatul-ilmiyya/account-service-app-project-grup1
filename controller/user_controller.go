package controller

import (
	"account-service-app/entity"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func RegisterAccount(db *sql.DB, user entity.Users) (string, error) {
	result, err := db.Exec("INSERT INTO users(id, username, email, password, phone_number, date_of_birth, address, balance) VALUE (?, ?, ?, ?, ?, ?, ?, ?)", user.ID, user.Username, user.Email, user.Password, user.PhoneNumber, user.DateOfBirth, user.Address, user.Balance)
	defer db.Close()
	if err!= nil {
		return"", errors.err()
	} else {
		row, _ := result.RowsAffected()
		if row > 0 {
			fmt.Println("[SUCCESS] Account registered successfully.")
		} else {
			fmt.Println("[FAILED] Failed to register account")
		}
	}
	outputStr := "\n[SUCCESS] Account registered successfully.\n\n"
	return outputStr, nil
}
