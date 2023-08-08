package main

import (
	"account-service-app/controller"
	"account-service-app/entity"
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

	for loop {
		//MENU
		fmt.Println("Pilih Menu:\n1. Register.\n2. Login.\n3. Profile.\n4. Update Account.\n5. Delete Account.\n6. Top Up.\n7. Transfer.\n8. Top Up history.\n9. Transfer History.\n10. Other Contact.\n11. Log Out.")
		var pilihan int
		fmt.Println("input pilihan anda:")
		fmt.Scanln(&pilihan)
		switch pilihan {
		case 1:

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

			}
		case 3:
			if !isLogin {
				fmt.Println("You are not login")
				return
			} else {
				fmt.Println("you have login")
			}
		case 4:
			fmt.Println("Update Account")
		case 5:
			fmt.Println("Delete Account")
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
		case 11:
			fmt.Println("Log Out")
		}
	}
}
