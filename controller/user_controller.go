package controller

import (
	"account-service-app/entity"
	"account-service-app/helpers"
	"database/sql"

	"github.com/google/uuid"
)

func RegisterAccount(db *sql.DB, user entity.Users) (string, error) {
	Uuid := uuid.New()

	passHashing := helpers.HashPassword(user.Password)

	_, err := db.Exec("INSERT INTO users(id, username, email, password, phone_number, date_of_birth, address, balance) VALUE (?, ?, ?, ?, ?, ?, ?, ?)", Uuid, user.Username, user.Email, passHashing, user.PhoneNumber, user.DateOfBirth, user.Address, user.Balance)

	defer db.Close()

	if err != nil {
		return "", err
	}

	outputStr := "\n[SUCCESS] Account registered successfully.\n\n"
	return outputStr, nil
}
