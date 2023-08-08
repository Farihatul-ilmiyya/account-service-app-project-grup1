package controller

import (
	"account-service-app/entity"
	"account-service-app/helpers"
	"database/sql"
	"errors"
)

func RegisterAccount(db *sql.DB, user entity.Users) (string, error) {
	passHashing := helpers.HashPassword(user.Password)
	_, err := db.Exec("INSERT INTO users(id, username, email, password, phone_number, date_of_birth, address, balance) VALUE (?, ?, ?, ?, ?, ?, ?, ?)", user.ID, user.Username, user.Email, passHashing, user.PhoneNumber, user.DateOfBirth, user.Address, user.Balance)
	defer db.Close()
	if err != nil {
		return "", errors.New("[FAILED] Failed to register account")
	}

	outputStr := "\n[SUCCESS] Account registered successfully.\n\n"
	return outputStr, nil
}
