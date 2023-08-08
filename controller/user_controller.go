package controller

import (
	"account-service-app/entity"
	"database/sql"
	"errors"
	"log"
)

func Login(db *sql.DB, user entity.Users) (string, error) {
	var password string

	sqlQuery := "SELECT password FROM users WHERE phone_number = ? AND deleted_at IS NULL"

	err := db.QueryRow(sqlQuery, user.PhoneNumber).Scan(&password)
	if err == sql.ErrNoRows {
		return "", errors.New("User is not found")
	}
	if err != nil {
		log.Fatal("Error login", err.Error())
		return "", err
	}

	if user.Password != password {
		return "", errors.New("Passowrd is not match")
	}
	return "Login successfully", nil

}
