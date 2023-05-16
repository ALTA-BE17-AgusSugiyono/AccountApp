package controller

import (
	"AccountApp/config"
	"AccountApp/topup"
	"fmt"
)

func TopUp(userID int) {

	var amount int
	fmt.Print("Enter the top-up amount: ")
	fmt.Scan(&amount)

	topupModel := topup.TopupModel{DB: config.DB}

	// Call the method to perform the top-up
	err := topupModel.TopUpAccount(userID, amount)
	if err != nil {
		fmt.Println("Gagal Top Up:", err)
		return
	}

	fmt.Println("Top-up successful!")
}
