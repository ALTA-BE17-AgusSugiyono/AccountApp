package main

import (
	"AccountApp/config"
	"AccountApp/controller"
	"fmt"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}

	// fmt.Println("Successfully connected to database!")
	var input int
	var isLoggedIn bool
	var loggedInUserID int
	var loggedInUserPhoneNumber string

	for {
		if !isLoggedIn {
			fmt.Println("\nMenu:")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("0. Exit")
		} else {
			fmt.Println("\nMenu:")
			fmt.Println("3. Read account")
			fmt.Println("4. Update account")
			fmt.Println("5. Delete account")
			fmt.Println("6. Top up")
			fmt.Println("7. Transfer")
			fmt.Println("8. History top up")
			fmt.Println("9. History transfer")
			fmt.Println("10. View other user profile")
			fmt.Println("0. Logout")
		}

		fmt.Print("Input: ")
		fmt.Scan(&input)

		if !isLoggedIn {
			switch input {
			case 1:
				fmt.Println("Register Menu")
				controller.Register()
			case 2:
				fmt.Println("Login Menu")
				isLoggedIn, loggedInUserID, loggedInUserPhoneNumber = controller.Login()
			case 0:
				fmt.Println("Kamu Keluar Dari Program...")
				return
			default:
				fmt.Println("Invalid input. Please input a valid option.")
			}
		} else {
			switch input {
			case 3:
				fmt.Println("View Profile Menu")
				controller.ViewProfile(loggedInUserPhoneNumber, loggedInUserID)
			case 4:
				fmt.Println("Update Profile Menu")
				controller.UpdateAccount(loggedInUserPhoneNumber, loggedInUserID)
			case 5:
				fmt.Println("Delete Account Menu")
				controller.DeleteAccount(loggedInUserPhoneNumber, loggedInUserID)
			case 6:
				fmt.Println("Top-up Menu")
				controller.TopUp(loggedInUserID)
			case 7:

			case 8:

			case 9:

			case 10:

			case 0:
				fmt.Println("Logging out...")
				isLoggedIn = false
				loggedInUserPhoneNumber = ""
				loggedInUserID = 0
			default:
				fmt.Println("Invalid input. Please input a valid option.")
			}
		}
	}
}
