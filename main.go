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
	password := ""
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
			userNameLoop := false
			for !userNameLoop {
				fmt.Print("\nUsername\t: ")
				fmt.Scanln(&newUser.Username)
				if err != nil {
					log.Fatal("Error: ", err.Error())
				}
				if newUser.Address == "" {
					fmt.Print("Address cannot empty")
					userNameLoop = false
				} else {
					userNameLoop = true
				}
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
				fmt.Print("\nDate of Birth(YYYY-MM-DD)\t: ")
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
			addressLoop := false
			for !addressLoop {
				fmt.Print("\nAddress\t\t: ")
				newUser.Address, err = helpers.Readline()
				if err != nil {
					log.Fatal("Error: ", err.Error())
				}
				if newUser.Address == "" {
					fmt.Print("Address cannot empty")
					addressLoop = false
				} else {
					addressLoop = true
				}
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
			userLogin := entity.Users{}

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
			userUpdate := entity.Users{PhoneNumber: phoneNumber}
			if !isLogin {
				fmt.Println("You are not login")
				return
			}
			updateMenu := `
				Select the section you want to update:

				[1].	User Name
				[2].	Password
				[3].	Email
				[4].	Birth Date
				[5].	Address
				[6].	Finish Update
					
				`

			updateLoop := true
			for updateLoop {
				fmt.Println(updateMenu)
				fmt.Print("\nEnter update menu option: ")
				var option int
				fmt.Scanln(&option)

				switch option {
				case 1: //update username

					fmt.Print("\nUsername\t: ")
					userUpdate.Username, err = helpers.Readline()
					if err != nil {
						log.Fatal("Error: ", err.Error())
					}

				case 2: //update password
					passLoop := false
					for !passLoop {
						fmt.Print("\nPassword\t: ")
						fmt.Scanln(&userUpdate.Password)
						passwordIsValid, err := helpers.ValidationPassword(userUpdate.Password)
						if passwordIsValid {
							passLoop = true
						} else {
							log.Printf("Error: %s", err.Error())
						}
					}
				case 3: //Update Email
					emailLoop := false
					for !emailLoop {
						fmt.Print("\nEmail\t\t: ")
						fmt.Scanln(&userUpdate.Email)
						emailIsValid, err := helpers.ValidationEmail(userUpdate.Email)
						if emailIsValid {
							emailLoop = true
						} else {
							log.Printf("Error: %s", err.Error())
						}
					}
				case 4: //update Birth Date
					dateLoop := false
					for !dateLoop {
						fmt.Print("\nDate of Birth(YYYY-MM-DD)\t: ")
						fmt.Scanln(&userUpdate.DateOfBirth)

						if userUpdate.DateOfBirth == "" {
							fmt.Print("dateofbirth cannot empty")
							dateLoop = false
						} else {
							birthdateIsValid, err := helpers.ValidationDateofBirth(userUpdate.DateOfBirth)
							if birthdateIsValid {
								dateLoop = true
							} else {
								log.Printf("Error: %s", err.Error())
							}
						}
					}
				case 5: //Update Address

					fmt.Print("\nAddress\t\t: ")
					userUpdate.Address, err = helpers.Readline()
					if err != nil {
						log.Fatal("Error: ", err.Error())
					}
				case 6: //Final update
					updateLoop = false
					fmt.Printf("\nUpdate complete.\n")

				}
				outputStr, err := controller.UpdateAccount(db, userUpdate)
				if err != nil {
					log.Printf("Error: %s", err.Error())
				} else {
					log.Printf("%s", outputStr)
				}

			}

		case 5:
			fmt.Println("Delete Account")
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
			if !isLogin {
				fmt.Println("You are not Login")
				return
			} else {
				var topupAmount float64 = 0
				fmt.Print("\nEnter top amount: ")
				fmt.Scanln(&topupAmount)
				str, err := controller.TopUp(db, phoneNumber, topupAmount)
				if err != nil {
					fmt.Println("error dia")
					fmt.Printf("\n")
					log.Printf("\033[91mError: %s\033[0m\n", err.Error())
				} else {
					fmt.Printf("\n")
					log.Printf("\033[92m%s\033[0m\n", str)
				}
			}

		case 7:
			fmt.Println("Transfer")
			if !isLogin {
				fmt.Println("You are not Login")
				return
			}
			var phoneRecipient string
			var tranferAmount float64

			fmt.Print("\nPhone Number\t: ")
			fmt.Scanln(&phoneRecipient)

			//Entering Password
			fmt.Print("\nAmount\t\t: ")
			fmt.Scanln(&tranferAmount)

			outputStr, err := controller.Transfer(db, phoneNumber, phoneRecipient, tranferAmount)
			if err != nil {
				log.Fatal("[FAILED] failed transfer", err.Error())
				return
			} else {
				log.Printf("%s", outputStr)
			}
		case 8:
			fmt.Println("Top Up History")

			if !isLogin {
				fmt.Println("You are not login")
				return
			}

			histories, err := controller.HistoryTopUp(db, phoneNumber)

			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf("\n")
			fmt.Println("-----------------------------------------")
			fmt.Printf("Your top-up history: \n")
			fmt.Println("-----------------------------------------")
			topupCounter := 0
			// Print top-up histories
			for _, history := range histories {
				topupCounter++
				fmt.Printf("User ID\t: %s\n", history.Id)
				fmt.Printf("Amount\t: %.2f\n", history.Amount)
				fmt.Printf("Time\t: %s\n", history.CreatedAt.Format("2006-01-02 15:04:05"))
				fmt.Println("-----------------------------------------")
			}
			fmt.Println("Count:", topupCounter)
		case 9:
			fmt.Println("Transfer History")
			if !isLogin {
				fmt.Println("You are not Login")
				return
			}
			tfHistory, err := controller.TransferHistory(db, phoneNumber)
			if err != nil {
				log.Fatal(err)
			}

			for _, History := range tfHistory {
				if History.IsSender {
					fmt.Println("Transaction Type: Sender")
				} else {
					fmt.Println("Transaction Type: Recipient")
				}
				fmt.Printf("Transfer ID\t: %s\n", History.ID)
				fmt.Printf("User Name\t: %s\n", History.Username)
				fmt.Printf("Phone Number\t: %s\n", History.PhoneNumber)
				fmt.Printf("Transfer Amount\t: %.2f\n", History.Amount)
				fmt.Printf("Transaction Time: %s\n", History.CreatedAt)

				fmt.Println("-----------------------------------")
			}
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
			str := controller.LogOutAccount(&phoneNumber, &password)
			fmt.Printf("\033[92m%s\033[0m\n", str)
			isLogin = false
		}
	}
}
