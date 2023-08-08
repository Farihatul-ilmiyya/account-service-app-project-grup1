package controller

import (
	"account-service-app/entity"
	"account-service-app/helpers"
	"database/sql"
	"errors"

	"fmt"
	"log"

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

func Profile(db *sql.DB, user entity.Users) (entity.Users, error) {
	sqlQuery := "select id ,username,phone_number,email ,date_of_birth,balance,address from users where phone_number = ? and deleted_at is null"
	fmt.Println(user.PhoneNumber)
	err := db.QueryRow(sqlQuery, user.PhoneNumber).Scan(&user.ID, &user.Username, &user.PhoneNumber, &user.Email, &user.DateOfBirth, &user.Balance, &user.Address)

	if err != nil {
		log.Fatal("Error check profile", err.Error())
		return entity.Users{}, err
	}
	outputStr := fmt.Sprintln("-----------------------------------------")
	outputStr += fmt.Sprintln("Your Account Information")
	outputStr += fmt.Sprintln("-----------------------------------------")
	outputStr += fmt.Sprintf("ID\t\t: %s\n", user.ID)
	outputStr += fmt.Sprintf("User Name\t: %s\n", user.Username)
	outputStr += fmt.Sprintf("Birth Date\t: %s\n", user.DateOfBirth)
	outputStr += fmt.Sprintf("Address\t\t: %s\n", user.Address)
	outputStr += fmt.Sprintf("Email\t\t: %s\n", user.Email)
	outputStr += fmt.Sprintf("Phone Number\t: %s\n", user.PhoneNumber)
	outputStr += fmt.Sprintf("Balance\t\t: %.2f\n", user.Balance)
	outputStr += fmt.Sprintln("-----------------------------------------")
	fmt.Println(outputStr)
	return user, nil

}

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
