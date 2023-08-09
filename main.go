package main

import (
	"account-service-app/controller"
	"account-service-app/entity"
	"account-service-app/helpers"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	ConnectionDB := os.Getenv("ConnectionDB")
	db, err := sql.Open("mysql", ConnectionDB)

	if err != nil {
		log.Fatal("connection failed to db", err.Error())
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	errPing := db.Ping()
	if errPing != nil {
		log.Fatal("error connect to db ", errPing.Error())
	} else {
		fmt.Println("success connect to db")
	}
	//close db conn
	defer db.Close()

	loop := true
	isLogin := false
	phoneNumber := ""
	for loop {
		//MENU
		fmt.Println("Pilih Menu:\n1. Register.\n2. Login.\n3. Profile.\n4. Update Account.\n5. Delete Account.\n6. Top Up.\n7. Transfer.\n8. Top Up history.\n9. Transfer History.\n10. Other Contact.\n11. Log Out.")
		var pilihan int
		fmt.Println("input pilihan anda:")
		fmt.Scanln(&pilihan)
		switch pilihan {

		case 1:
			fmt.Print("Register Account")
			newUser := entity.Users{Balance: 0}
			fmt.Print("\nEnter the data below:")

			//Entering Username
			fmt.Print("\nUsername\t: ")
			newUser.Username, err = helpers.Readline()
			if err != nil {
				log.Fatal("Error: ", err.Error())
			}

			//Entering email
			emailLoop := false
			for !emailLoop {
				fmt.Print("\nEmail\t\t: ")
				fmt.Scanln(&newUser.Email)
				emailIsValid, err := helpers.ValidationEmail(newUser.Email)
				if emailIsValid {
					emailLoop = true
				} else {
					log.Printf("Error: %s", err.Error())
				}
			}

			//Entering Password
			passLoop := false
			for !passLoop {
				fmt.Print("\nPassword\t: ")
				fmt.Scanln(&newUser.Password)
				passwordIsValid, err := helpers.ValidationPassword(newUser.Password)
				if passwordIsValid {
					passLoop = true
				} else {
					log.Printf("Error: %s", err.Error())
				}
			}

			//Entering Phone Number
			phoneLoop := false
			for !phoneLoop {
				fmt.Print("\nPhone Number\t: ")
				fmt.Scanln(&newUser.PhoneNumber)
				phoneNumberIsValid, err := helpers.ValidationPhoneNumber(newUser.PhoneNumber)
				if phoneNumberIsValid {
					phoneLoop = true
				} else {
					log.Printf("Error: %s", err.Error())
				}
			}

			//Entering Date of birth
			dateLoop := false
			for !dateLoop {
				fmt.Print("\nDate of Birth(YYYY-MM-DD): ")
				fmt.Scanln(&newUser.DateOfBirth)

				if newUser.DateOfBirth == "" {
					fmt.Print("dateofbirth cannot empty")
					dateLoop = false
				} else {
					birthdateIsValid, err := helpers.ValidationDateofBirth(newUser.DateOfBirth)
					if birthdateIsValid {
						dateLoop = true
					} else {
						log.Printf("Error: %s", err.Error())
					}
				}
			}

			//Entering address
			fmt.Print("\nAddress\t\t: ")
			fmt.Scanln(&newUser.Address)
			newUser.Address, err = helpers.Readline()
			if err != nil {
				log.Fatal("Error: ", err.Error())
			}

			//registering new user
			outputStr, err := controller.RegisterAccount(db, newUser)
			if err != nil {
				log.Printf("Error: %s", err.Error())
			} else {
				log.Printf("%s", outputStr)
			}

		case 2:
			fmt.Println("Login")
			userLogin := entity.Users{Balance: 0}

			fmt.Print("\nPhone Number\t: ")
			fmt.Scanln(&userLogin.PhoneNumber)

			//Entering Password
			fmt.Print("\nPassword\t: ")
			fmt.Scanln(&userLogin.Password)

			str, err := controller.Login(db, userLogin)
			if err != nil {

				log.Fatal("[FAILED] Failed to login account", err.Error())
				return
			} else {
				fmt.Println("")
				log.Print("succes", str)
				isLogin = true
				phoneNumber = userLogin.PhoneNumber

			}
		case 3:
			userProfile := entity.Users{PhoneNumber: phoneNumber}
			if !isLogin {
				fmt.Println("You are not login")
				return
			}
			_, err := controller.Profile(db, userProfile)
			if err != nil {

				log.Fatal("[FAILED] failed check users profile ", err.Error())
				return
			} else {
				fmt.Println("")

			}

		case 4:
			fmt.Println("Update Account")
		case 5:
			userDelete := entity.Users{PhoneNumber: phoneNumber}
			if !isLogin {
				fmt.Println("You are not Login")
				return
			}
			outputStr, err := controller.DeleteUser(db, userDelete)
			if err != nil {
				log.Fatal("[FAILED] failed delete account ", err.Error())
				return
			} else {
				log.Printf("%s", outputStr)
			}

		case 6:
			fmt.Println("Top Up")
		case 7:
			fmt.Println("Transfer")
		case 8:
			fmt.Println("Top Up History")
		case 9:
			fmt.Println("Transfer History")
		case 10:
			fmt.Println("Other Contact")

			if !isLogin {
				fmt.Println("You are not login")
				return
			}
			var otherUserPhoneNumber string
			otherUser := entity.Users{PhoneNumber: otherUserPhoneNumber}
			fmt.Print("\nEnter other user's phone number\t: ")
			fmt.Scanln(&otherUser.PhoneNumber)

			user, err := controller.ReadOtherUser(db, otherUser)
			if otherUser.PhoneNumber != phoneNumber {
				if err != nil {
					if err == sql.ErrNoRows {
						log.Println("User not found")
					}
					log.Fatal("[FAILED] failed to check users profile ", err.Error())
					return
				} else {
					outputStr := fmt.Sprintln("-----------------------------------------")
					outputStr += fmt.Sprintln("Other User Information")
					outputStr += fmt.Sprintln("-----------------------------------------")
					outputStr += fmt.Sprintf("User Name\t: %s\n", user.Username)
					outputStr += fmt.Sprintf("Birth Date\t: %s\n", user.DateOfBirth)
					outputStr += fmt.Sprintf("Address\t\t: %s\n", user.Address)
					outputStr += fmt.Sprintf("Email\t\t: %s\n", user.Email)
					outputStr += fmt.Sprintf("Phone Number\t: %s\n", user.PhoneNumber)
					outputStr += fmt.Sprintln("-----------------------------------------")
					fmt.Println(outputStr)
				}
			} else {
				log.Printf("Choose option 3 to see your account information")
			}

		case 11:
			fmt.Println("Log Out")
		}
	}
}
