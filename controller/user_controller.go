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

	if !helpers.ComparePassword(user.Password, password) {
		return "", errors.New("Passowrd is not match")
	}
	return "Login successfully", nil

}

func DeleteUser(db *sql.DB, user entity.Users) (string, error) {
	_, err := db.Exec("UPDATE users SET deleted_at = now() WHERE phone_number = ?", user.PhoneNumber)

	if err != nil {
		return "", err
	}

	outputStr := "\n[SUCCESS] Account deleted successfully.\n\n"
	return outputStr, nil
}

func ReadOtherUser(db *sql.DB, user entity.Users) (entity.Users, error) {
	sqlQuery := "SELECT username, email, phone_Number, date_of_birth, address FROM Users WHERE phone_number = ? AND deleted_at IS NULL "

	err := db.QueryRow(sqlQuery, user.PhoneNumber).Scan(&user.Username, &user.Email, &user.PhoneNumber, &user.DateOfBirth, &user.Address)

	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Users{}, fmt.Errorf("user not found")
		}

	}
	return user, err
}

func UpdateAccount(db *sql.DB, user entity.Users) (string, error) {
	var (
		field string
		args  []any
		where string
	)

	if user.Username != "" {
		field += " username=? "
		args = append(args, user.Username)
	}

	if user.Password != "" {
		field += " password=? "
		args = append(args, user.Password)
	}

	if user.Email != "" {
		field += " email=? "
		args = append(args, user.Email)
	}

	if user.DateOfBirth != "" {
		field += " date_of_birth=? "
		args = append(args, user.DateOfBirth)
	}

	if user.Address != "" {
		field += " address=? "
		args = append(args, user.Address)
	}

	// baris ini khusu wheres di sql
	if user.PhoneNumber != "" {
		where += " AND phone_number= ? "
		args = append(args, user.PhoneNumber)
	}

	fmt.Println(args...)
	sqlQuerry := "update users set " + field + " WHERE    deleted_at IS NULL" + where
	fmt.Println(sqlQuerry)
	_, err := db.Exec(sqlQuerry, args...)
	if err != nil {
		return "", err
	}
	outputStr := "\n[SUCCESS] Update Account successfully.\n\n"
	return outputStr, nil
}

func LogOutAccount(phoneNumber, password *string) string {
	*phoneNumber = ""
	*password = ""

	return "\n[SUCCESS] Log out success"
}
