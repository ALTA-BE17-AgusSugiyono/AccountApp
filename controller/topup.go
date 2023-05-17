package controller

import (
	"AccountApp/config"
	"AccountApp/topup"
	"AccountApp/users"
	"fmt"
)

func TopUp(userID int) {
	usersModel := users.UsersModel{DB: config.DB}

	// Get the logged-in user based on the provided userID
	loggedInUser, err := usersModel.GetUserByID(userID)
	if err != nil {
		fmt.Println("", err)
		return
	}

	var phoneNumber string
	fmt.Print("Enter your phone number: ")
	fmt.Scan(&phoneNumber)

	var amount int
	fmt.Print("Enter top up amount: ")
	fmt.Scan(&amount)

	if amount <= 0 {
		fmt.Println("saldo tidak boleh Kosong.")
		return
	}

	if phoneNumber != loggedInUser.PhoneNumber {
		fmt.Println("Nomer yang anda Masukan Tidak Sesuai.")
		return
	}

	topupModel := topup.TopupModel{DB: config.DB}
	err = topupModel.TopUpAccount(phoneNumber, amount)
	if err != nil {
		fmt.Println("Failed to top-up:", err)
		return
	}

	fmt.Println("Top-up successful!")
}
