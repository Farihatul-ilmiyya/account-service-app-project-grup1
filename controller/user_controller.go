package controller

import (
	"account-service-app/entity"
	"account-service-app/helpers"
	"database/sql"

	"github.com/google/uuid"
)

func RegisterAccount(db *sql.DB, user entity.Users) (string, error) {
	Uuid := uuid.New()
	//validating email
	emailIsValid, err := helpers.ValidationEmail(user.Email)
	if !emailIsValid {
		return "", err
	}

	//validating password
	passwordIsValid, err := helpers.ValidationPassword(user.Password)
	if !passwordIsValid {
		return "", err
	}
	passHashing := ""
	if passwordIsValid {
		passHashing = helpers.HashPassword(user.Password)
	} else {
		return "", err
	}

	//validating phonenumber
	phoneNumberIsValid, err := helpers.ValidationPhoneNumber(user.PhoneNumber)
	if !phoneNumberIsValid {
		return "", err
	}

	//validating dateofbirth
	isValid, birthdate, err := helpers.ValidationDateofBirth(user.DateOfBirth)
	if !isValid {
		return "", err
	}
	_, err = db.Exec("INSERT INTO users(id, username, email, password, phone_number, date_of_birth, address, balance) VALUE (?, ?, ?, ?, ?, ?, ?, ?)", Uuid, user.Username, user.Email, passHashing, user.PhoneNumber, birthdate, user.Address, user.Balance)

	defer db.Close()

	if err != nil {
		return "", err
	}

	outputStr := "\n[SUCCESS] Account registered successfully.\n\n"
	return outputStr, nil
}
