package helpers

import (
	"fmt"
	"regexp"
	"time"
)

func ValidationEmail(email string) (bool, error) {
	if email == "" {
		return false, fmt.Errorf("email cannot empty")
	}

	patern := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

	regex := regexp.MustCompile(patern)
	if !regex.MatchString(email) {
		return false, fmt.Errorf("invalid email format")
	}
	return true, nil
}

func ValidationPassword(password string) (bool, error) {
	if password == "" {
		return false, fmt.Errorf("password cannot empty")
	}

	if len(password) < 5 {
		return false, fmt.Errorf("password must be at least 5 characters")
	}
	return true, nil
}

func ValidationPhoneNumber(phonenumber string) (bool, error) {
	if phonenumber == "" {
		return false, fmt.Errorf("phone number cannot empty")
	}
	charLowerCase := `abcdefghijklmnopqrstuvwxyz`
	charUpperCase := `ABCDEFGHIJKLMNOPQRSTUVWXYZ`
	if phonenumber == charLowerCase || phonenumber == charUpperCase {
		return false, fmt.Errorf("phone number cannot contain letters")
	}
	pattern := `[a-zA-Z0-9]`
	regex := regexp.MustCompile(pattern)
	if regex.MatchString(phonenumber) {
		return true, fmt.Errorf("phone number cannot contain special characters")
	}
	return true, nil
}

func ValidationDateofBirth(dateString string) (bool, error) {
	if len(dateString) > 0 {
		_, err := time.Parse("2006-01-02", dateString)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}
